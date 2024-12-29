package main

import (
	"fmt"
	"github.com/goschool/crud/routes"
	"net/http"
)

func main() {
	fmt.Println("Program starting...")

	r := routes.SetupRoutes()

	http.ListenAndServe(":8081", r)

	//Start the app here:
}
