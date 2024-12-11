package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Nutchanon28/file-sharing-system/graph"
	"github.com/Nutchanon28/file-sharing-system/graph/model"
	"github.com/Nutchanon28/file-sharing-system/internal/middlewares"
)

const defaultPort = "4000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// ... our fake DB
	sharedFiles := []*model.SharedFile{}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(
		graph.Config{
			Resolvers: &graph.Resolver{
				SharedFilesList: sharedFiles,
			},
		},
	))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	// Wrap the default mux with the CORS middleware
	handler := middlewares.Cors(http.DefaultServeMux)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
