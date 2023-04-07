package middleware

import (
	"fmt"
	"net/http"
	"strconv"
	"uc-shop/database"
	"uc-shop/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func IsRoleAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.GetDB()
		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		User := models.User{}
		Role := models.Role{}

		err := db.Preload("Roles", func(ctx *gorm.DB) *gorm.DB {
			return ctx.First(&Role, 2).Select("id")
		}).First(&User, uint(userID)).Error

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Roles not found",
				"message": "data Roles didn't exist",
			})
			return
		}

		if len(User.Roles) == 0 {
			fmt.Println("no user roles")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you are not allowed to access this services",
			})
			return
		}

		ctx.Next()
	}
}

func ProductAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.GetDB()
		productID, err := strconv.Atoi(ctx.Param("productID"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "bad request",
				"message": "invalid parameter",
			})
			return
		}
		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Product := models.Product{}
		User := models.User{}
		Role := models.Role{}

		err = db.Select("user_id").First(&Product, uint(productID)).Error

		err2 := db.Preload("Roles", func(ctx *gorm.DB) *gorm.DB {
			return ctx.First(&Role, 2).Select("id")
		}).First(&User, uint(userID)).Error

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data not found",
				"message": "data didn't exist",
			})
			return
		}

		if err2 != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data not found",
				"message": "data didn't exist",
			})
			return
		}

		if Product.UserID == userID || len(User.Roles) != 0 {
			ctx.Next()
			return
		} else {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you are not allowed to access this data",
			})
		}
	}
}
