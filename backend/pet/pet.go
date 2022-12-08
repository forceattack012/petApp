package pet

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Pet struct {
	gorm.Model
	Name        string `json:"name" binding:"required"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Age         string `json:"age"`
	// Birthdate   string `json:"birthdate"`
	// Age         int64  `json:"age"`
	// Image       []byte `json:"image"`
	// gorm.Model
}

func (Pet) Tablename() string {
	return "pets"
}

type PetHandler struct {
	db *gorm.DB
}

func NewPetHandler(db *gorm.DB) *PetHandler {
	return &PetHandler{
		db: db,
	}
}

func (h *PetHandler) NewPet(c *gin.Context) {
	var pet Pet
	if err := c.ShouldBindJSON(&pet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result := h.db.Create(&pet)
	if err := result.Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"ID": pet.Model.ID,
	})
}

func (h *PetHandler) GetPets(c *gin.Context) {
	var pets []Pet
	result := h.db.Find(&pets)

	if err := result.Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
		return
	}
	c.JSON(http.StatusOK, pets)
}
