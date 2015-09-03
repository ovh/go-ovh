package govh

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestSig(t *testing.T) {
	// Sig should be: "$1$" + SHA1_HEX(AS+"+"+CK+"+"+METHOD+"+"+QUERY+"+"+BODY+"+"+TSTAMP)
	as := "EXEgWIz07P0HYwtQDs7cNIqCiQaWSuHF"
	ck := "MtSwSrPpNjqfVSmJhLbPyr2i45lSwPU1"
	method := "GET"
	query := "https://eu.api.ovh.com/1.0/domains/"
	var body string
	ts := int64(1366560945)

	// Get sig
	got := sig(as, ck, method, query, body, ts)

	// Expected
	expected := "$1$d3705e8afb27a0d2970a322b96550abfc67bb798"

	if got != expected {
		t.Fatalf("invalid sig: got %q, expected %q", got, expected)
	}
}

func TestTime(t *testing.T) {
	var apiCallCount int
	expectedDelay := 10 * time.Second
	expectedTime := time.Now().Add(expectedDelay)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiCallCount++
		fmt.Fprintf(w, "%d", expectedTime.Unix())
	}))
	defer ts.Close()

	// Create a fake client
	endpoint := Endpoint(ts.URL)
	client := NewClient(endpoint, "ak", "as", "ck")

	// Check that the time is correct
	gotTime, err := client.Time()
	if err != nil {
		t.Fatalf("expected not error, gotTime %q", err)
	}

	if expectedTime.Unix() != gotTime.Unix() {
		t.Errorf("expeted %q, gotTime %q", expectedTime, gotTime)
	}

	// Get the delay between the current time and the API
	gotDelay, err := delay(client.endpoint, &client.once)
	if err != nil {
		t.Fatalf("expected not error, got %q", err)
	}

	// The delay should be negative, we should use math.Floor to avoid dealing
	// with nanoseconds
	gotDelaySec := int(math.Floor(gotDelay.Seconds()))
	expectedDelaySec := int(math.Floor(-1 * expectedDelay.Seconds()))
	if gotDelaySec != expectedDelaySec {
		t.Errorf("expeted %d, gotTime %d", expectedDelaySec, gotDelaySec)
	}

	// The delay functionn is used inside a once methhod, let's check if the
	// call on the API to get the delay is really called once and only once.
	// Let's call delay one more time and check that the time API as only been
	// called twice (for the previous tests)
	_, err = delay(client.endpoint, &client.once)
	if err != nil {
		t.Fatalf("expected not error, got %q", err)
	}

	if apiCallCount != 2 {
		t.Errorf("expected 2 calls on time api, got %d", apiCallCount)
	}
}

func TestGetResponse(t *testing.T) {
	ak := "fakeKey"
	as := "fakeAppSecret"
	ck := "fakeCk"
	now := time.Now()

	// overwritte the time functions
	delay = func(e Endpoint, o *sync.Once) (time.Duration, error) {
		// No delay
		return 0 * time.Second, nil
	}
	getTime = func(endpoint Endpoint) (*time.Time, error) {
		return &now, nil
	}

	// Received headers
	var receivedHeaders *http.Header
	var receivedBody []byte
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		receivedHeaders = &r.Header

		// Get the content of the request
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "failed to read body", http.StatusInternalServerError)
		}
		defer r.Body.Close()
		receivedBody = body

		// Return a valid JSON result
		fmt.Fprintln(w, `{"test_status":"passing"}`)
	}))
	defer ts.Close()

	// Create a new client
	endpoint := Endpoint(ts.URL)
	client := NewClient(endpoint, ak, as, ck)

	// Fake post params
	type postParam struct {
		What string `json:"what"`
	}
	postParams := &postParam{What: "Some data"}
	expectedBody := []byte(`{"what":"Some data"}`)

	// Expected response
	type responseType struct {
		TestStatus string `json:"test_status"`
	}
	resp := &responseType{}

	// Make a post
	if err := client.Post("/fakeURL", postParams, resp); err != nil {
		t.Fatalf("expected no error, got %q", err)
	}

	// Check response body
	if !bytes.Equal(expectedBody, receivedBody) {
		t.Errorf("expected body %q, got %q", string(expectedBody), string(receivedBody))
	}

	// Check the response
	expectedResponse := &responseType{
		TestStatus: "passing",
	}
	if !reflect.DeepEqual(expectedResponse, resp) {
		t.Fatalf("invalid response, got %+v, expected %+v", expectedResponse, resp)
	}

	// Check headers
	for k, v := range map[string]string{
		"Content-Type":      "application/json",
		"X-Ovh-Timestamp":   strconv.FormatInt(time.Now().Unix(), 10),
		"X-Ovh-Application": ak,
		"X-Ovh-Consumer":    ck,
	} {
		h := receivedHeaders.Get(k)
		if v != h {
			t.Errorf("expected header %q received %q", v, h)
		}
	}

	// The values of these headers is tested in other functions, we will only
	// test their presence here
	for _, k := range []string{"X-Ovh-Signature", "X-Ovh-Timestamp"} {
		h := receivedHeaders.Get(k)
		if h == "" {
			t.Errorf("expected header %q received nothing", h)
		}
	}
}
