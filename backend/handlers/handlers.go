package handlers

import (
    "net/http"
    "url-shortener/models"
    "github.com/gin-gonic/gin"
)

func ShortenURL(c *gin.Context) {
    var req models.URLRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    shortCode := models.GenerateShortCode()
    models.SaveURL(shortCode, req.LongURL)
    c.JSON(http.StatusOK, gin.H{"short_code": shortCode, "clicks": 0})
}

func RedirectURL(c *gin.Context) {
    shortCode := c.Param("shortcode")
    longURL, err := models.GetURL(shortCode)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
        return
    }
    models.IncrementClicks(shortCode)
    c.Redirect(http.StatusMovedPermanently, longURL)
}

func GetAllURLs(c *gin.Context) {
    urls, err := models.GetAllURLs()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, urls)
}

func TrackClick(c *gin.Context) {
    shortCode := c.Param("shortcode")
    models.IncrementClicks(shortCode)
    c.JSON(http.StatusOK, gin.H{"message": "Click tracked"})
}