package services

import (
	"shiva/db"
	"shiva/errs"
	"shiva/models"
)

type PostService struct {
	Repo *db.MongoRepo
}

func (p *PostService) GetPosts(username string, postSearchRequest *models.PostSearchRequest) (*[]models.Post, *errs.AppErr) {
	posts, err := p.Repo.GetPosts(username, postSearchRequest)
	if err != nil {
		return nil, errs.NewTechErr()
	}
	return posts, nil
}

func (p *PostService) MyPosts(userID string) (*[]models.Post, *errs.AppErr) {
	posts, err := p.Repo.MyPosts(userID)
	if err != nil {
		return nil, errs.NewTechErr()
	}
	return posts, nil
}
func (p *PostService) Create(userID string, postCreateRequest *models.PostCreateRequest) *errs.AppErr {
	post := new(models.Post)
	post.FromPostCreateRequest(*postCreateRequest)
	post.FromUser = userID
	err := p.Repo.CreatePost(post)
	if err != nil {
		return errs.NewTechErr()
	}
	return nil
}

func (p *PostService) Edit(userID string, postEditRequest *models.PostEditRequest) *errs.AppErr {

	err := p.Repo.EditPost(postEditRequest)
	if err != nil {
		return errs.NewTechErr()
	}
	return nil
}

func (p *PostService) Delete(userID, postID string) *errs.AppErr {

	err := p.Repo.DeletePost(userID, postID)
	if err != nil {
		return errs.NewTechErr()
	}
	return nil
}

func (p *PostService) Comment(userID string, postCommentRequest *models.PostCommentRequest) *errs.AppErr {

	comment := new(models.Comment)
	comment.FromCommentRequest(postCommentRequest)
	comment.FromUser = userID

	err := p.Repo.Comment(comment)
	if err != nil {
		return errs.NewTechErr()
	}

	return nil
}

func (p *PostService) Like(postID string) {
	p.Repo.Like(postID)
}
