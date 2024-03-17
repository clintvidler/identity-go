package data

import "time"

func (p *TokenProvider) CreateOne(uid, aid, t string, d time.Duration) (err error) {
	query := `
	INSERT INTO tokens (user_id, app_id, token, expired_at)
	VALUES ($1, $2, $3, $4);
	`

	_, err = p.Exec(query, uid, aid, t, time.Now().Add(d))

	return
}
