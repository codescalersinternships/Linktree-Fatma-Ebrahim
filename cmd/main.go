package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

type link struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Name   string             `bson:"name"`
	Link   string             `bson:"link"`
	Visits int                `bson:"visits"`
}
type linktree struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
	Fullname string             `bson:"fullname"`
	Bio      string             `bson:"bio"`
	Links    []link             `bson:"links"`
}

func addLinktree_db(linktree linktree) (*mongo.InsertOneResult, error) {
	linktree.ID = primitive.NewObjectID()
	result, err := collection.InsertOne(ctx, linktree)
	return result, err
}

func getLinktrees_db() ([]linktree, error) {
	var linktrees []linktree
	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return linktrees, err
	}
	for cur.Next(ctx) {
		var l linktree
		if err := cur.Decode(&l); err != nil {
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

func addLink_db(id string, link link) (linktree, error) {
	var linktree linktree
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return linktree, err
	}
	link.ID = primitive.NewObjectID()
	update := bson.M{
		"$push": bson.M{
			"links": link,
		},
	}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err = collection.FindOneAndUpdate(ctx, bson.M{"_id": ID}, update, opts).Decode(&linktree)
	return linktree, err
}
func addBio_db(id string, bio string) (linktree, error) {
	var linktree linktree
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return linktree, err
	}
	update := bson.M{"$set": bson.M{"bio": bio}}
	err = collection.FindOneAndUpdate(ctx, bson.M{"_id": ID}, update).Decode(&linktree)
	return linktree, err
}

func addFullname_db(id string, fullname string) (linktree, error) {
	var linktree linktree
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return linktree, err
	}
	update := bson.M{"$set": bson.M{"fullname": fullname}}
	err = collection.FindOneAndUpdate(ctx, bson.M{"_id": ID}, update).Decode(&linktree)
	return linktree, err
}

func updateLinkByName_db(id string, link link) (linktree, error) {
	var linktree linktree
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return linktree, err
	}
	update := bson.M{
		"$set": bson.M{
			"links.$[link].name":   link.Name,
			"links.$[link].link":   link.Link,
			"links.$[link].visits": link.Visits,
		},
	}
	arrayFilters := options.ArrayFilters{
		Filters: []interface{}{
			bson.M{"link.name": link.Name},
		},
	}
	opts := options.FindOneAndUpdate().SetArrayFilters(arrayFilters).SetReturnDocument(options.After)
	err = collection.FindOneAndUpdate(ctx, bson.M{"_id": ID}, update, opts).Decode(&linktree)
	return linktree, err
}

func getLinktrees(c *gin.Context) {
	linktrees, err := getLinktrees_db()
	if err != nil {
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
	id := c.Param("id")
	fmt.Println(id)
	linktree, err := getLinktreebyID_db(id)
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%v", err)})
		return
	}
	c.IndentedJSON(http.StatusOK, linktree)

}

func addLinktoTree(c *gin.Context) {
	id := c.Param("id")
	var link link

	if err := c.BindJSON(&link); err != nil {
		log.Fatal(err)
	}
	result, err := addLink_db(id, link)
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%v", err)})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

func addBiotoTree(c *gin.Context) {
	id := c.Param("id")
	var bio string

	if err := c.BindJSON(&bio); err != nil {
		log.Fatal(err)
	}
	result, err := addBio_db(id, bio)
	if err != nil {
		fmt.Println(err)
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
	result, err := addFullname_db(id, fullname)
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%v", err)})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

func updateLinkByName(c *gin.Context) {
	id := c.Param("id")
	var link link

	if err := c.BindJSON(&link); err != nil {
		log.Fatal(err)
	}

	result, err := updateLinkByName_db(id, link)
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%v", err)})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

func main() {

	router := gin.Default()
	router.GET("/linktrees", getLinktrees)
	router.POST("/linktree", addLinktree)
	router.GET("linktrees/:id", getLinktreeByID)
	router.POST("/linktree/:id/addlink", addLinktoTree)
	router.POST("/linktree/:id/addbio", addBiotoTree)
	router.POST("/linktree/:id/addfullname", addFullnametoTree)
	router.POST("/linktree/:id/updatelink", updateLinkByName)

	router.Run("localhost:8080")
}
