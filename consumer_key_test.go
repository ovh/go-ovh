package govh

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

var fakeAppKey = "7kbG7Bk7S9Nt7ZSV"

func TestNewCkReqest(t *testing.T) {
	// Expected headers
	expectedHeaders := map[string]string{
		"Content-Type":      "application/json",
		"X-Ovh-Application": fakeAppKey,
	}
	// Received headers
	receivedHeaders := map[string]string{}

	// Received request
	var receivedBody []byte

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the headers of the request
		for h := range expectedHeaders {
			receivedHeaders[h] = r.Header.Get(h)
		}

		// Get the content of the request
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "failed to read body", http.StatusInternalServerError)
		}
		defer r.Body.Close()
		receivedBody = body

		// Return a valid JSON result
		fmt.Fprintln(w, `{
			"validationUrl":"https://validation.url",
			"consumerKey":"MtSwSrPpNjqfVSmJhLbPyr2i45lSwPU1",
			"state":"pendingValidation"
		}`)
	}))
	defer ts.Close()

	// Create a new ck request
	endpoint := Endpoint(ts.URL)
	ckRequest := NewCkRequest(endpoint, fakeAppKey)
	ckRequest.AddRule("GET", "/me")
	ckRequest.AddRule("GET", "/xdsl/*")

	// Run the request
	got, err := ckRequest.Do()
	if err != nil {
		t.Fatalf("expected no error, got %q", err)
	}

	// Check the headers
	if !reflect.DeepEqual(expectedHeaders, receivedHeaders) {
		t.Fatalf("invalid headers\nwanted: %+v\ngot: %+v\n", expectedHeaders, receivedHeaders)
	}

	// Check the request
	receivedRequest := &CkRequest{}
	if err = json.Unmarshal(receivedBody, receivedRequest); err != nil {
		t.Fatalf("error during request unmarshal: %q", err)
	}
	expectedRequest := &CkRequest{
		AccessRules: []*AccessRule{
			{Method: "GET", Path: "/me"},
			{Method: "GET", Path: "/xdsl/*"},
		},
	}
	if !reflect.DeepEqual(receivedRequest, expectedRequest) {
		t.Fatalf("invalid ck request")
	}

	// Check the response
	expected := &CkValidationState{
		ValidationURL: "https://validation.url",
		State:         "pendingValidation",
		ConsumerKey:   "MtSwSrPpNjqfVSmJhLbPyr2i45lSwPU1",
	}

	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("unexpected response from new consumerKey requet")
	}
}

func TestInvalidCkReqest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, `{"message":"Invalid application key"}`, http.StatusUnauthorized)
	}))
	defer ts.Close()

	// Create a new ck request
	endpoint := Endpoint(ts.URL)
	ckRequest := NewCkRequest(endpoint, fakeAppKey)
	ckRequest.AddRule("GET", "/me")

	// Run the request
	_, err := ckRequest.Do()
	if err == nil {
		t.Fatal("expected an error, got none")
	}

	expected := &APIOvhError{
		Code:    http.StatusUnauthorized,
		Message: "Invalid application key",
	}

	if !reflect.DeepEqual(err, expected) {
		t.Fatalf("unexpected error from get ck request")
	}
}
