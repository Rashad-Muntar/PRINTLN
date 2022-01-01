package http

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Rashad-Muntar/println/graph"
	"github.com/gin-gonic/gin"
)

const defaultPort = ":8080"

func GraphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	h.AddTransport(&transport.Websocket{})
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}

}

func PlaygroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
