package iface

import (
	"net/http"
	"database/sql"
)

var GetConnection func(r *http.Request) (*sql.DB, error)

