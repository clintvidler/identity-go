package data

func (p *UserProvider) ReadPendingRegistration(k string) (email string, err error) {
	query := `
	SELECT email
	FROM pending_registration
	WHERE key = $1 AND expired_at >= NOW();
	`

	err = p.QueryRow(query, k).Scan(&email)

	return
}

// Upsert email and generate confirmation key
func (p *UserProvider) UpsertPendingRegistration(e string) (key string, err error) {
	// Note:
	// - md5(random()::text) has an infinitesimally small chance of generating
	//   a collision.
	// - Upon collision, an unexpected error will be thrown due to the
	//   uniqueness constraint, thereby preventing storage of a duplicate key.
	// - TODO: [SECURITY] Are these keys predictible and risking unauthorised account access?
	query := `
	INSERT INTO pending_registration (email, key, expired_at)
	VALUES ($1, md5(random()::text), NOW() + INTERVAL '24 hour')
	ON CONFLICT (email) DO
	UPDATE
	SET key=md5(random()::text)
	RETURNING key;
	`

	err = p.QueryRow(query, e).Scan(&key)

	return
}

func (p *UserProvider) Register(u User) (id uint, err error) {
	query := `
	INSERT INTO users (email, display_name, password)
	VALUES ($1, $2, $3)
	RETURNING id
	`

	if err = p.QueryRow(query, u.Email, u.DisplayName, u.Password).Scan(&id); err != nil {
		return
	}

	query2 := `
	DELETE
	FROM pending_registration
	WHERE email = $1
	`

	_, err = p.Exec(query2, u.Email)

	return
}
