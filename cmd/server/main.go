package main

import (
	"database/sql"
	"log"
	"../pgk/store"
	"github.com/go-sql-driver/mysql"
	"github.com/gin-gonic/gin"
)

func main() {

	db, err := sql.Open("mysql", "root:root.@admin(localhost:3306)/tp_clinica_odontologica")
	if err != nil {
		log.Fatal(err)
	}

	storageDB := store.NewSqlStore(db)

	//Repositorios
	rep_Odontologo := odontologo.newRepository(storageDB)
	
	

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	odontologoGroup := router.Group("/odontologo")
	{
		odontologoGroup.GET(":id", func() {})
		odontologoGroup.POST(":id", func() {})
		odontologoGroup.PUT(":id", func() {})
		odontologoGroup.PATCH(":id", func() {})
		odontologoGroup.DELETE(":id", func() {})
	}
}
