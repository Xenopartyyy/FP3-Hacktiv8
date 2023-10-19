package controller

import (
	"FP3-Hacktiv8/dto"
	db "FP3-Hacktiv8/infra/database"
	"FP3-Hacktiv8/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Create(c *gin.Context) {
	var category model.Category
	userData, _ := c.Get("userData")
	userClaim, ok := userData.(jwt.MapClaims)

	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		return
	}

	userROLE := userClaim["role"].(string)

	if userROLE != "admin" {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "You don't have permission to perform this action"})
		return
	}

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
