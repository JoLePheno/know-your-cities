package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	fmt.Println("init 1_create_cities_table")
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table cities...")
		_, err := db.Exec(`CREATE TABLE IF NOT EXISTS cities (
			id                  SERIAL PRIMARY KEY,
			external_id			BIGINT NOT NULL,
			created_at          TIMESTAMP NOT NULL,
			updated_at          TIMESTAMP NOT NULL,
			
			name				VARCHAR,
			post_code			VARCHAR NOT NULL
		);`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table cities...")
		_, err := db.Exec(`DROP TABLE cities`)
		return err
	})
}
