package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Product struct {
	Id        string   `json:"_id" bson:"_id"`
	Address   string   `json:"address" bson:"address"`
	CreatedBy string   `json:"createdBy" bson:"createdBy"`
	Email     string   `json:"email" bson:"email"`
	Gender    string   `json:"gender" bson:"gender"`
	Imgs      []string `json:"imgs" bson:"imgs"`
	Item      string   `json:"item" bson:"item"`
	Name      string   `json:"name" bson:"name"`
	Phone     string   `json:"phone" bson:"phone"`
	Price     int      `json:"price" bson:"price"`
	Status    string   `json:"status" bson:"status"`
	Tags      []string `json:"tags" bson:"tags"`
}

func main() {
	godotenv.Load()
	pw := os.Getenv("PP")
	str := fmt.Sprintf("mongodb+srv://mpoapostolis:%s@cluster0.mrxks.mongodb.net", pw)

	var ctx = context.TODO()
	clientOptions := options.Client().ApplyURI(str)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}
	app := fiber.New(
		fiber.Config{
			Prefork:       true,
			CaseSensitive: true,
			BodyLimit:     20 * 1024 * 1024,
			StrictRouting: true,
			ServerHeader:  "Fiber",
			AppName:       "Thrift",
		})
	app.Use(logger.New())
	app.Use(cors.New())

	app.Get("/0d517520c1f6878/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		docID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			log.Println("Invalid id")
		}

		var result Product
		collection := client.Database("thrift").Collection("products")
		if err = collection.FindOne(ctx, bson.M{"_id": docID}).Decode(&result); err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)
		return c.Status(200).JSON(result)
	})

	app.Post("/", func(c *fiber.Ctx) error {
		fileheader, err := c.FormFile("picture")
		if err != nil {
			return c.SendStatus(fiber.StatusBadGateway)
		}

		file, err := fileheader.Open()
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		defer file.Close()

		buffer, err := io.ReadAll(file)
		if err != nil {
			return c.SendStatus(fiber.StatusBadGateway)
		}

		errDir := createFolder("uploads")
		if errDir != nil {
			return c.SendStatus(fiber.StatusBadGateway)

		}

		filename, err := imageProcessing(buffer, 40, "uploads")
		if err != nil {
			return c.SendStatus(fiber.StatusBadGateway)

		}
		return c.JSON(fiber.Map{
			"picture": "/uploads/" + filename,
		})
	})

	app.Post("/", func(c *fiber.Ctx) error {
		fileheader, err := c.FormFile("picture")
		if err != nil {
			return c.SendStatus(fiber.StatusBadGateway)

		}

		file, err := fileheader.Open()
		if err != nil {
			return c.SendStatus(fiber.StatusBadGateway)

		}
		defer file.Close()

		buffer, err := io.ReadAll(file)
		if err != nil {
			return c.SendStatus(fiber.StatusBadGateway)

		}

		errDir := createFolder("uploads")
		if errDir != nil {
			return c.SendStatus(fiber.StatusBadGateway)

		}

		filename, err := imageProcessing(buffer, 40, "uploads")
		if err != nil {
			return c.SendStatus(fiber.StatusBadGateway)

		}
		return c.JSON(fiber.Map{
			"picture": "/uploads/" + filename,
		})
	})

	app.Listen(":3939")
}
