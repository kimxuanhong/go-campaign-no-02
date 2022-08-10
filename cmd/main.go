package main

import (
	"github.com/joho/godotenv"
	"github.com/kimxuanhong/go-campaign-no-02/pkg/handler"
	"github.com/labstack/gommon/log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Router := handler.NewRouter(":8080")
	Router.Start()
}
