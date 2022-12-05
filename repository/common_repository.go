package repository

import (
	"github.com/Rajanhub/goapi/infrastructure"
	"github.com/Rajanhub/goapi/lib"
	"gorm.io/gorm"
)

// PostRepository database structure
type Repository struct {
	infrastructure.Database
	logger lib.Logger
}

// NewPostRepository creates a new user repository
func NewRepository(db infrastructure.Database, logger lib.Logger) Repository {
	return Repository{
		Database: db,
		logger:   logger,
	}
}

// WithTrx delegate transaction from user repository
func (r Repository) WithTrx(trxHandle *gorm.DB) Repository {
	if trxHandle != nil {
		r.logger.Debug("using WithTrx as trxHandle is not nil")
		r.Database.DB = trxHandle
	}
	return r
}
