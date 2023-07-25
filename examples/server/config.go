package main

import (
	flag "github.com/spf13/pflag"

	"github.com/spf13/viper"
)

var (
	config      = ""      // path to config file
	address     = ":7070" // server listen address
	prefix_path = "/"     // path to app
	loglevel    = "error"
)

func init() {
	flag.StringVar(&config, "config", config, "configuration file")
	flag.StringVar(&address, "address", address, "listen address")
	flag.StringVar(&prefix_path, "prefix_path", prefix_path, "prefix path")
	flag.StringVar(&loglevel, "loglevel", loglevel, "loglevel: error, info, debug")
}

func initConfig() {

	flag.Parse()

	// defaults
	viper.SetDefault("config", config)
	viper.SetDefault("address", address)
	viper.SetDefault("prefix_path", prefix_path)
	viper.SetDefault("loglevel", loglevel)

	// parse environment variables
	viper.SetEnvPrefix("II")
	viper.AutomaticEnv()

	// read config file
	if config != "" {
		viper.SetConfigFile(config)
	}

	err := viper.ReadInConfig()
	if err != nil {
		if viper.ConfigFileUsed() != "" {
			log.Fatal().Err(err).Msgf("unable to read config")
		}
	} else {
		log.Info().Msgf("using config file: %s", viper.ConfigFileUsed())
	}

	// set vars
	address = viper.GetString("address")
	prefix_path = viper.GetString("prefix_path")
	loglevel = viper.GetString("loglevel")

}
