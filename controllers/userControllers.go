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

	// err := db.Find(&user).Error

	// err := db.Joins("JOIN products ON products.user_id = users.id").Find(&user).Error

	// err := db.Table("users").
	// 	Select("users.id, users.first_name, users.last_name, users.email, users.password, products.title, products.description").
	// 	Joins("JOIN products ON products.user_id = users.id").
	// 	Find(&user).Error

	err := db.Preload("Roles").Preload("Products").Find(&user).Error

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

	err := db.Debug().Where("email = ?", User.Email).Preload("Roles").Take(&User).Error

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

	token := helpers.TokenGenerator(User.ID, User.Email, User.Roles)

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
