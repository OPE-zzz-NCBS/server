package main

import (
	"github.com/OPENCBS/server/iface"
	"github.com/OPENCBS/server/factory"
)

func bootstrap() {
	iface.GetConnection = factory.GetMsSqlConnection
}
