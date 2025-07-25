package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getAllItems(c *gin.Context) {
	items, err := h.Services.ItemService.GetAllItems(c.Request.Context())
	if err != nil {
		h.Logger.Error("error in getting all items", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}
	c.JSON(http.StatusOK, items)
}
