package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasamadeuu/bookstore-API/controllers"
	"github.com/lucasamadeuu/bookstore-API/models"
)

func main() {

	router := gin.Default()

	models.ConnectDatabase()

	router.GET("/books", controllers.EncontrarLivros)
	router.POST("/books", controllers.CriarLivros)
	router.GET("/books/:id", controllers.EncontrarLivro)
	router.PATCH("/books/:id", controllers.AtualizarLivro)
	router.DELETE("/books/:id", controllers.DeletarLivro)

	router.Run()

}
