package main

import (
	"fmt"
	"github.com/OPENCBS/server/iface"
)

func main() {
	bootstrap()
	fmt.Println(iface.GetConnectionString(nil))
}

