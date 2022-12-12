package owner

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Owner struct {
	gorm.Model
	UserID int64 `json:"user_id"`
	PetID  int64 `json:"pet_id"`
}

type OwnerHandler struct {
	db *gorm.DB
}

func NewOwnerHandler(db *gorm.DB) *OwnerHandler {
	return &OwnerHandler{
		db: db,
	}
}

func (h *OwnerHandler) CreateOwner(c *gin.Context) {
	var owner []Owner
	if err := c.ShouldBindJSON(&owner); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result := h.db.Create(&owner)
	if err := result.Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "sucess",
	})
}

func (h *OwnerHandler) RemoveOwner(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	result := h.db.Delete(&Owner{}, id)
	if err := result.Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
