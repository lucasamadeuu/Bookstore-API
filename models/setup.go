package models

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "teste.db")

	if err != nil {
		log.Println(err)
		panic("Erro ao conectar no banco de dados.")
	}

	database.AutoMigrate(&Livro{})

	DB = database
}
