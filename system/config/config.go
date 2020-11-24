package config

import (
	"encoding/json"
	"fmt"

	fwConfig "github.com/d3ta-go/system/system/config"
	"github.com/spf13/viper"
)

// NewConfig create new Config
func NewConfig(path string) (*Config, *viper.Viper, error) {

	defaultConfigFile, err := fwConfig.GetConfigFilePath(path)
	if err != nil {
		return nil, nil, err
	}

	v := viper.New()
	v.SetConfigFile(defaultConfigFile)
	err = v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err.Error()))
	}
	v.WatchConfig()

	c := new(Config)
	err = v.Unmarshal(&c)
	if err != nil {
		return nil, nil, err
	}

	return c, v, err
}

// Config represent Config
type Config struct {
	// Environment refers to fwConfig.Environment
	Environment fwConfig.Environment `json:"environment" yaml:"environment"`

	// DirLocations refers to fwConfig.Config
	DirLocations fwConfig.DirLocations `json:"dirLocations" yaml:"dirLocations"`

	// Add your custome Config
}

// ToJSON convert Config to JSON
func (c *Config) ToJSON() []byte {
	JSON, err := json.Marshal(c)
	if err != nil {
		return nil
	}
	return JSON
}

// CanRunTest can Run Test
func (c *Config) CanRunTest() bool {
	can := false
	for _, v := range c.Environment.RunTestEnvironment {
		if string(v) == c.Environment.Stage {
			can = true
			break
		}
	}
	return can
}
