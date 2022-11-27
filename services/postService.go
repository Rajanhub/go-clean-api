package services

import (
	"github.com/Rajanhub/goapi/models"
	"github.com/Rajanhub/goapi/repository"
	"github.com/gin-gonic/gin"
)

type PostService struct {
	repository repository.PostRepository
}

func NewPostService(
	postRepository repository.PostRepository,
) *PostService {
	return &PostService{
		repository: postRepository,
	}
}

func (s PostService) Create(post *models.Post) error {
	return s.repository.Create(&post).Error
}

func (s PostService) GetAllPost() (response map[string]interface{}, err error) {
	var posts []models.Post
	var count int64

	err = s.repository.Find(&posts).Offset(-1).Limit(-1).Count(&count).Error
	if err != nil {
		return nil, err
	}

	return gin.H{"data": posts, "count": count}, nil
}
