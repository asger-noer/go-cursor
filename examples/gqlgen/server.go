package main

import (
	"log"
	"net/http"
	"os"

	"examples/gqlgen/graph"
	"examples/gqlgen/graph/model"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

//go:generate go run github.com/99designs/gqlgen generate

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		Data: []model.Todo{
			{ID: "1", Text: "Do the laundry", Done: false},
			{ID: "2", Text: "Do the dishes", Done: true},
			{ID: "3", Text: "Do the cleaning", Done: false},
			{ID: "4", Text: "Do the shopping", Done: true},
			{ID: "5", Text: "Do the cooking", Done: false},
			{ID: "6", Text: "Do the gardening", Done: true},
			{ID: "7", Text: "Do the homework", Done: false},
		},
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
