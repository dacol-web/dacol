package main

import (
	"fmt"
	"net/http"

	ctrl "github.com/Hy-Iam-Noval/dacol/src/controllers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	// testing
	r := gin.New()
	// store
	store := cookie.NewStore([]byte("dwdwd"))
	store.Options(sessions.Options{
		Path:     "/",
		Domain:   "localhost",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	})

	r.Use(sessions.Sessions("web_session", store))

	// Route user
	r.POST("/login", ctrl.Login)
	r.POST("/register", ctrl.Register)

	// Home
	auth := r.Group("/auth").Use(ctrl.AuthReq)
	{
		auth.GET("/user", ctrl.User)
		auth.GET("/home", ctrl.Home)

		// Route Product
		auth.POST("/add-product", ctrl.ProdAdd)
		auth.GET("/product/:id/", ctrl.ProdDetail)
		auth.DELETE("/product/:id/delete", ctrl.ProdDelete)

		auth.DELETE("/logout", ctrl.Logout)
	}
	r.GET("/get", func(c *gin.Context) {
		s := sessions.Default(c)
		c.String(200, fmt.Sprintf("%v", s.Get("c")))
	})
	r.GET("/set", func(c *gin.Context) {
		s := sessions.Default(c)
		if c := s.Get("c"); c != nil {
			s.Set("c", c.(int)+1)
		} else {
			s.Set("c", 1)
		}
		s.Save()
		c.Done()
	})
	r.Run()
}
