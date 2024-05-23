package main

import (
	"net/http"

	handler "github.com/putragabrielll/fwg17-cinematix-be/api"
)

// func noLink(c *gin.Context) {
// 	c.JSON(http.StatusNotFound, &services.ResponseBack{
// 		Success: false,
// 		Message: "Resource not found!",
// 	})
// }

func main() {
	// r := gin.Default()
	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"http://localhost:5173", "https://astonishing-cheesecake-81f941.netlify.app", "http://cinematix.pasukanhosting.my.id", "https://cinematix-app.netlify.app"},
	// 	AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE"},
	// 	AllowHeaders:     []string{"Authorization", "Content-Type", "Accept", "access-control-allow-origin", "access-control-allow-headers"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// }))
	// routers.Combine(r)
	// r.NoRoute(noLink)
	// r.Run(":8080")
	http.HandleFunc("/", handler.Handler)
}
