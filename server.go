package main

import (
	"Tiny_Vote/db"
	"Tiny_Vote/graph"
	"Tiny_Vote/utils"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"log"
	"net/http"
)

const defaultPort = ":8080"
const defaultTitle = "GraphQL playground"
const defaultEndpoint = "/query"

func main() {
	db.InitDB()
	go utils.GenerateTicket()

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler(defaultTitle, defaultEndpoint))
	http.Handle(defaultEndpoint, srv)

	log.Fatal(http.ListenAndServe(defaultPort, nil))
}
