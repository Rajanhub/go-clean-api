package services

import (
	"github.com/Rajanhub/goapi/lib"
	"github.com/Rajanhub/goapi/models"
	"github.com/Rajanhub/goapi/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserService struct {
	repository      repository.Repository
	paginationScope *gorm.DB
}

func NewUserService(
	repository repository.Repository,
) *UserService {
	return &UserService{
		repository: repository,
	}
}

// WithTrx delegates transaction to repository database
func (s UserService) WithTrx(trxHandle *gorm.DB) UserService {
	s.repository = s.repository.WithTrx(trxHandle)
	return s
}

// PaginationScope
func (s UserService) SetPaginationScope(scope func(*gorm.DB) *gorm.DB) UserService {
	s.paginationScope = s.repository.WithTrx(s.repository.Scopes(scope)).DB
	return s
}

func (s UserService) Create(user *models.User) error {
	return s.repository.Create(&user).Error
}

func (s UserService) GetAllUser() (response map[string]interface{}, err error) {
	var users []models.User
	var count int64

	err = s.repository.WithTrx(s.paginationScope).Find(&users).Offset(-1).Limit(-1).Count(&count).Error
	if err != nil {
		return nil, err
	}

	return gin.H{"data": users, "count": count}, nil
}

// GetOneUser gets one user
func (s UserService) GetOneUser(userID lib.BinaryUUID) (user models.User, err error) {
	return user, s.repository.First(&user, "id = ?", userID).Error
}

func (s UserService) UpdateUser(user *models.User) error {
	return s.repository.Save(&user).Error
}

// DeleteUser deletes the user
func (s UserService) DeleteUser(uuid lib.BinaryUUID) error {
	return s.repository.Where("id = ?", uuid).Delete(&models.User{}).Error
}
