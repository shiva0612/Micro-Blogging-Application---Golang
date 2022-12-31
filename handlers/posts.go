package handlers

import (
	"net/http"
	"shiva/models"
	"shiva/services"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	Service *services.PostService
}

func (p *PostHandler) GetPosts(c *gin.Context) {
	postSearchRequest := new(models.PostSearchRequest)
	postSearchRequest.Tags = make([]string, 0)
	err := c.ShouldBindJSON(postSearchRequest)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid post search request")
		return
	}

	posts, appErr := p.Service.GetPosts(c.Query("username"), postSearchRequest)
	if appErr != nil {
		c.JSON(appErr.Code, appErr)
		return
	}
	c.JSON(http.StatusOK, posts)
}

func (p *PostHandler) MyPosts(c *gin.Context) {
	posts, appErr := p.Service.MyPosts(c.GetString("userID"))
	if appErr != nil {
		c.JSON(appErr.Code, appErr)
		return
	}
	c.JSON(http.StatusOK, posts)
}
func (p *PostHandler) Create(c *gin.Context) {
	postCreateRequest := new(models.PostCreateRequest)
	err := c.ShouldBindJSON(postCreateRequest)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid post create request")
		return
	}
	appErr := p.Service.Create(c.GetString("userID"), postCreateRequest)
	if appErr != nil {
		c.JSON(appErr.Code, appErr)
		return
	}
	c.String(http.StatusCreated, "post is created...")

}

func (p *PostHandler) Edit(c *gin.Context) {
	postEditRequest := new(models.PostEditRequest)
	err := c.ShouldBindJSON(postEditRequest)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid post edit request")
		return
	}
	appErr := p.Service.Edit(c.GetString("userID"), postEditRequest)
	if appErr != nil {
		c.JSON(appErr.Code, appErr)
		return
	}
	c.String(http.StatusOK, "post is updated...")

}

func (p *PostHandler) Delete(c *gin.Context) {
	postID, present := c.GetQuery("postID")
	if !present {
		c.String(http.StatusBadRequest, "empty postID")
	}
	appErr := p.Service.Delete(c.GetString("userID"), postID)
	if appErr != nil {
		c.JSON(appErr.Code, appErr)
		return
	}
	c.String(http.StatusOK, "post is deleted...")

}

func (p *PostHandler) Comment(c *gin.Context) {
	postCommentRequest := new(models.PostCommentRequest)
	err := c.ShouldBindJSON(postCommentRequest)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid post comment request")
		return
	}

	appErr := p.Service.Comment(c.GetString("userID"), postCommentRequest)
	if appErr != nil {
		c.JSON(appErr.Code, appErr)
		return
	}
	c.String(http.StatusCreated, "comment is added to the post...")
}

func (p *PostHandler) Like(c *gin.Context) {
	p.Service.Like(c.Query("postID"))
}
