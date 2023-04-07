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

// func A() gin.HandlerFunc {
// 	fmt.Println("is role admin?")
// 	return func(ctx *gin.Context) {
// 		db := database.GetDB()

// 		userData := ctx.MustGet("userData").(jwt.MapClaims)
// 		userID := uint(userData["id"].(float64))
// 		User := models.User{}
// 		Role := models.Role{}

// 		fmt.Println("user data:", userID)

// 		// err := db.Select("Roles").First(&User, uint(userID)).Preload("Roles").Error

// 		err := db.Preload("Roles", func(ctx *gorm.DB) *gorm.DB {
// 			return ctx.First(&Role, 2).Select("id")
// 		}).First(&User, uint(userID)).Error

// 		if err != nil {
// 			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
// 				"error":   "Data Roles not found",
// 				"message": "data roles didn't exist",
// 			})
// 			return
// 		}

// 		fmt.Println(User.Roles[len(User.Roles)-1].ID)

// 		if User.Roles == nil {
// 			fmt.Println("no user roles")
// 			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
// 				"error":   "Unauthorized",
// 				"message": "you are not allowed to access this services",
// 			})
// 			return
// 		}
// 		ctx.Next()
// 	}
// }

func IsRoleAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.GetDB()
		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		User := models.User{}
		Role := models.Role{}

		fmt.Println("user data", userData)

		// err = db.Select("user_id").First(&Product, uint(productID)).Error

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

		// if User.Roles[len(User.Roles)-1].ID != 2 {
		// 	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		// 		"error":   "Unauthorized",
		// 		"message": "you are not allowed to access this data",
		// 	})
		// 	return
		// }

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

		fmt.Println("user data", userData)

		err = db.Select("user_id").First(&Product, uint(productID)).Error

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data not found",
				"message": "data didn't exist",
			})
			return
		}

		if Product.UserID != userID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you are not allowed to access this data",
			})
			return
		}
		ctx.Next()
	}
}
