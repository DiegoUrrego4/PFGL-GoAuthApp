package main

import (
	"fmt"
	"github.com/goschool/crud/db"
	"github.com/goschool/crud/routes"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Program starting...")

	r := routes.SetupRoutes()

	_, err := db.Open()
	if err != nil {
		log.Println("There was an error in database open")
	}

	fmt.Println("Program running at port :8081")
	http.ListenAndServe(":8081", r)

	//Start the app here:
}
