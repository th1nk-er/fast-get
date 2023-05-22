package main

import (
	"github.com/gin-contrib/cors"
	"github.com/th1nk-er/fast-get/server/routers"
)

func main() {
	r := routers.InitRouters()

	// cors middleware
	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}

	r.Use(cors.New(corsConfig))

	_ = r.Run(":8080")
}
