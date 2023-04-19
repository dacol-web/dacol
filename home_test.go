package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/assert.v1"
)

func rw() *gin.Engine {
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

	return r
}

var r *gin.Engine = rw()

func TestSet(t *testing.T) {

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/set", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

}

func TestGet(t *testing.T) {

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/get", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "1", w.Body)
}
