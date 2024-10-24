package server

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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var user_id primitive.ObjectID

func signup(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%v", err)})
	}

	tokenMaker := token.NewJWTMaker(os.Getenv("SECRET_KEY"))
	user.ID = primitive.NewObjectID()
	user_id = user.ID
	tokenStr, _, err := tokenMaker.CreateToken(user.ID, user.Username, user.Email, user.Password, time.Hour*24)
	user.Token = tokenStr
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%v", err)})
		return
	}
	err = database.AddUser(&user)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%v", err)})
		return
	}

	c.IndentedJSON(http.StatusCreated, user)

}

func authentication(c *gin.Context) {
	clientToken := c.Request.Header.Get("token")
	if clientToken == "" {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("No Authorization header provided")})
		c.Abort()
		return
	}
	tokenMaker := token.NewJWTMaker(os.Getenv("SECRET_KEY"))
	claims, err := tokenMaker.VerifyToken(clientToken)

	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Invalid token")})
		c.Abort()
		return
	}

	c.Set("id", claims.ID)
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
	c.IndentedJSON(http.StatusOK, id)
	user_id = user.ID

}

func addLinktree(c *gin.Context) {
	var newLinktree models.Linktree
	if err := c.BindJSON(&newLinktree); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%v", err)})
	}
	for i := range newLinktree.Links {
		newLinktree.Links[i].ID = primitive.NewObjectID()
	}
	err := database.AddLinktree(&newLinktree)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%v", err)})
		return
	}
	user, err := database.AddTreeIDToUser(user_id, newLinktree.ID)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%v", err)})
		return
	}
	// c.IndentedJSON(http.StatusCreated, newLinktree)
	c.IndentedJSON(http.StatusCreated, user)
}

func getLinktreeByID(c *gin.Context) {
	id := c.Param("id")
	linktree, err := database.GetLinktreebyID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%v", err)})
		return
	}
	c.IndentedJSON(http.StatusOK, linktree)

}

func getLinktrees(c *gin.Context) {
	linktrees, err := database.GetLinktrees()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
		return
	}
	c.IndentedJSON(http.StatusOK, linktrees)
}

func addLinktoTree(c *gin.Context) {
	id := c.Param("id")
	var link models.Link

	if err := c.BindJSON(&link); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%v", err)})
		return
	}
	result, err := database.AddLink(id, link)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%v", err)})
		return
	}
	c.IndentedJSON(http.StatusCreated, result)
}

func addBiotoTree(c *gin.Context) {
	id := c.Param("id")
	var bio string

	if err := c.BindJSON(&bio); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%v", err)})
		return
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

func updateLinkByID(c *gin.Context) {
	id := c.Param("id")
	var link models.Link

	if err := c.BindJSON(&link); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%v", err)})
		return
	}

	result, err := database.UpdateLinkByID(id, link)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%v", err)})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}
func deleteLinkByID(c *gin.Context) {

	id := c.Param("id")
	var link models.Link

	if err := c.BindJSON(&link); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%v", err)})
		return
	}

	result, err := database.DeleteLinkByID(id, link)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%v", err)})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

func Linktreeserver() *gin.Engine {

	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	router.POST("/linktree/signup", signup)
	router.POST("/linktree/login", login)
	router.Use(authentication)
	router.POST("/linktree", addLinktree)
	router.GET("/linktree/:id", getLinktreeByID)
	router.POST("/linktree/:id/addlink", addLinktoTree)
	router.POST("/linktree/:id/addbio", addBiotoTree)
	router.POST("/linktree/:id/addfullname", addFullnametoTree)
	router.POST("/linktree/:id/updatelink", updateLinkByID)
	router.DELETE("/linktree/:id/deletelink", deleteLinkByID)

	router.GET("/linktrees", getLinktrees)

	return router
}
