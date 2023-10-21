package controller

import (
	"FP3-Hacktiv8/dto"
	db "FP3-Hacktiv8/infra/database"
	"FP3-Hacktiv8/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func GetIdToken(c *gin.Context) uint {

	userData, _ := c.Get("userData")
	tokenClaims := userData.(jwt.MapClaims)

	userID := uint(tokenClaims["id"].(float64))

	return userID
}

func CreateTask(c *gin.Context) {
	var (
		task     model.Task
		category model.Category
	)

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	task.UserID = GetIdToken(c)
	categoryID := task.CategoryID

	if err := db.DB.First(&category, categoryID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "Kategori tidak ditemukan"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal mencari kategori"})
		return
	}

	task.Status = false

	if err := db.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal membuat tugas", "error": err.Error()})
		return
	}

	taskdto := dto.Task{
		ID:          task.ID,
		Title:       task.Title,
		Status:      task.Status,
		Description: task.Description,
		UserID:      task.UserID,
		CategoryID:  task.CategoryID,
		CreatedAt:   task.CreatedAt,
	}

	c.JSON(http.StatusCreated, taskdto)
}

func PutPatchTask(c *gin.Context) {
	var task model.Task

	userID := GetIdToken(c)

	taskID, err := strconv.Atoi(c.Param("taskID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid taskID"})
		return
	}

	if err := db.DB.First(&task, taskID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve task"})
		return
	}

	if task.UserID != userID {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "this task not belong to you"})
		return
	}

	if err := c.ShouldBindJSON(&task); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if db.DB.Model(&task).Where("id = ?", taskID).Updates(&task).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Failed to update category"})
		return
	}

	taskdto := dto.UpdateTask{
		ID:          task.ID,
		Title:       task.Title,
		Status:      task.Status,
		Description: task.Description,
		UserID:      task.UserID,
		CategoryID:  task.CategoryID,
		UpdatedAt:   task.UpdatedAt,
	}

	c.JSON(http.StatusOK, taskdto)

}

func DeleteTask(c *gin.Context) {
	var task model.Task
	userID := GetIdToken(c)

	taskID, err := strconv.Atoi(c.Param("taskID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid taskID"})
		return
	}

	if err := db.DB.Where("id = ?", taskID).Find(&task).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Taskid not found"})
		return
	}

	if userID != task.UserID {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "this task not belong to you"})
		return
	}

	if err := db.DB.Delete(&task, taskID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete task"})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Task has been succesfully deleted",
	})

}
