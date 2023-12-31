package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"gitlab.com/programticreviews/golang-gin-poc/controller"
	"gitlab.com/programticreviews/golang-gin-poc/middlewares"
	"gitlab.com/programticreviews/golang-gin-poc/service"
)
var (
	videoService service.VideoService = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput(){
	f, _ := os.Create("gin.log")
    gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
func main(){

	setupLogOutput()

	server := gin.New()

	server.Use(gin.Recovery() , middlewares.Logger())

	server.GET("/videos",func(ctx *gin.Context) {
		ctx.JSON(200,videoController.FindAll())
	})
	server.POST("/videos",func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})
	server.Run(":8080")
}