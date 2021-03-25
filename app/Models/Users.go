package Models

import (
	".."
	"../Classes"
)

type User struct {
	Id            int64
	username      string
	email         string
	Password      string
	AccountStatus int
	Token         string
	success       bool
}

func IsUsernameExist(username string) bool {
	db := app.Database()
	var total int
	err := db.QueryRow("SELECT COUNT(*) AS total FROM users WHERE username =?", username).Scan(&total)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	if total > 0 {
		return true
	}
	return false
}

func IsEmailExist(email string) bool {
	db := app.Database()
	var total int
	err := db.QueryRow("SELECT COUNT(*) AS total FROM users WHERE email =?", email).Scan(&total)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	if total > 0 {
		return true
	}
	return false
}

func CreateUser(username string, email string, password string,account_status int) User {
	var user User
	user.success = false
	if !IsUsernameExist(username) || !IsEmailExist(email) {
		var hash Classes.Hash
		hashedPassword := hash.HashString(password)
		db := app.Database()
		res, err := db.Exec("INSERT INTO users (username,email,password,token,api_token,account_status) VALUES (?,?,?,?,?)", username, email, hashedPassword,hash.GenerateToken(),hash.GenerateToken(), account_status)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}else{
			id, err := res.LastInsertId()
			if err == nil {
				var emailHelper Classes.EmailHelper
				emailHelper.SendEmail("test@test.com",email,"verify account","")
				user.success = true
				user.Id = id
			}
		}
		return user
	}
	return user
}

func ActivateAccount(token string) bool {
	db := app.Database()
	var total int
	err := db.QueryRow("SELECT COUNT(*) AS total FROM users WHERE token =?", token).Scan(&total)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	if total > 0 {
		db.QueryRow("UPDATE users SET account_status = 1 WHERE token =?", token)
		return true
	}
	return false
}

func GetUserByUOE(username_or_email string) User {
	db := app.Database()
	var user User
	err := db.QueryRow("SELECT id,username,password,email,account_status,token AS total FROM users WHERE username =? OR email =?", username_or_email,username_or_email).Scan(&user.Id,&user.username,&user.Password,&user.email,&user.AccountStatus,&user.Token)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	return user
}

func GetUserByUsername()  {

}

func GetUserByEmail() {

}