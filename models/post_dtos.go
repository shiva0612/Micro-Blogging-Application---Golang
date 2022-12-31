package models

import (
	"log"

	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostSearchRequest struct {
	Search string   `bson:"search,omitempty" json:"search,omitempty"`
	Tags   []string `bson:"tags,omitempty" json:"tags,omitempty"`
}

type PostCreateRequest struct {
	Title string   `bson:"title,omitempty" json:"title,omitempty"`
	Body  string   `bson:"body,omitempty" json:"body,omitempty"`
	Tags  []string `bson:"tags,omitempty" json:"tags,omitempty"`
}
type PostEditRequest struct {
	PostID string   `bson:"post_id,omitempty" json:"post_id,omitempty"`
	Title  string   `bson:"title,omitempty" json:"title,omitempty"`
	Body   string   `bson:"body,omitempty" json:"body,omitempty"`
	Tags   []string `bson:"tags,omitempty" json:"tags,omitempty"`
}

type PostCommentRequest struct {
	PostID string `bson:"post_id,omitempty" json:"post_id,omitempty"`
	Body   string `bson:"body,omitempty" Json:"body,omitempty"`
}

func (p *Post) FromPostCreateRequest(postCreateRequest PostCreateRequest) error {
	err := copier.Copy(p, postCreateRequest)
	if err != nil {
		log.Println("error while converting PostCreateRequest -> post: ", err.Error())
		return err
	}
	p.ID = primitive.NewObjectID()
	p.PostID = p.ID.Hex()
	p.Comments = make([]Comment, 0)
	if p.Tags == nil {
		p.Tags = make([]string, 0)
	}
	return nil
}

func (c *Comment) FromCommentRequest(postCommentRequest *PostCommentRequest) error {
	err := copier.Copy(c, postCommentRequest)
	if err != nil {
		log.Println("error while converting postCommentRequest -> comment: ", err.Error())
		return err
	}
	c.CommentID = primitive.NewObjectID().Hex()
	return nil
}
