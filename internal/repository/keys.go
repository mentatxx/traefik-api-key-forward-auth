package repository

import (
	"github.com/mentatxx/traefik-api-key-forward-auth/internal/database"
	"github.com/mentatxx/traefik-api-key-forward-auth/models"
)

type KeyFilter struct {
	OrganizationID *string `json:"organizationId,omitempty"`
	ProjectID      *string `json:"projectId,omitempty"`
	UserID         *string `json:"userId,omitempty"`
}

type KeysRepository interface {
	GetByFilter(filter KeyFilter, after *string, limit *float64) ([]models.Key, error)
	Create(key *models.Key) error
	MarkAsDeletedByID(id string) error
}

type KeysRepositoryImpl struct {
	db *database.Database
}

func NewKeysRepository(db *database.Database) KeysRepository {
	return &KeysRepositoryImpl{db: db}
}

// Retrieve keys from the database based on the provided filter
// Return the retrieved keys and any potential error
func (r *KeysRepositoryImpl) GetByFilter(filter KeyFilter, after *string, limit *float64) ([]models.Key, error) {
	var keys []models.Key

	chain := r.db.Where(filter).Order("id")
	if (after != nil) && (*after != "") {
		chain = chain.Where("id > ?", *after)
	}
	if (limit != nil) && (*limit > 0) {
		chain = chain.Limit(int(*limit))
	} else {
		chain = chain.Limit(50)
	}

	result := chain.Find(&keys)
	if result.Error != nil {
		return nil, result.Error
	} else {
		return keys, nil
	}
}

// Create a key in the database
func (r *KeysRepositoryImpl) Create(key *models.Key) error {
	result := r.db.Create(key)
	return result.Error
}

// Mark a key as deleted in the database
func (r *KeysRepositoryImpl) MarkAsDeletedByID(id string) error {
	result := r.db.Model(&models.Key{}).Where("id = ?", id).Update("deleted_at", "now()")
	return result.Error
}
