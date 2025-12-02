package db

import (
	"Product-Service/config"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)


func ConnectPostgres() *sqlx.DB {
	dsn := config.LoadEnvData().DATABASE_URL

	if dsn == "" {
		log.Fatalln("Error: DATABASE_URL environment variable is not set.")
	}

	const driverName = "postgres"

	db, err := sqlx.Connect(driverName, dsn)
	if err != nil {
		log.Fatalf("Failed to connect to the database using driver %s: %v", driverName, err)
	}
	// defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	log.Println("Successfully connected to the database!")
	log.Println("Database driver in use:", driverName)
	return db
}