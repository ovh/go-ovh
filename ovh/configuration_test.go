package ovh

import (
	"testing"

	"github.com/maxatome/go-testdeep/td"
)

const (
	systemConf       = "testdata/system.ini"
	userPartialConf  = "testdata/userPartial.ini"
	userConf         = "testdata/user.ini"
	localPartialConf = "testdata/localPartial.ini"
	localWithURLConf = "testdata/localWithURL.ini"
	doesNotExistConf = "testdata/doesNotExist.ini"
	invalidINIConf   = "testdata/invalid.ini"
	errorConf        = "testdata"
)

func setConfigPaths(t testing.TB, paths ...string) {
	old := configPaths
	configPaths = paths
	t.Cleanup(func() { configPaths = old })
}

func TestConfigForbidsTrailingSlash(t *testing.T) {
	client := Client{}
	err := client.loadConfig("https://example.org/")
	td.Require(t).String(err, "endpoint name cannot have a tailing slash")
}

func TestConfigFromFiles(t *testing.T) {
	setConfigPaths(t, systemConf, userPartialConf, localPartialConf)

	client := Client{}
	err := client.loadConfig("ovh-eu")
	td.Require(t).CmpNoError(err)
	td.Cmp(t, client, td.Struct(Client{
		AppKey:      "system",
		AppSecret:   "user",
		ConsumerKey: "local",
	}))
}

func TestConfigFromOnlyOneFile(t *testing.T) {
	setConfigPaths(t, userConf)

	client := Client{}
	err := client.loadConfig("ovh-eu")
	td.Require(t).CmpNoError(err)
	td.Cmp(t, client, td.Struct(Client{
		AppKey:      "user",
		AppSecret:   "user",
		ConsumerKey: "user",
	}))
}

func TestConfigFromNonExistingFile(t *testing.T) {
	setConfigPaths(t, doesNotExistConf)

	client := Client{}
	err := client.loadConfig("ovh-eu")
	td.CmpString(t, err, `missing application key, please check your configuration or consult the documentation to create one`)
}

func TestConfigFromInvalidINIFile(t *testing.T) {
	setConfigPaths(t, invalidINIConf)

	client := Client{}
	err := client.loadConfig("ovh-eu")
	td.CmpString(t, err, "cannot load configuration: unclosed section: [ovh\n")
}

func TestConfigFromInvalidFile(t *testing.T) {
	setConfigPaths(t, errorConf)

	client := Client{}
	err := client.loadConfig("ovh-eu")
	td.CmpString(t, err, "cannot load configuration: BOM: read testdata: is a directory")
}

func TestConfigFromEnv(t *testing.T) {
	setConfigPaths(t, userConf)

	t.Setenv("OVH_ENDPOINT", "ovh-eu")
	t.Setenv("OVH_APPLICATION_KEY", "env")
	t.Setenv("OVH_APPLICATION_SECRET", "env")
	t.Setenv("OVH_CONSUMER_KEY", "env")

	client := Client{}
	err := client.loadConfig("")
	td.Require(t).CmpNoError(err)
	td.Cmp(t, client, td.Struct(Client{
		AppKey:      "env",
		AppSecret:   "env",
		ConsumerKey: "env",
		endpoint:    OvhEU,
	}))
}

func TestConfigFromArgs(t *testing.T) {
	setConfigPaths(t, userConf)

	client := Client{AppKey: "param", AppSecret: "param", ConsumerKey: "param"}
	err := client.loadConfig("ovh-eu")
	td.Require(t).CmpNoError(err)
	td.Cmp(t, client, td.Struct(Client{
		AppKey:      "param",
		AppSecret:   "param",
		ConsumerKey: "param",
		endpoint:    OvhEU,
	}))
}

func TestEndpoint(t *testing.T) {
	assert, require := td.AssertRequire(t)

	setConfigPaths(t, localWithURLConf)

	// Test: by name
	client := Client{}
	err := client.loadConfig("ovh-eu")
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

func TestConfigPaths(t *testing.T) {
	home, err := currentUserHome()
	td.Require(t).CmpNoError(err)

	setConfigPaths(t, "", "file", "file.ini", "dir/file.ini", "~/file.ini", "~typo.ini")

	td.Cmp(t, home, td.Not(td.HasSuffix("/")))
	td.Cmp(t,
		expandConfigPaths(),
		[]interface{}{"", "file", "file.ini", "dir/file.ini", home + "/file.ini", "~typo.ini"},
	)
}
