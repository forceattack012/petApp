package user

import (
	"net/http"

	"github.com/forceattack012/petAppApi/auth"
	"github.com/forceattack012/petAppApi/entities"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string           `json:"username"`
	Password string           `json:"password"`
	Owners   []entities.Owner `json:"owners" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type UserHandler struct {
	db *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{db: db}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result := h.db.Create(&user)
	if err := result.Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "sucess",
	})
}

func (h *UserHandler) Login(signature []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var userLogin User
		result := h.db.Where("username = ? AND password = ?", user.Username, user.Password).Find(&userLogin)
		if err := result.Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}

		if userLogin.Username == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "username doesn't exist",
			})
			return
		}

		token, err := auth.AccessToken(signature, userLogin.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"id":    userLogin.Model.ID,
			"name":  userLogin.Username,
			"token": token,
		})
	}
}
