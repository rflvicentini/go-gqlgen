package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/handler"
	generated "github.com/rflvicentini/go-gqlgen/api/graph/generated"
	resolvers "github.com/rflvicentini/go-gqlgen/api/graph/resolvers"
	repo_user "github.com/rflvicentini/go-gqlgen/pkg/repo/user"
	repo_meetup "github.com/rflvicentini/go-gqlgen/pkg/repo/meetup"
	"github.com/rflvicentini/go-gqlgen/pkg/net"
)

const defaultPort = "8080"

type Server struct {
	UserRepo	repo_user.UserRepo
	MeetupRepo	repo_meetup.MeetupRepo
	Port			string
}

func New() Server {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	userAPIHTTPClient := GetDefaultHTTPClient("User API")
	UserRepo := repo_user.UserRepo{
		HTTPClient: &userAPIHTTPClient,
	}

	meetupAPIHTTPClient := GetDefaultHTTPClient("Repo API")
	MeetupRepo := repo_meetup.MeetupRepo{
		HTTPClient: &meetupAPIHTTPClient,
	}

	s := Server{
		UserRepo: UserRepo,
		MeetupRepo: MeetupRepo,
		Port: port,
	}

	return s
}

func (s Server) Run() {
	c := generated.Config{Resolvers: &resolvers.Resolver{
		UserRepo: &s.UserRepo,
		MeetupRepo: &s.MeetupRepo,
	}}
	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(generated.NewExecutableSchema(c)))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", s.Port)
	log.Fatal(http.ListenAndServe(":"+s.Port, nil))
}

func GetDefaultHTTPClient(circuitName string) http.Client {
	return net.Client(net.WithClientTimeout(time.Second*5))
}