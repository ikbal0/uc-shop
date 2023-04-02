package controllers

import (
	"net/http"
	"strconv"
	"uc-shop/database"
	"uc-shop/helpers"
	"uc-shop/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func UpdateProduct(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(ctx)
	Product := models.Product{}
	productID, _ := strconv.Atoi(ctx.Param("productID"))
	userID := uint(userData["id"].(float64))

	if contentType == appJson {
		ctx.ShouldBindJSON(&Product)
	} else {
		ctx.ShouldBind(&Product)
	}

	Product.UserID = userID
	Product.ID = uint(productID)

	err := db.Model(&Product).Where("id = ?", productID).Updates(models.Product{Title: Product.Title, Description: Product.Description}).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, Product)
}

func CreateProduct(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(ctx)
	Product := models.Product{}
	userID := uint(userData["id"].(float64))

	if contentType == appJson {
		ctx.ShouldBindJSON(&Product)
	} else {
		ctx.ShouldBind(&Product)
	}

	Product.UserID = userID

	err := db.Debug().Create(&Product).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, Product)
}
