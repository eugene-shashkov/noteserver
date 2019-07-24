package handlers

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"notema/app"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type createNoteJSON struct {
	Status string `json:"status"`
}

type readNoteJSON struct {
	Status      string  `json:"status"`
	Data        []*Note `json:"data"`
	LastPage    int64   `json:"last_page"`
	CurrentPage int64   `json:"current_page"`
}

// Note data type for notes
type Note struct {
	ID   int64  `json:"id"`
	Note string `json:"note"`
}

type updateNotesJSON struct {
	Status string `json:"status"`
}

func isInteger(val float64) bool {
	return val == float64(int(val))
}

// CreateNoteHandler is a handler for creation new note
func CreateNoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	note := vars["note"]
	token := vars["token"]
	time := int32(time.Now().Unix())

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	j := createNoteJSON{}

	db := app.Db()

	var isTokenValid bool
	var uid int

	isTokenValid, uid = IsAuth(token)

	if isTokenValid {
		prNote, err := db.Prepare("INSERT INTO notes (note,uid,time) VALUES (?,?,?) ")
		if err != nil {
			panic(err.Error())
		}
		prNote.Exec(note, uid, time)
		w.WriteHeader(http.StatusOK)
		j.Status = "success"

	} else {
		w.WriteHeader(http.StatusInternalServerError)
		j.Status = "error"
	}

	// send response
	resp, _ := json.Marshal(j)
	w.Write(resp)
}

// ReadNotesHandler is a handler to view notes
func ReadNotesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	token := vars["token"]
	page, _ := strconv.ParseFloat(vars["page"], 64)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	j := readNoteJSON{}

	db := app.Db()

	var isTokenValid bool
	var uid int

	isTokenValid, uid = IsAuth(token)

	if isTokenValid {

		// implementing pagination
		var total float64
		prNote, err := db.Prepare("SELECT COUNT(*) AS positions FROM notes WHERE uid=? ")
		if err != nil {
			panic(err.Error())
		}
		prNote.QueryRow(uid).Scan(&total)

		var pagesCount float64
		var perPage int
		var perPageEnv string
		perPageEnv = os.Getenv("NOTES_PER_PAGE")

		perPage, _ = strconv.Atoi(perPageEnv)

		pagesCount = total / float64(perPage)
		pagesCount = math.Ceil(pagesCount)

		var fromPosition float64
		var toPosition float64

		if page <= pagesCount {
			w.WriteHeader(http.StatusOK)
			j.Status = "success"

			notes := make([]*Note, 0)

			fromPosition = page*float64(perPage) - float64(perPage)
			toPosition = page * float64(perPage)

			pNoteDataRows, err := db.Query("SELECT id,note FROM notes WHERE uid=? LIMIT ?,? ", uid, fromPosition, toPosition)
			if err != nil {
				log.Fatal(err)
			}
			defer pNoteDataRows.Close()

			for pNoteDataRows.Next() {
				note := new(Note)
				pNoteDataRows.Scan(&note.ID, &note.Note)
				notes = append(notes, note)
			}
			j.Data = notes
			j.LastPage = int64(pagesCount)
			j.LastPage = int64(pagesCount)
			j.CurrentPage = int64(page)
		}
		w.WriteHeader(http.StatusInternalServerError)
		j.Status = "wrong page parameter"
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		j.Status = "error"
	}

	// send response
	resp, _ := json.Marshal(j)
	w.Write(resp)
}

// UpdateNotesHandler is a handler for note updating
func UpdateNotesHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id := vars["id"]
	token := vars["token"]
	note := vars["note"]

	db := app.Db()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	responsejson := updateNotesJSON{}

	var isTokenValid bool
	var uid int

	isTokenValid, uid = IsAuth(token)

	if isTokenValid {
		w.WriteHeader(http.StatusOK)

		n, _ := db.Prepare("UPDATE notes SET note=? WHERE id=? AND uid=? ")
		n.Exec(note, id, uid)

		responsejson.Status = "note updated"

	} else {
		w.WriteHeader(http.StatusInternalServerError)
		responsejson.Status = "token is not valid"
	}
	resp, _ := json.Marshal(responsejson)
	w.Write(resp)
}

// DeleteNoteHandler is a handler func to delete note
func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	token := vars["token"]

	db := app.Db()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	responsejson := updateNotesJSON{}

	var isTokenValid bool
	var uid int

	isTokenValid, uid = IsAuth(token)

	if isTokenValid {
		w.WriteHeader(http.StatusOK)

		n, _ := db.Prepare("DELETE FROM notes WHERE id=? AND uid=? ")
		n.Exec(id, uid)

		responsejson.Status = "note deleted"

	} else {
		w.WriteHeader(http.StatusInternalServerError)
		responsejson.Status = "token is not valid"
	}
	resp, _ := json.Marshal(responsejson)
	w.Write(resp)
}
