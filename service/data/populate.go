package data

func (s *Store) Populate() (err error) {
	u := User{
		Email: "x@x",
	}

	u.SetPassword("x")

	query := `
		INSERT INTO users (email, password)
		VALUES ($1, $2)
		ON CONFLICT DO NOTHING;
	`

	if _, err = s.db.Exec(query, u.Email, u.Password); err != nil {
		return
	}

	u2 := User{
		Email: "y@y",
	}

	u2.SetPassword("y")

	query2 := `
		INSERT INTO users (email, password)
		VALUES ($1, $2)
		ON CONFLICT DO NOTHING;
	`

	if _, err = s.db.Exec(query2, u2.Email, u2.Password); err != nil {
		return
	}

	return
}
