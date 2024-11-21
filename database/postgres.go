package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

func NewPostgresDB(dbURL string, maxOpenConn, maxIdleConns int, maxIdleTime string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	duration, err := time.ParseDuration(maxIdleTime)

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConn)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	err = db.PingContext(ctx)

	if err != nil {
		return nil, err
	}
	fmt.Printf("The DB with URL::%v is connected successfully!\n", dbURL)
	return db, nil
}
