package data

func (p *UserProvider) UpdateOne(u User) (err error) {
	query := `
		UPDATE users 
		SET email = $2, password = $3, display_name = $4
		WHERE id = $1;
	`

	_, err = p.Exec(query, u.Id, u.Email, u.Password, u.DisplayName)

	return
}
