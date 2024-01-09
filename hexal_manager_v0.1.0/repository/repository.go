package repository

import "context"

type Repository interface {
	Insert(ctx context.Context, model interface{}) error
	GetById(ctx context.Context, id interface{}) (interface{}, error)
	GetByProperti(ctx context.Context, properti interface{}) (interface{}, error)
	Update(ctx context.Context, model interface{}) error
	Delete(ctx context.Context) error
	Close() error
}

var implementationRepository Repository

func SetRepository(repository Repository) {
	implementationRepository = repository
}
func Close() error {
	return implementationRepository.Close()
}
func Insert(ctx context.Context, model interface{}) error {
	return implementationRepository.Insert(ctx, model)
}
func GetById(ctx context.Context, id interface{}) (interface{}, error) {
	return implementationRepository.GetById(ctx, id)
}
func GetByProperti(ctx context.Context, properti interface{}) (interface{}, error) {
	return implementationRepository.GetByProperti(ctx, properti)
}
func Update(ctx context.Context, model interface{}) error {
	return implementationRepository.Update(ctx, model)
}
func Delete(ctx context.Context) error {
	return implementationRepository.Delete(ctx)
}
