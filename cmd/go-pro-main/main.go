package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/md/go-pro-main/internal/config"
	"github.com/md/go-pro-main/internal/routes"
)

func main() {
	// DB Migration :
	config.DataMigration()

	// Routes
	r := mux.NewRouter()
	routes.Routes(r)

	// Go Server
	fmt.Println("Serving application on port 9090.")
	log.Fatal(http.ListenAndServe(":9090", r))
}
