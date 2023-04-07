package controllers

import (
	"fmt"
	"net/http"
	"uc-shop/database"
	"uc-shop/helpers"
	"uc-shop/models"

	"github.com/gin-gonic/gin"
)

var (
	appJson = "application/json"
)

func GetUser(c *gin.Context) {
	db := database.GetDB()

	var user []models.User
	// User := models.User{}
	// Role := models.Role{}

	err := db.Preload("Roles").Preload("Products").Find(&user).Error

	// err := db.Preload("Roles", func(ctx *gorm.DB) *gorm.DB {
	// 	return ctx.Select("id")
	// }).Preload("Products").Find(&user).Error

	// err := db.Find(&user).Error

	// err := db.Joins("JOIN products ON products.user_id = users.id").Find(&user).Error

	// err := db.Table("users").
	// 	Select("users.id, users.first_name, users.last_name, users.email, users.password, products.title, products.description").
	// 	Joins("JOIN products ON products.user_id = users.id").
	// 	Find(&user).Error

	// err := db.Preload("Roles", func(ctx *gorm.DB) *gorm.DB {
	// 	return ctx.Omit("name", "description")
	// }).Preload("Products").Find(&user).Error

	// err := db.Preload("Roles", func(ctx *gorm.DB) *gorm.DB {
	// 	return ctx.Select("id").Where("id = ?", 2)
	// }).Preload("Products").First(&user, 2).Error

	// err := db.Preload("Roles", func(ctx *gorm.DB) *gorm.DB {
	// 	return ctx.Select("id").Where("id = ?", 3)
	// }).First(&User, 2).Error

	// err := db.Preload("Roles", func(ctx *gorm.DB) *gorm.DB {
	// 	return ctx.First(&Role, 2)
	// }).First(&User, 3).Error

	// type APIRole struct {
	// 	ID uint
	// }

	// err := db.Preload("Roles", func(ctx *gorm.DB) *gorm.DB {
	// 	return ctx.Find(&APIRole{})
	// }).Preload("Products").Find(&user).Error

	// fmt.Println(User.Roles[len(User.Roles)-1].ID)

	// if User.Roles[len(User.Roles)-1].ID != 2 {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"user": nil,
	// 	})
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{
	// 	"user": User,
	// })

	// c.JSON(http.StatusOK, gin.H{
	// 	"user": Role.ID,
	// })

	if err != nil {
		fmt.Println("Error getting user data:", err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})

}

func UserLogin(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}
	password := ""

	if contentType == appJson {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password
	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	// err := db.Debug().Where("email = ?", User.Email).Preload("Roles", func(ctx *gorm.DB) *gorm.DB {
	// 	return ctx.Select([]string{"role_id"})
	// }).Take(&User).Error

	// err := db.Debug().Where("email = ?", User.Email).Preload("Roles").Take(&User).Error

	// type APIRole struct {
	// 	ID uint
	// }

	// err := db.Debug().Where("email = ?", User.Email).Preload("Roles", func(ctx *gorm.DB) *gorm.DB {
	// 	return ctx.Find(&APIRole{})
	// }).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorize",
			"message": "Invalid email/password",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorize",
			"message": "Invalid email/password",
		})
		return
	}

	token := helpers.TokenGenerator(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}

	if contentType == appJson {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":        User.ID,
		"email":     User.Email,
		"full_name": User.FirstName + " " + User.LastName,
		"password":  User.Password,
		"roles":     User.Roles,
	})
}
