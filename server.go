package main

import (
	"github.com/Rashad-Muntar/println/config"
	"github.com/Rashad-Muntar/println/http"
	"github.com/gin-gonic/gin"
	
)

const defaultPort = ":8080"



func init() {
	config.LoadInitializers()
	config.ConnectDB()
}

func main() {
	server := gin.Default()
	server.GET("/", http.PlaygroundHandler())
	server.POST("/query", http.GraphqlHandler())
	server.Run(defaultPort)
}
