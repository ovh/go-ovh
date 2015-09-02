package govh

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// AccessRule represents a method allowed for a path
type AccessRule struct {
	// Allowed HTTP Method for the requested AccessRule.
	// Can be set to GET/POST/PUT/DELETE.
	Method string `json:"method"`
	// Allowed path.
	// Can be an exact string or a string with '*' char.
	// Example :
	// 		/me : only /me is authorized
	//		/* : all calls are authorized
	Path string `json:"path"`
}

// CkValidationState represents the response when asking a new consumerKey.
type CkValidationState struct {
	// Consumer key, which need to be validated by customer.
	ConsumerKey string `json:"consumerKey"`
	// Current status, should be always "pendingValidation".
	State string `json:"state"`
	// URL to redirect user in order to log in.
	ValidationURL string `json:"validationUrl"`
}

// CkRequest represents the parameters to fill in order to ask a new
// consumerKey.
type CkRequest struct {
	endpoint    Endpoint
	appKey      string
	AccessRules []*AccessRule `json:"accessRules"`
}

func (ck *CkValidationState) String() string {
	return fmt.Sprintf("CK: %q\nStatus: %q\nValidation URL: %q\n",
		ck.ConsumerKey,
		ck.State,
		ck.ValidationURL,
	)
}

// NewCkRequest helps create a new ck request
func NewCkRequest(endpoint Endpoint, appKey string) *CkRequest {
	return &CkRequest{
		endpoint:    endpoint,
		appKey:      appKey,
		AccessRules: []*AccessRule{},
	}
}

// AddRule adds a new rule to the ckRequest
func (ck *CkRequest) AddRule(method, path string) {
	ck.AccessRules = append(ck.AccessRules, &AccessRule{
		Method: method,
		Path:   path,
	})
}

// Do runs executes the request
func (ck *CkRequest) Do() (*CkValidationState, error) {
	params, err := json.Marshal(ck)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/auth/credential", ck.endpoint)
	request, err := http.NewRequest("POST", url, bytes.NewReader(params))
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-OVH-Application", ck.appKey)

	result, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer result.Body.Close()

	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}

	if result.StatusCode != http.StatusOK {
		apiError := &APIOvhError{Code: result.StatusCode}
		if err = json.Unmarshal(body, apiError); err != nil {
			return nil, err
		}

		return nil, apiError
	}

	state := &CkValidationState{}
	if err := json.Unmarshal(body, state); err != nil {
		return nil, err
	}

	return state, nil
}
