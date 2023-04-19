package models

import (
	_ "database/sql"
)

type User struct {
	Id       uint   `json:"id"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"`
}

type Product struct {
	Id       uint   `db:"id"`
	Qty      uint   `db:"qty" json:"qty"`
	Price    uint   `db:"price" json:"price"`
	IdUser   uint   `db:"id_user"`
	Name     string `db:"name" json:"name"`
	Descript string `db:"descript" json:"descript"`
}
