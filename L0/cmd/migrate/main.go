package main

import (
	"l0/internal/migrate"
	"l0/internal/repository"
)

func main() {
	db, err := repository.NewRepository("internal/config/config.json")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	if err = migrate.CreateSchema(db); err != nil {
		panic(err)
	}
}
