package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/viky1sr/go_cache.git/app/providers"
	"github.com/viky1sr/go_cache.git/app/routes"
)

func main() {
	// define flags for host and port
	host := flag.String("h", "localhost", "server host")
	port := flag.String("p", "8080", "server port")

	// parse the flags
	flag.Parse()

	// create app provider instance
	appProvider := &providers.AppProvider{Host: *host, Port: *port}

	// initialize database connection
	db, err := appProvider.ProvideDB()
	if err != nil {
		log.Fatalf("Failed to connect to DB: %s", err)
	}
	appProvider.Db = db

	// migrate database
	appProvider.Migrate()

	// create router
	router := appProvider.ProvideRouter()
	routes.RegisterRoutes(router, appProvider)

	// start server
	fmt.Printf("Starting server on %s:%s\n", appProvider.Host, appProvider.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", appProvider.Host, appProvider.Port), router))
}
