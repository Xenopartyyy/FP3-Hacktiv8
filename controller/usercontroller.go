package controller

import (
	"FP3-Hacktiv8/dto"
	db "FP3-Hacktiv8/infra/database"
	"FP3-Hacktiv8/middleware"
	"FP3-Hacktiv8/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if db.DB.Where("email = ?", user.Email).Find(&user).RowsAffected > 0 {
		c.JSON(http.StatusConflict, gin.H{"message": "The email is already in use"})
		return
	}

	user.Role = "member"

	if err := user.HashPassword(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to hash password"})
		return
	}

	if err := db.DB.Create(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to create user"})
		return
	}

	userdto := dto.User{
		ID:         user.ID,
		Full_name:  user.Full_name,
		Email:      user.Email,
		Created_at: user.CreatedAt,
	}

	c.JSON(http.StatusCreated, userdto)
}

func Login(c *gin.Context) {

	var body struct {
		Email    string
		Password string
	}

	if err := c.Bind(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var user model.User

	db.DB.Find(&user, "email = ?", body.Email)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Incorrect email or password"})
		return
	}

	signedToken, err := middleware.GenerateToken(user.ID, user.Email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Incorrect email or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": signedToken,
	})
}

func Update(c *gin.Context) {
	userData, _ := c.Get("userData")
	userClaim, ok := userData.(jwt.MapClaims)

	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		return
	}

	userID := uint(userClaim["id"].(float64))

	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if db.DB.Model(&user).Where("id = ?", userID).Updates(&user).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Failed to update user"})
		return
	}

	userdto := dto.User{
		ID:         userID,
		Full_name:  user.Full_name,
		Email:      user.Email,
		Created_at: user.UpdatedAt,
	}

	c.JSON(http.StatusOK, userdto)
}

func Delete(c *gin.Context) {
	userData, _ := c.Get("userData")
	userClaim, ok := userData.(jwt.MapClaims)

	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		return
	}

	userID := uint(userClaim["id"].(float64))

	var user model.User

	if err := db.DB.Delete(&user, userID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your account has been succesfully deleted",
	})
}
