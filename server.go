package main

import (
	"fmt"
	"github.com/OPENCBS/server/iface"
)

func main() {
	bootstrap()

	db, err := iface.GetDb(nil)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("OK")
}

