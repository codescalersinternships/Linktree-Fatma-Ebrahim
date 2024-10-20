package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)
type linktree struct{
	ID string `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	Fullname string `json:"fullName"`
	Bio string `json:"bio"`
	Links []string `json:"links"`
}
var linktrees = []linktree{
    {ID: "1", Username: "username1", Email: "email1@example.com", Password:"password_1",Fullname: "First1 Second1", Bio:"Bio bio1",Links:[]string{"https://www.google.com/","https://www.facebook.com/"}},
    {ID: "2", Username: "username2", Email: "email2@example.com", Password:"password_2",Fullname: "First2 Second2", Bio:"Bio bio2",Links:[]string{"https://www.youtube.com/","https://www.twitter.com/"}},
    {ID: "3", Username: "username3", Email: "email3@example.com", Password:"password_3",Fullname: "First3 Second3", Bio:"Bio bio3",Links:[]string{"https://www.github.com/","https://www.linkedin.com/"}},
}

//gin.context It carries request details, validates and serializes JSON, and more.
func getLinktrees(c *gin.Context) {
	// Context.IndentedJSON to serialize the struct into JSON and add it to the response.
    c.IndentedJSON(http.StatusOK, linktrees)
}

func postLinktree(c *gin.Context) {
    var newLinktree linktree

    // Call BindJSON to bind the received JSON to newLinktree
	if err:=c.BindJSON(&newLinktree); err!=nil{
		log.Fatal(err)
	}
	linktrees=append(linktrees, newLinktree)

    c.IndentedJSON(http.StatusCreated, newLinktree)
}

func getLinktreeByID(c *gin.Context) {
    id:=c.Param("id")

    for _, l := range linktrees{
        if l.ID == id {
            c.IndentedJSON(http.StatusOK, l)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "linktree not found"})
}


func main() {
    router := gin.Default()
	//associate the GET HTTP method and /albums path with a handler function.
    router.GET("/linktrees", getLinktrees)
	router.POST("/linktree",postLinktree)
	router.GET("linktrees/:id",getLinktreeByID)

	//Run function attach the router to an http.Server and start the server.
    router.Run("localhost:8080")
}
