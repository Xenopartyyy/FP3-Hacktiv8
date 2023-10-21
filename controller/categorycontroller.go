package controller

import (
	"FP3-Hacktiv8/dto"
	db "FP3-Hacktiv8/infra/database"
	"FP3-Hacktiv8/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	var category model.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := db.DB.Create(&category).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to create category"})
		return
	}

	categorydto := dto.Category{
		ID:        category.ID,
		Type:      category.Type,
		CreatedAt: category.CreatedAt,
	}

	c.JSON(http.StatusCreated, categorydto)
}

// func GetCategory(c *gin.Context) {
// 	var category model.Category

// 	if err := db.DB.Find(category); err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Data doesnt exist!"})
// 	}

// 	response := dto.Category{}

// 	for_, category := range err

// 	c.JSON(http.StatusOK, categorydto)
// }

func PatchCategory(c *gin.Context) {
	categoryID, err := strconv.Atoi(c.Param("categoryID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid categoryID"})
		return
	}
	var category model.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if db.DB.Model(&category).Where("id = ?", categoryID).Updates(&category).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Failed to update category"})
		return
	}

	userdto := dto.Category{
		ID:        uint(categoryID),
		Type:      category.Type,
		CreatedAt: category.UpdatedAt,
	}

	c.JSON(http.StatusOK, userdto)
}

func DeleteCategory(c *gin.Context) {

	categoryID, err := strconv.Atoi(c.Param("categoryID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid categoryID"})
		return
	}

	var category model.Category

	if err := db.DB.Where("id = ?", categoryID).Find(&category).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Category not found"})
		return
	}

	if err := db.DB.Delete(&category, categoryID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Category has been succesfully deleted",
	})
}
