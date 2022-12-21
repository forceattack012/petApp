package user

import (
	"net/http"

	"github.com/forceattack012/petAppApi/auth"
	"github.com/forceattack012/petAppApi/domain"
	"github.com/forceattack012/petAppApi/entities"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	domain.UserDomain
}

func NewUserHandler(d domain.UserDomain) *UserHandler {
	return &UserHandler{d}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user entities.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result := h.UserDomain.Create(&user)
	if err := result.Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "sucess",
	})
}

func (h *UserHandler) Login(signature []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user entities.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var userLogin *entities.User
		userLogin, result := h.UserDomain.GetUser(user.Username, user.Password)
		if err := result.Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
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
			"id":    userLogin.ID,
			"name":  userLogin.Username,
			"token": token,
		})
	}
}
