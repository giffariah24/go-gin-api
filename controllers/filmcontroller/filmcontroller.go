package filmcontroller

import (
	"encoding/json"
	"github.com/giffariah666/go-gin-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func GetFilm(c *gin.Context) {
	var films []models.Film

	models.DB.Find(&films)
	c.JSON(http.StatusOK, gin.H{"films": films})
}

func GetFilmById(c *gin.Context) {
	var film models.Film
	id := c.Param("id")

	if err := models.DB.First(&film, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data film tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"film": film})
}

func CreateFilm(c *gin.Context) {
	var film models.Film

	if err := c.ShouldBindJSON(&film); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	models.DB.Create(&film)
	c.JSON(http.StatusOK, gin.H{"film": film})
}

func UpdateFilm(c *gin.Context) {
	var film models.Film
	id := c.Param("id")

	if err := c.ShouldBindJSON(&film); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if models.DB.Model(&film).Where("id = ?", id).Updates(&film).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Failed to update"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data update success"})
}

func DeleteFilm(c *gin.Context) {
	var film models.Film

	var input struct {
		Id json.Number
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	id, _ := input.Id.Int64()
	if models.DB.Delete(&film, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus film!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil di hapus"})
}
