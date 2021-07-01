package repository

type Transaction interface {
	Rollback() error
	Commit() error
}
