package data

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Store struct {
	db    *sql.DB
	User  *UserProvider
	Token *TokenProvider
}

func NewStore(host, port, username, password, name string) (*Store, error) {
	// Prepare the database connection string

	ds := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, name)

	// Create the database connection

	db, err := sql.Open("pgx", ds)

	if err != nil {
		return nil, err
	}

	// Test the database connection

	err = db.Ping()

	// Finalise

	return &Store{db: db, User: NewUserProvider(db), Token: NewTokenProvider(db)}, err
}
