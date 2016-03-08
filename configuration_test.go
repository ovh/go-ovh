package govh

import (
	"io/ioutil"
	"os"
	"testing"
)

//
// Utils
//

var home string

func setup() {
	systemConfigPath = "./ovh.unittest.global.conf"
	userConfigPath = "/.ovh.unittest.user.conf"
	localConfigPath = "./ovh.unittest.local.conf"
	home, _ = currentUserHome()
}

func teardown() {
	os.Remove(systemConfigPath)
	os.Remove(home + userConfigPath)
	os.Remove(localConfigPath)
}

//
// Tests
//

func TestConfigFromFiles(t *testing.T) {
	// Write each parameter to one different configuration file
	// This is a simple way to test precedence

	// Prepare
	ioutil.WriteFile(systemConfigPath, []byte(`
[ovh-eu]
application_key=system
application_secret=system
consumer_key=system
`), 0660)

	ioutil.WriteFile(home+userConfigPath, []byte(`
[ovh-eu]
application_secret=user
consumer_key=user
`), 0660)

	ioutil.WriteFile(localConfigPath, []byte(`
[ovh-eu]
consumer_key=local
`), 0660)

	// Clear
	defer ioutil.WriteFile(systemConfigPath, []byte(``), 0660)
	defer ioutil.WriteFile(home+userConfigPath, []byte(``), 0660)
	defer ioutil.WriteFile(localConfigPath, []byte(``), 0660)

	// Test
	client := Client{}
	err := client.loadConfig("ovh-eu")

	// Validate
	if err != nil {
		t.Fatalf("loadConfig failed with: '%v'", err)
	}
	if client.appKey != "system" {
		t.Fatalf("client.appKey should be 'system'. Got '%s'", client.appKey)
	}
	if client.appSecret != "user" {
		t.Fatalf("client.appSecret should be 'user'. Got '%s'", client.appSecret)
	}
	if client.consumerKey != "local" {
		t.Fatalf("client.consumerKey should be 'local'. Got '%s'", client.consumerKey)
	}
}

func TestConfigFromOnlyOneFile(t *testing.T) {
	// ini package has a bug causing it to ignore all subsequent configuration
	// files if any could not be loaded. Make sure that workaround... works.

	// Prepare
	os.Remove(systemConfigPath)
	ioutil.WriteFile(home+userConfigPath, []byte(`
[ovh-eu]
application_key=user
application_secret=user
consumer_key=user
`), 0660)

	// Clear
	defer ioutil.WriteFile(home+userConfigPath, []byte(``), 0660)

	// Test
	client := Client{}
	err := client.loadConfig("ovh-eu")

	// Validate
	if err != nil {
		t.Fatalf("loadConfig failed with: '%v'", err)
	}
	if client.appKey != "user" {
		t.Fatalf("client.appKey should be 'user'. Got '%s'", client.appKey)
	}
	if client.appSecret != "user" {
		t.Fatalf("client.appSecret should be 'user'. Got '%s'", client.appSecret)
	}
	if client.consumerKey != "user" {
		t.Fatalf("client.consumerKey should be 'user'. Got '%s'", client.consumerKey)
	}
}

func TestConfigFromEnv(t *testing.T) {
	// Prepare
	ioutil.WriteFile(systemConfigPath, []byte(`
[ovh-eu]
application_key=fail
application_secret=fail
consumer_key=fail
`), 0660)

	defer ioutil.WriteFile(systemConfigPath, []byte(``), 0660)
	os.Setenv("OVH_ENDPOINT", "ovh-eu")
	os.Setenv("OVH_APPLICATION_KEY", "env")
	os.Setenv("OVH_APPLICATION_SECRET", "env")
	os.Setenv("OVH_CONSUMER_KEY", "env")

	// Clear
	defer os.Unsetenv("OVH_ENDPOINT")
	defer os.Unsetenv("OVH_APPLICATION_KEY")
	defer os.Unsetenv("OVH_APPLICATION_SECRET")
	defer os.Unsetenv("OVH_CONSUMER_KEY")

	// Test
	client := Client{}
	err := client.loadConfig("")

	// Validate
	if err != nil {
		t.Fatalf("loadConfig failed with: '%v'", err)
	}
	if client.endpoint != OvhEU {
		t.Fatalf("client.appKey should be 'env'. Got '%s'", client.appKey)
	}
	if client.appKey != "env" {
		t.Fatalf("client.appKey should be 'env'. Got '%s'", client.appKey)
	}
	if client.appSecret != "env" {
		t.Fatalf("client.appSecret should be 'env'. Got '%s'", client.appSecret)
	}
	if client.consumerKey != "env" {
		t.Fatalf("client.consumerKey should be 'env'. Got '%s'", client.consumerKey)
	}
}

func TestConfigFromArgs(t *testing.T) {
	// Test
	client := Client{
		appKey:      "param",
		appSecret:   "param",
		consumerKey: "param",
	}
	err := client.loadConfig("ovh-eu")

	// Validate
	if err != nil {
		t.Fatalf("loadConfig failed with: '%v'", err)
	}
	if client.endpoint != OvhEU {
		t.Fatalf("client.appKey should be 'param'. Got '%s'", client.appKey)
	}
	if client.appKey != "param" {
		t.Fatalf("client.appKey should be 'param'. Got '%s'", client.appKey)
	}
	if client.appSecret != "param" {
		t.Fatalf("client.appSecret should be 'param'. Got '%s'", client.appSecret)
	}
	if client.consumerKey != "param" {
		t.Fatalf("client.consumerKey should be 'param'. Got '%s'", client.consumerKey)
	}
}

func TestEndpoint(t *testing.T) {
	// Prepare
	ioutil.WriteFile(systemConfigPath, []byte(`
[ovh-eu]
application_key=ovh
application_secret=ovh
consumer_key=ovh

[https://api.example.com:4242]
application_key=example.com
application_secret=example.com
consumer_key=example.com
`), 0660)

	// Clear
	defer ioutil.WriteFile(systemConfigPath, []byte(``), 0660)

	// Test: by name
	client := Client{}
	err := client.loadConfig("ovh-eu")
	if err != nil {
		t.Fatalf("loadConfig should not fail for endpoint 'ovh-eu'. Got '%v'", err)
	}
	if client.appKey != "ovh" {
		t.Fatalf("configured value should be 'ovh' for endpoint 'ovh-eu'. Got '%s'", client.appKey)
	}

	// Test: by URL
	client = Client{}
	err = client.loadConfig("https://api.example.com:4242")
	if err != nil {
		t.Fatalf("loadConfig should not fail for endpoint 'https://api.example.com:4242'. Got '%v'", err)
	}
	if client.appKey != "example.com" {
		t.Fatalf("configured value should be 'example.com' for endpoint 'https://api.example.com:4242'. Got '%s'", client.appKey)
	}

}

func TestMissingParam(t *testing.T) {
	// Setup
	var err error
	client := Client{
		appKey:      "param",
		appSecret:   "param",
		consumerKey: "param",
	}

	// Test
	client.endpoint = Endpoint("")
	if err = client.loadConfig(""); err == nil {
		t.Fatalf("loadConfig should fail when client.endpoint is missing. Got '%s'", client.endpoint)
	}

	client.appKey = ""
	if err = client.loadConfig("ovh-eu"); err == nil {
		t.Fatalf("loadConfig should fail when client.appKey is missing. Got '%s'", client.appKey)
	}
	client.appKey = "param"

	client.appSecret = ""
	if err = client.loadConfig("ovh-eu"); err == nil {
		t.Fatalf("loadConfig should fail when client.appSecret is missing. Got '%s'", client.appSecret)
	}
	client.appSecret = "param"
}

//
// Main
//

// TestMain changes the location of configuration files. We need
// this to avoid any interference with existing configuration
// and non-root users
func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
