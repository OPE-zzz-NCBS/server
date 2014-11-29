package api

import (
	"net/http"
	"encoding/json"
	"encoding/base64"
	"crypto/rand"
	"log"
	"github.com/OPENCBS/server/app"
	"github.com/OPENCBS/server/model"
	"github.com/OPENCBS/server/repo"
)

func getRandomToken() (string, error) {
	rb := make([]byte, 32)
	_, err := rand.Read(rb)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(rb), nil
}

func AddSession(ctx *app.AppContext, w http.ResponseWriter, r *http.Request) {
	var session model.Session
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&session)
	if err != nil {
		fail(w, err)
		return
	}

	if session.Username == "" || session.Password == "" {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Find user
	repo := repo.NewUserRepo(ctx.DbProvider)
	user, err := repo.GetByUsernameAndPassword(session.Username, session.Password)
	if err != nil {
		fail(w, err)
		return
	}

	if user == nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Create session
	token, err := getRandomToken()
	if err != nil {
		fail(w, err)
		return
	}

	// Cache session and return to the client
	session = model.Session{token, user.Username, ""}
	model.SetSession(token, &session)
	log.Printf("created a session for user \"%s\"", user.Username)

	sendJson(w, session)
}

func DeleteSession(ctx *app.AppContext, w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("X-Access-Token")
	session := model.GetSession(token)
	if session != nil {
		model.DeleteSession(token)
		log.Printf("deleted the session for user \"%s\"", session.Username)
	}
	w.WriteHeader(200)
}
