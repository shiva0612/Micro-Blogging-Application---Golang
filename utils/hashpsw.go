package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(psw string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(psw), bcrypt.MinCost)
	if err != nil {
		log.Println("error while generating hash for password: ", err.Error())
		return ""
	}
	return string(bytes)
}
