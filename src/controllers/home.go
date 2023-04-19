package controllers

import (
	"encoding/json"

	"github.com/Hy-Iam-Noval/dacol/src/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	// get session user
	var user models.User
	s := sessions.Default(c)

	jsonData := s.Get(UserSessionKey).(string)

	if err := json.Unmarshal([]byte(jsonData), &user); err != nil {
		panic(err)
	}

	// query
	datas := []models.Product{}

	for q, _ := Conn().Query("SELECT * FROM product WHERE id_user = ?", user.Id); q.Next(); {
		data := models.Product{}
		if err := q.Scan(&data.Id, &data.Name, &data.Qty); err != nil {
			panic(err)
		}
		datas = append(datas, data)
	}
	Conn().Close()

	// res
	c.JSON(200, datas)
}
