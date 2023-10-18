package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/programticreviews/golang-gin-poc/entity"
	"gitlab.com/programticreviews/golang-gin-poc/service"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}
type Controller struct {
	service service.VideoService
}

// FindAll implements VideoController.
func (c *Controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

// Save implements VideoController.
func (c *Controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video
}

func New(service service.VideoService) VideoController {
	return &Controller{
		service: service,
	}
}
