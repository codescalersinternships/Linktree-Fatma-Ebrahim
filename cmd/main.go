package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codescalersinternships/Linktree-Fatma-Ebrahim/database"
	"github.com/codescalersinternships/Linktree-Fatma-Ebrahim/models"
	"github.com/codescalersinternships/Linktree-Fatma-Ebrahim/token"
	"github.com/gin-gonic/gin"
)

func signup(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%v", err)})
	}
	tokenMaker := token.NewJWTMaker(os.Getenv("SECRET_KEY"))
	tokenStr,_, err := tokenMaker.CreateToken(user.ID, user.Username, user.Email, user.Password, time.Hour*24)
	user.Token = tokenStr
	result, err := database.AddUser(user)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%v", err)})
		return
	}
	fmt.Println(user.Token)
	c.IndentedJSON(http.StatusCreated, result)

}

func authentication(c *gin.Context) {
	clientToken := c.Request.Header.Get("token")
	if clientToken == "" {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No Authorization header provided")})
		c.Abort()
		return
	}
	tokenMaker := token.NewJWTMaker(os.Getenv("SECRET_KEY"))
	claims, err := tokenMaker.VerifyToken(clientToken)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		c.Abort()
		return
	}
	c.Set("id", claims.ID)
	c.Set("email", claims.Email)
	c.Set("username", claims.Username)
	c.Set("password", claims.Password)
	c.Next()

}

func login(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%v", err)})
	}
	id, err := database.CheckUser(user)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%v", err)})
		return
	}
	c.IndentedJSON(http.StatusCreated, id)

}

func getLinktrees(c *gin.Context) {
	linktrees, err := database.GetLinktrees()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
		return
	}
	c.IndentedJSON(http.StatusOK, linktrees)
}

func addLinktree(c *gin.Context) {
	var newLinktree models.Linktree
	if err := c.BindJSON(&newLinktree); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%v", err)})
	}
	result, err := database.AddLinktree(newLinktree)
	//database.AddTreeIDToUser()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%v", err)})
		return
	}
	c.IndentedJSON(http.StatusCreated, result)
}



func getLinktreeByID(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	linktree, err := database.GetLinktreebyID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%v", err)})
		return
	}
	c.IndentedJSON(http.StatusOK, linktree)

}

func addLinktoTree(c *gin.Context) {
	id := c.Param("id")
	var link models.Link

	if err := c.BindJSON(&link); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%v", err)})
	}
	result, err := database.AddLink(id, link)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%v", err)})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

func addBiotoTree(c *gin.Context) {
	id := c.Param("id")
	var bio string

	if err := c.BindJSON(&bio); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%v", err)})
	}
	result, err := database.AddBio(id, bio)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%v", err)})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

func addFullnametoTree(c *gin.Context) {
	id := c.Param("id")
	var fullname string

	if err := c.BindJSON(&fullname); err != nil {
		log.Fatal(err)
	}
	result, err := database.AddFullname(id, fullname)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%v", err)})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

func updateLinkByName(c *gin.Context) {
	id := c.Param("id")
	var link models.Link

	if err := c.BindJSON(&link); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%v", err)})
	}

	result, err := database.UpdateLinkByName(id, link)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%v", err)})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

func main() {

	router := gin.Default()

	router.POST("/linktree/signup", signup)
	router.POST("/linktree/login", login)
	router.Use(authentication)

	router.GET("/linktrees", getLinktrees)
	router.POST("/linktree", addLinktree)
	router.GET("linktrees/:id", getLinktreeByID)
	router.POST("/linktree/:id/addlink", addLinktoTree)
	router.POST("/linktree/:id/addbio", addBiotoTree)
	router.POST("/linktree/:id/addfullname", addFullnametoTree)
	router.POST("/linktree/:id/updatelink", updateLinkByName)

	router.Run("localhost:8080")
}
