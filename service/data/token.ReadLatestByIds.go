package data

func (p *TokenProvider) ReadLatestByIds(uid, aid string) (token string, err error) {
	q := `
	SELECT token
	FROM tokens
	WHERE user_id = $1 AND app_id = $2
	ORDER BY created_at
	DESC LIMIT 1;
	`

	err = p.QueryRow(q, uid, aid).Scan(&token)

	return
}
