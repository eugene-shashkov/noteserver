package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"notema/app"
	"notema/utils"
	"time"

	"github.com/gorilla/mux"
)

type loginjson struct {
	Status string `json:"status"`
}

// LoginHandler executes login operation with email and password
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	Email := vars["email"]
	Password := vars["password"]
	var uid int64
	db := app.Db()
	u, err := db.Prepare("SELECT id, password FROM users WHERE email = ? LIMIT 1")
	if err != nil {
		panic(err.Error())
	}
	defer u.Close()
	var hashedPassword sql.NullString
	responsejson := loginjson{}
	err = u.QueryRow(Email).Scan(&uid, &hashedPassword)

	// send response
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if hashedPassword.Valid && utils.CheckPasswordHash(Password, hashedPassword.String) {
		w.WriteHeader(http.StatusOK)
		responsejson.Status = "success"
		token := utils.GenerateToken()
		time := int32(time.Now().Unix())
		TokenToDb(token, uid, 1, time)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		responsejson.Status = "error"
	}
	resp, _ := json.Marshal(responsejson)
	w.Write(resp)

}
