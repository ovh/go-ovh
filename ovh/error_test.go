package ovh

import (
	"net/http"
	"testing"

	"github.com/maxatome/go-testdeep/td"
)

func TestErrorString(t *testing.T) {
	err := APIError{
		Code:    http.StatusBadRequest,
		Message: "An input error occurred",
	}
	expected := `OVHcloud API error (status code 400): "An input error occurred"`
	td.Cmp(t, err.Error(), expected)
	td.Cmp(t, err.String(), expected)

	err = APIError{
		Code:    http.StatusConflict,
		Message: `the cart id "foobar" already exists`,
		Class:   "CartAlreadyExists",
		QueryID: "EU.ext-99.foobar",
	}
	expected = `OVHcloud API error (status code 409): CartAlreadyExists: "the cart id \"foobar\" already exists" (X-OVH-Query-Id: EU.ext-99.foobar)`
	td.Cmp(t, err.Error(), expected)
	td.Cmp(t, err.String(), expected)

	err.Class = ""
	expected = `OVHcloud API error (status code 409): "the cart id \"foobar\" already exists" (X-OVH-Query-Id: EU.ext-99.foobar)`
	td.Cmp(t, err.Error(), expected)
	td.Cmp(t, err.String(), expected)

	err = APIError{
		Code:    http.StatusForbidden,
		Message: `User not granted for this request`,
		Class:   "Client::Forbidden",
		QueryID: "EU.ext-99.foobar",
		Details: map[string]string{
			"unauthorizedActionsByAuthentication": "",
			"unauthorizedActionsByIAM":            "account:apiovh:me/installationTemplate/get",
		},
	}

	expected = `OVHcloud API error (status code 403): Client::Forbidden: "User not granted for this request (missing IAM permissions: account:apiovh:me/installationTemplate/get)" (X-OVH-Query-Id: EU.ext-99.foobar)`
	td.Cmp(t, err.Error(), expected)
	td.Cmp(t, err.String(), expected)

	err = APIError{
		Code:    http.StatusForbidden,
		Message: `User not granted for this request`,
		Class:   "Client::Forbidden",
		QueryID: "EU.ext-99.foobar",
		Details: map[string]string{
			"unauthorizedActionsByAuthentication": "account:apiovh:me/accessRestriction/ip/get",
			"unauthorizedActionsByIAM":            "account:apiovh:me/installationTemplate/get",
		},
	}

	expected = `OVHcloud API error (status code 403): Client::Forbidden: "User not granted for this request (missing IAM permissions: account:apiovh:me/accessRestriction/ip/get, account:apiovh:me/installationTemplate/get)" (X-OVH-Query-Id: EU.ext-99.foobar)`
	td.Cmp(t, err.Error(), expected)
	td.Cmp(t, err.String(), expected)
}
