package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/caarlos0/env"
	"github.com/go-pg/migrations"
	"github.com/go-pg/pg"
)

// Copy/Pasta from https://github.com/go-pg/migrations/

const usageText = `This program runs command on the db. Supported commands are:
  - init - creates version info table in the database
  - up - runs all available migrations.
  - up [target] - runs available migrations up to the target one.
  - down - reverts last migration.
  - reset - reverts all migrations.
  - version - prints current db version.
  - set_version [version] - sets db version without running migrations.
Usage:
  go run *.go <command> [args]
`

type config struct {
	PostgresUser     string `env:"POSTGRES_USER" envDefault:"postgres"`
	PostgresDatabase string `env:"POSTGRES_DATABASE" envDefault:"postgres"`
	PostgresPassword string `env:"POSTGRES_PASSWORD" envDefault:"password"`
	PostgresAddr     string `env:"POSTGRES_ADDR" envDefault:"localhost:5432"`
}

func main() {
	var cfg config
	if err := env.Parse(&cfg); err != nil {
		log.Fatal("parsing envs failed: %w", err)
	}

	flag.Usage = usage
	flag.Parse()

	db := pg.Connect(&pg.Options{
		User:     cfg.PostgresUser,
		Database: cfg.PostgresDatabase,
		Password: cfg.PostgresPassword,
		Addr:     cfg.PostgresAddr,
	})
	waitForPostgres(db)

	oldVersion, newVersion, err := migrations.Run(db, flag.Args()...)
	if err != nil {
		exitf(err.Error())
	}
	if newVersion != oldVersion {
		fmt.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		fmt.Printf("version is %d\n", oldVersion)
	}
}

func waitForPostgres(db *pg.DB) {
	var err error
	waitTime := 15 * time.Second
	waitUntil := time.Now().Add(waitTime)
start:
	if time.Now().After(waitUntil) {
		log.Fatalf("Waited %s for database to get ready but wasn't: %v", waitTime, err)
	}

	_, err = db.ExecOne("SELECT 1;")
	if err != nil {
		time.Sleep(time.Second)
		fmt.Print(".")
		goto start
	}
}

func usage() {
	fmt.Print(usageText)
	flag.PrintDefaults()
	os.Exit(2)
}

func errorf(s string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, s+"\n", args...)
}

func exitf(s string, args ...interface{}) {
	errorf(s, args...)
	os.Exit(1)
}
