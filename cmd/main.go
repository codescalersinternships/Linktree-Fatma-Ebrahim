package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection *mongo.Collection
var ctx = context.TODO()

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database("linktree").Collection("linktrees")
}

var (
	outfile, _ = os.Create("logs.log")
	logger     = log.New(outfile, "", 0)
)

type linktree struct {
	Username string   `bson:"username"`
	Email    string   `bson:"email"`
	Password string   `bson:"password"`
	Fullname string   `bson:"fullName"`
	Bio      string   `bson:"bio"`
	Links    []string `bson:"links"`
}


func addLinktree_db(linktree linktree) (*mongo.InsertOneResult, error) {
	result, err := collection.InsertOne(ctx, linktree)	
	return  result ,err
}

func getLinktrees_db() ([]linktree, error) {
    var linktrees []linktree
	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return linktrees, err
	}
	for cur.Next(ctx) {
		var l linktree
		if err:= cur.Decode(&l) ;err!=nil{
			return linktrees, err
		}
		linktrees = append(linktrees, l)
	}
	if err := cur.Err(); err != nil {
		return linktrees, err
	}
	cur.Close(ctx)
    return linktrees, err
}

func getLinktreebyID_db(id string) (linktree, error) {
	var linktree linktree
    ID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return linktree, err
    }
    
    err = collection.FindOne(ctx, bson.M{"_id": ID}).Decode(&linktree)
    return linktree, err
}




func getLinktrees(c *gin.Context) {
	linktrees,err:=getLinktrees_db()
	if err!=nil{
		c.IndentedJSON(http.StatusNotFound, err)
	}
	c.IndentedJSON(http.StatusOK, linktrees)
}

func addLinktree(c *gin.Context) {
	var newLinktree linktree
	if err := c.BindJSON(&newLinktree); err != nil {
		log.Fatal(err)
	}
	result, err := addLinktree_db(newLinktree)
	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusCreated, result)
}

func getLinktreeByID(c *gin.Context) {
	id:=c.Param("id")
	fmt.Println(id)
	linktree,err:=getLinktreebyID_db(id)
	if err!=nil{
		fmt.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%v", err)})
		return
	}
	c.IndentedJSON(http.StatusOK, linktree)

}






func addLink(c *gin.Context) {
	var link string
	if err := c.BindJSON(&link); err != nil {
		log.Fatal(err)
	}
	fmt.Println(link)
	
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "linktree not found"})
}

func deleteLink(c *gin.Context) {
	var link string
	if err := c.BindJSON(&link); err != nil {
		log.Fatal(err)
	}
	fmt.Println(link)

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "linktree not found"})
}

func editLink(c *gin.Context) {

}



func main() {
	
	router := gin.Default()
	//associate the GET HTTP method and /albums path with a handler function.
	router.GET("/linktrees", getLinktrees)
	router.POST("/linktree", addLinktree)
	router.GET("linktrees/:id", getLinktreeByID)
	router.POST("/linktree/:id", addLink)
	router.DELETE("/linktree/:id", deleteLink)

	//Run function attach the router to an http.Server and start the server.
	router.Run("localhost:8080")
}
