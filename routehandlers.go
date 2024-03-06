package goauthultimate

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type Auth struct {
	Username string `valid:"Required; MaxSize(50)" json:"email"`
	Password string `valid:"Required; MaxSize(50)" json:"password"`
}

type LoginResponse struct {
	Status       customErrors `json:"status"`
	Message      string       `json:"message"`
	Token        string       `json:"token,omitempty"`
	RefreshToken string       `json:"refresh_token,omitempty"`
}

func LoginUser(a AuthType, u *UtilityFuncs, whitelist *[]string) gin.HandlerFunc {
	return func(c *gin.Context) {

		valid := validation.Validation{}
		var auth Auth
		err := c.BindJSON(&auth)
		if err != nil {
			c.JSON(
				http.StatusOK,
				ErrorResponseMessage{Status: ERROR, Error: "Invalid request."},
			)

			return
		}

		ok, _ := valid.Valid(&auth)

		if !ok {
			c.JSON(
				http.StatusOK,
				ErrorResponseMessage{Status: ERROR_EXIST_USER, Error: "Invalid request."},
			)

			return
		}

		authService := IdentifierPasswordAuth{Username: auth.Username, Password: auth.Password}
		// check for auth type
		for _, v := range *whitelist {

			if authService.Username == v {
				isExist := u.CheckIdentifierPassword(&authService.Username, &authService.Password)
				if !isExist {
					user := u.CreateUser(&authService.Username)
					if user == nil {
						c.JSON(
							http.StatusOK,
							ErrorResponseMessage{
								Status: ERROR_AUTH,
								Error:  "couldnt make user",
							},
						)
						c.Abort()
					}
				}
				token, refresh, err := genTokenRefresh(authService.Username, c)
				if err != nil {
					return
				}
				c.JSON(
					http.StatusOK,
					LoginResponse{
						Status:       SUCCESS,
						Message:      "User logged in successfully.",
						Token:        *token,
						RefreshToken: *refresh,
					},
				)
				c.Abort()
				return
			}
		}
		switch a {
		case UserPassword:
			if !checkIdentifierPassword(
				&authService.Username,
				&authService.Password,
				u.CheckIdentifierPassword,
				u.CreateUser,
				c,
			) {
				return
			}
			token, refresh, err := genTokenRefresh(authService.Username, c)
			if err != nil {
				return
			}
			c.JSON(
				http.StatusOK,
				LoginResponse{
					Status:       SUCCESS,
					Message:      "User logged in successfully.",
					Token:        *token,
					RefreshToken: *refresh,
				},
			)

		case MailPassword:
			if !checkIdentifierPassword(
				&authService.Username,
				&authService.Password,
				u.CheckIdentifierPassword,
				u.CreateUser,
				c,
			) {
				return
			}
			token, refresh, err := genTokenRefresh(authService.Username, c)
			if err != nil {
				return
			}
			c.JSON(
				http.StatusOK,
				LoginResponse{
					Status:       SUCCESS,
					Message:      "User logged in successfully.",
					Token:        *token,
					RefreshToken: *refresh,
				},
			)
		case PhoneNumberPassword:

			if !checkIdentifierPassword(
				&authService.Username,
				&authService.Password,
				u.CheckIdentifierPassword,
				u.CreateUser,

				c,
			) {
				return
			}
			token, refresh, err := genTokenRefresh(authService.Username, c)
			if err != nil {
				return
			}
			c.JSON(
				http.StatusOK,
				LoginResponse{
					Status:       SUCCESS,
					Message:      "User logged in successfully.",
					Token:        *token,
					RefreshToken: *refresh,
				},
			)
		case Phone:
			if !sendCodeNoPassword(&authService.Username, u.SendCode, c) {
				return

			}
			c.JSON(
				http.StatusOK,
				ResponseMessage{
					Status:  SUCCESS,
					Message: "a code has been sent for you please proceed to input it",
				},
			)
		case Mail:
			if !sendCodeNoPassword(&authService.Username, u.SendCode, c) {
				return

			}
			c.JSON(
				http.StatusOK,
				ResponseMessage{
					Status:  SUCCESS,
					Message: "a code has been sent for you please proceed to input it",
				},
			)
		case Wallet:
			if !sendCodeNoPassword(&authService.Username, u.SendCode, c) {
				return

			}
			// authenticate
		default:
			if !checkIdentifierPassword(
				&authService.Username,
				&authService.Password,
				u.CheckIdentifierPassword,
				u.CreateUser,
				c,
			) {
				return
			}
			token, refresh, err := genTokenRefresh(authService.Username, c)
			if err != nil {
				return
			}
			c.JSON(
				http.StatusOK,
				LoginResponse{
					Status:       SUCCESS,
					Message:      "User logged in successfully.",
					Token:        *token,
					RefreshToken: *refresh,
				},
			)
		}

	}
}

type RegisterResponse struct {
	Status  customErrors `json:"status"`
	Message string       `json:"message"`
	Token   string       `json:"token,omitempty"`
}

func Register(u *UtilityFuncs) gin.HandlerFunc {
	return func(c *gin.Context) {
		valid := validation.Validation{}
		var auth Auth
		err := c.BindJSON(&auth)
		if err != nil {
			c.JSON(
				http.StatusOK,
				ErrorResponseMessage{Status: ERROR, Error: "Invalid request."},
			)

			return
		}

		ok, _ := valid.Valid(&auth)

		if !ok {
			c.JSON(
				http.StatusOK,
				ErrorResponseMessage{Status: ERROR_EXIST_USER, Error: "Invalid request."},
			)

			return
		}

		authService := IdentifierPasswordAuth{Username: auth.Username, Password: auth.Password}
		isExist := u.CheckIdentifierPassword(&authService.Username, &authService.Password)
		if !isExist {
			c.JSON(
				http.StatusOK,
				ErrorResponseMessage{
					Status: ERROR_AUTH,
					Error:  "Invalid email or password.",
				},
			)

			// c.JSON(http.StatusBadRequest, RegisterResponse{Status: http.StatusBadRequest, Message: "Invalid phone number or password."})
			return
		}

		fmt.Println(isExist)
		if isExist {
			user := u.GetUser(&auth.Username)
			if user == nil {
				c.JSON(
					http.StatusOK,
					ErrorResponseMessage{
						Status: ERROR_AUTH,
						Error:  "Invalid email or password.",
					},
				)

				return
			}

		}

		user := u.CreateUser(&auth.Username)
		if user == nil {
			c.JSON(
				http.StatusOK,
				ErrorResponseMessage{
					Status: ERROR_EXIST_USER,
					Error:  "User already exists.",
				},
			)

			return
		}
		suc := u.SendCode(&auth.Username)
		if !suc {
			c.JSON(
				http.StatusOK,
				ErrorResponseMessage{
					Status: ERROR_SENDING_MAIL_FAILED,
					Error:  "error sending the code.",
				},
			)

			return
		}
		c.JSON(
			http.StatusOK,
			RegisterResponse{
				Status:  SUCCESS,
				Message: "Account created successfully, A confirmation code has been sent to your identifier",
			},
		)
	}
}

type ConfirmRegisterReq struct {
	Username string `valid:"Required; MaxSize(50)" json:"email"`
	Code     string `valid:"Required; MaxSize(50)" json:"code"`
}

type ConfirmRegisterResponse struct {
	Code string `valid:"Required; MaxSize(50)" json:"code"`
}

func ConfirmCode(u *UtilityFuncs) gin.HandlerFunc {
	return func(c *gin.Context) {
		valid := validation.Validation{}
		var auth ConfirmRegisterReq
		err := c.BindJSON(&auth)
		if err != nil {
			c.JSON(
				http.StatusOK,
				ErrorResponseMessage{Status: ERROR, Error: "Invalid request."},
			)

			return
		}

		ok, _ := valid.Valid(&auth)

		if !ok {
			c.JSON(
				http.StatusOK,
				ErrorResponseMessage{Status: ERROR_EXIST_USER, Error: "Invalid request."},
			)

			return
		}

		ok = u.ValidateCode(&auth.Username, &auth.Code)
		if !ok {
			c.JSON(
				http.StatusOK,
				ErrorResponseMessage{
					Status: ERROR_AUTH_CHECK_CODE_FAIL,
					Error:  "Invalid code.",
				},
			)
			return

		}
		exists := u.UserExists(&auth.Username)
		if !exists {

			user := u.CreateUser(&auth.Username)
			if user == nil {
				c.JSON(
					http.StatusOK,
					ErrorResponseMessage{
						Status: ERROR_EXIST_USER,
						Error:  "User already exists.",
					},
				)

				return
			}
		}

		token, err := CreateToken(auth.Username, time.Duration(24*time.Hour))

		if err != nil {
			c.JSON(
				http.StatusOK,
				ErrorResponseMessage{
					Status: ERROR,
					Error:  "Something went wrong, please try again later.",
				},
			)

			return
		}
		refreshtoken, err := CreateRefreshToken(auth.Username)

		if err != nil {
			c.JSON(
				http.StatusOK,
				ErrorResponseMessage{
					Status: ERROR,
					Error:  "Something went wrong, please try again later.",
				},
			)

			return
		}
		c.JSON(
			http.StatusOK,
			LoginResponse{
				Status:       SUCCESS,
				Message:      "identifier verified successfully.",
				Token:        token,
				RefreshToken: refreshtoken,
			},
		)

	}
}

type RefreshTokenRes struct {
	Token   string       `valid:"Required; MaxSize(50)" json:"token"`
	Status  customErrors `valid:"Required; MaxSize(50)" json:"status"`
	Message string       `valid:"Required; MaxSize(50)" json:"message"`
}

func RefreshToken(u *UtilityFuncs) gin.HandlerFunc {
	return func(c *gin.Context) {
		headerToken := c.GetHeader("Authorization")
		if headerToken == "" {
			c.JSON(
				http.StatusOK,
				ErrorResponseMessage{Status: ERROR_AUTH_TOKEN, Error: "Invalid token."},
			)
			return
		}
		token := strings.Split(headerToken, " ")[1]
		valid, err := ValidateToken(token)
		username, usernameerr := UsernameFromToken(&token)
		if ok := u.UserExists(&username); !ok {
			c.JSON(
				http.StatusOK,
				ErrorResponseMessage{
					Status: ERROR_EXIST_USER,
					Error:  "user doesnt exist",
				},
			)
			c.Abort()
			return
		}
		if usernameerr != nil {
			c.JSON(
				http.StatusOK,
				ErrorResponseMessage{Status: ERROR_AUTH_TOKEN, Error: "token is curropt"},
			)
			c.Abort()
			return
		}
		if !valid {
			c.JSON(
				http.StatusOK,
				ErrorResponseMessage{Status: ERROR_AUTH_TOKEN, Error: "token not valid"},
			)
			c.Abort()
			return
		}

		if err != nil {
			c.JSON(
				http.StatusOK,
				ErrorResponseMessage{Status: ERROR_AUTH_TOKEN, Error: "Invalid token."},
			)
			return
		}
		token, err = CreateRefreshToken(username)
		if err != nil {
			c.JSON(
				http.StatusOK,
				ErrorResponseMessage{
					Status: ERROR,
					Error:  "Something went wrong, please try again later.",
				},
			)
			return
		}
		c.JSON(
			http.StatusOK,
			RefreshTokenRes{
				Status:  SUCCESS,
				Message: "Token refreshed successfully.",
				Token:   token,
			},
		)

	}
}
