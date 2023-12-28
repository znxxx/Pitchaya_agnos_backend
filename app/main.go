package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/znxxx/Pitchaya_agnos_backend/router"
)

func main() {
	db, err := sql.Open("postgres", "user=root password=root host=postgres port=5432 dbname=test_db sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := router.SetupRouter(db)

	r.Run(":8080")
}
