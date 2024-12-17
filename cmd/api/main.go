package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Nutchanon28/file-sharing-system/config"
	"github.com/Nutchanon28/file-sharing-system/graph"
	"github.com/Nutchanon28/file-sharing-system/graph/model"
	"github.com/Nutchanon28/file-sharing-system/internal/middlewares"
	"github.com/Nutchanon28/file-sharing-system/pkg/database/minio"
	minioUseCase "github.com/Nutchanon28/file-sharing-system/pkg/database/minio/usecase"
)

const defaultPort = "4000"

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	minioClient, err := minio.NewMinioClient(config)
	if err != nil {
		fmt.Printf("Failed to connect minio: %v", err)
	}

	minioUseCase := minioUseCase.NewMinioUseCase(
		config,
		minioClient,
		// logger,
	)

	// ... our fake DB
	sharedFiles := []*model.SharedFile{}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(
		graph.Config{
			Resolvers: &graph.Resolver{
				SharedFilesList: sharedFiles,
				MinioUseCase:    minioUseCase,
				MinioClient:     minioClient,
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
