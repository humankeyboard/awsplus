package types

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"log"
	"os"
	"os/user"
	"path"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type CliConfig struct {
	CredentialsFile string
	ConfigFile      *goconfig.ConfigFile
}

func NewCliConfig() *CliConfig {
	c := new(CliConfig)
	c.CredentialsFile = c.getCredentialsFile()
	c.ReadConfig()
	return c
}

// returns the credentials file for the user:
//  ~/.aws/credentials
func (c *CliConfig) getCredentialsFile() string {
	usr, err := user.Current()
	check(err)

	cdPath := path.Join(usr.HomeDir, ".aws")
	cfPath := path.Join(usr.HomeDir, ".aws", "credentials")

	file, err := os.Open(cfPath) // check if file exists
	file.Close()
	if err != nil { // if it doesn't, create the file
		os.MkdirAll(cdPath, 0777)
		file, err := os.Create(cfPath)
		if err != nil {
			check(err)
		}
		file.Close()
	}
	return cfPath
}

func (c *CliConfig) ReadConfig() {
	cfg, err := goconfig.LoadConfigFile(c.CredentialsFile)
	check(err)
	c.ConfigFile = cfg
}

// this is not used and to be removed
func (c *CliConfig) ListProfiles() {
	for _, section := range c.ConfigFile.GetSectionList() {
		fmt.Printf("-profile->%s\n", section)
	}
}

// each section defines a profile
func (c *CliConfig) GetKeyID(profile string) string {
	cfg := c.ConfigFile
	value, err := cfg.GetValue(profile, "aws_access_key_id")
	check(err)
	return value
}

// each section defines a profile
func (c *CliConfig) GetAccessKey(profile string) string {
	cfg := c.ConfigFile
	value, err := cfg.GetValue(profile, "aws_secret_access_key")
	check(err)
	return value
}

func (c *CliConfig) GetProfiles() []string {
	return c.ConfigFile.GetSectionList()
}

func (c *CliConfig) WriteProfile(p, id, key string) {
	cfg := c.ConfigFile
	cfg.SetValue(p, "aws_access_key_id", id)
	cfg.SetValue(p, "aws_secret_access_key", key)
	goconfig.SaveConfigFile(c.ConfigFile, c.CredentialsFile)
}
