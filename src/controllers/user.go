package controllers

import (
	"encoding/json"

	"golang.org/x/crypto/bcrypt"

	"github.com/Hy-Iam-Noval/dacol/src/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

var validate *validator.Validate

type RegisterForm struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,eqfield=Password2,min=8"`
	Password2 string `json:"password2"`
}

const UserSessionKey = "user"

func Login(c *gin.Context) {
	var req, user models.User

	// get request
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(err)
	}

	// query
	if err := Conn().
		QueryRow("SELECT * FROM user WHERE email = ?", req.Email).
		Scan(&user.Id, &user.Email, &user.Password); err != nil {
		panic(err)
	}

	defer Conn().Close()

	if res := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); res != nil && (user == models.User{}) {
		c.JSON(401, struct{ err string }{err: "Password or Email wrong"})
		return
	}

	s := sessions.Default(c)
	jsonData, err := json.Marshal(user)
	if err != nil {
		ErrCtrl("Login", err)
	}
	s.Set(UserSessionKey, string(jsonData))
	s.Save()

	c.Done()

}

func User(c *gin.Context) {
	s := sessions.Default(c)
	c.JSON(200, s.Get("user"))
}

func Logout(c *gin.Context) {
	s := sessions.Default(c)
	s.Delete("user")
	c.Done()
}

// PASS
func Register(c *gin.Context) {
	var req RegisterForm

	if err := c.ShouldBindJSON(&req); err != nil {
		ErrCtrl("Register", err)
	}
	// validate
	validate = validator.New()
	if err := validate.Struct(req); err != nil {
		c.JSON(401, ValidateMsg(err))
		return
	}

	newPass, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if _, err := Conn().Query("INSERT INTO user(email, password) VALUES(?, ?)", req.Email, string(newPass)); err != nil {
		ErrCtrl("Register", err)
	}
	c.Done()
}
