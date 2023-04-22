package controllers

import (
	"fmt"
	"os"
	"time"

	"github.com/Hy-Iam-Noval/dacol/src/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var (
	JWT_TOKEN         string                 = os.Getenv("JWT_TOKEN")
	JWT_SIGNED_METHOD *jwt.SigningMethodHMAC = jwt.SigningMethodHS256
)

type JWTClaim struct {
	models.User
	jwt.RegisteredClaims
}

func ValidateToken(signToken string) bool {
	token, err := jwt.ParseWithClaims(
		signToken,
		&JWTClaim{},
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected method %v", t.Method.Alg())
			}
			return []byte(JWT_TOKEN), nil
		},
	)
	if err != nil {
		ErrCtrl("ValidateToken", err)
	}

	_, ok := token.Claims.(*JWTClaim)
	return ok && token.Valid
}

func GenerateToken(user models.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaim{
		user,
		jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().AddDate(0, 0, 1),
			},
		},
	})

	// Sign and get the complete encoded token as a string using the secret
	s, err := token.SignedString([]byte(JWT_TOKEN))
	if err != nil {
		ErrCtrl("GenerateToken", err)
	}
	return s
}

func AuthReq(c *gin.Context) {
	sess := sessions.Default(c)
	if data := sess.Get(UserSessionKey); data == nil {
		c.AbortWithStatus(401)
	} else {
		c.Next()
	}
}
