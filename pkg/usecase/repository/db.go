package repository

type DBRepository interface {
	TransactionTX(fn func(interface{}) (interface{}, error)) (interface{}, error)
}
