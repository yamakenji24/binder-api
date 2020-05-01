package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.POST("/login", func(c *gin.Context) {
		fmt.Println("called login cai")
	})
}
