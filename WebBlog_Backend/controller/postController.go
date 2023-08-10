package controller

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"example.com/blog/database"
	"example.com/blog/models"
	"example.com/blog/util"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreatePost(c *fiber.Ctx) error {
	var blogPost models.Blog
	err := c.BodyParser(&blogPost)
	if err != nil {
		fmt.Println("Unable to parse body")
	}

	if err := database.DB.Create(&blogPost).Error; err != nil{
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid payload",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Congratulation, Your post is live",
	})
}

func AllPost(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 5
	offset := (page-1)*limit
	var total int64
	var getblog []models.Blog

	database.DB.Preload("User").Offset(offset).Limit(limit).Find(&getblog)
	database.DB.Model(&models.Blog{}).Count(&total)
	return c.JSON(fiber.Map{
		"data":getblog,
		"meta":fiber.Map{
			"total":total,
			"page":page,
			"last_page":math.Ceil(float64(int(total)/limit)),
		},
	})
}

func DetailPost(c *fiber.Ctx)error{
	id, _ := strconv.Atoi(c.Params("id"))
	var blogPost models.Blog
	database.DB.Where("id=?", id).Preload("User").First(&blogPost)
	return c.JSON(fiber.Map{
		"data":blogPost,
	})
}

func UpdatePost(c *fiber.Ctx)error{
	id, _ := strconv.Atoi(c.Params("id"))
	blog := models.Blog{
		Id:int(id),
	}
	err := c.BodyParser(&blog)
	if err != nil {
		fmt.Println("Unable to parse body")
	}
	database.DB.Model(&blog).Updates(blog)
	return c.JSON(fiber.Map{
		"message": "Post updated successfully",
	})
}

func UniquePost(c *fiber.Ctx)error{
	cookie := c.Cookies("jwt")
	id,_ := util.Parsejwt(cookie)
	var blog []models.Blog
	database.DB.Model(&blog).Where("user_id=?",id).Preload("User").Find(&blog)

	return c.JSON(blog)
}

func DeletePost(c *fiber.Ctx)error{
	id,_ := strconv.Atoi(c.Params("id"))
	blog := models.Blog{
		Id:int(id),
	}
	deleteQuery := database.DB.Delete(&blog)
	if errors.Is(deleteQuery.Error, gorm.ErrRecordNotFound){
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Opps!, record Not found",
		})
	}
	
	return c.JSON(fiber.Map{
		"message": "Delete Successfully",
	})
}