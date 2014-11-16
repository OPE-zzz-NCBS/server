package iface

import (
	"github.com/OPENCBS/server/repo"
)

var NewUserRepo func() *repo.UserRepo
var NewClientRepo func() *repo.ClientRepo

