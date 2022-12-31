package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	ID       primitive.ObjectID `bson:"_id" json:"-"`
	PostID   string             `bson:"post_id" json:"post_id"`
	FromUser string             `bson:"from_user" json:"from_user"`
	Title    string             `bson:"title" json:"title"`
	Body     string             `bson:"body" json:"body"`
	Tags     []string           `bson:"tags" json:"tags"`
	Likes    int                `bson:"likes" json:"likes"`
	Comments []Comment          `bson:"comments" json:"comments"`
}

type Comment struct {
	PostID    string `bson:"post_id" json:"post_id"`
	Body      string `bson:"body" Json:"body"`
	CommentID string `bson:"comment_id" json:"comment_id"`
	FromUser  string `bson:"from_user" json:"from_user"`
}
