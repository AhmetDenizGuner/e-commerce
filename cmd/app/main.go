package main

import (
	"adg/e-commerce/internal/api"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	api.RegisterHandlers(r)

	srv := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	log.Println("Server is running...")

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()
}
