package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/caarlos0/env/v6"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"github.com/jasongauvin/zebi-scraper/api/database"
	"github.com/jasongauvin/zebi-scraper/api/routes"
)

func main() {
	cfg := database.Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	database.Connect(cfg)
	database.MakeMigrations()

	router := gin.Default()

	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Content-Type",
		Credentials:     true,
		ValidateHeaders: false,
	}))

	routes.InitializeRoutes(router)

	log.Fatal(router.Run(":8000"))
}
