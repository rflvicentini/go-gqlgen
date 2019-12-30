gqlgen:
	@rm -rf ./api/graph/generated/resolver.go
	@go run github.com/99designs/gqlgen -v

run:
	@go run ./server/server.go