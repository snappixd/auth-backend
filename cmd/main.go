package main

import (
	"auth-back/internal/config"
	"auth-back/internal/handler"
	"auth-back/internal/repository"
	"auth-back/internal/server"
	"auth-back/internal/service"
	"auth-back/pkg/auth"
	"auth-back/pkg/db/mongodb"
	"auth-back/pkg/hash"

	"log"
)

func main() {
	cfg := config.InitConfig()

	mongoClient, err := mongodb.NewClient(cfg)
	if err != nil {
		log.Println("ERROR!!! " + err.Error())
	}

	db := mongoClient.Database(cfg.Mongo.DBName)

	hasher := hash.NewSHA1Hasher(cfg.Auth.PasswordSalt)

	tokenManager, err := auth.NewManager(cfg.Auth.JWT.SigningKey)
	if err != nil {
		log.Println(err.Error())
	}

	repos := repository.NewRepositories(db)
	services := service.NewServices(service.Deps{
		Repos:        repos,
		Hasher:       hasher,
		TokenManager: tokenManager,
		TokenTTL:     cfg.Auth.JWT.TokenTTL,
	})
	handlers := handler.NewHandler(services, tokenManager)

	srv := server.NewServer(cfg, handlers.InitHandler(cfg))
	if err := srv.Run(); err != nil {
		log.Println(err.Error())
	}
}
