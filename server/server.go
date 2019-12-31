package server

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	generated "github.com/rflvicentini/go-gqlgen/api/graph/generated"
	resolvers "github.com/rflvicentini/go-gqlgen/api/graph/resolvers"
)

const defaultPort = "8080"

type Server struct {
	port	string
}

func New() Server {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	s := Server{
		port: port,
	}

	return s
}

func (s Server) Run() {
	c := generated.Config{Resolvers: &resolvers.Resolver{}}
	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(generated.NewExecutableSchema(c)))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", s.port)
	log.Fatal(http.ListenAndServe(":"+s.port, nil))
}
