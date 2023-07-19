package config

import (
	"os"

	"github.com/deka-microservices/go-bcrypt-service/internal/consts"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func init() {
	// Start-up log configuration
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	viper.SetEnvPrefix("BCRYPT_SERVICE")
	viper.AutomaticEnv()

	viper.SetDefault(consts.CONFIG_PORT, 9000)
	viper.SetDefault(consts.CONFIG_IP, "0.0.0.0")
	viper.SetDefault(consts.CONFIG_COST, 14)

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/go-bcrypt-service")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Warn().Err(err).Msg("failed to find config file. using default value or env")
		} else {
			log.Fatal().Err(err).Msg("failed to open config file")
		}
	}
}
