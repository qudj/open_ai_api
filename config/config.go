package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var DEBUG = os.Getenv("DEBUG") != ""
var Global = initConfig()

type Config struct {
	ENV       string
	Debug     bool
	AuthHosts struct {
		HanHaiHost AuthHost `mapstructure:"han_hai_host"`
	} `mapstructure:"auth_hosts"`
}

type AuthHost struct {
	Host   string   `mapstructure:"host"`
	Tokens []string `mapstructure:"tokens"`
}

func initConfig() *Config {
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}

	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("../config") // add for tests package
	viper.SetConfigName(env)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
	config.ENV = env
	fmt.Printf("config: %+v\n", config)
	return &config
}
