package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Rashad-Muntar/println/config"
	"github.com/Rashad-Muntar/println/models"
	// "github.com/Rashad-Muntar/println/http"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Rashad-Muntar/println/graph"
	// "github.com/gin-gonic/gin"
	// "github.com/nats-io/nats.go"
)

const defaultPort = "8080"



func init() {
	config.LoadInitializers()
	config.ConnectDB()
}

func main() {
	// server := gin.Default()
	// nc, err := nats.Connect(nats.DefaultURL)
	// if err != nil {
	// 	panic(err)
	// }
	
	// defer nc.Close()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	srv := 	handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		// CreatedJob: &models.Job{},
		JobObservers: map[string]chan *models.Job{},
	}}))
	srv.AddTransport(&transport.Websocket{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.Use(extension.Introspection{})
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
