package repository

import "context"

type DBRepository interface {
	TransactionTX(ctx context.Context, fn func(interface{}) (interface{}, error)) (interface{}, error)
}
