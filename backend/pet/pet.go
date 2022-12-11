package pet

import (
	"math"
	"net/http"
	"strconv"

	"github.com/forceattack012/petAppApi/file"
	"github.com/forceattack012/petAppApi/pagination"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Pet struct {
	gorm.Model
	Name        string      `json:"name" binding:"required"`
	Type        string      `json:"type"`
	Description string      `json:"description"`
	Age         string      `json:"age"`
	Files       []file.File `json:"files" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
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
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"ID": pet.Model.ID,
	})
}

func (h *PetHandler) GetPets(c *gin.Context) {
	q := c.Request.URL.Query()
	pageSize, _ := strconv.Atoi(q.Get("limit"))
	page, _ := strconv.Atoi(q.Get("page"))
	var pets []Pet
	var totalRows int64
	h.db.Model(Pet{}).Count(&totalRows)

	totalPages := int(math.Ceil(float64(totalRows) / float64(pageSize)))

	result := h.db.Scopes(Paginate(c)).Preload("Files").Find(&pets)
	if err := result.Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	p := &pagination.Pagination{
		Limit:      pageSize,
		Page:       page,
		Sort:       "",
		TotalRows:  totalRows,
		TotalPages: totalPages,
		Result:     pets,
	}

	c.JSON(http.StatusOK, p)
}

func (h *PetHandler) DeletePet(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result := h.db.Preload("Files").Delete(&Pet{}, id)
	if err := result.Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "sucess",
	})
}

func (h *PetHandler) GetPet(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var pet Pet
	result := h.db.Preload("Files").Find(&pet, id)
	if err := result.Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, pet)
}

func (h *PetHandler) UpdatePet(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var newPet Pet
	if err := c.ShouldBindJSON(&newPet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var pet Pet
	result := h.db.Find(&pet, id)
	if err := result.Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	pet.Name = newPet.Name
	pet.Type = newPet.Type
	pet.Description = newPet.Description
	pet.Age = newPet.Age
	resultUpdate := h.db.Updates(pet)

	if err := resultUpdate.Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "sucess",
	})
}

func Paginate(r *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		q := r.Request.URL.Query()
		page, _ := strconv.Atoi(q.Get("page"))
		if page == 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(q.Get("limit"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
