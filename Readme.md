# micro-blogging

- user can do the following:
    - Signup - Login - Edit Profile - Delete profile
    - Create - Edit - Delete → posts
    - **Like** - **Comment** on Posts
    - **search** posts based on : title, body + tags
    - search posts of a specific user and also using above filters
    - search other users using username
    - get other user details using username
    - **follow** other users
    - check **Homefeed** which contains(posts of the user he/she is following)

```bash
clone the repo
docker-compose up #for running mongo in docker
go run main.go 

folder structure :

├── 1.mongodb
│   └── docker-compose.yml
├── Readme.md
├── config
│   ├── config.go
│   └── config.json
├── db
│   ├── mongoRepo.go
│   ├── posts.go
│   └── users.go
├── errs
│   └── appErr.go
├── go.mod
├── go.sum
├── handlers
│   ├── middleware.go
│   ├── posts.go
│   └── users.go
├── main.go
├── models
│   ├── post.go
│   ├── post_dtos.go
│   ├── user.go
│   ├── user_claims.go
│   └── user_dtos.go
├── services
│   ├── posts.go
│   └── users.go
└── utils
    └── hashpsw.go
```

```go
[GIN-debug] POST   /users/signup             --> shiva/handlers.(*UserHandler).Signup-fm (3 handlers)
[GIN-debug] POST   /users/login              --> shiva/handlers.(*UserHandler).Login-fm (3 handlers)
[GIN-debug] POST   /users/search             --> shiva/handlers.(*UserHandler).GetUsers-fm (4 handlers)
[GIN-debug] POST   /users/getuser            --> shiva/handlers.(*UserHandler).GetUser-fm (4 handlers)
[GIN-debug] GET    /users/view               --> shiva/handlers.(*UserHandler).ViewUser-fm (4 handlers)
[GIN-debug] PATCH  /users/edit               --> shiva/handlers.(*UserHandler).Edit-fm (4 handlers)
[GIN-debug] POST   /users/logout             --> shiva/handlers.(*UserHandler).Logout-fm (4 handlers)
[GIN-debug] POST   /users/delete             --> shiva/handlers.(*UserHandler).Delete-fm (4 handlers)
[GIN-debug] PATCH  /users/follow             --> shiva/handlers.(*UserHandler).Follow-fm (4 handlers)
[GIN-debug] GET    /home                     --> shiva/handlers.(*UserHandler).HomeFeed-fm (4 handlers)
[GIN-debug] GET    /posts/self               --> shiva/handlers.(*PostHandler).MyPosts-fm (4 handlers)
[GIN-debug] GET    /posts/search             --> shiva/handlers.(*PostHandler).GetPosts-fm (4 handlers)
[GIN-debug] POST   /posts/create             --> shiva/handlers.(*PostHandler).Create-fm (4 handlers)
[GIN-debug] PATCH  /posts/edit               --> shiva/handlers.(*PostHandler).Edit-fm (4 handlers)
[GIN-debug] POST   /posts/delete             --> shiva/handlers.(*PostHandler).Delete-fm (4 handlers)
[GIN-debug] PATCH  /posts/comment            --> shiva/handlers.(*PostHandler).Comment-fm (4 handlers)
[GIN-debug] PATCH  /posts/like               --> shiva/handlers.(*PostHandler).Like-fm (4 handlers)
main.go:51: starting the server...
```