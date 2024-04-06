package data

// Upsert email and generate confirmation key
func (p *UserProvider) UpsertPendingUpdateEmail(uid, e string) (key string, err error) {
	// Note:
	// - md5(random()::text) has an infinitesimally small chance of generating
	//   a collision.
	// - Upon collision, an unexpected error will be thrown due to the
	//   uniqueness constraint, thereby preventing storage of a duplicate key.
	query := `
	INSERT INTO pending_update_email (user_id, email, key, expired_at)
	VALUES ($1, $2, md5(random()::text), NOW() + INTERVAL '24 hour')
	ON CONFLICT (email) DO
	UPDATE
	SET key=md5(random()::text)
	RETURNING key;
	`

	err = p.QueryRow(query, uid, e).Scan(&key)

	return
}

func (p *UserProvider) ReadPendingUpdateEmail(k string) (user_id, email string, err error) {
	query := `
	SELECT user_id, email
	FROM pending_update_email
	WHERE key = $1 AND expired_at >= NOW();
	`

	err = p.QueryRow(query, k).Scan(&user_id, &email)

	return
}

func (p *UserProvider) UpdateEmail(uid, e string) (err error) {
	query := `
	UPDATE users 
	SET email = $2
	WHERE id = $1;
	`

	_, err = p.Exec(query, uid, e)

	if err != nil {
		return
	}

	query2 := `
	DELETE
	FROM pending_update_email
	WHERE user_id = $1
	`

	_, err = p.Exec(query2, uid)

	return
}
