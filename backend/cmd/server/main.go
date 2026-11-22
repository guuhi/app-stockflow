package main

import (
	"log"
	"stockflow/internal/lib"
	"stockflow/internal/router"

	_ "github.com/lib/pq"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	db, err := lib.OpenDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	return router.NewRouter(db).Run(":8081")
}
