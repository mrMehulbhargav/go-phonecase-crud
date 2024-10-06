
package main

import (
	"log"
	"net/http"
	"go-phonecase-crud/routes"  // Import path matches the module name
)

func main() {
	router := routes.RegisterRoutes()

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
