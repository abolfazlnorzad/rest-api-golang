package main

import (
	"app/database"
	"app/router"
	"fmt"
)

func main() {
	fmt.Println("start")
	connection, err := database.CreateConnection()
	if err != nil {
		return
	}
	errs := router.CreateRouter(connection)
	if errs != nil {
		return
	}
}
