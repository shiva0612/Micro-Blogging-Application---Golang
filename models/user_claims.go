package models

import (
	"log"
	"shiva/config"
	"shiva/errs"
	"time"

	"github.com/golang-jwt/jwt"
)

type UserClaims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

func Generate_singed_token(user *User) (string, *errs.AppErr) {

	user_claims := new(UserClaims)
	user_claims.UserID = user.UserID
	user_claims.Email = user.Email
	user_claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: time.Now().Add(config.Cfg.App.Token_time * time.Minute).Unix(),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, user_claims).SignedString([]byte(config.Cfg.App.Jwt_secret))
	if err != nil {
		log.Println("could not create token: err = ", err.Error())
		return "", errs.NewTechErr()
	}

	return token, nil
}

func Get_claims_from_token(token string) (*UserClaims, *errs.AppErr) {
	user_claims := new(UserClaims)
	token_parse, err := jwt.ParseWithClaims(token, user_claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Cfg.App.Jwt_secret), nil
	})

	if err != nil || !token_parse.Valid {
		log.Println("error while parsing claims: err = ", err)
		return nil, errs.NewUnAuthorizedErr()
	}
	return user_claims, nil
}
