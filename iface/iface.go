package iface

import (
	"net/http"
	"database/sql"
)

var GetDb func(r *http.Request) (*sql.DB, error)

