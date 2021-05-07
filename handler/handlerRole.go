package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project-individu-go-react/entity"
)

func CreateNewRole(c *gin.Context) {
	var getRole entity.Role

	if err := c.ShouldBindJSON(&getRole); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":        "error bad request",
			"message_error": err.Error(),
		})
		return
	}

	var newRole = entity.Role{
		ID:       getRole.ID,
		Admin:    getRole.Admin,
		User:     getRole.User,
		Psikolog: getRole.Psikolog,
	}

	if err := DB.Create(&newRole).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":        "error in internal server",
			"message_error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, newRole)

}

func GetAllRole(c *gin.Context) {
	var roles []entity.Role

	if err := DB.Find(&roles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":        "error in internal server",
			"message_error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, roles)
}

func GetRoleByID(c *gin.Context) {
	var GetRoleID entity.Role

	id := c.Param("role_id")

	if err := DB.Where("id = ?", id).Preload("Users").Find(&GetRoleID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":        "error bad request",
			"message_error": err.Error(),
		})
		return
	}

	if GetRoleID.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":        "error not found",
			"message_error": "user id " + id + " not found in database",
		})
		return
	}

	c.JSON(http.StatusOK, GetRoleID)
}

func UpdateRoleByID(c *gin.Context) {
	var updateRole entity.Role

	id := c.Params.ByName("role_id")

	if err := DB.Where("id = ?", id).Preload("Users").Find(&updateRole).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":        "error in internal server",
			"message_error": err.Error(),
		})
		return
	}

	if updateRole.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":        "error not found",
			"message_error": "user id " + id + " not found in database",
		})
		return
	}

	if err := c.ShouldBindJSON(&updateRole); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":        "error bad request",
			"message_error": err.Error(),
		})
		return
	}

	var newRole = entity.Role{
		Admin:    updateRole.Admin,
		User:     updateRole.User,
		Psikolog: updateRole.Psikolog,
	}

	if err := DB.Where("id = ?", id).Save(&newRole).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":        "error in internal server",
			"message_error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updateRole)
}

func DeleteByRoleID(c *gin.Context) {
	id := c.Param("role_id")

	var DeleteRole entity.Role

	DB.Where("id = ?", id).Find(&DeleteRole)

	if DeleteRole.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":        "error not found",
			"message_error": "user id " + id + " not found in database",
		})
		return
	}

	if err := DB.Where("id = ?", id).Delete(&entity.Role{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":        "error in internal server",
			"message_error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "delete success",
		"message": "user id " + id + " successfull delete",
	})
}
