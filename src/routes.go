package src

import (
	"net/http"

	ctrl "github.com/Hy-Iam-Noval/dacol/src/controllers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Route() *gin.Engine {
	// testing
	r := gin.New()
	// store
	store := cookie.NewStore([]byte("dwdwd"))
	store.Options(sessions.Options{
		Path:     "/",
		Domain:   "localhost:3000",
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

	return r
}
