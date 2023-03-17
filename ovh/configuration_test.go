package ovh

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/maxatome/go-testdeep/td"
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
	assert, require := td.AssertRequire(t)

	// Write each parameter to one different configuration file
	// This is a simple way to test precedence

	// Prepare
	err := ioutil.WriteFile(systemConfigPath, []byte(`
[ovh-eu]
application_key=system
application_secret=system
consumer_key=system
`), 0660)
	require.CmpNoError(err)

	err = ioutil.WriteFile(home+userConfigPath, []byte(`
[ovh-eu]
application_secret=user
consumer_key=user
`), 0660)
	require.CmpNoError(err)

	err = ioutil.WriteFile(localConfigPath, []byte(`
[ovh-eu]
consumer_key=local
`), 0660)
	require.CmpNoError(err)

	// Clear
	t.Cleanup(func() {
		_ = ioutil.WriteFile(systemConfigPath, []byte(``), 0660)
		_ = ioutil.WriteFile(home+userConfigPath, []byte(``), 0660)
		_ = ioutil.WriteFile(localConfigPath, []byte(``), 0660)
	})

	client := Client{}
	err = client.loadConfig("ovh-eu")
	require.CmpNoError(err)
	assert.Cmp(client, td.Struct(Client{
		AppKey:      "system",
		AppSecret:   "user",
		ConsumerKey: "local",
	}))
}

func TestConfigFromOnlyOneFile(t *testing.T) {
	assert, require := td.AssertRequire(t)

	// ini package has a bug causing it to ignore all subsequent configuration
	// files if any could not be loaded. Make sure that workaround... works.

	// Prepare
	err := os.Remove(systemConfigPath)
	require.CmpNoError(err)

	err = ioutil.WriteFile(home+userConfigPath, []byte(`
[ovh-eu]
application_key=user
application_secret=user
consumer_key=user
`), 0660)
	require.CmpNoError(err)

	// Clear
	t.Cleanup(func() {
		_ = ioutil.WriteFile(home+userConfigPath, []byte(``), 0660)
	})

	client := Client{}
	err = client.loadConfig("ovh-eu")
	require.CmpNoError(err)
	assert.Cmp(client, td.Struct(Client{
		AppKey:      "user",
		AppSecret:   "user",
		ConsumerKey: "user",
	}))
}

func TestConfigFromEnv(t *testing.T) {
	assert, require := td.AssertRequire(t)

	err := ioutil.WriteFile(systemConfigPath, []byte(`
[ovh-eu]
application_key=fail
application_secret=fail
consumer_key=fail
`), 0660)
	require.CmpNoError(err)

	t.Cleanup(func() {
		_ = ioutil.WriteFile(systemConfigPath, []byte(``), 0660)
	})

	t.Setenv("OVH_ENDPOINT", "ovh-eu")
	t.Setenv("OVH_APPLICATION_KEY", "env")
	t.Setenv("OVH_APPLICATION_SECRET", "env")
	t.Setenv("OVH_CONSUMER_KEY", "env")

	// Test
	client := Client{}
	err = client.loadConfig("")
	require.CmpNoError(err)
	assert.Cmp(client, td.Struct(Client{
		AppKey:      "env",
		AppSecret:   "env",
		ConsumerKey: "env",
		endpoint:    OvhEU,
	}))
}

func TestConfigFromArgs(t *testing.T) {
	assert, require := td.AssertRequire(t)

	client := Client{AppKey: "param", AppSecret: "param", ConsumerKey: "param"}
	err := client.loadConfig("ovh-eu")
	require.CmpNoError(err)
	assert.Cmp(client, td.Struct(Client{
		AppKey:      "param",
		AppSecret:   "param",
		ConsumerKey: "param",
		endpoint:    OvhEU,
	}))
}

func TestEndpoint(t *testing.T) {
	assert, require := td.AssertRequire(t)

	err := ioutil.WriteFile(systemConfigPath, []byte(`
[ovh-eu]
application_key=ovh
application_secret=ovh
consumer_key=ovh

[https://api.example.com:4242]
application_key=example.com
application_secret=example.com
consumer_key=example.com
`), 0660)
	require.CmpNoError(err)

	// Clear
	t.Cleanup(func() {
		_ = ioutil.WriteFile(systemConfigPath, []byte(``), 0660)
	})

	// Test: by name
	client := Client{}
	err = client.loadConfig("ovh-eu")
	require.CmpNoError(err)
	assert.Cmp(client, td.Struct(Client{
		AppKey: "ovh",
	}))

	// Test: by URL
	client = Client{}
	err = client.loadConfig("https://api.example.com:4242")
	require.CmpNoError(err)
	assert.Cmp(client, td.Struct(Client{
		AppKey: "example.com",
	}))
}

func TestMissingParam(t *testing.T) {
	client := Client{AppKey: "param", AppSecret: "param", ConsumerKey: "param"}

	client.endpoint = ""
	err := client.loadConfig("")
	td.CmpString(t, err, `unknown endpoint '', consider checking 'Endpoints' list of using an URL`)

	client.AppKey = ""
	err = client.loadConfig("ovh-eu")
	td.CmpString(t, err, `missing application key, please check your configuration or consult the documentation to create one`)
	client.AppKey = "param"

	client.AppSecret = ""
	err = client.loadConfig("ovh-eu")
	td.CmpString(t, err, `missing application secret, please check your configuration or consult the documentation to create one`)
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
