package service

import (
	"gitlab.com/programticreviews/golang-gin-poc/entity"
)

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

// FindAll implements VideoServece.
func (service *videoService) FindAll() []entity.Video {
	return service.videos
}

// Save implements VideoServece.
func (service *videoService) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}

func New() VideoService {
	return &videoService{}
}
