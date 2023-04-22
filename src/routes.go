package src

import (
	"net/http"

	ctrl "github.com/Hy-Iam-Noval/dacol/src/controllers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Route() *gin.Engine {
	// All PASS
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

	// Pass
	// Not need auth

	r.POST("/login", ctrl.Login)
	r.POST("/register", ctrl.Register)

	// Need auth first
	auth := r.Group("/auth").Use(ctrl.AuthReq)
	{
		// Get user from session
		auth.GET("/user", ctrl.User) // Pass

		auth.GET("/home", ctrl.Home) // Pass

		// Route Product
		// Pass
		auth.POST("/add-product", ctrl.ProdAdd)
		auth.GET("/product/:id", ctrl.ProdDetail)
		auth.DELETE("/product/:id/delete", ctrl.ProdDelete)

		auth.DELETE("/logout", ctrl.Logout) // Pass
	}

	return r
}
