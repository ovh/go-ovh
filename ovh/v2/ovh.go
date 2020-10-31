// Package ovh provides a HTTP wrapper for the OVH API.
package ovh

import (
	"bytes"
	"context"
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"sync"
	"time"
)

// DefaultTimeout api requests after 180s
const DefaultTimeout = 180 * time.Second

const (
	OvhTimestampHeader      = "X-Ovh-Timestamp"
	OvhConsumerKeyHeader    = "X-Ovh-Consumer"
	OvhApplicationKeyHeader = "X-Ovh-Application"
	OvhSignatureHeader      = "X-Ovh-Signature"
	OvhRequestIDHeader      = "X-Ovh-QueryID"
)

// Endpoints
const (
	OvhEU        = "https://eu.api.ovh.com/1.0"
	OvhCA        = "https://ca.api.ovh.com/1.0"
	OvhUS        = "https://api.us.ovhcloud.com/1.0"
	KimsufiEU    = "https://eu.api.kimsufi.com/1.0"
	KimsufiCA    = "https://ca.api.kimsufi.com/1.0"
	SoyoustartEU = "https://eu.api.soyoustart.com/1.0"
	SoyoustartCA = "https://ca.api.soyoustart.com/1.0"
)

// Endpoints conveniently maps endpoints names to their URI for external configuration
var Endpoints = map[string]string{
	"ovh-eu":        OvhEU,
	"ovh-ca":        OvhCA,
	"ovh-us":        OvhUS,
	"kimsufi-eu":    KimsufiEU,
	"kimsufi-ca":    KimsufiCA,
	"soyoustart-eu": SoyoustartEU,
	"soyoustart-ca": SoyoustartCA,
}

// Errors
var (
	ErrAPIDown = errors.New("go-vh: the OVH API is down, it does't respond to /time anymore")
)

// Client represents a client to call the OVH API
type Client struct {
	// Self generated tokens. Create one by visiting
	// https://eu.api.ovh.com/createApp/
	// AppKey holds the Application key
	AppKey string

	// AppSecret holds the Application secret key
	AppSecret string

	// ConsumerKey holds the user/app specific token. It must have been validated before use.
	ConsumerKey string

	// API endpoint
	endpoint string

	// Client is the underlying HTTP client used to run the requests. It may be overloaded but a default one is instanciated in ``NewClient`` by default.
	Client *http.Client

	// Logger is used to log HTTP requests and responses.
	Logger Logger

	// Ensures that the timeDelta function is only ran once
	// sync.Once would consider init done, even in case of error
	// hence a good old flag
	timeDeltaMutex *sync.Mutex
	timeDeltaDone  bool
	timeDelta      time.Duration
	Timeout        time.Duration
}

// NewClient represents a new client to call the API
func NewClient(endpoint, appKey, appSecret, consumerKey string) (*Client, error) {
	client := Client{
		AppKey:         appKey,
		AppSecret:      appSecret,
		ConsumerKey:    consumerKey,
		Client:         &http.Client{},
		timeDeltaMutex: &sync.Mutex{},
		timeDeltaDone:  false,
		Timeout:        time.Duration(DefaultTimeout),
	}

	// Get and check the configuration
	if err := client.loadConfig(endpoint); err != nil {
		return nil, err
	}
	return &client, nil
}

// NewEndpointClient will create an API client for specified
// endpoint and load all credentials from environment or
// configuration files
func NewEndpointClient(endpoint string) (*Client, error) {
	return NewClient(endpoint, "", "", "")
}

// NewDefaultClient will load all it's parameter from environment
// or configuration files
func NewDefaultClient() (*Client, error) {
	return NewClient("", "", "", "")
}

//
// High level helpers
//

// Ping performs a ping to OVH API.
// In fact, ping is just a /auth/time call, in order to check if API is up.
func (c *Client) Ping() error {
	_, err := c.getTime()
	return err
}

// TimeDelta represents the delay between the machine that runs the code and the
// OVH API. The delay shouldn't change, let's do it only once.
func (c *Client) TimeDelta() (time.Duration, error) {
	return c.getTimeDelta()
}

// Time returns time from the OVH API, by asking GET /auth/time.
func (c *Client) Time() (*time.Time, error) {
	return c.getTime()
}

//
// Common request wrappers
//

// Get is a wrapper for the GET method
func (c *Client) Get(path string, queryParams url.Values, resType interface{}) error {
	return c.CallAPI(http.MethodGet, path, queryParams, nil, resType, true)
}

// GetUnAuth is a wrapper for the unauthenticated GET method
func (c *Client) GetUnAuth(path string, queryParams url.Values, resType interface{}) error {
	return c.CallAPI(http.MethodGet, path, queryParams, nil, resType, false)
}

// Post is a wrapper for the POST method
func (c *Client) Post(path string, queryParams url.Values, reqBody interface{}, resType interface{}) error {
	return c.CallAPI(http.MethodPost, path, queryParams, reqBody, resType, true)
}

// PostUnAuth is a wrapper for the unauthenticated POST method
func (c *Client) PostUnAuth(path string, queryParams url.Values, reqBody interface{}, resType interface{}) error {
	return c.CallAPI(http.MethodPost, path, queryParams, reqBody, resType, false)
}

// Put is a wrapper for the PUT method
func (c *Client) Put(path string, queryParams url.Values, reqBody interface{}, resType interface{}) error {
	return c.CallAPI(http.MethodPut, path, queryParams, reqBody, resType, true)
}

// PutUnAuth is a wrapper for the unauthenticated PUT method
func (c *Client) PutUnAuth(path string, queryParams url.Values, reqBody interface{}, resType interface{}) error {
	return c.CallAPI(http.MethodPut, path, queryParams, reqBody, resType, false)
}

// Delete is a wrapper for the DELETE method
func (c *Client) Delete(path string, queryParams url.Values, reqBody interface{}, resType interface{}) error {
	return c.CallAPI(http.MethodDelete, path, queryParams, reqBody, resType, true)
}

// DeleteUnAuth is a wrapper for the unauthenticated DELETE method
func (c *Client) DeleteUnAuth(path string, queryParams url.Values, reqBody interface{}, resType interface{}) error {
	return c.CallAPI(http.MethodDelete, path, queryParams, reqBody, resType, false)
}

// GetWithContext is a wrapper for the GET method
func (c *Client) GetWithContext(ctx context.Context, path string, queryParams url.Values, resType interface{}) error {
	return c.CallAPIWithContext(ctx, http.MethodGet, path, queryParams, nil, resType, true)
}

// GetUnAuthWithContext is a wrapper for the unauthenticated GET method
func (c *Client) GetUnAuthWithContext(ctx context.Context, path string, queryParams url.Values, resType interface{}) error {
	return c.CallAPIWithContext(ctx, http.MethodGet, path, queryParams, nil, resType, false)
}

// PostWithContext is a wrapper for the POST method
func (c *Client) PostWithContext(ctx context.Context, path string, queryParams url.Values, reqBody interface{}, resType interface{}) error {
	return c.CallAPIWithContext(ctx, http.MethodPost, path, queryParams, reqBody, resType, true)
}

// PostUnAuthWithContext is a wrapper for the unauthenticated POST method
func (c *Client) PostUnAuthWithContext(ctx context.Context, path string, queryParams url.Values, reqBody interface{}, resType interface{}) error {
	return c.CallAPIWithContext(ctx, http.MethodPost, path, queryParams, reqBody, resType, false)
}

// PutWithContext is a wrapper for the PUT method
func (c *Client) PutWithContext(ctx context.Context, path string, queryParams url.Values, reqBody interface{}, resType interface{}) error {
	return c.CallAPIWithContext(ctx, http.MethodPut, path, queryParams, reqBody, resType, true)
}

// PutUnAuthWithContext is a wrapper for the unauthenticated PUT method
func (c *Client) PutUnAuthWithContext(ctx context.Context, path string, queryParams url.Values, reqBody interface{}, resType interface{}) error {
	return c.CallAPIWithContext(ctx, http.MethodPut, path, queryParams, reqBody, resType, false)
}

// DeleteWithContext is a wrapper for the DELETE method
func (c *Client) DeleteWithContext(ctx context.Context, path string, queryParams url.Values, reqBody interface{}, resType interface{}) error {
	return c.CallAPIWithContext(ctx, http.MethodDelete, path, queryParams, reqBody, resType, true)
}

// DeleteUnAuthWithContext is a wrapper for the unauthenticated DELETE method
func (c *Client) DeleteUnAuthWithContext(ctx context.Context, path string, queryParams url.Values, reqBody interface{}, resType interface{}) error {
	return c.CallAPIWithContext(ctx, http.MethodDelete, path, queryParams, reqBody, resType, false)
}

// timeDelta returns the time  delta between the host and the remote API
func (c *Client) getTimeDelta() (time.Duration, error) {

	if !c.timeDeltaDone {
		// Ensure only one thread is updating
		c.timeDeltaMutex.Lock()

		// Ensure that the mutex will be released on return
		defer c.timeDeltaMutex.Unlock()

		// Did we wait ? Maybe no more needed
		if !c.timeDeltaDone {
			ovhTime, err := c.getTime()
			if err != nil {
				return 0, err
			}

			c.timeDelta = time.Since(*ovhTime)
			c.timeDeltaDone = true
		}
	}

	return c.timeDelta, nil
}

// getTime t returns time from for a given api client endpoint
func (c *Client) getTime() (*time.Time, error) {
	var timestamp int64

	err := c.GetUnAuth("/auth/time", nil, &timestamp)
	if err != nil {
		return nil, err
	}

	serverTime := time.Unix(timestamp, 0)
	return &serverTime, nil
}

// getLocalTime is a function to be overwritten during the tests, it return the time
// on the the local machine
var getLocalTime = func() time.Time {
	return time.Now()
}

// getEndpointForSignature is a function to be overwritten during the tests, it returns a
// the endpoint
var getEndpointForSignature = func(c *Client) string {
	return c.endpoint
}

// NewRequest returns a new HTTP request
func (c *Client) NewRequest(method, partialPath string, queryParams url.Values, reqBody interface{}, needAuth bool) (*http.Request, error) {
	var body []byte
	var err error

	//var reader io.Reader
	if reqBody != nil {
		body, err = json.Marshal(reqBody)
		if err != nil {
			return nil, err
		}
		//reader = bytes.NewReader(body)
	}
	target, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	partialPathParsed, err := url.Parse(partialPath)
	for k, values := range partialPathParsed.Query() {
		for _, value := range values {
			queryParams.Add(k, value)
		}
	}
	target.Path = path.Join(target.Path, partialPathParsed.Path)
	target.RawQuery = queryParams.Encode()

	req, err := http.NewRequest(method, target.String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	// Inject headers
	if body != nil {
		req.Header.Add("Content-Type", "application/json;charset=utf-8")
	}
	req.Header.Add(OvhApplicationKeyHeader, c.AppKey)
	req.Header.Add("Accept", "application/json")

	// Inject signature. Some methods do not need authentication, especially /time,
	// /auth and some /order methods are actually broken if authenticated.
	if needAuth {
		timeDelta, err := c.TimeDelta()
		if err != nil {
			return nil, err
		}

		timestamp := getLocalTime().Add(-timeDelta).Unix()

		req.Header.Add(OvhTimestampHeader, strconv.FormatInt(timestamp, 10))
		req.Header.Add(OvhConsumerKeyHeader, c.ConsumerKey)

		h := sha1.New()
		h.Write([]byte(fmt.Sprintf("%s+%s+%s+%s%s+%s+%d",
			c.AppSecret,
			c.ConsumerKey,
			method,
			getEndpointForSignature(c),
			partialPath,
			body,
			timestamp,
		)))
		req.Header.Add(OvhSignatureHeader, fmt.Sprintf("$1$%x", h.Sum(nil)))
	}

	// Send the request with requested timeout
	c.Client.Timeout = c.Timeout

	return req, nil
}

// Do sends an HTTP request and returns an HTTP response
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	if c.Logger != nil {
		c.Logger.LogRequest(req)
	}
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	if c.Logger != nil {
		c.Logger.LogResponse(resp)
	}
	return resp, nil
}

// CallAPI is the lowest level call helper. If needAuth is true,
// inject authentication headers and sign the request.
//
// Request signature is a sha1 hash on following fields, joined by '+':
// - applicationSecret (from Client instance)
// - consumerKey (from Client instance)
// - capitalized method (from arguments)
// - full request url, including any query string argument
// - full serialized request body
// - server current time (takes time delta into account)
//
// Call will automatically assemble the target url from the endpoint
// configured in the client instance and the path argument. If the reqBody
// argument is not nil, it will also serialize it as json and inject
// the required Content-Type header.
//
// If everything went fine, unmarshall response into resType and return nil
// otherwise, return the error
func (c *Client) CallAPI(method, path string, queryParams url.Values, reqBody interface{}, resType interface{}, needAuth bool) error {
	return c.CallAPIWithContext(context.Background(), method, path, queryParams, reqBody, resType, needAuth)
}

// CallAPIWithContext is the lowest level call helper. If needAuth is true,
// inject authentication headers and sign the request.
//
// Request signature is a sha1 hash on following fields, joined by '+':
// - applicationSecret (from Client instance)
// - consumerKey (from Client instance)
// - capitalized method (from arguments)
// - full request url, including any query string argument
// - full serialized request body
// - server current time (takes time delta into account)
//
// Context is used by http.Client to handle context cancelation
//
// Call will automatically assemble the target url from the endpoint
// configured in the client instance and the path argument. If the reqBody
// argument is not nil, it will also serialize it as json and inject
// the required Content-Type header.
//
// If everything went fine, unmarshall response into resType and return nil
// otherwise, return the error
func (c *Client) CallAPIWithContext(ctx context.Context, method, path string, queryParams url.Values, reqBody interface{}, resType interface{}, needAuth bool) error {
	req, err := c.NewRequest(method, path, queryParams, reqBody, needAuth)

	if err != nil {
		return err
	}
	req = req.WithContext(ctx)
	response, err := c.Do(req)
	if err != nil {
		return err
	}
	return c.UnmarshalResponse(response, resType)
}

// UnmarshalResponse checks the response and unmarshals it into the response
// type if needed Helper function, called from CallAPI
func (c *Client) UnmarshalResponse(response *http.Response, resType interface{}) error {
	// Read all the response body
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	// < 200 && >= 300 : API error
	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusMultipleChoices {
		apiError := &APIError{Code: response.StatusCode}
		if err = json.Unmarshal(body, apiError); err != nil {
			apiError.Message = string(body)
		}
		apiError.QueryID = response.Header.Get(OvhRequestIDHeader)

		return apiError
	}

	// Nothing to unmarshal
	if len(body) == 0 || resType == nil {
		return nil
	}

	d := json.NewDecoder(bytes.NewReader(body))
	d.UseNumber()
	return d.Decode(&resType)
}
