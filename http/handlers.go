package http

import (
	"github.com/99designs/gqlgen/graphql/handler"
	// "github.com/nats-io/nats.go"

	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	// "github.com/99designs/gqlgen/graphql/handler"
	// "github.com/Rashad-Muntar/println/graph"
	"github.com/Rashad-Muntar/println/graph"
	// "github.com/Rashad-Muntar/println/models"
	// "github.com/Rashad-Muntar/println/graph/generated"
	"github.com/gin-gonic/gin"
)

const defaultPort = ":8080"
// nc *nats.Conn
// func GraphqlHandler() gin.HandlerFunc {
// 	handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
// 		JobCreated: &models.Job,
// 		JobObservers: map[string]chan *models.Job{},
// 	}}))
// 	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
// 	h.AddTransport(&transport.Websocket{})
// 	return func(c *gin.Context) {
// 		h.ServeHTTP(c.Writer, c.Request)
// 	}

// }

func GraphqlHandler() gin.HandlerFunc {
	
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	h.AddTransport(&transport.Websocket{})
	h.AddTransport(transport.GET{})
	h.AddTransport(transport.POST{})
	h.Use(extension.Introspection{})
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}

}


func PlaygroundHandler() gin.HandlerFunc {
	// gin.WrapH(handler.Playground("GraphQL playground", "/api"))
	h := playground.Handler("GraphQL", "/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
