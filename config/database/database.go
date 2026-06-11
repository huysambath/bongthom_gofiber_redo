package database

import (
	"log"
	"os"
	"sync"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

var(
	onec sync.Once
	db_pool *sqlx.DB
	listener *pq.Listener
)

func initializeDB() {
	DATABASE_URL := os.Getenv("DATABASE_URL")

	var err error
	db_pool, err = sqlx.Connect("postgres", DATABASE_URL)
	if err != nil {
		log.Fatalln("Error connecting to the database", err)
	}

	if err := db_pool.Ping(); err != nil {
		defer db_pool.Close()
		log.Fatalf("Failed to ping the datebase: %v",err)
	}
	log.Println("database connected successfully")

	db_pool.SetMaxIdleConns(10)
	db_pool.SetMaxOpenConns(10)
	db_pool.SetConnMaxLifetime(0)
}

func GetDB() *sqlx.DB {
	onec.Do(initializeDB)
	return db_pool
}