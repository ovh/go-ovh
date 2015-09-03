// Package govh provides a HTTP wrapper for the OVH API.
package govh

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Endpoint reprensents an API endpoint
type Endpoint string

// Endpoints
const (
	OvhEU    Endpoint = "https://api.ovh.com/1.0"
	OvhCA             = "https://ca.api.ovh.com/1.0"
	Runabove          = "https://api.runabove.com/1.0"
)

// Errors
var (
	ErrAPIDown = errors.New("govh: the OVH API is down, it does't respond to /time anymore")
)

// Client represents a client to call the OVH API
type Client struct {
	// Self generated tokens. Create one by visiting
	// https://eu.api.ovh.com/createApp/
	appKey    string
	appSecret string

	// Token that must me validated
	consumerKey string

	// API endpoint
	endpoint Endpoint

	// Ensures that the delay function is only called once
	once sync.Once
}

// NewClient represents a new client to call the API
func NewClient(endpoint Endpoint, appKey, appSecret, consumerKey string) *Client {
	return &Client{
		appKey:      appKey,
		appSecret:   appSecret,
		consumerKey: consumerKey,
		endpoint:    endpoint,
	}
}

// Ping performs a ping to OVH API.
// In fact, ping is just a /auth/time call, in order to check if API is up.
func (c *Client) Ping() error {
	_, err := getTime(c.endpoint)
	return err
}

// Delay represents the delay between the machine that runs the code and the
// OVH API. The delay shouldn't change, let's do it only once.
func (c *Client) Delay() (time.Duration, error) {
	return delay(c.endpoint, &c.once)
}

// Time returns time from the OVH API, by asking GET /auth/time.
func (c *Client) Time() (*time.Time, error) {
	return getTime(c.endpoint)
}

// Get is a wrapper for the GET method
func (c *Client) Get(url string, resType interface{}) error {
	return c.callAPI(url, "GET", nil, resType)
}

// Post is a wrapper for the POST method
func (c *Client) Post(url string, reqBody, resType interface{}) error {
	return c.callAPI(url, "POST", reqBody, resType)
}

// Put is a wrapper for the PUT method
func (c *Client) Put(url string, reqBody, resType interface{}) error {
	return c.callAPI(url, "PUT", reqBody, resType)
}

// Delete is a wrapper for the DELETE method
func (c *Client) Delete(url string, resType interface{}) error {
	return c.callAPI(url, "DELETE", nil, resType)
}

// CallAPI makes a new call to the OVH API
// ApplicationKey, ApplicationSecret and ConsumerKey must be set on Client
// Returns the unmarshal json object or error if any occured
func (c *Client) callAPI(url, method string, reqBody, resType interface{}) error {
	url = string(c.endpoint) + url

	// Encode the request params
	params, err := c.encodeParams(reqBody)
	if err != nil {
		return err
	}

	// Create a new request
	request, err := http.NewRequest(method, url, bytes.NewReader(params))
	if err != nil {
		return err
	}

	// Set the headers and sign the request
	if err := c.setHeaders(request, method, url, string(params)); err != nil {
		return err
	}

	// Run the request
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}

	// Unmarshal the result into the resType if possible
	if err := c.getResponse(response, resType); err != nil {
		return err
	}

	return nil
}

// encodeParams encode the body params as json
func (c *Client) encodeParams(reqBody interface{}) (params []byte, err error) {
	if reqBody == nil {
		return
	}

	params, err = json.Marshal(reqBody)

	return
}

// setHeaders set the request with the proper headers
func (c *Client) setHeaders(r *http.Request, method, url, params string) error {
	// Get the timestamp with the delay
	delay, err := delay(c.endpoint, &c.once)
	if err != nil {
		return err
	}

	timestamp := time.Now().Add(delay).Unix()

	// Sig the header
	sig := sig(c.appSecret, c.consumerKey, method, url, params, timestamp)
	for h, v := range map[string]string{
		"Content-Type":      "application/json",
		"X-Ovh-Timestamp":   strconv.FormatInt(timestamp, 10),
		"X-Ovh-Application": c.appKey,
		"X-Ovh-Consumer":    c.consumerKey,
		"X-Ovh-Signature":   sig,
	} {
		r.Header.Add(h, v)
	}

	return nil
}

// getResult check the response and unmarshals it into the response type if needed
func (c *Client) getResponse(response *http.Response, resType interface{}) error {
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
			return err
		}

		return apiError
	}

	// Nothing to unmarshal
	if len(body) == 0 || resType == nil {
		return nil
	}

	return json.Unmarshal(body, &resType)
}

// sig is used to compute the signature of a request
func sig(as, ck, method, query, body string, timestamp int64) string {
	h := sha1.New()
	sig := strings.Join([]string{
		as,
		ck,
		method,
		query,
		body,
		strconv.FormatInt(timestamp, 10),
	}, "+")
	io.WriteString(h, sig)
	return "$1$" + hex.EncodeToString(h.Sum(nil))
}

// delay is a function to be overwritten during the tests, it return duration
// delay between the current machine and the given API
var delay = func(e Endpoint, o *sync.Once) (time.Duration, error) {
	var d time.Duration
	var err error

	o.Do(func() {
		ovhTime, err := getTime(e)
		if err != nil {
			return
		}

		d = time.Since(*ovhTime)
	})

	return d, err
}

// getTime is a function to be overwritten during the tests, it returns time
// from for a given endpoint
var getTime = func(endpoint Endpoint) (*time.Time, error) {
	url := string(endpoint) + "/auth/time"

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/json")

	result, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	if result.StatusCode != http.StatusOK {
		return nil, ErrAPIDown
	}

	defer result.Body.Close()
	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}

	ts, err := strconv.Atoi(string(body))
	if err != nil {
		return nil, err
	}

	t := time.Unix(int64(ts), 0)

	return &t, nil
}
