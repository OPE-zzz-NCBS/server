package iface

import (
	"github.com/OPENCBS/server/repo"
)

var NewUserRepo func() *repo.UserRepo
var NewClientRepo func() *repo.ClientRepo
var NewPersonRepo func() *repo.PersonRepo
var NewActivityRepo func() *repo.ActivityRepo
var NewBranchRepo func() *repo.BranchRepo
var NewCityRepo func() *repo.CityRepo

