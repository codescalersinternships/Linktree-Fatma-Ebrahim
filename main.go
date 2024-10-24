package main

import server "github.com/codescalersinternships/Linktree-Fatma-Ebrahim/pkg"

func main(){

	server:=server.Linktreeserver()
	server.Run(":8080")

}