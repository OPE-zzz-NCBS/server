package api

import (
	"net/http"
	"database/sql"
	"encoding/json"
	"encoding/base64"
	"crypto/rand"
	"log"
	"github.com/OPENCBS/server/iface"
	"github.com/OPENCBS/server/model"
)

func getRandomToken() (string, error) {
	rb := make([]byte, 32)
	_, err := rand.Read(rb)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(rb), nil
}

func AddSession(w http.ResponseWriter, r *http.Request) {
	var repo iface.UserRepo
	var db *sql.DB
	var user *model.User
	var err error
	var js []byte
	var token string
	var session *model.Session

	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")
	if username == "" || password == "" {
		goto Unauthorized
	}

	db, err = iface.GetDb(r)
	if err != nil {
		goto Error
	}

	// Find user
	repo = iface.NewUserRepo()
	user, err = repo.FindByUsernameAndPassword(db, username, password)
	if err != nil {
		goto Error
	}

	if user == nil {
		goto Unauthorized
	}

	// Create session
	token, err = getRandomToken()
	if err != nil {
		goto Error
	}

	// Cache session and return to the client
	session = &model.Session{token, user}
	model.SetSession(token, session)
	log.Printf("created a session for user \"%s\"", user.UserName)

	js, err = json.MarshalIndent(session, "", "  ")
	if err != nil {
		goto Error
	}
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(js)
	return

Unauthorized:
	http.Error(w, "Invalid username or password", http.StatusUnauthorized)
	return

Error:
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func DeleteSession(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("X-Authentication-Token")
	session := model.GetSession(token)
	if session != nil {
		model.DeleteSession(token)
		log.Printf("deleted the session for user \"%s\"", session.User.UserName)
	}
	w.WriteHeader(200)
}

