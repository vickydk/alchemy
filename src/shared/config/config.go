package config

import (
	"fmt"

	"github.com/spf13/viper"

	Database "alchemy/src/shared/database"
)

type Config struct {
	Apps     Apps                    `json:"apps"`
	Database Database.ConfigDatabase `json:"database"`
}

type Apps struct {
	Name     string `json:"name"`
	HttpPort int    `json:"httpPort"`
	GRPCPort int    `json:"grpcPort"`
	Version  string `json:"version"`
}

func (c *Config) AppAddress() string {
	return fmt.Sprintf(":%v", c.Apps.HttpPort)
}

func NewConfig(path string) *Config {
	fmt.Println("Try NewConfig ... ")

	viper.SetConfigFile(path)
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	conf := Config{}
	err := viper.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}

	return &conf
}
