package data

// Match against id or email
func (p *UserProvider) ReadOne(id uint, email string) (u User, err error) {
	query := `
	SELECT id, email, password
	FROM users
	WHERE id = $1 OR email = $2
	`

	err = p.QueryRow(query, id, email).Scan(&u.Id, &u.Email, &u.Password)

	return
}
