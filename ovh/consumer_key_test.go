package ovh

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/maxatome/go-testdeep/td"
)

func (ms *MockSuite) TestNewCkRequest(assert, require *td.T) {
	httpmock.RegisterResponder("POST", "https://eu.api.ovh.com/1.0/auth/credential", func(req *http.Request) (*http.Response, error) {
		assert.Cmp(req.Header["Accept"], []string{"application/json"})
		assert.Cmp(req.Header["X-Ovh-Application"], []string{MockApplicationKey})
		assert.Cmp(req.Body, td.Smuggle(json.RawMessage{}, td.JSON(`{"accessRules":[{"method":"GET","path":"/me"},{"method":"GET","path":"/xdsl/*"}]}`)))
		return httpmock.NewStringResponse(200, `{
            "validationUrl":"https://validation.url",
            "ConsumerKey":"`+MockConsumerKey+`",
            "state":"pendingValidation"
        }`), nil
	})

	ms.client.ConsumerKey = ""
	ckRequest := ms.client.NewCkRequest()
	ckRequest.AddRule("GET", "/me")
	ckRequest.AddRule("GET", "/xdsl/*")

	got, err := ckRequest.Do()
	require.CmpNoError(err)
	assert.Cmp(got, &CkValidationState{
		ConsumerKey:   MockConsumerKey,
		ValidationURL: "https://validation.url",
		State:         "pendingValidation",
	})
	assert.Cmp(ms.client.ConsumerKey, MockConsumerKey, "CkRequest.Do() sets client.ConsumerKey")
}

func (ms *MockSuite) TestInvalidCkRequest(assert, require *td.T) {
	httpmock.RegisterResponder("POST", "https://eu.api.ovh.com/1.0/auth/credential",
		httpmock.NewStringResponder(http.StatusForbidden, `{"message":"Invalid application key"}`))

	ckRequest := ms.client.NewCkRequest()
	ckRequest.AddRule("GET", "/me")
	ckRequest.AddRule("GET", "/xdsl/*")

	_, err := ckRequest.Do() // Returns 0 value, not nil
	assert.Cmp(err, &APIError{
		Code:    http.StatusForbidden,
		Message: "Invalid application key",
	})
}

func TestAddRules(t *testing.T) {
	client := Client{}

	// Test: allow all
	ckRequest := client.NewCkRequest()
	ckRequest.AddRecursiveRules(ReadWrite, "/")
	td.Cmp(t, ckRequest.AccessRules, []AccessRule{
		{Method: "GET", Path: "/*"},
		{Method: "POST", Path: "/*"},
		{Method: "PUT", Path: "/*"},
		{Method: "DELETE", Path: "/*"},
	})

	// Test: allow exactly /sms, RO
	ckRequest = client.NewCkRequest()
	ckRequest.AddRules(ReadOnly, "/sms")
	td.Cmp(t, ckRequest.AccessRules, []AccessRule{
		{Method: "GET", Path: "/sms"},
	})

	// Test: allow /sms/*, RW, no delete
	ckRequest = client.NewCkRequest()
	ckRequest.AddRecursiveRules(ReadWriteSafe, "/sms")
	td.Cmp(t, ckRequest.AccessRules, []AccessRule{
		{Method: "GET", Path: "/sms"},
		{Method: "POST", Path: "/sms"},
		{Method: "PUT", Path: "/sms"},

		{Method: "GET", Path: "/sms/*"},
		{Method: "POST", Path: "/sms/*"},
		{Method: "PUT", Path: "/sms/*"},
	})
}

func TestCkRequestString(t *testing.T) {
	td.CmpString(t, &CkValidationState{
		ConsumerKey:   "ck",
		State:         "pending",
		ValidationURL: "fakeURL",
	}, "CK: \"ck\"\nStatus: \"pending\"\nValidation URL: \"fakeURL\"\n")
}

func TestCkRequestRedirection(t *testing.T) {
	client := Client{}
	redirection := "http://localhost/api/auth/callback?token=123456"
	ckRequest := client.NewCkRequestWithRedirection(redirection)
	td.Cmp(t, ckRequest.Redirection, redirection)
}
