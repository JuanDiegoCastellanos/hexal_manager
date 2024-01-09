package repository

import (
	"context"
	"hexal_manager/hexal_manager_v0.1.0/models"
)

type UserRepository interface {
	Repository
	// InsertUser(ctx context.Context, user *models.User) error
	// GetUserById(ctx context.Context, id string) (*models.User, error)
	// GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	// UpdateUser(ctx context.Context, user *models.User) error
	// DeteletUser(ctx context.Context, id string) error
	// ListAllUser(ctx context.Context, page uint64) ([]*models.User, error)
}

var userRepositoryImplementation UserRepository

func SetUserRepository(repository UserRepository) {
	userRepositoryImplementation = repository
}

func InsertUser(ctx context.Context, user *models.User) error {
	return userRepositoryImplementation.Insert(ctx, user)
}
func GetUserById(ctx context.Context, id string) (*models.User, error) {
	return userRepositoryImplementation.GetById(ctx, id)
}
func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return userRepositoryImplementation.GetUserByEmail(ctx, email)
}
func UpdateUser(ctx context.Context, user *models.User) error {
	return userRepositoryImplementation.UpdateUser(ctx, user)
}
func DeteletUser(ctx context.Context, id string) error {
	return userRepositoryImplementation.DeteletUser(ctx, id)
}
func ListAllUser(ctx context.Context, page uint64) ([]*models.User, error) {
	return userRepositoryImplementation.ListAllUser(ctx, page)
}
