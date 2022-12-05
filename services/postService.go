package services

import (
	"github.com/Rajanhub/goapi/lib"
	"github.com/Rajanhub/goapi/models"
	"github.com/Rajanhub/goapi/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostService struct {
	repository      repository.Repository
	paginationScope *gorm.DB
}

func NewPostService(
	postRepository repository.Repository,
) *PostService {
	return &PostService{
		repository: postRepository,
	}
}

// WithTrx delegates transaction to repository database
func (s PostService) WithTrx(trxHandle *gorm.DB) PostService {
	s.repository = s.repository.WithTrx(trxHandle)
	return s
}

// PaginationScope
func (s PostService) SetPaginationScope(scope func(*gorm.DB) *gorm.DB) PostService {
	s.paginationScope = s.repository.WithTrx(s.repository.Scopes(scope)).DB
	return s
}

func (s PostService) Create(post *models.Post) error {
	return s.repository.Create(&post).Error
}

func (s PostService) GetAllPost() (response map[string]interface{}, err error) {
	var posts []models.Post
	var count int64

	err = s.repository.Preload("User").Find(&posts).Offset(-1).Limit(-1).Count(&count).Error
	if err != nil {
		return nil, err
	}

	return gin.H{"data": posts, "count": count}, nil
}

// GetOnePost gets one post
func (s PostService) GetOnePost(postID lib.BinaryUUID) (post models.Post, err error) {
	return post, s.repository.First(&post, "id = ?", postID).Error
}

func (s PostService) UpdatePost(post *models.Post) error {
	return s.repository.Save(&post).Error
}

// DeletePost deletes the post
func (s PostService) DeletePost(uuid lib.BinaryUUID) error {
	return s.repository.Where("id = ?", uuid).Delete(&models.Post{}).Error
}
