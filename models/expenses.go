package models

type expenses struct {
        data map[int]int
}

// TODO: instantiate from transaction model
func NewExpenses(d []int) Transactions {
        return NewTransactions(d)
}