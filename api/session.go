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
	var js []byte
	var token string
	var session model.Session

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&session)
	if err != nil {
		goto Error
	}

	if session.Username == "" || session.Password == "" {
		goto Unauthorized
	}

	db, err = iface.GetDb(r)
	if err != nil {
		goto Error
	}

	// Find user
	repo = iface.NewUserRepo()
	user, err = repo.FindByUsernameAndPassword(db, session.Username, session.Password)
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
	session = model.Session{token, user.Username, ""}
	model.SetSession(token, &session)
	log.Printf("created a session for user \"%s\"", user.Username)

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
	log.Println(err.Error())
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func DeleteSession(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("X-Access-Token")
	session := model.GetSession(token)
	if session != nil {
		model.DeleteSession(token)
		log.Printf("deleted the session for user \"%s\"", session.Username)
	}
	w.WriteHeader(200)
}

