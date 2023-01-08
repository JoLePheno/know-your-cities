package postgres

import (
	"github.com/JoLePheno/know-your-cities/internal/port"
	"github.com/go-pg/pg"
)

var _ port.Store = (*CitiesStore)(nil)

type CitiesStore struct {
	db *pg.DB
}

func NewCitiesStore() *CitiesStore {
	cfg := Config()

	db := pg.Connect(&pg.Options{
		User:     cfg.PostgresUser,
		Database: cfg.PostgresDatabase,
		Password: cfg.PostgresPassword,
		Addr:     cfg.PostgresAddr,
	})

	return &CitiesStore{db: db}
}
