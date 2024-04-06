package data

// Upsert email and generate confirmation key
func (p *UserProvider) UpsertPendingResetPassword(e string) (key string, err error) {
	// Note:
	// - md5(random()::text) has an infinitesimally small chance of generating
	//   a collision.
	// - Upon collision, an unexpected error will be thrown due to the
	//   uniqueness constraint, thereby preventing storage of a duplicate key.
	query := `
	INSERT INTO pending_reset_password (email, key, expired_at)
	VALUES ($1, md5(random()::text), NOW() + INTERVAL '24 hour')
	ON CONFLICT (email) DO
	UPDATE
	SET key=md5(random()::text)
	RETURNING key;
	`

	err = p.QueryRow(query, e).Scan(&key)

	return
}

func (p *UserProvider) ReadPendingResetPassword(k string) (email string, err error) {
	query := `
	SELECT email
	FROM pending_reset_password
	WHERE key = $1 AND expired_at >= NOW();
	`

	err = p.QueryRow(query, k).Scan(&email)

	return
}

func (p *UserProvider) ResetPassword(u User) (err error) {
	query := `
	UPDATE users 
	SET password = $2
	WHERE email = $1;
	`

	_, err = p.Exec(query, u.Email, u.Password)

	if err != nil {
		return
	}

	query2 := `
	DELETE
	FROM pending_reset_password
	WHERE email = $1
	`

	_, err = p.Exec(query2, u.Email)

	return
}
