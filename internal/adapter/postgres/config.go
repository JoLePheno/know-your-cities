package postgres

import (
	"log"

	"github.com/caarlos0/env"
)

type config struct {
	PostgresUser     string `env:"POSTGRES_USER" envDefault:"postgres"`
	PostgresDatabase string `env:"POSTGRES_DATABASE" envDefault:"postgres"`
	PostgresPassword string `env:"POSTGRES_PASSWORD" envDefault:"password"`
	PostgresAddr     string `env:"POSTGRES_ADDR" envDefault:"localhost:5432"`
}

func Config() *config {
	var cfg config
	if err := env.Parse(&cfg); err != nil {
		log.Fatal("parsing envs failed: %w", err)
	}
	return &cfg
}
