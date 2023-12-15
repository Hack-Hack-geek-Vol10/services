package config

import (
	"github.com/caarlos0/env"
	"github.com/schema-creator/services/sql-service/pkg/logger"
)

func LoadEnv(l logger.Logger) {
	Config = &config{}

	if err := env.Parse(&Config.Server); err != nil {
		l.Panic(err)
	}

	if err := env.Parse(&Config.NewRelic); err != nil {
		l.Panic(err)
	}
}
