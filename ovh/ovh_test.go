package ovh

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/maxatome/go-testdeep/helpers/tdsuite"
	"github.com/maxatome/go-testdeep/td"
)

const (
	// In case you wonder, these are real *revoked* credentials
	MockApplicationKey    = "TDPKJdwZwAQPwKX2"
	MockApplicationSecret = "9ufkBmLaTQ9nz5yMUlg79taH0GNnzDjk"
	MockConsumerKey       = "5mBuy6SUQcRw2ZUxg0cG68BoDKpED4KY"

	MockTime = 1457018875
)

//
// Utils
//

func sbody(s string) io.ReadCloser {
	return ioutil.NopCloser(strings.NewReader(s))
}

//
// Tests
//

func TestClientEndpoint(t *testing.T) {
	require := td.Require(t)

	client, err := NewClient("ovh-eu", MockApplicationKey, MockApplicationSecret, MockConsumerKey)
	require.CmpNoError(err)
	td.Cmp(t, client.Endpoint(), OvhEU)

	client, err = NewClient("ovh-ca", MockApplicationKey, MockApplicationSecret, MockConsumerKey)
	require.CmpNoError(err)
	td.Cmp(t, client.Endpoint(), OvhCA)

	client, err = NewClient("https://example.org", MockApplicationKey, MockApplicationSecret, MockConsumerKey)
	require.CmpNoError(err)
	td.Cmp(t, client.Endpoint(), "https://example.org")
}

type MockSuite struct {
	client *Client
}

func (ms *MockSuite) Setup(t *td.T) error {
	httpmock.Activate()
	return nil
}

func (ms *MockSuite) PreTest(t *td.T, testName string) error {
	client, err := NewClient("ovh-eu", MockApplicationKey, MockApplicationSecret, MockConsumerKey)
	if err != nil {
		return err
	}
	ms.client = client
	return nil
}

func (ms *MockSuite) PostTest(t *td.T, testName string) error {
	httpmock.Reset()
	return nil
}

func (ms *MockSuite) Destroy(t *td.T) error {
	httpmock.DeactivateAndReset()
	return nil
}

func TestMockSuite(t *testing.T) {
	tdsuite.Run(t, &MockSuite{})
}

func (ms *MockSuite) TestPing(assert *td.T) {
	httpmock.RegisterResponder("GET", "https://eu.api.ovh.com/1.0/auth/time",
		httpmock.NewStringResponder(200, "0"))

	assert.CmpNoError(ms.client.Ping())
	assert.Cmp(httpmock.GetCallCountInfo()["GET https://eu.api.ovh.com/1.0/auth/time"], 1)
}

func (ms *MockSuite) TestTime(assert, require *td.T) {
	httpmock.RegisterResponder("GET", "https://eu.api.ovh.com/1.0/auth/time",
		httpmock.NewStringResponder(200, strconv.Itoa(MockTime)))

	serverTime, err := ms.client.Time()
	require.CmpNoError(err)
	assert.CmpLax(serverTime.Unix(), MockTime)
	assert.Cmp(httpmock.GetCallCountInfo()["GET https://eu.api.ovh.com/1.0/auth/time"], 1)
}

func (ms *MockSuite) TestGetTimeDelta(assert, require *td.T) {
	httpmock.RegisterResponder("GET", "https://eu.api.ovh.com/1.0/auth/time",
		httpmock.NewStringResponder(200, strconv.FormatInt(time.Now().Unix()-10, 10)))

	delta, err := ms.client.TimeDelta()
	require.CmpNoError(err)
	assert.Between(delta.Seconds(), 9.0, 11.0, td.BoundsInIn)
	assert.Cmp(httpmock.GetCallCountInfo()["GET https://eu.api.ovh.com/1.0/auth/time"], 1)
}

func (ms *MockSuite) TestError500HTML(assert, require *td.T) {
	errHTML := `<html><body><p>test</p></body></html>`
	httpmock.RegisterResponder("GET", "https://eu.api.ovh.com/1.0/test",
		httpmock.NewStringResponder(http.StatusServiceUnavailable, errHTML))

	err := ms.client.CallAPI("GET", "/test", nil, nil, false)
	assert.Cmp(err, &APIError{
		Code:    http.StatusServiceUnavailable,
		Message: errHTML,
	})
}

func (ms *MockSuite) TestAllAPIMethods(assert, require *td.T) {
	const payloadAuth = `{"call":"auth"}`
	const payloadUnAuth = `{"call":"unauth"}`
	var body = json.RawMessage(`{"a":"b","c":"d"}`)

	previous := getLocalTime
	getLocalTime = func() time.Time {
		return time.Unix(MockTime, 0)
	}
	assert.Cleanup(func() { getLocalTime = previous })

	checkAuthHeaders := func(assert *td.T, req *http.Request, signature string) {
		assert.Helper()
		if signature == "" {
			assert.Cmp(req.Header, td.Not(td.ContainsKey("X-Ovh-Timestamp")), "No X-Ovh-Timestamp for %s unauth call", req.Method)
			assert.Cmp(req.Header, td.Not(td.ContainsKey("X-Ovh-Consumer")), "No X-Ovh-Consumer for %s unauth call", req.Method)
			assert.Cmp(req.Header, td.Not(td.ContainsKey("X-Ovh-Signature")), "No X-Ovh-Signature for %s unauth call", req.Method)
		} else {
			assert.Cmp(req.Header["X-Ovh-Timestamp"], []string{strconv.Itoa(MockTime)}, "Right X-Ovh-Timestamp for %s auth call", req.Method)
			assert.Cmp(req.Header["X-Ovh-Consumer"], []string{MockConsumerKey}, "Right X-Ovh-Consumer for %s auth call", req.Method)
			assert.Cmp(req.Header["X-Ovh-Signature"], []string{signature}, "Right X-Ovh-Signature for %s auth call", req.Method)
		}
	}
	checkBody := func(assert *td.T, req *http.Request) {
		assert.Helper()
		if req.Method != "POST" && req.Method != "PUT" {
			assert.Cmp(req.Body, td.Smuggle(io.ReadAll, td.Empty()), "Body is empty for %s call", req.Method)
		} else {
			assert.Cmp(req.Body, td.Smuggle(json.RawMessage{}, td.JSON(`{"a":"b","c":"d"}`)), "Right body for %s call", req.Method)
		}
	}

	httpmock.RegisterResponder("GET", "https://eu.api.ovh.com/1.0/auth/time",
		httpmock.NewStringResponder(200, strconv.Itoa(MockTime)))

	mockSignatures := map[string]struct{ authSig, timeoutSig string }{
		"GET":    {authSig: "$1$e9556054b6309771395efa467c22e627407461ad", timeoutSig: "$1$1f0958be70f095ddaba525778a9ac1dcffac89f3"},
		"POST":   {authSig: "$1$ec2fb5c7a81f64723c77d2e5b609ae6f58a84fc1", timeoutSig: "$1$b592effcb3bc2d37860eceb06a1b17670fbe49c6"},
		"PUT":    {authSig: "$1$8a75a9e7c8e7296c9dbeda6a2a735eb6bd58ec4b", timeoutSig: "$1$6b27c2a693a0eb4980217046b2fe10d74ba796f0"},
		"DELETE": {authSig: "$1$a1eecd00b3b02b6cf5708b84b9ff42059a950d85", timeoutSig: "$1$bd59b15361548c388058009e00c508081e991e8b"},
	}
	buildMock := func(assert *td.T, method string) {
		method = strings.ToUpper(method)
		httpmock.RegisterResponder(method, "https://eu.api.ovh.com/1.0/auth", func(req *http.Request) (*http.Response, error) {
			checkAuthHeaders(assert, req, mockSignatures[method].authSig)
			checkBody(assert, req)
			return httpmock.NewStringResponse(200, payloadAuth), nil
		})
		httpmock.RegisterResponder(method, "https://eu.api.ovh.com/1.0/unauth", func(req *http.Request) (*http.Response, error) {
			checkAuthHeaders(assert, req, "")
			checkBody(assert, req)
			return httpmock.NewStringResponse(200, payloadUnAuth), nil
		})
		httpmock.RegisterResponder(method, "https://eu.api.ovh.com/1.0/authTO", func(req *http.Request) (*http.Response, error) {
			checkAuthHeaders(assert, req, mockSignatures[method].timeoutSig)
			checkBody(assert, req)
			time.Sleep(200 * time.Millisecond)
			return httpmock.NewStringResponse(200, `{"call":"authTO"}`), nil
		})
		httpmock.RegisterResponder(method, "https://eu.api.ovh.com/1.0/unauthTO", func(req *http.Request) (*http.Response, error) {
			checkAuthHeaders(assert, req, "")
			checkBody(assert, req)
			time.Sleep(200 * time.Millisecond)
			return httpmock.NewStringResponse(200, `{"call":"unauthTO"}`), nil
		})
	}

	// Tests without body: GET and DELETE
	for _, test := range []struct {
		method                                 string
		call, callUnAuth                       func(string, interface{}) error
		callWithContext, callUnAuthWithContext func(context.Context, string, interface{}) error
	}{
		{"Get", ms.client.Get, ms.client.GetUnAuth, ms.client.GetWithContext, ms.client.GetUnAuthWithContext},
		{"Delete", ms.client.Delete, ms.client.DeleteUnAuth, ms.client.DeleteWithContext, ms.client.DeleteUnAuthWithContext},
	} {
		assert.RunAssertRequire(test.method+" method", func(assert, require *td.T) {
			buildMock(assert, test.method)

			var res json.RawMessage
			err := test.call("/auth", &res)
			require.CmpNoError(err, "No errors for method %s with auth", test.method)
			assert.Cmp(res, td.JSON(payloadAuth), "Got expected payload for method %s with auth", test.method)

			res = json.RawMessage{}
			err = test.callWithContext(context.Background(), "/auth", &res)
			require.CmpNoError(err, "No errors for method %s with auth and context", test.method)
			assert.Cmp(res, td.JSON(payloadAuth), "Got expected payload for method %s with auth and context", test.method)

			res = json.RawMessage{}
			err = test.callUnAuth("/unauth", &res)
			require.CmpNoError(err, "No errors for method %s without auth", test.method)
			assert.Cmp(res, td.JSON(payloadUnAuth), "Got expected payload for method %s without auth", test.method)

			res = json.RawMessage{}
			err = test.callUnAuthWithContext(context.Background(), "/unauth", &res)
			require.CmpNoError(err, "No errors for method %s without auth and with context", test.method)
			assert.Cmp(res, td.JSON(payloadUnAuth), "Got expected payload for method %s without auth and with context", test.method)

			res = json.RawMessage{}
			ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
			assert.Cleanup(cancel)
			err = test.callWithContext(ctx, "/authTO", &res)
			assert.Empty(res, "Empty result after timeout for method %s with auth", test.method)
			assert.String(err,
				test.method+` "https://eu.api.ovh.com/1.0/authTO": context deadline exceeded`,
				"Timeout messsage for method %s with auth", test.method,
			)

			res = json.RawMessage{}
			ctx, cancel = context.WithTimeout(context.Background(), 100*time.Millisecond)
			assert.Cleanup(cancel)
			err = test.callUnAuthWithContext(ctx, "/unauthTO", &res)
			assert.Empty(res, "Empty result after timeout for method %s without auth", test.method)
			assert.String(err,
				test.method+` "https://eu.api.ovh.com/1.0/unauthTO": context deadline exceeded`,
				"Timeout messsage for method %s without auth", test.method,
			)
		})
	}

	// Tests with body: POST and PUT
	for _, test := range []struct {
		method                                 string
		call, callUnAuth                       func(string, interface{}, interface{}) error
		callWithContext, callUnAuthWithContext func(context.Context, string, interface{}, interface{}) error
	}{
		{"Post", ms.client.Post, ms.client.PostUnAuth, ms.client.PostWithContext, ms.client.PostUnAuthWithContext},
		{"Put", ms.client.Put, ms.client.PutUnAuth, ms.client.PutWithContext, ms.client.PutUnAuthWithContext},
	} {
		assert.RunAssertRequire(test.method+" method", func(assert, require *td.T) {
			buildMock(assert, test.method)

			var res json.RawMessage
			err := test.call("/auth", body, &res)
			require.CmpNoError(err, "No errors for method %s with auth", test.method)
			assert.Cmp(res, td.JSON(payloadAuth), "Got expected payload for method %s with auth", test.method)

			res = json.RawMessage{}
			err = test.callWithContext(context.Background(), "/auth", body, &res)
			require.CmpNoError(err, "No errors for method %s with auth and context", test.method)
			assert.Cmp(res, td.JSON(payloadAuth), "Got expected payload for method %s with auth and context", test.method)

			res = json.RawMessage{}
			err = test.callUnAuth("/unauth", body, &res)
			require.CmpNoError(err, "No errors for method %s without auth", test.method)
			assert.Cmp(res, td.JSON(payloadUnAuth), "Got expected payload for method %s without auth", test.method)

			res = json.RawMessage{}
			err = test.callUnAuthWithContext(context.Background(), "/unauth", body, &res)
			require.CmpNoError(err, "No errors for method %s without auth and with context", test.method)
			assert.Cmp(res, td.JSON(payloadUnAuth), "Got expected payload for method %s without auth and with context", test.method)

			res = json.RawMessage{}
			ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
			assert.Cleanup(cancel)
			err = test.callWithContext(ctx, "/authTO", body, &res)
			assert.String(err,
				test.method+` "https://eu.api.ovh.com/1.0/authTO": context deadline exceeded`,
				"Timeout messsage for method %s with auth", test.method,
			)
			assert.Empty(res, "Empty result after timeout for method %s with auth", test.method)

			res = json.RawMessage{}
			ctx, cancel = context.WithTimeout(context.Background(), 100*time.Millisecond)
			assert.Cleanup(cancel)
			err = test.callUnAuthWithContext(ctx, "/unauthTO", body, &res)
			assert.String(err,
				test.method+` "https://eu.api.ovh.com/1.0/unauthTO": context deadline exceeded`,
				"Timeout messsage for method %s without auth", test.method,
			)
			assert.Empty(res, "Empty result after timeout for method %s without auth", test.method)
		})
	}
}

// Mock ReadCloser, always failing
type ErrorReadCloser struct{}

func (ErrorReadCloser) Read(p []byte) (int, error) {
	return 0, fmt.Errorf("ErrorReader")
}
func (ErrorReadCloser) Close() error { return nil }

func TestGetResponse(t *testing.T) {
	client := Client{}

	// Nominal
	var apiInt int
	err := client.UnmarshalResponse(&http.Response{
		StatusCode: 200,
		Body:       sbody(`42`),
	}, &apiInt)
	td.CmpNoError(t, err, "Can Client.UnmarshalResponse with status 200 and int body")
	td.Cmp(t, apiInt, 42)

	// Nominal: empty body
	err = client.UnmarshalResponse(&http.Response{
		StatusCode: 200,
		Body:       sbody(``),
	}, nil)
	td.CmpNoError(t, err, "Can Client.UnmarshalResponse with status 200 and empty body")

	// Error
	apiInt = 0
	err = client.UnmarshalResponse(&http.Response{
		StatusCode: 400,
		Body:       sbody(`{"code": 400, "message": "Ooops..."}`),
	}, &apiInt)
	td.Cmp(t, err, &APIError{
		Code:    400,
		Message: "Ooops...",
	}, "Can parse a API error")
	td.Cmp(t, apiInt, 0)

	// Error: body read error
	err = client.UnmarshalResponse(&http.Response{
		Body: ErrorReadCloser{},
	}, nil)
	td.CmpString(t, err, "ErrorReader")

	// Error: HTTP Error + broken json
	err = client.UnmarshalResponse(&http.Response{
		StatusCode: 400,
		Body:       sbody(`{"code": 400, "mes`),
	}, nil)
	td.Cmp(t, err, &APIError{Code: 400, Message: `{"code": 400, "mes`})

	// Error with QueryID
	responseHeaders := http.Header{}
	responseHeaders.Add("X-Ovh-QueryID", "FR.ws-8.5860f657.4632.0180")
	err = client.UnmarshalResponse(&http.Response{
		StatusCode: 400,
		Body:       sbody(`{"code": 400, "message": "Ooops..."}`),
		Header:     responseHeaders,
	}, &apiInt)
	td.Cmp(t, err, &APIError{
		Code:    400,
		Message: "Ooops...",
		QueryID: "FR.ws-8.5860f657.4632.0180",
	})
}

func TestGetResponseUnmarshalNumber(t *testing.T) {
	assert, require := td.AssertRequire(t)
	client := Client{}

	call := func(output interface{}) {
		t.Helper()
		err := client.UnmarshalResponse(&http.Response{
			StatusCode: 200,
			Body:       sbody(`{"orderId": 1234567890}`),
		}, output)
		require.CmpNoError(err)
	}

	// with map[string]interface{} as output
	var output map[string]interface{}
	call(&output)
	assert.Cmp(output, map[string]interface{}{"orderId": json.Number("1234567890")})

	// with map[string]int64 as output
	var outputInt map[string]int64
	call(&outputInt)
	assert.Cmp(outputInt, map[string]int64{"orderId": 1234567890})

	// with map[string]int64 as output
	var outputFloat map[string]float64
	call(&outputFloat)
	assert.Cmp(outputFloat, map[string]float64{"orderId": 1234567890})
}

func TestConstructors(t *testing.T) {
	assert, require := td.AssertRequire(t)

	// Error: missing Endpoint
	client, err := NewClient("", MockApplicationKey, MockApplicationSecret, MockConsumerKey)
	assert.Nil(client)
	assert.String(err, `unknown endpoint '', consider checking 'Endpoints' list of using an URL`)

	// Error: missing ApplicationKey
	client, err = NewClient("ovh-eu", "", MockApplicationSecret, MockConsumerKey)
	assert.Nil(client)
	assert.String(err, `missing application key, please check your configuration or consult the documentation to create one`)

	// Error: missing ApplicationSecret
	client, err = NewClient("ovh-eu", MockConsumerKey, "", MockConsumerKey)
	assert.Nil(client)
	assert.String(err, `missing application secret, please check your configuration or consult the documentation to create one`)

	// Next: success cases
	expected := td.Struct(&Client{
		AppKey:      MockApplicationKey,
		AppSecret:   MockApplicationSecret,
		ConsumerKey: MockConsumerKey,
		endpoint:    "https://eu.api.ovh.com/1.0",
	})

	// Nominal: full constructor
	client, err = NewClient("ovh-eu", MockApplicationKey, MockApplicationSecret, MockConsumerKey)
	require.CmpNoError(err)
	assert.Cmp(client, expected)

	// Nominal: Endpoint constructor
	t.Setenv("OVH_APPLICATION_KEY", MockApplicationKey)
	t.Setenv("OVH_APPLICATION_SECRET", MockApplicationSecret)
	t.Setenv("OVH_CONSUMER_KEY", MockConsumerKey)

	client, err = NewEndpointClient("ovh-eu")
	require.CmpNoError(err)
	assert.Cmp(client, expected)

	// Nominal: Default constructor
	t.Setenv("OVH_ENDPOINT", "ovh-eu")

	client, err = NewDefaultClient()
	require.CmpNoError(err)
	assert.Cmp(client, expected)
}

func (ms *MockSuite) TestVersionInURL(assert, require *td.T) {
	// Signature checking mocks
	httpmock.RegisterResponder("GET", "https://eu.api.ovh.com/1.0/call", func(req *http.Request) (*http.Response, error) {
		assert.Cmp(req.Header["X-Ovh-Signature"], []string{"$1$7f2db49253edfc41891023fcd1a54cf61db05fbb"}, "Right X-Ovh-Signature for /1.0 auth call")
		return httpmock.NewStringResponse(200, "{}"), nil
	})
	httpmock.RegisterResponder("GET", "https://eu.api.ovh.com/v1/call", func(req *http.Request) (*http.Response, error) {
		assert.Cmp(req.Header["X-Ovh-Signature"], []string{"$1$e6e7906d385eb28adcbfbe6b66c1528a42d741ad"}, "Right X-Ovh-Signature for /v1 auth call")
		return httpmock.NewStringResponse(200, "{}"), nil
	})
	httpmock.RegisterResponder("GET", "https://eu.api.ovh.com/v2/call", func(req *http.Request) (*http.Response, error) {
		assert.Cmp(req.Header["X-Ovh-Signature"], []string{"$1$bb63b132a6f84ad5433d0c534d48d3f7c3804285"}, "Right X-Ovh-Signature for /v2 auth call")
		return httpmock.NewStringResponse(200, "{}"), nil
	})

	// Mock local and distant time
	previous := getLocalTime
	getLocalTime = func() time.Time {
		return time.Unix(MockTime, 0)
	}
	assert.Cleanup(func() { getLocalTime = previous })

	httpmock.RegisterResponder("GET", "https://eu.api.ovh.com/1.0/auth/time",
		httpmock.NewStringResponder(200, strconv.Itoa(MockTime)))

	assertCallCount := func(assert *td.T, ccNoVersion, ccV1, ccV2 int) {
		assert.Helper()
		assert.Cmp(httpmock.GetCallCountInfo(), map[string]int{
			"GET https://eu.api.ovh.com/1.0/auth/time": 1,
			"GET https://eu.api.ovh.com/1.0/call":      ccNoVersion,
			"GET https://eu.api.ovh.com/v1/call":       ccV1,
			"GET https://eu.api.ovh.com/v2/call":       ccV2,
		})
	}

	require.Cmp(ms.client.endpoint, "https://eu.api.ovh.com/1.0")

	require.CmpNoError(ms.client.Get("/call", nil))
	assertCallCount(assert, 1, 0, 0)

	require.CmpNoError(ms.client.Get("/v1/call", nil))
	assertCallCount(assert, 1, 1, 0)

	require.CmpNoError(ms.client.Get("/v2/call", nil))
	assertCallCount(assert, 1, 1, 1)
}
