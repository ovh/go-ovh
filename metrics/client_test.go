package metrics

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ovh/go-ovh/ovh"
)

var (
	ovhClient   *ovh.Client
	clientIsSet = false
)

func init() {
	fmt.Println("Setup ovh client for tests")

	var err error
	ovhClient, err = ovh.NewEndpointClient("ovh-eu")
	if err != nil {
		fmt.Println("Cannot init ovh Client, skip it")
	} else {
		clientIsSet = true
	}
}

func TestClient_Services(t *testing.T) {
	type fields struct {
		ovhClient *ovh.Client
	}
	tests := []struct {
		name    string
		fields  fields
		want    ClientServices
		wantErr bool
	}{{
		name: "fetch services",
		fields: fields{
			ovhClient: ovhClient,
		},
		want: ClientServices{&ClientService{Service: &Service{
			Name: "cjlgifbfdjbhj",
		}}},
		wantErr: false,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if ovhClient == nil {
				return
			}

			c := &Client{
				ovhClient: tt.fields.ovhClient,
			}
			got, err := c.Services()
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Services() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Services() = %v, want %v", got, tt.want)
			}
		})
	}
}
