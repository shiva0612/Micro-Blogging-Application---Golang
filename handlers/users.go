package handlers

import (
	"net/http"
	"shiva/models"
	services "shiva/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service *services.UserService
}

func (u *UserHandler) Signup(c *gin.Context) {
	userSignupRequest := new(models.UserSignupRequest)
	err := c.ShouldBindJSON(userSignupRequest)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid user signup request")
		return
	}

	appErr := u.Service.Signup(userSignupRequest)
	if appErr != nil {
		c.String(appErr.Code, appErr.Message)
		return
	}
	c.String(http.StatusCreated, "user created, please login...")
}
func (u *UserHandler) Login(c *gin.Context) {
	userLoginRequest := new(models.UserLoginRequest)
	err := c.ShouldBindJSON(userLoginRequest)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid user login request")
		return
	}
	token, appErr := u.Service.Login(userLoginRequest)
	if appErr != nil {
		c.JSON(appErr.Code, appErr)
		return
	}

	c.Header("token", token)
	c.String(http.StatusOK, "your login is successful...")
}
func (u *UserHandler) Logout(c *gin.Context) {
	c.Header("token", "")
	c.String(http.StatusOK, "you logged out")
}
func (u *UserHandler) Edit(c *gin.Context) {
	userEditRequest := new(models.UserEditRequest)
	err := c.ShouldBindJSON(userEditRequest)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid user edit request")
		return
	}

	appErr := u.Service.Edit(c.GetString("userID"), userEditRequest)
	if appErr != nil {
		c.JSON(appErr.Code, appErr)
		return
	}
	c.String(http.StatusOK, "edit successful...")
}

func (u *UserHandler) Delete(c *gin.Context) {
	appErr := u.Service.Delete(c.GetString("userID"))
	if appErr != nil {
		c.JSON(appErr.Code, appErr)
		return
	}
	c.String(http.StatusOK, "user deleted...")
}

func (u *UserHandler) ViewUser(c *gin.Context) {
	user, appErr := u.Service.ViewUser(c.GetString("userID"))
	if appErr != nil {
		c.JSON(appErr.Code, appErr)
		return
	}
	c.JSON(http.StatusOK, user)

}
func (u *UserHandler) GetUsers(c *gin.Context) {
	username := c.Query("username")
	users, appErr := u.Service.GetUsers(username)
	if appErr != nil {
		c.JSON(appErr.Code, appErr)
		return
	}
	c.JSON(http.StatusOK, users)
}
func (u *UserHandler) GetUser(c *gin.Context) {
	username := c.Query("username")
	user, appErr := u.Service.GetUser(username)
	if appErr != nil {
		c.JSON(appErr.Code, appErr)
		return
	}
	c.JSON(http.StatusOK, user)

}
func (u *UserHandler) Follow(c *gin.Context) {
	toFollow := c.Query("toFollow")
	userID := c.GetString("userID")
	appErr := u.Service.Follow(toFollow, userID)
	if appErr != nil {
		c.JSON(appErr.Code, appErr)
		return
	}
	c.String(http.StatusOK, "ok...followed")
}
func (u *UserHandler) HomeFeed(c *gin.Context) {
	posts, appErr := u.Service.HomeFeed(c.GetString("userID"))
	if appErr != nil {
		c.JSON(appErr.Code, appErr)
		return
	}
	c.JSON(http.StatusOK, posts)
}
