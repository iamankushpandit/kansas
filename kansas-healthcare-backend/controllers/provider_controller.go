package controllers

import (
	"kansas-healthcare-api/models"
	"kansas-healthcare-api/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProviderController struct {
	service services.ProviderServiceInterface
}

func NewProviderController(service services.ProviderServiceInterface) *ProviderController {
	return &ProviderController{service: service}
}

func (c *ProviderController) GetProviders(ctx *gin.Context) {
	providers, err := c.service.GetAllProviders()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, providers)
}

func (c *ProviderController) GetProviderNetwork(ctx *gin.Context) {
	networks, err := c.service.GetProviderNetworks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, networks)
}

func (c *ProviderController) GetFilteredData(ctx *gin.Context) {
	var filter models.FilterRequest
	if err := ctx.ShouldBindJSON(&filter); err != nil {
		log.Printf("[ERROR] Invalid filter request: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("[INFO] Filtering providers: specialty=%s, network=%s", filter.Specialty, filter.Network)
	providers, err := c.service.GetFilteredProviders(filter)
	if err != nil {
		log.Printf("[ERROR] Failed to filter providers: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Printf("[INFO] Successfully filtered %d providers", len(providers))
	ctx.JSON(http.StatusOK, providers)
}
