package pet

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/forceattack012/petAppApi/domain"
	"github.com/forceattack012/petAppApi/entities"
	"github.com/forceattack012/petAppApi/pagination"
	"github.com/gin-gonic/gin"
)

type PetResoponse struct {
	Id          string `json:"id" binding:"required"`
	PetId       string `json:"petId" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Content     string `json:"content"`
}

func Tablename() string {
	return "pets"
}

type PetHandler struct {
	domain.PetDomain
}

func NewPetHandler(d domain.PetDomain) *PetHandler {
	return &PetHandler{d}
}

func (h *PetHandler) NewPet(c *gin.Context) {
	var pet entities.Pet
	if err := c.ShouldBindJSON(&pet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result := h.Save(&pet)
	if result != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"ID": pet.ID,
	})
}

func (h *PetHandler) GetPets(c *gin.Context) {
	q := c.Request.URL.Query()
	pageSize, _ := strconv.Atoi(q.Get("limit"))
	page, _ := strconv.Atoi(q.Get("page"))
	var pets []entities.Pet
	var totalRows int64

	h.PetDomain.ExceuteSql("SELECT count(*) FROM pets p LEFT JOIN owners o ON p.id = o.pet_id WHERE o.id IS NULL AND o.deleted_at IS NULL", &totalRows)
	fmt.Printf("total %d \n", totalRows)

	if page == 0 {
		page = 1
	}

	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	totalPages := int(math.Ceil(float64(totalRows) / float64(pageSize)))
	resultErr := h.PetDomain.Paginate(page, pageSize, "Files", "LEFT JOIN owners o ON pets.id = o.pet_id AND o.deleted_at IS NULL", "o.id IS NULL", &pets)
	if resultErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": resultErr.Error(),
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

	resultErr := h.PetDomain.Delete(&entities.Pet{}, id)
	if resultErr != nil {
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

	var pet entities.Pet
	errDb := h.PetDomain.GetPet(&pet, id)
	if errDb != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": errDb.Error(),
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

	var newPet entities.Pet
	if err := c.ShouldBindJSON(&newPet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var pet entities.Pet
	resultErr := h.PetDomain.GetPet(&pet, id)
	if resultErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": resultErr.Error(),
		})
		return
	}

	pet.Name = newPet.Name
	pet.Type = newPet.Type
	pet.Description = newPet.Description
	pet.Age = newPet.Age
	resultUpdate := h.PetDomain.Update(&pet)

	if resultUpdate != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": resultUpdate.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "sucess",
	})
}

func (h *PetHandler) GetOwnerByUserId(c *gin.Context) {
	userId := c.Param("userId")

	var petResp []PetResoponse
	err := h.PetDomain.Raw("SELECT o.id, o.pet_id, p.name, p.type, p.description, f.content FROM pets p JOIN owners o ON p.id = o.pet_id JOIN files f ON f.pet_id = p.id WHERE o.user_id = ? AND o.deleted_at IS NULL GROUP BY f.pet_id", userId, &petResp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, petResp)
}
