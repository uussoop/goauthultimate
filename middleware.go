package goauthultimate

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func JWTMiddleware(u *UtilityFuncs) gin.HandlerFunc {
	return func(c *gin.Context) {
		//check if bearer token available if not return error
		token := extractBearerToken(c)

		if token == "" {
			c.JSON(
				http.StatusOK,
				ErrorResponseMessage{
					Status: ERROR_AUTH_TOKEN,
					Error:  "token not found",
				},
			)
			c.Abort()
			return
		} else {

			username, usernameerr := UsernameFromToken(&token)
			if usernameerr != nil {
				c.JSON(http.StatusOK, ErrorResponseMessage{Status: ERROR_AUTH_TOKEN, Error: "token is curropt"})
				c.Abort()
				return
			}
			if ok := u.UserExists(&username); !ok {
				c.JSON(http.StatusOK, ErrorResponseMessage{Status: ERROR_EXIST_USER, Error: "user doesnt exist"})
				c.Abort()
				return
			}

			valid, err := ValidateToken(token)
			if !valid {
				c.JSON(http.StatusOK, ErrorResponseMessage{Status: ERROR_AUTH_TOKEN, Error: "token not valid"})
				c.Abort()
				return
			}
			if err != nil {
				c.JSON(http.StatusOK, ErrorResponseMessage{Status: ERROR_AUTH_TOKEN, Error: "token check ended with error"})
				c.Abort()
				return
			}
			fmt.Println(username)
			user := u.GetUser(&username)
			if user == nil {

				fmt.Println(err)
				c.JSON(http.StatusOK, ErrorResponseMessage{Status: ERROR_AUTH_TOKEN, Error: "token check ended with error"})
				c.Abort()
				return
			}
			c.Set("user", user)

			c.Next()
		}

	}
}

func JwtSsoMiddleware(u *UtilityFuncs) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := extractBearerToken(c)
		// var token string
		// token = c.GetHeader("Authorization")
		if token == "" {

			c.JSON(
				http.StatusOK,
				ErrorResponseMessage{
					Status: ERROR_AUTH_TOKEN,
					Error:  "token not found",
				},
			)
			c.Abort()
			return

		} else {

			valid, username, err := ValidateTokenSso(token)
			if err != nil || !valid {
				c.JSON(
					http.StatusOK,
					ErrorResponseMessage{
						Status: ERROR_AUTH_TOKEN,
						Error:  "token not valid",
					},
				)
				c.Abort()
				return
			}
			user := u.GetUser(&username)
			if user != nil {

				user = u.CreateUser(&username)
				if err != nil {
					logrus.Warn(err)
					c.JSON(http.StatusOK, ErrorResponseMessage{Status: ERROR_USER_OPERATIONS_FAILED, Error: "user creation failed"})
					c.Abort()
					return
				}
			} else if err != nil {
				logrus.Warn(err)
				c.JSON(http.StatusOK, ErrorResponseMessage{Status: ERROR_AUTH_TOKEN, Error: "token check ended with error"})
				c.Abort()
				return
			}

			c.Set("user", user)

		}

	}
}
