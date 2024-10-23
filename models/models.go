package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Link struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Name   string             `bson:"name"`
	Link   string             `bson:"link"`
	Visits int                `bson:"visits"`
}
type Linktree struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Fullname string             `bson:"fullname"`
	Bio      string             `bson:"bio"`
	Links    []Link             `bson:"links"`
}
type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Username   string             `bson:"username"`
	Email      string             `bson:"email"`
	Password   string             `bson:"password"`
	Token      string             `bson:"token"`
	LinkTreeID primitive.ObjectID `bson:"linktree_id,omitempty"`
}
