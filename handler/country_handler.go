package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thisispj/autosuggestions/service"
)

type Handler struct {
	service *service.CountryService
}

func NewCountryHandler(service *service.CountryService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) AutocompleteHandler(c *gin.Context) {
	query := c.Query("query")
	suggestions, err := h.service.GetAutocompleteSuggestions(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, suggestions)
}
