package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func ConnectDB(dbUrl string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbUrl)

	if err != nil {
		return nil, fmt.Errorf("error opening database %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("database ping failed %w", err)
	}

	fmt.Println("Connected to the database")
	return db, nil
}
