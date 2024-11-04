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
// @Router /linktree/{id} [get]
func getLinktreeByID(c *gin.Context) {
	id := c.Param("id")
	linktree, err := database.GetLinktreebyID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, models.Error{Message: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, linktree)

}

// @Summary update tree given its id
// @Accept json
// @Produce json
// @Param   id path string true "linktree id"
// @Success 200 {object} models.Linktree
// @Failure 404 {object} models.Error
// @Security token
// @Router /linktree/{id}/updatetree [put]
func updateTree(c *gin.Context) {
	id := c.Param("id")
	var newLinktree models.Linktree
	if err := c.BindJSON(&newLinktree); err != nil {
		c.IndentedJSON(http.StatusBadRequest, models.Error{Message: err.Error()})
	}
	for i := range newLinktree.Links {
		newLinktree.Links[i].ID = primitive.NewObjectID()
	}

	result, err := database.UpdateTree(id,&newLinktree)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, models.Error{Message: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

// @Summary add link to linktree
// @Accept json
// @Produce json
// @Param   id path string true "linktree id"
// @Param  link body models.Link true "link"
// @Success 200 {object} models.Linktree
// @Failure 404 {object} models.Error
// @Security token
// @Router /linktree/{id}/addlink [put]
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
// @Router /linktree/{id}/addbio [put]
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
// @Router /linktree/{id}/addfullname [put]
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
// @Param  linktree body models.Linktree true "linktree"
// @Success 200 {object} models.Linktree
// @Failure 404 {object} models.Error
// @Security token
// @Router /linktree/{id} [put]
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

// @Summary update visits count for a link
// @Produce json
// @Param id path string true "tree id"
// @Param link_id header string true "link id"
// @Success 200 {object} models.Linktree
// @Failure 404 {object} models.Error
// @Router /linktree/{id}/addvisit [put]
func addVisit(c *gin.Context) {
	tree_id := c.Param("id")
	link_id := c.GetHeader("link_id")
	
	result, err := database.AddVisit(tree_id,link_id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	fmt.Println(result.Links[0].Visits)
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
	url := ginSwagger.URL("http://localhost:8000/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	router.POST("/linktree/signup", signup)
	router.POST("/linktree/login", login)
	router.GET("/linktree/:id", getLinktreeByID)
	router.PUT("/linktree/:id/addvisit", addVisit)
	authorized := router.Group("/")
	authorized.Use(authentication)
	{
		authorized.POST("/linktree", addLinktree)
		authorized.PUT("/linktree/:id", updateTree)
		authorized.PUT("/linktree/:id/addlink", addLinktoTree)
		authorized.PUT("/linktree/:id/addbio", addBiotoTree)
		authorized.PUT("/linktree/:id/addfullname", addFullnametoTree)
		authorized.PUT("/linktree/:id/updatelink", updateLinkByID)
		authorized.DELETE("/linktree/:id/deletelink", deleteLinkByID)
	}


	return router
}
