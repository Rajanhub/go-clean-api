package repository

import "github.com/Rajanhub/goapi/infrastructure"

// PostRepository database structure
type PostRepository struct {
	infrastructure.Database
}

// NewPostRepository creates a new user repository
func NewPostRepository(db infrastructure.Database) PostRepository {
	return PostRepository{
		Database: db,
	}
}
