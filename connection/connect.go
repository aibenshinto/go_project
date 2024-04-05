package connection

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {

	db, err := sql.Open("postgres", "postgres://vefceiis:pzTKYphme3DcRyDg3ertxItmrtc44HiQ@rain.db.elephantsql.com/vefceiis")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to PostgreSQL database successfully")
	return db
}
