package repository

type Transaction interface {
	Rollback() error
	Commit() error
}

type TransactionContextReader interface {
	GetContext() interface{}
}

type TransactionStack struct {
	committed bool
	transactions []Transaction
}

func NewTransactionStack() *TransactionStack{
	return &TransactionStack{
		committed: false,
	}
}

func (ts *TransactionStack) Add(t Transaction)  {
	ts.transactions = append(ts.transactions, t)
}

func (ts *TransactionStack) Rollback()  {
	for _, t := range ts.transactions {
		t.Rollback()
	}
}

func (ts *TransactionStack) Commit()  {
	for _, t := range ts.transactions {
		t.Commit()
	}
	ts.committed = true
}

func (ts *TransactionStack) RollbackUnlessCommitted() {
	if !ts.committed {
		ts.Rollback()
	}
}