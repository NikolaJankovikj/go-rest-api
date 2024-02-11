package handlers

import (
	"context"
	"github.com/NikolaJankovikj/go-rest-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/NikolaJankovikj/go-rest-api/database"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type gameDTO struct {
	Title     string `json:"title" bson:"title"`
	Genre     string `json:"genre" bson:"genre"`
	Developer string `json:"developer" bson:"developer"`
}

func GetGames(c *fiber.Ctx) error {
	gamesCollection := database.GetCollection("games")
	cursor, err := gamesCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		return err
	}

	var games []models.Game

	if err = cursor.All(context.TODO(), &games); err != nil {
		return err
	}

	return c.JSON(games)
}

func CreateGame(c *fiber.Ctx) error {
	newGame := new(gameDTO)

	if err := c.BodyParser(newGame); err != nil {
		return err
	}

	gamesCollection := database.GetCollection("games")
	newDoc, err := gamesCollection.InsertOne(context.TODO(), newGame)

	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"id": newDoc.InsertedID})
}

func DeleteGame(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ObjectID")
	}

	gamesCollection := database.GetCollection("games")

	_, err = gamesCollection.DeleteOne(context.TODO(), bson.M{"_id": objID})

	if err != nil {
		return err
	}

	return c.SendString("Game deleted successfully")
}

func UpdateGame(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ObjectID")
	}

	updates := new(gameDTO)
	if err = c.BodyParser(updates); err != nil {
		return err
	}

	gamesCollection := database.GetCollection("games")

	update := bson.M{}

	set := update["$set"].(bson.M)
	if updates.Title != "" {
		set["title"] = updates.Title
	}
	if updates.Genre != "" {
		set["genre"] = updates.Genre
	}
	if updates.Developer != "" {
		set["developer"] = updates.Developer
	}

	if len(update) == 0 {
		return c.Status(fiber.StatusBadRequest).SendString("No fields provided for update")
	}

	_, err = gamesCollection.UpdateOne(context.TODO(), bson.M{"_id": objID}, update)
	if err != nil {
		return err
	}

	return c.SendString("Game updated successfully")
}
