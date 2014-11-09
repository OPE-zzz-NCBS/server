package model

type Session struct {
	Token string `json:"token"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var sessions map[string]*Session = make(map[string]*Session)

func SetSession(key string, session *Session) {
	sessions[key] = session
}

func GetSession(key string) *Session{
	session, _ := sessions[key]
	return session
}

func DeleteSession(key string) {
	delete(sessions, key)
}

