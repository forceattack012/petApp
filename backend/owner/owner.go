package owner

import (
	"net/http"
	"strconv"

	"github.com/forceattack012/petAppApi/domain"
	"github.com/forceattack012/petAppApi/entities"
	"github.com/gin-gonic/gin"
)

type OwnerHandler struct {
	d domain.OwnerDomain
}

func NewOwnerHandler(d domain.OwnerDomain) *OwnerHandler {
	return &OwnerHandler{
		d: d,
	}
}

func (h *OwnerHandler) CreateOwner(c *gin.Context) {
	var owner []entities.Owner
	if err := c.ShouldBindJSON(&owner); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	resultErr := h.d.Save(&owner)
	if resultErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": resultErr,
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

	err := h.d.Remove(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
