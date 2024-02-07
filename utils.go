package goauthultimate

import (
	"crypto/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func RandomBytes(n int) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

	return b
}
func extractBearerToken(c *gin.Context) string {
	var token string
	token = c.GetHeader("Authorization")
	if token == "" {
		token = c.GetHeader("authorization")
		if token == "" {
			return ""
		}
		token = strings.Replace(token, "Bearer ", "", 1)
		return token
	}
	token = strings.Replace(token, "Bearer ", "", 1)
	return token
}

func checkIdentifierPassword(
	username, password *string,
	c func(identifier *string, password *string) bool,
	createuser func(identifier *string) any,
	gc *gin.Context,
) bool {
	if c != nil {

		isExist := c(username, password)
		if !isExist {
			user := createuser(username)
			if user == nil {

				gc.JSON(
					http.StatusOK,
					ErrorResponseMessage{
						Status: ERROR_AUTH,
						Error:  "couldnt make user",
					},
				)
				gc.Abort()
				return false
			}
			return true
		} else {
			return true
		}

	} else {
		gc.JSON(
			http.StatusOK,
			ErrorResponseMessage{
				Status: ERROR_AUTH,
				Error:  "no check handler provided",
			},
		)
		gc.Abort()

		return false

	}
}

func sendCodeNoPassword(username *string, s func(identifier *string) bool, gc *gin.Context) bool {
	if !s(username) {

		gc.JSON(
			http.StatusOK,
			ErrorResponseMessage{
				Status: ERROR_SENDING_MAIL_FAILED,
				Error:  "error sending the code.",
			},
		)

		return false

	}
	return true
}

func genTokenRefresh(username string, gc *gin.Context) (*string, *string, error) {
	token, err := CreateToken(username, time.Duration(24*time.Hour))
	if err != nil {
		gc.JSON(
			http.StatusOK,
			ErrorResponseMessage{
				Status: ERROR,
				Error:  "Something went wrong, please try again later.",
			},
		)

		return nil, nil, err

	}

	refreshtoken, err := CreateRefreshToken(username)

	if err != nil {
		gc.JSON(
			http.StatusOK,
			ErrorResponseMessage{
				Status: ERROR,
				Error:  "Something went wrong, please try again later.",
			},
		)

		return nil, nil, err
	}
	return &token, &refreshtoken, nil
}
