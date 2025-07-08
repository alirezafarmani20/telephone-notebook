package api

import (
	"fmt"
	"log"
	"telephone/internal/database"
	"telephone/internal/modules"

	"github.com/gofiber/fiber/v2"
)

// create CRUD api for user

func GetUsers(ctx *fiber.Ctx) error {
	// get all users
	var users []modules.User
	database.Database.Find(&users)
	return ctx.JSON(users)
}

func GetUserById(ctx *fiber.Ctx) error {
	// get user by id
	id := ctx.Params("id")
	var user modules.User
	if err := database.Database.First(id, &user).Error; err != nil {
		log.Fatal("error! user not found", err)
		return ctx.SendStatus(404)
	}
	database.Database.Select(&user)
	return ctx.JSON(user)
}

func CreateUser(ctx *fiber.Ctx) error {
	// create use
	user := new(modules.User)
	if err := ctx.BodyParser(user); err != nil {
		log.Fatal("error user is not created", err)
		return ctx.SendStatus(404)
	}
	database.Database.Create(&user)
	return ctx.JSON(user)
}

func UpdateUser(ctx *fiber.Ctx) error {
	// updating user
	id := ctx.Params("id")
	var user modules.User
	if err := database.Database.First(id, &user).Error; err != nil {
		log.Fatal("error ! user not updated", err)
		return ctx.SendStatus(404)
	}
	if err := ctx.BodyParser(&user); err != nil {
		fmt.Println("update user succesfuly")
		return ctx.SendStatus(400)
	}
	database.Database.Save(&user)
	return ctx.JSON(user)
}

func DeleteUser(ctx *fiber.Ctx) error {
	// delete user
	id := ctx.Params("id")
	var user modules.User
	if err := database.Database.First(id, &user).Error; err != nil {
		log.Fatal("error! user not deleted")
		return ctx.SendStatus(404)
	}
	database.Database.Delete(&user)
	return ctx.SendString("user deleted succesfuly")
}
