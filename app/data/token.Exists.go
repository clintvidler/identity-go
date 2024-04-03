package data

func (p *TokenProvider) Exists(token string) (err error) {
	q := `
	SELECT token
	FROM tokens
	WHERE token = $1
	`

	var t string

	err = p.QueryRow(q, token).Scan(&t)

	return
}
