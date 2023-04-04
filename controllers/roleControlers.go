package controllers

import (
	"net/http"
	"uc-shop/database"
	"uc-shop/helpers"
	"uc-shop/models"

	"github.com/gin-gonic/gin"
)

func SetRole(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	_, _ = db, contentType
	Role := models.Role{}

	if contentType == appJson {
		ctx.ShouldBindJSON(&Role)
	} else {
		ctx.ShouldBind(&Role)
	}

	err := db.Debug().Create(&Role).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":          Role.ID,
		"name":        Role.Name,
		"description": Role.Description,
	})
}

func AddRole(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	_, _ = db, contentType
	Role := models.Role{}

	if contentType == appJson {
		ctx.ShouldBindJSON(&Role)
	} else {
		ctx.ShouldBind(&Role)
	}

	err := db.Debug().Create(&Role).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":          Role.ID,
		"name":        Role.Name,
		"description": Role.Description,
	})
}
