package iface

import (
	"net/http"
	"database/sql"
	"github.com/OPENCBS/server/model"
)

// Repositories
type UserRepo interface {
	Find(db *sql.DB) ([]model.User, error)
}

var GetDb func(r *http.Request) (*sql.DB, error)
var NewUserRepo func() UserRepo 

