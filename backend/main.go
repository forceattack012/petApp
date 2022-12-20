package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ReneKroon/ttlcache"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/forceattack012/petAppApi/auth"
	"github.com/forceattack012/petAppApi/basket"
	"github.com/forceattack012/petAppApi/entities"
	"github.com/forceattack012/petAppApi/file"
	"github.com/forceattack012/petAppApi/owner"
	"github.com/forceattack012/petAppApi/pet"
	"github.com/forceattack012/petAppApi/store"
	"github.com/forceattack012/petAppApi/user"
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

	var cache = ttlcache.NewCache()

	db.AutoMigrate(&entities.Pet{})
	db.AutoMigrate(&file.File{})
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&owner.Owner{})
	fileHandler := file.NewFileHandler(db)
	userHandler := user.NewUserHandler(db)
	ownerHandler := owner.NewOwnerHandler(db)
	basketHandler := basket.NewBasketHandler(cache)

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

	signature := os.Getenv("signature")
	bytes := []byte(signature)

	store := store.NewPetStore(db)
	petHandler := pet.NewPetHandler(store)

	r.POST("/api/register", userHandler.CreateUser)
	r.POST("/api/login", userHandler.Login(bytes))
	r.GET("/api/pet", petHandler.GetPets)

	protect := r.Group("", auth.Protect(bytes))
	protect.POST("/api/pet", petHandler.NewPet)
	protect.GET("/api/pet/:id", petHandler.GetPet)
	protect.PATCH("/api/pet/:id", petHandler.UpdatePet)
	protect.DELETE("/api/pet/:id", petHandler.DeletePet)

	protect.POST("/api/basket", basketHandler.Add)
	protect.GET("/api/basket/:username", basketHandler.Get)

	protect.POST("/api/owner", ownerHandler.CreateOwner)
	protect.GET("/api/owner/:userId", petHandler.GetOwnerByUserId)
	protect.DELETE("/api/owner/:id", ownerHandler.RemoveOwner)

	r.POST("/api/upload/:id", fileHandler.Upload)
	r.GET("/api/download", fileHandler.Download)

	port := ":" + os.Getenv("PORT")
	r.Run(port)
}
