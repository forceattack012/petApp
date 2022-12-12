package basket

import (
	"log"
	"net/http"

	"github.com/ReneKroon/ttlcache"
	"github.com/forceattack012/petAppApi/pet"
	"github.com/gin-gonic/gin"
)

type Basket struct {
	Username string    `json:"name"`
	Pets     []pet.Pet `json:"pets"`
}

type BasketHandler struct {
	Cache *ttlcache.Cache
}

func NewBasketHandler(cache *ttlcache.Cache) *BasketHandler {
	return &BasketHandler{Cache: cache}
}

func (h *BasketHandler) Add(c *gin.Context) {
	var basket Basket
	if err := c.ShouldBindJSON(&basket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	h.Cache.Set(basket.Username, basket)

	c.JSON(http.StatusCreated, gin.H{
		"message": "sucess",
	})
}

func (h *BasketHandler) Get(c *gin.Context) {
	userName := c.Param("username")
	log.Printf("userName: %s", userName)

	if basket, found := h.Cache.Get(userName); found {
		c.JSON(http.StatusOK, basket)
		return
	}
	c.JSON(http.StatusNotFound, nil)
}
