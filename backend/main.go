package main

import (
	"fmt"

	server "github.com/codescalersinternships/Linktree-Fatma-Ebrahim/pkg"
	"github.com/gin-gonic/gin"
)

// @securityDefinitions.apikey token
// @in header
// @name token

func main() {
	gin.SetMode(gin.ReleaseMode)
	server := server.Linktreeserver()
	fmt.Println("server running on http://localhost:8000")
	server.Run(":8000")

}
