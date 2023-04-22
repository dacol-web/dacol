package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Hy-Iam-Noval/dacol/src/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// POST /add-product/
func ProdAdd(c *gin.Context) {
	req := models.Product{}
	session := sessions.Default(c)
	user := models.User{}

	sUser, ok := session.Get("user").(string)
	if !ok {
		panic("Invalid type in ProdAdd")
	}

	if err := json.Unmarshal([]byte(sUser), &user); err != nil {
		ErrCtrl("ProdAdd", err)
	}

	if err := c.BindJSON(&req); err != nil {
		ErrCtrl("Prodadd", err)
	}

	// insert
	_, err := Conn().
		Query(`
			INSERT INTO 
				product(name, price, qty, descript, id_user) 
				VALUES(?, ?, ?, ?, ?)`,
			req.Name, req.Price, req.Qty, req.Descript,
			user.Id,
		)
	Conn().Close()

	if err != nil {
		ErrCtrl("ProdAdd", err)
	}

	c.Done()
}

// DELETE /auth/product/:id/delete
func ProdDelete(c *gin.Context) {
	req, ok := c.Params.Get("id")
	if !ok {
		ErrCtrl("ProdDelete", errors.New("Error"))
	}
	_, err := Conn().Query("DELETE FROM product WHERE id=?", req)

	Conn().Close()
	if err != nil {
		ErrCtrl("ProdDelete", err)
	}
	c.Done()
}

// GET /product/:ids
func ProdDetail(c *gin.Context) {
	req := c.Param("id")
	row := Conn().QueryRow("SELECT FROM product WHERE id=?", req)
	Conn().Close()

	c.JSON(http.StatusOK, row)
}
