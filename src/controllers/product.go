package controllers

import (
	"net/http"

	"github.com/Hy-Iam-Noval/dacol/src/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// POST /add-product/
func ProdAdd(c *gin.Context) {
	req := models.Product{}
	session := sessions.Default(c)

	c.BindJSON(req)

	// insert
	_, err := Conn().
		Query(`
			INSERT INTO 
				product(name, price, qty, descript, id_user) 
				VALUES(?, ?, ?, ?, ?)`,
			req.Name, req.Price, req.Qty, req.Descript,
			session.Get("user").(models.User).Id,
		)
	Conn().Close()

	if err != nil {
		panic(err)
	}

	c.Done()
}

// DELETE /product/:id/delete
func ProdDelete(c *gin.Context) {
	req, _ := c.Params.Get("id")
	_, err := Conn().Query("DELETE FROM product WHERE id=?", req)

	Conn().Close()
	if err != nil {
		panic(err)
	}
	c.Done()
}

// GET /product/:id
func ProdDetail(c *gin.Context) {
	req, _ := c.Params.Get("id")
	row := Conn().QueryRow("SELECT FROM product WHERE id=?", req)

	Conn().Close()
	c.JSON(http.StatusOK, row)
}
