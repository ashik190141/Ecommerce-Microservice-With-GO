package db

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)


func ConnectPostgres() *sqlx.DB {
	dsn := os.Getenv("DATABASE_URL")

	// Build DSN from docker-compose environment variables
	if dsn == "" {
		host := os.Getenv("POSTGRES_HOST")
		port := os.Getenv("POSTGRES_PORT")
		user := os.Getenv("POSTGRES_USER")
		password := os.Getenv("POSTGRES_PASSWORD")
		dbname := os.Getenv("POSTGRES_DB")
		
		if host != "" {
			dsn = "postgresql://" + user +":" + password + "@" + host +":" + port + "/" + dbname + "?sslmode=disable"
		} else {
			log.Fatalln("Error: DATABASE_URL or POSTGRES_* environment variables are not set.")
		}
	}

	// Check if DSN is still empty in locally
	// if dsn == "" {
	// 	log.Fatalln("Error: DATABASE_URL environment variable is not set.")
	// }

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