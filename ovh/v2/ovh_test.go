package ovh

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"
	"unicode"
	"unicode/utf8"
)

const (
	// In case you wonder, these are real *revoked* credentials
	MockApplicationKey    = "TDPKJdwZwAQPwKX2"
	MockApplicationSecret = "9ufkBmLaTQ9nz5yMUlg79taH0GNnzDjk"
	MockConsumerKey       = "5mBuy6SUQcRw2ZUxg0cG68BoDKpED4KY"

	MockTime = 1457018875
)

type SomeData struct {
	IntValue    int    `json:"i_val,omitempty"`
	StringValue string `json:"s_val,omitempty"`
}

//
// Utils
//

func initMockServer(InputRequest **http.Request, status int, responseBody string, handlerSleep time.Duration) (*httptest.Server, *Client) {
	// Mock time
	getLocalTime = func() time.Time {
		return time.Unix(MockTime, 0)
	}

	// Mock hostname, in signature only
	getEndpointForSignature = func(c *Client) string {
		return "http://localhost"
	}

	// Create a fake API server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Save input parameters
		body, _ := ioutil.ReadAll(r.Body)
		r.Body = ioutil.NopCloser(bytes.NewReader(body))
		*InputRequest = r
		if handlerSleep != 0 {
			time.Sleep(handlerSleep)
		}

		// Respond
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		fmt.Fprint(w, responseBody)
	}))

	// Create client
	client, _ := NewClient(ts.URL, MockApplicationKey, MockApplicationSecret, MockConsumerKey)
	client.timeDeltaDone = true

	return ts, client
}

func ensureHeaderPresent(t *testing.T, h http.Header, name string, value []string) {
	val, present := h[name]

	if !present {
		t.Fatalf("request should include a %s header with %s value.", name, value)
	}
	if !reflect.DeepEqual(val, value) {
		t.Fatalf("request should include a %s header with %v value. Got %v", name, value, val[0])
	}
}

func Capitalize(s string) string {
	if s == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + strings.ToLower(s[n:])
}

//
// Tests
//

func TestTime(t *testing.T) {
	// Init test
	var InputRequest *http.Request
	ts, client := initMockServer(&InputRequest, 200, fmt.Sprintf("%d", MockTime), time.Duration(0))
	defer ts.Close()

	// Test
	serverTime, err := client.Time()
	if err != nil {
		t.Fatalf("Unexpected error while retrieving server time: %v\n", err)
	}

	// Validate
	if InputRequest.Method != "GET" || InputRequest.URL.String() != "/auth/time" {
		t.Fatalf("Time should be retrieved using GET /auth/time. Got %s %s", InputRequest.Method, InputRequest.URL.String())
	}
	if serverTime.Unix() != MockTime {
		t.Fatalf("Server time should be %d. Got %d", MockTime, serverTime.Unix())
	}
}

func TestPing(t *testing.T) {
	// Init test
	var InputRequest *http.Request
	ts, client := initMockServer(&InputRequest, 200, `0`, time.Duration(0))
	defer ts.Close()

	// Test
	err := client.Ping()

	// Validate
	if err != nil {
		t.Fatalf("Unexpected error while pinging server: %v\n", err)
	}
}

func TestError500HTML(t *testing.T) {
	// Init test
	var InputRequest *http.Request
	errHTML := `<html><body><p>test</p></body></html>`
	ts, client := initMockServer(&InputRequest, http.StatusServiceUnavailable, errHTML, time.Duration(0))
	defer ts.Close()

	// Test
	var res struct{}
	err := client.CallAPI("GET", "/test", nil, nil, &res, false)

	// Validate
	if err == nil {
		t.Fatal("Expected error")
	}
	apiError := &APIError{
		Code:    http.StatusServiceUnavailable,
		Message: errHTML,
	}
	if err.Error() != apiError.Error() {
		t.Fatalf("Missmatch errors : \n%s\n%s", err, apiError)
	}
}

func TestPingUnreachable(t *testing.T) {
	// Init test
	var InputRequest *http.Request
	ts, client := initMockServer(&InputRequest, 200, `0`, time.Duration(0))
	defer ts.Close()

	// Test
	client.endpoint = "https://localhost:1/does not exist"
	err := client.Ping()

	// Validate
	if err == nil {
		t.Fatalf("Unexpected success while pinging server\n")
	}
}

func initServerAndClient(inputRequest **http.Request, requestTime time.Duration) (ts *httptest.Server, client *Client) {
	ts, client = initMockServer(inputRequest, 200, `"success"`, requestTime)
	return ts, client
}

func getExpectedRequest(method string, queryParams url.Values, body interface{}, signature string) *http.Request {
	var bodyBytes []byte
	if body != nil {
		bodyBytes, _ = json.Marshal(body)
	}
	path, _ := url.Parse("/some/resource")
	if queryParams != nil {
		path.RawQuery = queryParams.Encode()
	}
	expected := httptest.NewRequest(method, path.String(), bytes.NewReader(bodyBytes))
	expected.Header.Add("Accept", "application/json")
	expected.Header.Add("X-Ovh-Application", MockApplicationKey)
	if signature != "" {
		expected.Header.Add("X-Ovh-Consumer", MockConsumerKey)
		expected.Header.Add("X-Ovh-Timestamp", strconv.Itoa(MockTime))
		expected.Header.Add("X-Ovh-Signature", signature)
	}
	if body != nil {
		expected.Header.Add("Content-Type", "application/json;charset=utf-8")
	}
	return expected
}

func TestMergeQueryPath(t *testing.T) {
	var inputRequest *http.Request
	server, client := initServerAndClient(&inputRequest, time.Duration(0))
	expected := getExpectedRequest(http.MethodGet, url.Values{"foo": []string{"bar"}, "hello": []string{"world"}}, nil, "")
	defer server.Close()
	var res interface{}
	err := client.GetUnAuth("/some/resource?hello=world", url.Values{"foo": []string{"bar"}}, &res)
	if err != nil {
		t.Fatalf("Expected nil error")
	}

	requestTester(t, inputRequest, expected)
}

func TestGetUnAuth(t *testing.T) {
	var inputRequest *http.Request
	server, client := initServerAndClient(&inputRequest, time.Duration(0))
	expected := getExpectedRequest(http.MethodGet, nil, nil, "")
	defer server.Close()
	var res interface{}
	err := client.GetUnAuth("/some/resource", nil, &res)
	if err != nil {
		t.Fatalf("Expected nil error")
	}

	requestTester(t, inputRequest, expected)
}

func TestGetUnAuthWithContext(t *testing.T) {
	var inputRequest *http.Request
	server, client := initServerAndClient(&inputRequest, time.Duration(0))
	expected := getExpectedRequest(http.MethodGet, nil, nil, "")
	defer server.Close()

	var res interface{}
	err := client.GetUnAuthWithContext(context.Background(), "/some/resource", nil, &res)
	if err != nil {
		t.Fatalf("Expected nil error. Got %v", err)
	}
	requestTester(t, inputRequest, expected)
}

func TestGetUnAuthWithTimeoutContext(t *testing.T) {
	var inputRequest *http.Request
	server, client := initServerAndClient(&inputRequest, time.Duration(2)*time.Second)
	defer server.Close()

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(200)*time.Millisecond)
	defer cancelFunc()

	var res interface{}
	err := client.GetUnAuthWithContext(ctx, "/some/resource", nil, &res)
	if err == nil {
		t.Fatalf("Expected not nil error")
	}
	if !strings.Contains(err.Error(), "context deadline exceeded") {
		t.Fatalf("error should be context deadline exceeded ")
	}
}

func TestGet(t *testing.T) {
	var inputRequest *http.Request
	server, client := initServerAndClient(&inputRequest, time.Duration(0))
	expected := getExpectedRequest(http.MethodGet, nil, nil, "$1$8a21169b341aa23e82192e07457ca978006b1ba9")
	defer server.Close()
	var res interface{}
	err := client.Get("/some/resource", nil, &res)
	if err != nil {
		t.Fatalf("Expected nil error")
	}

	requestTester(t, inputRequest, expected)
}

func TestGetWithContext(t *testing.T) {
	var inputRequest *http.Request
	server, client := initServerAndClient(&inputRequest, time.Duration(0))
	expected := getExpectedRequest(http.MethodGet, nil, nil, "$1$8a21169b341aa23e82192e07457ca978006b1ba9")
	defer server.Close()

	var res interface{}
	err := client.GetWithContext(context.Background(), "/some/resource", nil, &res)
	if err != nil {
		t.Fatalf("Expected nil error. Got %v", err)
	}
	requestTester(t, inputRequest, expected)
}

func TestGetWithTimeoutContext(t *testing.T) {
	var inputRequest *http.Request
	server, client := initServerAndClient(&inputRequest, time.Duration(2)*time.Second)
	defer server.Close()

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(200)*time.Millisecond)
	defer cancelFunc()

	var res interface{}
	err := client.GetWithContext(ctx, "/some/resource", nil, &res)
	if err == nil {
		t.Fatalf("Expected not nil error")
	}
	if !strings.Contains(err.Error(), "context deadline exceeded") {
		t.Fatalf("error should be context deadline exceeded ")
	}
}

func TestPostUnAuth(t *testing.T) {
	var inputRequest *http.Request

	body := SomeData{
		IntValue:    42,
		StringValue: "Hello World!",
	}

	server, client := initServerAndClient(&inputRequest, time.Duration(0))
	expected := getExpectedRequest(http.MethodPost, nil, body, "")
	defer server.Close()
	var res interface{}
	err := client.PostUnAuth("/some/resource", nil, body, &res)
	if err != nil {
		t.Fatalf("Expected nil error")
	}

	requestTester(t, inputRequest, expected)
}

func TestPostUnAuthWithContext(t *testing.T) {
	var inputRequest *http.Request
	server, client := initServerAndClient(&inputRequest, time.Duration(0))

	body := SomeData{
		IntValue:    42,
		StringValue: "Hello World!",
	}

	expected := getExpectedRequest(http.MethodPost, nil, body, "")
	defer server.Close()

	var res interface{}
	err := client.PostUnAuthWithContext(context.Background(), "/some/resource", nil, body, &res)
	if err != nil {
		t.Fatalf("Expected nil error. Got %v", err)
	}
	requestTester(t, inputRequest, expected)
}

func TestPostUnAuthWithTimeoutContext(t *testing.T) {
	var inputRequest *http.Request
	server, client := initServerAndClient(&inputRequest, time.Duration(2)*time.Second)
	defer server.Close()
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(200)*time.Millisecond)
	defer cancelFunc()
	body := SomeData{
		IntValue:    42,
		StringValue: "Hello World!",
	}
	var res interface{}
	err := client.PostUnAuthWithContext(ctx, "/some/resource", nil, body, &res)
	if err == nil {
		t.Fatalf("Expected not nil error")
	}
	if !strings.Contains(err.Error(), "context deadline exceeded") {
		t.Fatalf("error should be context deadline exceeded ")
	}
}

func TestPost(t *testing.T) {
	var inputRequest *http.Request

	body := SomeData{
		IntValue:    42,
		StringValue: "Hello World!",
	}

	server, client := initServerAndClient(&inputRequest, time.Duration(0))
	expected := getExpectedRequest(http.MethodPost, nil, body, "$1$6549d84e65be72f4ec0d7b6d7eaa19554a265990")
	defer server.Close()
	var res interface{}
	err := client.Post("/some/resource", nil, body, &res)
	if err != nil {
		t.Fatalf("Expected nil error")
	}

	requestTester(t, inputRequest, expected)
}

func TestPostWithContext(t *testing.T) {
	var inputRequest *http.Request
	server, client := initServerAndClient(&inputRequest, time.Duration(0))

	body := SomeData{
		IntValue:    42,
		StringValue: "Hello World!",
	}

	expected := getExpectedRequest(http.MethodPost, nil, body, "$1$6549d84e65be72f4ec0d7b6d7eaa19554a265990")
	defer server.Close()

	var res interface{}
	err := client.PostWithContext(context.Background(), "/some/resource", nil, body, &res)
	if err != nil {
		t.Fatalf("Expected nil error. Got %v", err)
	}
	requestTester(t, inputRequest, expected)
}

func TestPostWithTimeoutContext(t *testing.T) {
	var inputRequest *http.Request
	server, client := initServerAndClient(&inputRequest, time.Duration(2)*time.Second)
	defer server.Close()

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(200)*time.Millisecond)
	defer cancelFunc()

	body := SomeData{
		IntValue:    42,
		StringValue: "Hello World!",
	}
	var res interface{}
	err := client.PostWithContext(ctx, "/some/resource", nil, body, &res)
	if err == nil {
		t.Fatalf("Expected not nil error")
	}
	if !strings.Contains(err.Error(), "context deadline exceeded") {
		t.Fatalf("error should be context deadline exceeded ")
	}
}

func TestPutUnAuth(t *testing.T) {
	var inputRequest *http.Request

	body := SomeData{
		IntValue:    42,
		StringValue: "Hello World!",
	}

	server, client := initServerAndClient(&inputRequest, time.Duration(0))
	expected := getExpectedRequest(http.MethodPut, nil, body, "")
	defer server.Close()
	var res interface{}
	err := client.PutUnAuth("/some/resource", nil, body, &res)
	if err != nil {
		t.Fatalf("Expected nil error")
	}

	requestTester(t, inputRequest, expected)
}

func TestPutUnAuthWithContext(t *testing.T) {
	var inputRequest *http.Request
	server, client := initServerAndClient(&inputRequest, time.Duration(0))

	body := SomeData{
		IntValue:    42,
		StringValue: "Hello World!",
	}

	expected := getExpectedRequest(http.MethodPut, nil, body, "")
	defer server.Close()

	var res interface{}
	err := client.PutUnAuthWithContext(context.Background(), "/some/resource", nil, body, &res)
	if err != nil {
		t.Fatalf("Expected nil error. Got %v", err)
	}
	requestTester(t, inputRequest, expected)
}

func TestPutUnAuthWithTimeoutContext(t *testing.T) {
	var inputRequest *http.Request
	server, client := initServerAndClient(&inputRequest, time.Duration(2)*time.Second)
	defer server.Close()

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(200)*time.Millisecond)
	defer cancelFunc()

	body := SomeData{
		IntValue:    42,
		StringValue: "Hello World!",
	}
	var res interface{}
	err := client.PutUnAuthWithContext(ctx, "/some/resource", nil, body, &res)
	if err == nil {
		t.Fatalf("Expected not nil error")
	}
	if !strings.Contains(err.Error(), "context deadline exceeded") {
		t.Fatalf("error should be context deadline exceeded ")
	}
}

func TestPut(t *testing.T) {
	var inputRequest *http.Request

	body := SomeData{
		IntValue:    42,
		StringValue: "Hello World!",
	}

	server, client := initServerAndClient(&inputRequest, time.Duration(0))
	expected := getExpectedRequest(http.MethodPut, nil, body, "$1$983e2a9a213c99211edd0b32715ac1ace1a6a0ea")
	defer server.Close()
	var res interface{}
	err := client.Put("/some/resource", nil, body, &res)
	if err != nil {
		t.Fatalf("Expected nil error")
	}

	requestTester(t, inputRequest, expected)
}

func TestPutWithContext(t *testing.T) {
	var inputRequest *http.Request
	server, client := initServerAndClient(&inputRequest, time.Duration(0))

	body := SomeData{
		IntValue:    42,
		StringValue: "Hello World!",
	}

	expected := getExpectedRequest(http.MethodPut, nil, body, "$1$983e2a9a213c99211edd0b32715ac1ace1a6a0ea")
	defer server.Close()

	var res interface{}
	err := client.PutWithContext(context.Background(), "/some/resource", nil, body, &res)
	if err != nil {
		t.Fatalf("Expected nil error. Got %v", err)
	}
	requestTester(t, inputRequest, expected)
}

func TestPutWithTimeoutContext(t *testing.T) {
	var inputRequest *http.Request
	server, client := initServerAndClient(&inputRequest, time.Duration(2)*time.Second)
	defer server.Close()

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(200)*time.Millisecond)
	defer cancelFunc()

	body := SomeData{
		IntValue:    42,
		StringValue: "Hello World!",
	}
	var res interface{}
	err := client.PutWithContext(ctx, "/some/resource", nil, body, &res)
	if err == nil {
		t.Fatalf("Expected not nil error")
	}
	if !strings.Contains(err.Error(), "context deadline exceeded") {
		t.Fatalf("error should be context deadline exceeded ")
	}
}

func TestDeleteUnAuth(t *testing.T) {
	var inputRequest *http.Request

	body := SomeData{
		IntValue:    42,
		StringValue: "Hello World!",
	}

	server, client := initServerAndClient(&inputRequest, time.Duration(0))
	expected := getExpectedRequest(http.MethodDelete, nil, body, "")
	defer server.Close()
	var res interface{}
	err := client.DeleteUnAuth("/some/resource", nil, body, &res)
	if err != nil {
		t.Fatalf("Expected nil error")
	}

	requestTester(t, inputRequest, expected)
}

func TestDeleteUnAuthWithContext(t *testing.T) {
	var inputRequest *http.Request
	server, client := initServerAndClient(&inputRequest, time.Duration(0))

	body := SomeData{
		IntValue:    42,
		StringValue: "Hello World!",
	}

	expected := getExpectedRequest(http.MethodDelete, nil, body, "")
	defer server.Close()

	var res interface{}
	err := client.DeleteUnAuthWithContext(context.Background(), "/some/resource", nil, body, &res)
	if err != nil {
		t.Fatalf("Expected nil error. Got %v", err)
	}
	requestTester(t, inputRequest, expected)
}

func TestDeleteUnAuthWithTimeoutContext(t *testing.T) {
	var inputRequest *http.Request
	server, client := initServerAndClient(&inputRequest, time.Duration(2)*time.Second)
	defer server.Close()

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(200)*time.Millisecond)
	defer cancelFunc()

	body := SomeData{
		IntValue:    42,
		StringValue: "Hello World!",
	}
	var res interface{}
	err := client.DeleteUnAuthWithContext(ctx, "/some/resource", nil, body, &res)
	if err == nil {
		t.Fatalf("Expected not nil error")
	}
	if !strings.Contains(err.Error(), "context deadline exceeded") {
		t.Fatalf("error should be context deadline exceeded ")
	}
}

func TestDelete(t *testing.T) {
	var inputRequest *http.Request

	body := SomeData{
		IntValue:    42,
		StringValue: "Hello World!",
	}

	server, client := initServerAndClient(&inputRequest, time.Duration(0))
	expected := getExpectedRequest(http.MethodDelete, nil, body, "$1$793199d35c3baf66a1588bcc8c3c8d0f7f4d4bd3")
	defer server.Close()
	var res interface{}
	err := client.Delete("/some/resource", nil, body, &res)
	if err != nil {
		t.Fatalf("Expected nil error")
	}

	requestTester(t, inputRequest, expected)
}

func TestDeleteWithContext(t *testing.T) {
	var inputRequest *http.Request
	server, client := initServerAndClient(&inputRequest, time.Duration(0))

	body := SomeData{
		IntValue:    42,
		StringValue: "Hello World!",
	}

	expected := getExpectedRequest(http.MethodDelete, nil, body, "$1$793199d35c3baf66a1588bcc8c3c8d0f7f4d4bd3")
	defer server.Close()

	var res interface{}
	err := client.DeleteWithContext(context.Background(), "/some/resource", nil, body, &res)
	if err != nil {
		t.Fatalf("Expected nil error. Got %v", err)
	}
	requestTester(t, inputRequest, expected)
}

func TestDeleteWithTimeoutContext(t *testing.T) {
	var inputRequest *http.Request
	server, client := initServerAndClient(&inputRequest, time.Duration(2)*time.Second)
	defer server.Close()

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(200)*time.Millisecond)
	defer cancelFunc()

	body := SomeData{
		IntValue:    42,
		StringValue: "Hello World!",
	}
	var res interface{}
	err := client.DeleteWithContext(ctx, "/some/resource", nil, body, &res)
	if err == nil {
		t.Fatalf("Expected not nil error")
	}
	if !strings.Contains(err.Error(), "context deadline exceeded") {
		t.Fatalf("error should be context deadline exceeded ")
	}
}

func requestTester(t *testing.T, request, expected *http.Request) {
	// Validate Method
	if request.Method != expected.Method {
		t.Fatalf("Method %s of received request isn't the same that expected method %s", request.Method, expected.Method)
	}

	if request.URL.String() != expected.URL.String() {
		t.Fatalf("Query %s of received request isn't the same that expected query %s", request.URL.String(), expected.URL.String())
	}
	isBodyEquals(t, request.Body, expected.Body)
	isHeaderContains(t, request.Header, expected.Header)

}
func isHeaderContains(t *testing.T, header, expected http.Header) {
	for k, v := range expected {
		ensureHeaderPresent(t, header, k, v)
	}
}

func isBodyEquals(t *testing.T, body, expected io.ReadCloser) {

	expectedBytes, err := ioutil.ReadAll(expected)
	if err != nil {
		t.Fatalf("Unable to read expected body. Got %v", err)
	}
	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		t.Fatalf("Unable to read body. Got %v", err)
	}
	res := bytes.Compare(bodyBytes, expectedBytes)
	if res != 0 {
		t.Fatalf("body received \"%s\" isn't equal to expected body \"%s\"", string(bodyBytes), string(expectedBytes))
	}
}

// Mock ReadCloser, always failing
type ErrorCloseReader struct{}

func (ErrorCloseReader) Read(p []byte) (int, error) {
	return 0, fmt.Errorf("ErrorReader")
}
func (ErrorCloseReader) Close() error {
	return nil
}

func TestGetResponse(t *testing.T) {
	var err error
	var apiInt int
	mockClient := Client{}

	// Nominal
	err = mockClient.UnmarshalResponse(&http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(strings.NewReader(`42`)),
	}, &apiInt)
	if err != nil {
		t.Fatalf("Client.UnmarshalResponse should be able to decode int when status is 200. Got %v", err)
	}

	// Nominal: empty body
	err = mockClient.UnmarshalResponse(&http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(strings.NewReader(``)),
	}, nil)
	if err != nil {
		t.Fatalf("UnmarshalResponse should not return an error when reponse is empty or target type is nil. Got %v", err)
	}

	// Error
	err = mockClient.UnmarshalResponse(&http.Response{
		StatusCode: 400,
		Body:       ioutil.NopCloser(strings.NewReader(`{"code": 400, "message": "Ooops..."}`)),
	}, &apiInt)
	if err == nil {
		t.Fatalf("Client.UnmarshalResponse should be able to decode an error when status is 400")
	}
	if _, ok := err.(*APIError); !ok {
		t.Fatalf("Client.UnmarshalResponse error should be an APIError when status is 400. Got '%s' of type %s", err, reflect.TypeOf(err))
	}

	// Error: body read error
	err = mockClient.UnmarshalResponse(&http.Response{
		Body: ErrorCloseReader{},
	}, nil)
	if err == nil {
		t.Fatalf("UnmarshalResponse should return an error when failing to read HTTP Response body. %v", err)
	}

	// Error: HTTP Error + broken json
	err = mockClient.UnmarshalResponse(&http.Response{
		StatusCode: 400,
		Body:       ioutil.NopCloser(strings.NewReader(`{"code": 400, "mes`)),
	}, nil)
	if err == nil {
		t.Fatalf("UnmarshalResponse should return an error when failing to decode HTTP Response body. %v", err)
	}

	// Error with QueryID
	responseHeaders := http.Header{}
	responseHeaders.Add("X-Ovh-QueryID", "FR.ws-8.5860f657.4632.0180")
	err = mockClient.UnmarshalResponse(&http.Response{
		StatusCode: 400,
		Body:       ioutil.NopCloser(strings.NewReader(`{"code": 400, "message": "Ooops..."}`)),
		Header:     responseHeaders,
	}, &apiInt)
	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("Client.UnmarshalResponse error should be an APIError when status is 400 and header QueryID is found. Got '%s' of type %s", err, reflect.TypeOf(err))
	}
	if apiErr.QueryID != "FR.ws-8.5860f657.4632.0180" {
		t.Fatalf("APIError should be filled with a correct QueryID. Got '%s' instead of '%s'", apiErr.QueryID, "FR.ws-8.5860f657.4632.0180")
	}
}

func TestGetResponseUnmarshalNumber(t *testing.T) {
	var err error
	var output map[string]interface{}
	mockClient := Client{}

	// with map[string]interface{} as output
	err = mockClient.UnmarshalResponse(&http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(strings.NewReader(`{"orderId": 1234567890}`)),
	}, &output)
	if err != nil {
		t.Fatalf("Client.UnmarshalResponse should be able to decode the body")
	}
	if "1234567890" != fmt.Sprint(output["orderId"]) {
		t.Fatalf("Client.UnmarshalResponse should unmarshal long integer as json.Number instead of float64, stringified incorrectly")
	}

	var outputInt map[string]int64

	// with map[string]int64 as output
	err = mockClient.UnmarshalResponse(&http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(strings.NewReader(`{"orderId": 1234567890}`)),
	}, &outputInt)
	if err != nil {
		t.Fatalf("Client.UnmarshalResponse should be able to decode the body")
	}
	if int64(1234567890) != outputInt["orderId"] {
		t.Fatalf("Client.UnmarshalResponse should unmarshal long integer as json.Number instead of float64, incorrectly casted as int64")
	}

	var outputFloat map[string]float64

	// with map[string]int64 as output
	err = mockClient.UnmarshalResponse(&http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(strings.NewReader(`{"orderId": 1234567890}`)),
	}, &outputFloat)
	if err != nil {
		t.Fatalf("Client.UnmarshalResponse should be able to decode the body")
	}
	if float64(1234567890) != outputFloat["orderId"] {
		t.Fatalf("Client.UnmarshalResponse should unmarshal long integer as json.Number instead of float64, incorrectly casted as float64")
	}
}

func TestConstructors(t *testing.T) {
	// Nominal: full constructor
	client, err := NewClient("ovh-eu", MockApplicationKey, MockApplicationSecret, MockConsumerKey)
	if err != nil {
		t.Fatalf("NewClient should not return an error in the nominal case. Got: %v", err)
	}
	if client.Client == nil {
		t.Fatalf("client.Client should be a valid HTTP client")
	}
	if client.AppKey != MockApplicationKey {
		t.Fatalf("client.AppKey should be '%s'. Got '%s'", MockApplicationKey, client.AppKey)
	}
	if client.AppSecret != MockApplicationSecret {
		t.Fatalf("client.AppSecret should be '%s'. Got '%s'", MockApplicationSecret, client.AppSecret)
	}
	if client.ConsumerKey != MockConsumerKey {
		t.Fatalf("client.ConsumerKey should be '%s'. Got '%s'", MockConsumerKey, client.ConsumerKey)
	}

	// Nominal: Endpoint constructor
	os.Setenv("OVH_APPLICATION_KEY", MockApplicationKey)
	os.Setenv("OVH_APPLICATION_SECRET", MockApplicationSecret)
	os.Setenv("OVH_CONSUMER_KEY", MockConsumerKey)

	client, err = NewEndpointClient("ovh-eu")
	if err != nil {
		t.Fatalf("NewEndpointClient should not return an error in the nominal case. Got: %v", err)
	}
	if client.Client == nil {
		t.Fatalf("client.Client should be a valid HTTP client")
	}
	if client.AppKey != MockApplicationKey {
		t.Fatalf("client.AppKey should be '%s'. Got '%s'", MockApplicationKey, client.AppKey)
	}
	if client.AppSecret != MockApplicationSecret {
		t.Fatalf("client.AppSecret should be '%s'. Got '%s'", MockApplicationSecret, client.AppSecret)
	}
	if client.ConsumerKey != MockConsumerKey {
		t.Fatalf("client.ConsumerKey should be '%s'. Got '%s'", MockConsumerKey, client.ConsumerKey)
	}

	// Nominal: Default constructor
	os.Setenv("OVH_ENDPOINT", "ovh-eu")

	client, err = NewDefaultClient()
	if err != nil {
		t.Fatalf("NewEndpointClient should not return an error in the nominal case. Got: %v", err)
	}
	if client.Client == nil {
		t.Fatalf("client.Client should be a valid HTTP client")
	}
	if client.endpoint != "https://eu.api.ovh.com/1.0" {
		t.Fatalf("client.Endpoint should be 'https://eu.api.ovh.com/1.0'. Got '%s'", client.endpoint)
	}

	// Clear
	os.Unsetenv("OVH_ENDPOINT")
	os.Unsetenv("OVH_APPLICATION_KEY")
	os.Unsetenv("OVH_APPLICATION_SECRET")
	os.Unsetenv("OVH_CONSUMER_KEY")

	// Error: missing Endpoint
	_, err = NewClient("", MockApplicationKey, MockApplicationSecret, MockConsumerKey)
	if err == nil {
		t.Fatalf("NewClient should return an error when missing Endpoint")
	}
	// Error: missing ApplicationKey
	_, err = NewClient("ovh-eu", "", MockApplicationSecret, MockConsumerKey)
	if err == nil {
		t.Fatalf("NewClient should return an error when missing ApplicationKey")
	}
	// Error: missing ApplicationSecret
	_, err = NewClient("ovh-eu", MockConsumerKey, "", MockConsumerKey)
	if err == nil {
		t.Fatalf("NewClient should return an error when missing ApplicationSecret")
	}
}

func TestGetTimeDelta(t *testing.T) {
	MockDelta := 747

	// Init test
	var InputRequest *http.Request
	ts, client := initMockServer(&InputRequest, 200, fmt.Sprintf("%d", int(time.Now().Unix())-MockDelta), time.Duration(0))
	defer ts.Close()

	// Test
	client.timeDeltaDone = false
	delta, err := client.getTimeDelta()

	if err != nil {
		t.Fatalf("getTimeDelta should not return an error. Got %v", err)
	}
	// Hack: take races into account, avoid mocking whole earth
	if math.Abs(float64(delta/time.Second-time.Duration(MockDelta))) > 2 {
		t.Fatalf("getTimeDelta should return a delta of %d. Got %d", time.Duration(MockDelta)*time.Second, delta)
	}
}
