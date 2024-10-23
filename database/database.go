package database

import (
	"context"
	"fmt"
	"log"

	"github.com/codescalersinternships/Linktree-Fatma-Ebrahim/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var linktree_col *mongo.Collection
var user_col *mongo.Collection
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
	user_col = client.Database("linktree").Collection("users")
	linktree_col = client.Database("linktree").Collection("linktrees")
}

func AddUser(user models.User) (*mongo.InsertOneResult, error) {
	user.ID = primitive.NewObjectID()
	var existingUser models.User
	user_col.FindOne(ctx, bson.M{"username": user.Username}).Decode(&existingUser)

	if existingUser.Username == user.Username {
		return nil, fmt.Errorf("user already exists")
	}
	result, err := user_col.InsertOne(ctx, user)
	return result, err
}

func CheckUser(user models.User) (primitive.ObjectID, error) {
	var existingUser models.User
	user_col.FindOne(ctx, bson.M{"username": user.Username}).Decode(&existingUser)
	if existingUser.Username != user.Username {
		return primitive.NewObjectID(), fmt.Errorf("Unauthorized user")
	}
	if existingUser.Password != user.Password {
		return primitive.NewObjectID(), fmt.Errorf("Wrong password")
	}
	return existingUser.ID, nil
}

func AddLinktree(linktree models.Linktree) (*mongo.InsertOneResult, error) {
	linktree.ID = primitive.NewObjectID()
	result, err := linktree_col.InsertOne(ctx, linktree)
	return result, err
}

func AddTreeIDToUser(id string, linktree_id string) (models.Linktree, error) {
	var linktree models.Linktree
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return linktree, err
	}
	update := bson.M{"$set": bson.M{"linktree_id": linktree_id}}
	err = user_col.FindOneAndUpdate(ctx, bson.M{"_id": ID}, update).Decode(&linktree)
	return linktree, err
}

func GetLinktrees() ([]models.Linktree, error) {
	var linktrees []models.Linktree
	cur, err := linktree_col.Find(ctx, bson.M{})
	if err != nil {
		return linktrees, err
	}
	for cur.Next(ctx) {
		var l models.Linktree
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

func GetLinktreebyID(id string) (models.Linktree, error) {
	var linktree models.Linktree
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return linktree, err
	}

	err = linktree_col.FindOne(ctx, bson.M{"_id": ID}).Decode(&linktree)
	return linktree, err
}

func AddLink(id string, link models.Link) (models.Linktree, error) {
	var linktree models.Linktree
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
	err = linktree_col.FindOneAndUpdate(ctx, bson.M{"_id": ID}, update, opts).Decode(&linktree)
	return linktree, err
}
func AddBio(id string, bio string) (models.Linktree, error) {
	var linktree models.Linktree
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return linktree, err
	}
	update := bson.M{"$set": bson.M{"bio": bio}}
	err = linktree_col.FindOneAndUpdate(ctx, bson.M{"_id": ID}, update).Decode(&linktree)
	return linktree, err
}

func AddFullname(id string, fullname string) (models.Linktree, error) {
	var linktree models.Linktree
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return linktree, err
	}
	update := bson.M{"$set": bson.M{"fullname": fullname}}
	err = linktree_col.FindOneAndUpdate(ctx, bson.M{"_id": ID}, update).Decode(&linktree)
	return linktree, err
}

func UpdateLinkByName(id string, link models.Link) (models.Linktree, error) {
	var linktree models.Linktree
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
	err = linktree_col.FindOneAndUpdate(ctx, bson.M{"_id": ID}, update, opts).Decode(&linktree)
	return linktree, err
}
