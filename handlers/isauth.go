package handlers

import (
	"database/sql"
	"noteserver/app"
)

// IsAuth check if user token valid
func IsAuth(token string) (bool, int) {
	db := app.Db()
	findtoken, err := db.Prepare("SELECT token, uid FROM tokens WHERE token=? ")
	if err != nil {
		panic(err.Error())
	}
	defer findtoken.Close()
	var fToken sql.NullString
	var uid int
	findtoken.QueryRow(token).Scan(&fToken, &uid)
	if fToken.Valid {
		return true, uid
	}
	return false, uid
}
