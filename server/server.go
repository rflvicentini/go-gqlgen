package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	generated "github.com/rflvicentini/go-gqlgen/api/graph/generated"
	resolvers "github.com/rflvicentini/go-gqlgen/api/graph/resolvers"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	c := generated.Config{Resolvers: &resolvers.Resolver{}}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(generated.NewExecutableSchema(c)))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
