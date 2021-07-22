package repository

type Transaction interface {
	Rollback() error
	Commit() error
}

type TransactionContextReader interface {
	GetContext() interface{}
}