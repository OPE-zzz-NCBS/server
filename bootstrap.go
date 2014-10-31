package main

import (
	"github.com/OPENCBS/server/iface"
	"github.com/OPENCBS/server/factory"
	"github.com/OPENCBS/server/mssql/repo"
)

func bootstrap() {
	iface.GetDb = factory.GetMsSqlDb
	iface.NewUserRepo = repo.NewMsSqlUserRepo
	iface.NewClientRepo = repo.NewMsSqlClientRepo
}
