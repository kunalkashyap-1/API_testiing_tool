package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" //postgresSQL driver
	"log"
)
//DB is the global database connection pool
var DB *sql.DB

func Initialize(connStr string)(*sql.DB, error){
	db, err := sql.Open("postgres", connStr)

	if err!= nil {
		return nil, fmt.Errorf("Error opening databse connection: %w", err)
	}

	if err := db.Ping(); err != nil{
		return nil, fmt.Errorf("error pinging databse: %w", err)
	}

	DB = db
	log.Println("Successfully connected to the database")
	return DB, nil
}

func Close(){
	if DB != nil{
		err := DB.Close()
		if err != nil{
			log.Printf("Error closing the database connection: %v", err)
		}
	}
}