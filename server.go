package main

import (
	"context"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yamakenji24/binder-api/controllers"
	"github.com/yamakenji24/binder-api/crypt"
	"github.com/yamakenji24/binder-api/graph/generated"
	"github.com/yamakenji24/binder-api/resolver"
)

func graphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))

	return func(c *gin.Context) {
		r := c.Request
		userId, exist := c.Get("userId")
		if !exist {
			h.ServeHTTP(c.Writer, r)
		} else {
			ctx := context.WithValue(r.Context(), "userId", userId.(int))
			h.ServeHTTP(c.Writer, r.WithContext(ctx))
		}
	}
}
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL playground", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

func CheckJWTHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if len(header) == 0 {
			c.Abort()
		}

		token, err := jwt.Parse(header, func(token *jwt.Token) (interface{}, error) {
			return crypt.NewPublicKey(), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": http.StatusText(http.StatusUnauthorized),
			})
			c.Abort()
		} else {
			claims := token.Claims.(jwt.MapClaims)
			userId := int(claims["sub"].(float64))
			c.Set("userId", userId)
		}
	}
}

func main() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowHeaders = []string{"Authorization", "Content-Type"}

	router.Use(cors.New(config))

	router.POST("/login", func(c *gin.Context) {
		controllers.Login(c)
	})

	router.Use(GinContextToContextMiddleware())
	router.Use(CheckJWTHandler())

	router.POST("/query", graphqlHandler())
	router.GET("/", playgroundHandler())
	router.Run()
}
