package main

import (
	"go-grpc-example/handlers"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"go-grpc-example/config"
)

func main() {

	g := gin.Default()

	config.LoadConfig()

	// api to get a single user
	g.POST("user-management/get-user", handlers.GetSingleUserHandler)

	// api for multiple users
	g.POST("user-management/get-users", handlers.GetMultipleUserHandler)

	// microservice is run and exposed at external port
	if err := g.Run(viper.GetString("external.port")); err != nil {
		log.Fatal().AnErr("Failed to run server : %v", err)

	}

}
