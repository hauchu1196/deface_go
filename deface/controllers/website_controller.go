package controllers

import (
	"deface/models"
	"deface/services"
	_ "log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateWebsiteInput struct {
	URL string `json:"url" binding:"required,url"`
}

func CreateWebsite(c *gin.Context) {
	var input CreateWebsiteInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingWebsite models.Website
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("url = ?", input.URL).First(&existingWebsite).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Website already exists."})
		return
	}

	htmlHash, err := services.GetWebsiteHTMLHash(input.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch website HTML."})
		return
	}

	screenshotName := "screenshot_initial.png"
	err = services.TakeScreenshot(input.URL, screenshotName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to take screenshot."})
		return
	}

	website := models.Website{
		URL:        input.URL,
		HTMLHash:   htmlHash,
		Screenshot: screenshotName,
	}

	if err := db.Create(&website).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create website."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": website})
}

func GetWebsites(c *gin.Context) {
	var websites []models.Website
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Find(&websites).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch websites."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": websites})
}

func GetWebsiteByID(c *gin.Context) {
	var website models.Website
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	if err := db.First(&website, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Website not found."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": website})
}
