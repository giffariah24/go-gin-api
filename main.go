package main

import (
	"github.com/giffariah666/go-gin-api/controllers/filmcontroller"
	"github.com/giffariah666/go-gin-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDB()

	r.GET("/films", filmcontroller.GetFilm)
	r.GET("/film/:id", filmcontroller.GetFilmById)
	r.POST("/films", filmcontroller.CreateFilm)
	r.PUT("/film/:id", filmcontroller.UpdateFilm)
	r.DELETE("/films", filmcontroller.DeleteFilm)

	r.Run(":8081")
}
