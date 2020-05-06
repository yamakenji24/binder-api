package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/yamakenji24/binder-api/graph"
	"github.com/yamakenji24/binder-api/graph/generated"
)

func graphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL playground", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
func GinContextToContextMiddleware() gin.HandleFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

func main() {
	router := gin.Default()

	router.USE(GinContextToContextMiddleware())
	router.POST("/query", graphqlHandler())
	router.GET("/", playgroundHandler())
	router.Run()
}
