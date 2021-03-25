package Controllers

import (
	"../Classes"
	"../Models"
	"../Responses"
	"github.com/gin-gonic/gin"
	"html/template"
)

type AuthController struct {
}

func (AuthController) Login(c *gin.Context) {
	var response Responses.Auth
	response.Success = false
	username_or_email := c.PostForm("username_or_email")
	if Models.IsUsernameExist(username_or_email) || Models.IsEmailExist(username_or_email) {
		user := Models.GetUserByUOE(username_or_email)
		var hash Classes.Hash
		if hash.CompareHashes(c.PostForm("password"), user.Password) {
			if user.AccountStatus == 1 {
				c.SetCookie("userToken", user.Token, 60*60*30, "/", "localhost", true, true)
				response.Success = true
				response.Message = append(response.Message, "you logged in")
				response.Url = c.FullPath()
				c.JSON(200, response)
			} else {
				response.Message = append(response.Message, "account is deactive check your email")
				c.JSON(404, response)
			}
		} else {
			response.Message = append(response.Message, "wrong password")
			c.JSON(404, response)
		}

	} else {
		response.Message = append(response.Message, "account does not exist")
		c.JSON(404, response)
	}

}

func (AuthController) Register(c *gin.Context) {
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")
	confirmPassword := c.PostForm("confirm_password")
	var response Responses.Auth
	var Validation Classes.Validation
	response.Success = false
	if username == "" {
		response.Message = append(response.Message, "'username' field required")
	}
	if email == "" {
		response.Message = append(response.Message, "'email' field required")
	}
	if password == "" {
		response.Message = append(response.Message, "'password' field required")
	}
	if confirmPassword == "" {
		response.Message = append(response.Message, "'confirm password' field required ")
	}
	if len(response.Message) > 0 {
		c.JSON(400, response)
	} else {
		if Models.IsEmailExist(email) {
			response.Message = append(response.Message, "email exists")
		}
		if Validation.IsEmail(email) {
			response.Message = append(response.Message, "this is not an email!!!")
		}
		if Models.IsUsernameExist(username) {
			response.Message = append(response.Message, "username exists")
		}
		if Validation.IsUsername(username) {
			response.Message = append(response.Message, "username must be english characters you can add . and _ in username like -> jafar / jaf_ar")
		}
		if len(password) < 8 {
			response.Message = append(response.Message, "password length must be at least 8 characters")
		}
		if password != confirmPassword {
			response.Message = append(response.Message, "password and confirm password does not match")
		}
		if len(response.Message) > 0 {
			c.JSON(400, response)
		} else {
			Models.CreateUser(username, email, password, 0)
			response.Success = true
			response.Message = append(response.Message, "account successfully created")
			logMessage := "verify email sent to -> " + template.HTMLEscapeString(email)
			response.Message = append(response.Message, logMessage)
			c.JSON(400, response)
		}
	}

}

func (AuthController) ForgotPassword(c *gin.Context) {

}

func (AuthController) VerifyAccount(c *gin.Context) {
	token := c.Param("token")
	status := Models.ActivateAccount(token)
	if status {
		c.String(200, "activated")
	} else {
		c.String(200, "Wrong Token")
	}
}
