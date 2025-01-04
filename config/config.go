package config

import (
	"github.com/spf13/viper"
	"log"
	"sync"
)

var once sync.Once
var c *Config

type Config struct {
	PostgresURL    string `mapstructure:"POSTGRES_URL"`
	JwtSecret      string `mapstructure:"JWT_SECRET"`
	PwdMemory      uint32 `mapstructure:"PWD_MEMORY"`
	PwdIterations  uint32 `mapstructure:"PWD_ITERATIONS"`
	PwdParallelism uint8  `mapstructure:"PWD_PARALLELISM"`
	PwdSaltLength  uint32 `mapstructure:"PWD_SALT_LENGTH"`
	PwsKeyLength   uint32 `mapstructure:"PWD_KEY_LENGTH"`
}

func NewConfig() *Config {
	once.Do(func() {
		viper.SetConfigFile(".env")
		err := viper.ReadInConfig()
		if err != nil {
			log.Fatalf("Error reading config file, %s", err)
		}
		err = viper.Unmarshal(&c)
		if err != nil {
			log.Fatalf("Error unmarshalling config, %s", err)
		}
	})
	return c
}
