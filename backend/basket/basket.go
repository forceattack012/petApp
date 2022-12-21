package basket

import (
	"log"
	"net/http"

	"github.com/forceattack012/petAppApi/domain"
	"github.com/forceattack012/petAppApi/entities"
	"github.com/gin-gonic/gin"
)

type BasketHandler struct {
	domain.BasketDomain
}

func NewBasketHandler(d domain.BasketDomain) *BasketHandler {
	return &BasketHandler{d}
}

func (h *BasketHandler) Add(c *gin.Context) {
	var basket entities.Basket
	if err := c.ShouldBindJSON(&basket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	h.BasketDomain.Add(basket.Username, &basket)

	c.JSON(http.StatusCreated, gin.H{
		"message": "sucess",
	})
}

func (h *BasketHandler) Get(c *gin.Context) {
	userName := c.Param("username")
	log.Printf("userName: %s", userName)

	if basket, found := h.BasketDomain.Get(userName); found {
		c.JSON(http.StatusOK, basket)
		return
	}
	c.JSON(http.StatusNotFound, nil)
}
