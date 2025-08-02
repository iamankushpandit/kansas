package controllers

import (
	"kansas-healthcare-api/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AnalyticsController struct {
	service services.AnalyticsServiceInterface
}

func NewAnalyticsController(service services.AnalyticsServiceInterface) *AnalyticsController {
	return &AnalyticsController{service: service}
}

func (c *AnalyticsController) GetAllCountyData(ctx *gin.Context) {
	log.Printf("[INFO] Getting all county data")
	data, err := c.service.GetAllCountyData()
	if err != nil {
		log.Printf("[ERROR] Failed to get county data: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Printf("[INFO] Successfully retrieved %d counties data", len(data))
	ctx.JSON(http.StatusOK, data)
}

func (c *AnalyticsController) GetCountyData(ctx *gin.Context) {
	county := ctx.Param("county")
	log.Printf("[INFO] Getting data for county: %s", county)

	data, err := c.service.GetCountyData(county)
	if err != nil {
		log.Printf("[ERROR] Failed to get data for county %s: %v", county, err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if data == nil {
		log.Printf("[WARN] County not found: %s", county)
		ctx.JSON(http.StatusNotFound, gin.H{"error": "County not found"})
		return
	}

	log.Printf("[INFO] Successfully retrieved data for county: %s", county)
	ctx.JSON(http.StatusOK, data)
}

func (c *AnalyticsController) GetRecommendations(ctx *gin.Context) {
	county := ctx.Param("county")
	recommendations := c.service.GetRecommendations(county)
	ctx.JSON(http.StatusOK, recommendations)
}

func (c *AnalyticsController) GetActiveProviderCount(ctx *gin.Context) {
	count, err := c.service.GetActiveProviderCount()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"total_active_providers": count})
}

func (c *AnalyticsController) GetTerminatedNetworkAnalysis(ctx *gin.Context) {
	networkId := ctx.Query("network_id") // Commercial, Medicare, Tricare

	if networkId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "network_id query parameter is required"})
		return
	}

	result, err := c.service.GetTerminatedNetworkAnalysis(networkId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (c *AnalyticsController) GetCountyTerminatedNetworkAnalysis(ctx *gin.Context) {
	county := ctx.Param("county")
	networkId := ctx.Query("network_id")

	if networkId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "network_id query parameter is required"})
		return
	}

	result, err := c.service.GetCountyTerminatedNetworkAnalysis(county, networkId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (c *AnalyticsController) GetSpecialtyDensityAnalysis(ctx *gin.Context) {
	county := ctx.Param("county")

	result, err := c.service.GetSpecialtyDensityAnalysis(county)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (c *AnalyticsController) GetRadiusAnalysis(ctx *gin.Context) {
	county := ctx.Param("county")
	radius := ctx.DefaultQuery("radius", "25")
	network := ctx.Query("network")

	if network == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "network query parameter is required"})
		return
	}

	// Convert radius to int
	radiusInt := 25
	if radius != "" {
		if r, err := strconv.Atoi(radius); err == nil {
			radiusInt = r
		}
	}

	result, err := c.service.GetRadiusAnalysis(county, radiusInt, network)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, result)
}
