package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	api "simple-open-api-spec-example/generated"
	"simple-open-api-spec-example/handler"
)

func main() {

	newHandler := handler.NewServer(
		handler.NewServerOptions{},
	)

	r := gin.Default()

	api.RegisterHandlers(r, newHandler)

	// And we serve HTTP until the world ends.

	s := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8080",
	}

	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}
