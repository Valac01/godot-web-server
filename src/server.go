package main

import (
	"fmt"
	"net/http"
)

type application struct {
	config config
}

func (app *application) serve() {
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", app.config.port),
		Handler: app.handler(),
	}

	fmt.Printf("Server starting on: %s\n", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
