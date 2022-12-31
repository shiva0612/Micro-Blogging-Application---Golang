package main

import (
	"log"
	"shiva/config"
	"shiva/db"
	h "shiva/handlers"
	s "shiva/services"

	"github.com/gin-gonic/gin"
)

func init() {
	log.SetFlags(log.Lshortfile)
	config.NewConfig("./config/config.json")
}

func app() {
	mr := db.ConnectDB()
	defer mr.DisconnectDB()

	us := &s.UserService{Repo: mr}
	uh := &h.UserHandler{Service: us}

	ps := &s.PostService{Repo: mr}
	ph := &h.PostHandler{Service: ps}

	r := gin.Default()

	r.POST("/users/signup", uh.Signup)
	r.POST("/users/login", uh.Login)

	r.Use(h.Authenticate)
	r.POST("/users", uh.GetUsers)
	r.POST("/users", uh.GetUser) //from queryparam(username)
	r.GET("/users/view", uh.ViewUser)
	r.PATCH("/users/edit", uh.Edit)
	r.POST("/users/logout", uh.Logout)
	r.POST("/users/delete", uh.Delete)
	r.PATCH("/users/follow", uh.Follow)
	r.GET("/home", uh.HomeFeed)

	r.GET("/posts/self", ph.MyPosts)
	r.GET("/posts/search", ph.GetPosts)
	r.POST("/posts/create", ph.Create)
	r.PATCH("/posts/edit", ph.Edit)
	r.POST("/posts/delete", ph.Delete)
	r.PATCH("/posts/comment", ph.Comment)
	r.PATCH("/posts/like", ph.Like)

	log.Println("starting the server...")
	log.Println(r.Run(":" + config.Cfg.App.Port))

}

func main() {
	app()
}
