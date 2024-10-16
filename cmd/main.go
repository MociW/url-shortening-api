package main

import (
	"fmt"
	"log"
	"url-shortening-api/internal/config"
	"url-shortening-api/internal/database"
	"url-shortening-api/internal/http"
	"url-shortening-api/internal/link"
	"url-shortening-api/internal/user"
)

func main() {
	config := config.NewViper()
	db := database.NewDB(config)
	userRepository := user.NewUserRepository(db)
	linkRepository := link.NewLinkRepository(db)

	userService := user.NewUserService(userRepository)
	linkService := link.NewLinkService(linkRepository, userRepository)

	userController := user.NewUserController(userService)
	linkController := link.NewLinkController(linkService)

	routeConfig := http.RouteConfig{
		UserController: userController,
		LinkController: linkController,
	}

	app := routeConfig.NewRouter()
	webPort := config.GetInt("PORT")
	err := app.Listen(fmt.Sprintf(":%d", webPort))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
