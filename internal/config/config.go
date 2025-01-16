package config

import (
	"context"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	Host            string        `envconfig:"HOST" default:"0.0.0.0"`
	Port            string        `envconfig:"PORT" default:"3060"`
	ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" default:"10s"`
}

type Log struct {
	Level string `envconfig:"LEVEL" default:"info"`
}

type Config struct {
	ApplicationVersion string
	Server             Server
	Log 			   Log
}

func InitConfig(ctx context.Context) *Config {
	conf := &Config{}
	_ = godotenv.Load()

	err := envconfig.Process("", conf)
	if err != nil {
		log.WithError(err).Panic("Error loading .env file")
	}

	log.WithField("Config", conf).
		Info("Success on loading .env file")

	return conf

}
