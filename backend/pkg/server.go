package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codescalersinternships/Linktree-Fatma-Ebrahim/database"
	_ "github.com/codescalersinternships/Linktree-Fatma-Ebrahim/docs"
	"github.com/codescalersinternships/Linktree-Fatma-Ebrahim/models"
	"github.com/codescalersinternships/Linktree-Fatma-Ebrahim/token"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var user_id primitive.ObjectID

// @Summary signup a new user
// @Accept json
// @Produce json
// @Param user body models.User true "Signup a new user"
// @Success 201 {object} models.User
// @Failure 400 {object} models.Error
// @Router /linktree/signup [post]
func signup(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}

	tokenMaker := token.NewJWTMaker(os.Getenv("SECRET_KEY"))
	user.ID = primitive.NewObjectID()
	user_id = user.ID
	tokenStr, _, err := tokenMaker.CreateToken(user.ID, user.Username, user.Email, user.Password, time.Hour*24)
	user.Token = tokenStr
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	err = database.AddUser(&user)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, user)

}

// @Summary login a user
// @Accept json
// @Produce json
// @Param user body models.User true "login a user"
// @Success 200 {object} models.User
// @Failure 400 {object} models.Error
// @Router /linktree/login [post]
func login(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, models.Error{Message: err.Error()})
	}

	checkeduser, err := database.CheckUser(user)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, checkeduser)
	user_id = checkeduser.ID
	fmt.Println(user_id)

}

// @Summary add a new linktree to a user
// @Accept json
// @Produce json
// @Param linktree body models.Linktree true "add a new linktree"
// @Success 201 {object} models.Linktree
// @Failure 400 {object} models.Error
// @Security token
// @Router /linktree/ [post]
func addLinktree(c *gin.Context) {
	var newLinktree models.Linktree
	if err := c.BindJSON(&newLinktree); err != nil {
		c.IndentedJSON(http.StatusBadRequest, models.Error{Message: err.Error()})
	}
	for i := range newLinktree.Links {
		newLinktree.Links[i].ID = primitive.NewObjectID()
	}

	tokenMaker := token.NewJWTMaker(os.Getenv("SECRET_KEY"))
	claims, err := tokenMaker.VerifyToken(c.Request.Header.Get("token"))
	user, err := database.AddLinktree(&newLinktree, claims.ID)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, user)
}

// @Summary get a new linktree by ID
// @Accept json
// @Produce json
// @Param   id path string true "linktree id"
// @Success 200 {object} models.Linktree
// @Failure 404 {object} models.Error
// @Security token
// @Router /linktree/{id} [get]
func getLinktreeByID(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	linktree, err := database.GetLinktreebyID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, models.Error{Message: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, linktree)

}

// @Summary add link to linktree
// @Accept json
// @Produce json
// @Param   id path string true "linktree id"
// @Param  link body models.Link true "link"
// @Success 200 {object} models.Linktree
// @Failure 404 {object} models.Error
// @Security token
// @Router /linktree/{id}/addlink [post]
func addLinktoTree(c *gin.Context) {
	id := c.Param("id")
	var link models.Link

	if err := c.BindJSON(&link); err != nil {
		c.IndentedJSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}

	result, err := database.AddLink(id, link)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, models.Error{Message: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

// @Summary add bio to linktree
// @Accept json
// @Produce json
// @Param   id path string true "linktree id"
// @Param  bio body string true "bio"
// @Success 200 {object} models.Linktree
// @Failure 404 {object} models.Error
// @Security token
// @Router /linktree/{id}/addbio [post]
func addBiotoTree(c *gin.Context) {
	id := c.Param("id")
	var bio string

	if err := c.BindJSON(&bio); err != nil {
		c.IndentedJSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	result, err := database.AddBio(id, bio)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, models.Error{Message: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

// @Summary add fullname to linktree
// @Accept json
// @Produce json
// @Param   id path string true "linktree id"
// @Param  fullname body string true "fullname"
// @Success 200 {object} models.Linktree
// @Failure 404 {object} models.Error
// @Security token
// @Router /linktree/{id}/addfullname [post]
func addFullnametoTree(c *gin.Context) {
	id := c.Param("id")
	var fullname string

	if err := c.BindJSON(&fullname); err != nil {
		log.Fatal(err)
	}
	result, err := database.AddFullname(id, fullname)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, models.Error{Message: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

// @Summary update link in linktree
// @Accept json
// @Produce json
// @Param   id path string true "linktree id"
// @Param  link body models.Link true "link"
// @Success 200 {object} models.Linktree
// @Failure 404 {object} models.Error
// @Security token
// @Router /linktree/{id}/updatelink [put]
func updateLinkByID(c *gin.Context) {
	id := c.Param("id")
	var link models.Link

	if err := c.BindJSON(&link); err != nil {
		c.IndentedJSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}

	result, err := database.UpdateLinkByID(id, link)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, models.Error{Message: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

// @Summary delete link in linktree
// @Accept json
// @Produce json
// @Param id path string true "linktree id"
// @Param  link body models.Link true "link"
// @Success 200 {object} models.Linktree
// @Failure 404 {object} models.Error
// @Security token
// @Router /linktree/{id}/deletelink [delete]
func deleteLinkByID(c *gin.Context) {

	id := c.Param("id")
	var link models.Link

	if err := c.BindJSON(&link); err != nil {
		c.IndentedJSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}

	result, err := database.DeleteLinkByID(id, link)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
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

func Linktreeserver() *gin.Engine {

	router := gin.Default()
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	router.POST("/linktree/signup", signup)
	router.POST("/linktree/login", login)

	authorized := router.Group("/")
	authorized.Use(authentication)
	{
		authorized.POST("/linktree", addLinktree)
		authorized.GET("/linktree/:id", getLinktreeByID)
		authorized.POST("/linktree/:id/addlink", addLinktoTree)
		authorized.POST("/linktree/:id/addbio", addBiotoTree)
		authorized.POST("/linktree/:id/addfullname", addFullnametoTree)
		authorized.PUT("/linktree/:id/updatelink", updateLinkByID)
		authorized.DELETE("/linktree/:id/deletelink", deleteLinkByID)
	}

	return router
}
