package ovh

import (
	"testing"

	"github.com/maxatome/go-testdeep/td"
)

const (
	systemConf                   = "testdata/system.ini"
	userPartialConf              = "testdata/userPartial.ini"
	userConf                     = "testdata/user.ini"
	userOAuth2Conf               = "testdata/user_oauth2.ini"
	userOAuth2InvalidConf        = "testdata/user_oauth2_invalid.ini"
	userOAuth2IncompatibleConfig = "testdata/user_oauth2_incompatible.ini"
	userBothConf                 = "testdata/user_both.ini"
	localPartialConf             = "testdata/localPartial.ini"
	localWithURLConf             = "testdata/localWithURL.ini"
	doesNotExistConf             = "testdata/doesNotExist.ini"
	invalidINIConf               = "testdata/invalid.ini"
	errorConf                    = "testdata"
)

func setConfigPaths(t testing.TB, paths ...string) {
	old := configPaths
	configPaths = paths
	t.Cleanup(func() { configPaths = old })
}

func TestConfigForbidsTrailingSlash(t *testing.T) {
	client := Client{}
	err := client.LoadConfig("https://example.org/")
	td.Require(t).String(err, "endpoint name cannot have a tailing slash")
}

func TestConfigFromFiles(t *testing.T) {
	setConfigPaths(t, systemConf, userPartialConf, localPartialConf)

	client := Client{}
	err := client.LoadConfig("ovh-eu")
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
	err := client.LoadConfig("ovh-eu")
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
	err := client.LoadConfig("ovh-eu")
	td.CmpString(t, err, `missing authentication information, you need to provide one of the following: application_key/application_secret, client_id/client_secret, or access_token`)
}

func TestConfigFromInvalidINIFile(t *testing.T) {
	setConfigPaths(t, invalidINIConf)

	client := Client{}
	err := client.LoadConfig("ovh-eu")
	td.CmpString(t, err, "cannot load configuration: unclosed section: [ovh\n")
}

func TestConfigFromInvalidFile(t *testing.T) {
	setConfigPaths(t, errorConf)

	client := Client{}
	err := client.LoadConfig("ovh-eu")
	td.CmpString(t, err, "cannot load configuration: BOM: read testdata: is a directory")
}

func TestConfigFromEnv(t *testing.T) {
	setConfigPaths(t, userConf)

	t.Setenv("OVH_ENDPOINT", "ovh-eu")
	t.Setenv("OVH_APPLICATION_KEY", "env")
	t.Setenv("OVH_APPLICATION_SECRET", "env")
	t.Setenv("OVH_CONSUMER_KEY", "env")

	client := Client{}
	err := client.LoadConfig("")
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
	err := client.LoadConfig("ovh-eu")
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
	err := client.LoadConfig("ovh-eu")
	require.CmpNoError(err)
	assert.Cmp(client, td.Struct(Client{
		AppKey: "ovh",
	}))

	// Test: by URL
	client = Client{}
	err = client.LoadConfig("https://api.example.com:4242")
	require.CmpNoError(err)
	assert.Cmp(client, td.Struct(Client{
		AppKey: "example.com",
	}))
}

func TestMissingParam(t *testing.T) {
	client := Client{AppKey: "param", AppSecret: "param", ConsumerKey: "param"}

	client.endpoint = ""
	err := client.LoadConfig("")
	td.CmpString(t, err, `unknown endpoint '', consider checking 'Endpoints' list or using an URL`)

	client.AppKey = ""
	err = client.LoadConfig("ovh-eu")
	td.CmpString(t, err, `invalid authentication config, both application_key and application_secret must be given`)
	client.AppKey = "param"

	client.AppSecret = ""
	err = client.LoadConfig("ovh-eu")
	td.CmpString(t, err, `invalid authentication config, both application_key and application_secret must be given`)
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

func TestConfigOAuth2(t *testing.T) {
	setConfigPaths(t, userOAuth2Conf)

	client := Client{}
	err := client.LoadConfig("ovh-eu")
	td.Require(t).CmpNoError(err)
	td.Cmp(t, client, td.Struct(Client{
		ClientID:     "foo",
		ClientSecret: "bar",
	}))
}

func TestConfigInvalidBoth(t *testing.T) {
	setConfigPaths(t, userBothConf)

	client := Client{}
	err := client.LoadConfig("ovh-eu")
	td.CmpString(t, err, "can't use multiple authentication methods: application_key/application_secret, client_id/client_secret")
}

func TestConfigOAuth2Invalid(t *testing.T) {
	setConfigPaths(t, userOAuth2InvalidConf)

	client := Client{}
	err := client.LoadConfig("ovh-eu")
	td.CmpString(t, err, "invalid oauth2 config, both client_id and client_secret must be given")
}

func TestConfigOAuth2Incompatible(t *testing.T) {
	setConfigPaths(t, userOAuth2IncompatibleConfig)

	client := Client{}
	err := client.LoadConfig("kimsufi-eu")
	td.CmpString(t, err, `oauth2 authentication is not compatible with endpoint "https://eu.api.kimsufi.com/1.0"`)
}
