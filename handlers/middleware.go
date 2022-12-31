package handlers

import (
	"shiva/models"

	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	token := c.GetHeader("token")
	claims, appErr := models.Get_claims_from_token(token)
	if appErr != nil {
		c.JSON(appErr.Code, appErr)
		c.Abort()
		return
	}

	//either store in server c.set, or set in cookie by encrpyting it
	c.Set("userID", claims.UserID)
	c.Set("email", claims.Email)
	c.Next()
}
