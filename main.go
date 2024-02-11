package main

import (
	"os"

	"github.com/NikolaJankovikj/go-rest-api/database"
	"github.com/joho/godotenv"
)

func main() {
	err := initApp()
	if err != nil {
		panic(err)
	}

	defer database.CloseMongoDB()

	app := generateApp()

	port := os.Getenv("PORT")

	app.Listen(":" + port)
}

func initApp() error {
	err := loadEnv()
	if err != nil {
		return err
	}

	err = database.StartMongoDB()

	if err != nil {
		return err
	}

	return nil
}

func loadEnv() error {
	goEnv := os.Getenv("GO_ENV")
	if goEnv == "" {
		err := godotenv.Load()
		if err != nil {
			return err
		}
	}
	return nil
}
