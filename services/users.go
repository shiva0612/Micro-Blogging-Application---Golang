package services

import (
	"net/http"
	"shiva/db"
	"shiva/errs"
	"shiva/models"
	"shiva/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo *db.MongoRepo
}

func (u *UserService) Signup(userSignupRequest *models.UserSignupRequest) *errs.AppErr {
	user := new(models.User)
	user.FromSignupRequest(userSignupRequest)
	err := u.Repo.Signup(user)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return errs.NewUserErr(http.StatusBadRequest, "user already exists")
		}
		return errs.NewTechErr()
	}
	return nil
}
func (u *UserService) Login(userLoginRequest *models.UserLoginRequest) (string, *errs.AppErr) {

	user, err := u.Repo.Login(userLoginRequest.UserName)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", errs.NewUserErr(http.StatusBadRequest, "user not found, please signup")
		}
		return "", errs.NewTechErr()
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLoginRequest.Password)); err != nil {
		return "", errs.NewUserErr(http.StatusBadRequest, "incorrect password")
	}

	token, appErr := models.Generate_singed_token(user)
	if appErr != nil {
		return "", appErr
	}

	return token, nil
}
func (u *UserService) Edit(userID string, userEditRequest *models.UserEditRequest) *errs.AppErr {

	if userEditRequest.Password != "" {
		userEditRequest.Password = utils.HashPassword(userEditRequest.Password)
	}

	err := u.Repo.EditUser(userID, userEditRequest)
	if err != nil {
		return errs.NewTechErr()
	}
	return nil
}
func (u *UserService) Delete(userID string) *errs.AppErr {

	err := u.Repo.DeleteUser(userID)
	if err != nil {
		return errs.NewTechErr()
	}
	return nil
}
func (u *UserService) ViewUser(userID string) (*models.User, *errs.AppErr) {
	user, err := u.Repo.ViewUser(userID)
	if err != nil {
		return nil, errs.NewTechErr()
	}
	return user, nil
}
func (u *UserService) GetUsers(username string) (*[]models.User, *errs.AppErr) {

	users, err := u.Repo.GetUsers(username)
	if err != nil {
		return nil, errs.NewTechErr()
	}
	return users, nil
}
func (u *UserService) GetUser(username string) (*models.User, *errs.AppErr) {
	user, err := u.Repo.GetUser(username)
	if err != nil {
		if err.Error() == "user not found" {
			return nil, errs.NewUserNotFoundErr()
		}
		return nil, errs.NewTechErr()
	}
	return user, nil
}
func (u *UserService) Follow(toFollow, userID string) *errs.AppErr {
	err := u.Repo.Follow(toFollow, userID)
	if err != nil {
		return errs.NewTechErr()
	}
	return nil
}
func (u *UserService) HomeFeed(userID string) (*[]models.Post, *errs.AppErr) {
	posts, err := u.Repo.HomeFeed(userID)
	if err != nil {
		return nil, errs.NewTechErr()
	}
	return posts, nil
}
