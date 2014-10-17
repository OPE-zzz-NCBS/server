package iface

import (
	"net/http"
)

var GetConnectionString func(r *http.Request) string
