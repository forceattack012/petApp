package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/forceattack012/petAppApi/file"
	"github.com/forceattack012/petAppApi/pet"
)

func main() {
	if err := godotenv.Load("local.env"); err != nil {
		log.Fatal("cannot read configuration")
	}

	dsn := os.Getenv("dsn")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("connect database failed !!!")
	}

	db.AutoMigrate(&pet.Pet{})
	db.AutoMigrate(&file.File{})
	petHandler := pet.NewPetHandler(db)
	fileHandler := file.NewFileHandler(db)

	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://localhost:4200",
	}
	config.AllowHeaders = []string{
		"Origin",
	}
	config.AllowMethods = []string{
		"GET",
		"POST",
		"PUT",
		"PATCH",
		"DELETE",
	}
	r.Use(cors.New(config))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/api/pet", petHandler.NewPet)
	r.GET("/api/pet", petHandler.GetPets)
	r.POST("/api/upload", fileHandler.Upload)
	r.GET("/api/download", fileHandler.Download)

	port := ":" + os.Getenv("PORT")
	r.Run(port)
}
