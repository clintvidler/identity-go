package data

func (s *Store) Populate() (err error) {
	u := User{
		Email:       "x@x",
		DisplayName: "Mr.x",
	}

	u.SetPassword("x")

	query := `
		INSERT INTO users (email, password, display_name)
		VALUES ($1, $2, $3)
		ON CONFLICT DO NOTHING;
	`

	if _, err = s.db.Exec(query, u.Email, u.Password, u.DisplayName); err != nil {
		return
	}

	u2 := User{
		Email:       "y@y",
		DisplayName: "Mrs.y",
	}

	u2.SetPassword("y")

	query2 := `
		INSERT INTO users (email, password, display_name)
		VALUES ($1, $2, $3)
		ON CONFLICT DO NOTHING;
	`

	if _, err = s.db.Exec(query2, u2.Email, u2.Password, u2.DisplayName); err != nil {
		return
	}

	return
}
