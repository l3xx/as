package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

const (
	//EnvPrefix env base const
	EnvPrefix = "t"
	//EnvAPIHost host config string
	EnvAPIHost = "api_host"
	//EnvLevel env level config string
	EnvLevel = "env_level"
	//EnvDebug is debug
	EnvDebug = "debug"
)

//Config config struct
type Config struct {
	APIHost  string
	EnvLevel string
	Debug    bool
}

// Get builds and returns config from environment variables
func Get() (*Config, error) {
	err := godotenv.Load()
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}
	cfg := Config{}
	viper.AutomaticEnv()
	viper.SetEnvPrefix(EnvPrefix)

	cfg.EnvLevel = viper.GetString(EnvLevel)
	cfg.APIHost = viper.GetString(EnvAPIHost)
	cfg.Debug = viper.GetBool(EnvDebug)

	return &cfg, nil
}
