package main

import (
	"fmt"
	"github.com/OPENCBS/server/iface"
)

func main() {
	bootstrap()

	conn, err := iface.GetConnection(nil)
	if err != nil {
		panic(err)
	}

	err = conn.Ping()
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Println("OK")
}

