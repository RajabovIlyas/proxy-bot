package config

import (
	"errors"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig
}

type ServerConfig struct {
	TelegramApiToken string
	Port             string
	ChannelName      string
}

func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			return nil, errors.New("config files not found")
		}
		return nil, err
	}

	return v, nil
}

func ParseConfig(v *viper.Viper, logger zerolog.Logger) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		logger.Error().Err(err).Msgf("unable to decode into struct")
		return nil, err
	}

	return &c, nil
}
