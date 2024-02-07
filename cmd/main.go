package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/uussoop/goauthultimate"
)

func main() {
	a := goauthultimate.UtilityFuncs{
		GetUser:                 nil,
		UserExists:              nil,
		CreateUser:              nil,
		CheckIdentifierPassword: nil,
		SendCode:                nil,
		ValidateCode:            nil,
	}
	fmt.Println(a)
	r := gin.Default()
	conf := goauthultimate.AuthConfig{
		Utilities: a,
		Base:      "auth",
		Router:    r,
		Authtype:  goauthultimate.Mail,
	}

	r.Use(goauthultimate.GoAuthMiddleware(&conf))

	r.Run()
}
