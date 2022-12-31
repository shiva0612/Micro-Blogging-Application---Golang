package models

import (
	"log"
	"shiva/utils"

	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserEditRequest struct {
	Email    string `bson:"email,omitempty" json:"email,omitempty" `
	Password string `bson:"password,omitempty" json:"password,omitempty" `
}
type UserSignupRequest struct {
	Email    string `bson:"email" json:"email"`
	UserName string `bson:"user_name" json:"user_name"`
	Password string `bson:"password" json:"password"`
}
type UserLoginRequest struct {
	UserName string `bson:"user_name,omitempty" json:"user_name,omitempty"`
	Password string `bson:"password,omitempty" json:"password,omitempty"`
}

func (u *User) FromSignupRequest(userSignupRequest *UserSignupRequest) error {
	err := copier.Copy(u, userSignupRequest)
	if err != nil {
		log.Println("error while converting userSignupRequest -> user: ", err.Error())
		return err
	}
	u.ID = primitive.NewObjectID()
	u.UserID = u.ID.Hex()
	u.Password = utils.HashPassword(u.Password)
	u.Followers = make([]string, 0)
	u.Following = make([]string, 0)

	return nil
}
