package data

func (p *TokenProvider) DeleteAppFamily(uid, aid string) (err error) {
	q := `
	DELETE FROM tokens
	WHERE user_id = $1 AND app_id = $2;
	`

	_, err = p.Exec(q, uid, aid)

	return
}
