package data

import (
	"database/sql"
	"errors"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        uint      `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	UpdatedAt time.Time `json:"-"`
	CreatedAt time.Time `json:"-"`
	ExpiredAt time.Time `json:"-"`
}

type UserProvider struct {
	*sql.DB
}

func NewUserProvider(db *sql.DB) *UserProvider {
	return &UserProvider{db}
}

func (u *User) SetPassword(p string) error {
	if p == "" {
		return errors.New("password is required")
	}

	hp, err := bcrypt.GenerateFromPassword([]byte(p), 12)

	if err != nil {
		return err
	}

	u.Password = string(hp[:])

	return nil
}

func (u *User) ComparePassword(p string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p))
}

func (p *UserProvider) UniqueEmail(e string) (unique bool, err error) {
	query := `
	SELECT COUNT(*)
	FROM users
	WHERE email = $1
	`

	var count int

	err = p.DB.QueryRow(query, e).Scan(&count)

	unique = (count == 0)

	return
}
