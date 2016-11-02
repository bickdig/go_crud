// main.go
// Copyright (C) 2016 Reza Jatnika <rezajatnika@gmail.com>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"log"
	"net/http"

	"github.com/bickdig/go_crud/config"
)

// Global variables in this scope
var (
	loggedRouter = config.NewLoggedRouter()
)

func main() {
	// Start server at 127.0.0.1:8080
	log.Println("Server started at 0.0.0.0:8080")
	log.Fatalln(http.ListenAndServe("0.0.0.0:8080", loggedRouter))
}
