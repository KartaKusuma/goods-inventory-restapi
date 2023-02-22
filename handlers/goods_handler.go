package handlers

import (
	"goods-inventory-restapi/database"
	"goods-inventory-restapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllGoods(c *gin.Context) {
	var goods []models.Goods

	database.DB.Find(&goods)

	c.JSON(http.StatusOK, gin.H{"goods": goods})
}

func GetGoodsByID(c *gin.Context) {
	var goods models.Goods

	goodsID := c.Param("id")

	if err := database.DB.First(&goods, goodsID).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data is not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"goods": goods})
}

func StoreGoods(c *gin.Context) {
	var goods models.Goods

	if err := c.ShouldBindJSON(&goods); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	database.DB.Create(&goods)

	c.JSON(http.StatusOK, gin.H{"goods": goods})
}

func UpdateGoods(c *gin.Context) {
	var goods models.Goods

	goodsID := c.Param("id")

	if err := c.ShouldBindJSON(&goods); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if database.DB.Model(&goods).Where("id = ?", goodsID).Updates(&goods).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Can't update data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success, data has updated"})

}

func DeleteGoods(c *gin.Context) {
	var goods models.Goods

	goodsID := c.Param("id")

	if database.DB.Delete(&goods, goodsID).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Can't delete data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success, data has deleted"})
}
