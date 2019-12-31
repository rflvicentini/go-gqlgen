package main

import (
	"github.com/rflvicentini/go-gqlgen/server"
)

func main() {
	s := server.New()
	s.Run()
}