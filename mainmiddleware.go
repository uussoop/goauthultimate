package goauthultimate

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthType int

const (
	UserPassword AuthType = iota
	MailPassword
	PhoneNumberPassword
	Phone
	Mail
	Wallet
)

type IdentifierPasswordAuth struct {
	Username string
	Password string
}

type UtilityFuncs struct {
	GetUser                 func(identifier *string) any
	UserExists              func(identifier *string) bool
	CreateUser              func(identifier *string) any
	CheckIdentifierPassword func(identifier, password *string) bool
	SendCode                func(identifier *string) bool
	ValidateCode            func(code *string) bool
}
type AuthConfig struct {
	Utilities UtilityFuncs
	Router    *gin.Engine
	Authtype  AuthType
	Base      string
}

var isRoutesSet = false

func GoAuthMiddleware(
	authc *AuthConfig,
) gin.HandlerFunc {
	if !isRoutesSet {
		var base string
		strings.Replace(authc.Base, "/", base, -1)
		auth := authc.Router.Group(fmt.Sprintf("/%s", base))

		switch authc.Authtype {
		case UserPassword:
			auth.POST("/register", Register(&authc.Utilities))
			auth.POST("/login", LoginUser(authc.Authtype, &authc.Utilities))
			auth.GET("/refresh", RefreshToken(&authc.Utilities))
			auth.POST("/confirmcode", ConfirmCode(&authc.Utilities))

		case MailPassword:
			auth.POST("/register", Register(&authc.Utilities))
			auth.POST("/login", LoginUser(authc.Authtype, &authc.Utilities))
			auth.GET("/refresh", RefreshToken(&authc.Utilities))
			auth.POST("/confirmcode", ConfirmCode(&authc.Utilities))

		case PhoneNumberPassword:
			auth.POST("/register", Register(&authc.Utilities))
			auth.POST("/login", LoginUser(authc.Authtype, &authc.Utilities))
			auth.GET("/refresh", RefreshToken(&authc.Utilities))
			auth.POST("/confirmcode", ConfirmCode(&authc.Utilities))

		case Phone:
			auth.POST("/login", LoginUser(authc.Authtype, &authc.Utilities))
			auth.GET("/refresh", RefreshToken(&authc.Utilities))
			auth.POST("/confirmcode", ConfirmCode(&authc.Utilities))

		case Mail:
			auth.POST("/login", LoginUser(authc.Authtype, &authc.Utilities))
			auth.GET("/refresh", RefreshToken(&authc.Utilities))
			auth.POST("/confirmcode", ConfirmCode(&authc.Utilities))

		case Wallet:
			// auth.POST("/login", LoginUser(authc.Authtype, &authc.Utilities))
			// auth.GET("/refresh", RefreshToken(&authc.Utilities))
			// auth.POST("/confirmcode", ConfirmCode(&authc.Utilities))

		default:
			auth.POST("/register", Register(&authc.Utilities))
			auth.POST("/login", LoginUser(authc.Authtype, &authc.Utilities))
			auth.GET("/refresh", RefreshToken(&authc.Utilities))
			auth.POST("/confirmcode", ConfirmCode(&authc.Utilities))

		}
		isRoutesSet = true
	}

	switch authc.Authtype {
	case UserPassword:
		return JWTMiddleware(&authc.Utilities)

	case MailPassword:
		return JWTMiddleware(&authc.Utilities)

	case PhoneNumberPassword:
		return JWTMiddleware(&authc.Utilities)

	case Phone:
		return JWTMiddleware(&authc.Utilities)

	case Mail:
		return JWTMiddleware(&authc.Utilities)

	case Wallet:
		// return JWTMiddleware(&authc.Utilities)
		return func(c *gin.Context) {
			c.JSON(400, "this method has not been implemented yet")
		}

	default:
		return JWTMiddleware(&authc.Utilities)

	}

}
