package main

import (
	"fmt"

	server "github.com/codescalersinternships/Linktree-Fatma-Ebrahim/pkg"
	"github.com/gin-gonic/gin"
)

func main(){
	gin.SetMode(gin.ReleaseMode)
	server:=server.Linktreeserver()
	fmt.Println("server running on http://localhost:8080")
	server.Run(":8080")


}