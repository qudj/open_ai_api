package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var DEBUG = os.Getenv("DEBUG") != ""
var Global = initConfig()

type Config struct {
	ENV   string
	Debug bool
	Host  struct {
		SolanaPrivateHost string `mapstructure:"solana_private_host"`
	} `mapstructure:"host"`
	Redis RedisConfig `mapstructure:"redis"`
}

type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	Ca       string `mapstructure:"ca"`
	Region   string `mapstructure:"region"`
	AwsAuth  bool   `mapstructure:"aws_auth"`
}

type RedisConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
	DB   int    `mapstructure:"db"`
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
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetEnvPrefix("")
	viper.AutomaticEnv()
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
