package data

func (p TokenProvider) DeleteByUser(uid string) (err error) {
	q := `
	DELETE FROM tokens
	WHERE user_id = $1
	`

	_, err = p.Exec(q, uid)

	return
}
