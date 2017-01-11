package connect

import (
	"errors"
	"fmt"
	"os/user"
	"path/filepath"

	"github.com/go-ini/ini"
	"github.com/runabove/go-sdk/ovh"
	"github.com/runabove/go-sdk/ovh/ovhcli/common"
	"github.com/spf13/cobra"
)

var (
	userConfigPath  = "/.ovh.conf" // prefixed with homeDir
	localConfigPath = "./ovh.conf"
)

func helper(consumerKey string) error {
	return fmt.Errorf("Set environment variable OVH_CONSUMER_KEY=%s\n", consumerKey)
}

func writeConsumerKeyFile(filename string, consumerKey string) (err error) {
	var cfg *ini.File
	var section *ini.Section
	var endpoint string
	var endpointKey *ini.Key

	if cfg, err = ini.Load(filename); err != nil {
		return errors.New("Cannot load file " + filename)
	}

	if defaultSection, errSection := cfg.GetSection("default"); errSection == nil {
		if _, errKey := defaultSection.GetKey("endpoint"); errKey != nil {
			return errors.New("Cannot read endpoint from configuration")
		}
		endpoint = endpointKey.String()
	} else {
		return errors.New("Cannot read default section")
	}

	if section, err = cfg.GetSection(endpoint); err != nil {
		if section, err = cfg.NewSection(endpoint); err != nil {
			return errors.New("Cannot create section " + endpoint)
		}
	}

	if section.NewKey("consumer_key", consumerKey); err != nil {
		return errors.New("Cannot create key 'consumer_key'")
	}

	if err = cfg.SaveTo(filename); err != nil {
		return errors.New("Cannot save to " + filename)
	}
	return
}

func writeConsumerKey(consumerKey string) (err error) {
	currentUserConfigPath := localConfigPath
	if usr, err := user.Current(); err == nil {
		currentUserConfigPath = filepath.Join(usr.HomeDir, userConfigPath)
	}

	if err := writeConsumerKeyFile(localConfigPath, consumerKey); err != nil {
		if err := writeConsumerKeyFile(currentUserConfigPath, consumerKey); err != nil {
			return helper(consumerKey)
		}
	}
	return
}

// Cmd domain
var Cmd = &cobra.Command{
	Use:   "connect",
	Short: "Domain commands: ovhcli connect",
	Long:  `Domain commands: ovhcli connect`,
	Run: func(cmd *cobra.Command, args []string) {

		c, err := ovh.NewDefaultClient()
		common.Check(err)

		ckReq := c.NewCkRequest()

		// Allow GET method on /me
		ckReq.AddRules(ovh.ReadWrite, "/*")

		response, err := ckReq.Do()
		common.Check(err)

		// Print the validation URL
		fmt.Printf("Please visit %s to complete your login\n", response.ValidationURL)

		// set consumer key
		if err = writeConsumerKey(response.ConsumerKey); err != nil {
			common.Check(err)
		}
	},
}
