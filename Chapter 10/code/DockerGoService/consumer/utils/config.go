package utils

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"rest_api_mq/consumer/consts"
)

func init() {
	viper.SetConfigFile(consts.ENV_FILE)
	viper.AddConfigPath(consts.ENV_FILE_DIRECTORY)
	err := viper.ReadInConfig()
	if err != nil {
		log.Debug().Err(err).
			Msg("Error occurred while reading env file, might fallback to OS env config")
	}
	viper.AutomaticEnv()
}

func GetEnvVar(name string) string {
	if !viper.IsSet(name) {
		log.Debug().Msgf("Environment variable %s is not set", name)
		return ""
	}
	value := viper.GetString(name)
	return value
}
