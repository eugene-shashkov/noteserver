package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"noteserver/app"
	"noteserver/utils"

	"github.com/gorilla/mux"
)

type successjson struct {
	Status string `json:"status"`
}

// RegistrationHandler is a function where we register notema our user
func RegistrationHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	Email := vars["email"]
	Password := vars["password"]
	Name := vars["name"]
	time := int32(time.Now().Unix())
	var uid int64

	var passEncr string
	var err error
	passEncr, err = utils.HashPassword(Password)
	if err != nil {
		panic(err)
	}

	db := app.Db()
	u, err := db.Prepare("SELECT email FROM users WHERE email = ? LIMIT 1")
	if err != nil {
		panic(err.Error())
	}
	defer u.Close()
	var findEmail sql.NullString
	responsejson := successjson{}
	err = u.QueryRow(Email).Scan(&findEmail)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if findEmail.Valid {
		// User already registred
		w.WriteHeader(http.StatusInternalServerError)
		responsejson.Status = "error"
	} else {
		// Create new user profile
		w.WriteHeader(http.StatusOK)
		createUserInsert, err := db.Prepare("INSERT INTO users (name, email, password,time) VALUES(?,?,?,?) ")
		_, err = createUserInsert.Exec(Name, Email, passEncr, time)
		if err != nil {
			panic(err.Error())
		}

		// get user id into uid variable
		userData, err := db.Prepare("SELECT id FROM users WHERE email=? LIMIT 1 ")
		userData.QueryRow(Email).Scan(&uid)
		if err != nil {

			panic(err.Error())
		}

		// create random user token
		token := utils.GenerateToken()
		TokenQuery(token, uid, 1, time)

		responsejson.Status = token
	}
	// send response
	resp, _ := json.Marshal(responsejson)
	w.Write(resp)
}
