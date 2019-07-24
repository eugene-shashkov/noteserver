package handlers

import (
	"notema/app"
)

// SaveToken func that stores token to database
func SaveToken(token string, uid int64, Ttype int, time int32) bool {

	db := app.Db()
	t, err := db.Prepare("INSERT INTO tokens (token,uid,type,time) VALUES (?,?,?,?) ")
	if err != nil {
		panic(err.Error())
		return false
	}

	_, err = t.Exec(token, uid, Ttype, time)
	if err != nil {
		panic(err.Error())
	}
	return true
}
