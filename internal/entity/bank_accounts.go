package entity

type RepositoryBankAccount struct {
	Id       string
	Balance  float64
	Currency string
}

type MysqlEntityBankAccount struct {
	Id       string  `db:"id"`
	UserId   string  `db:"user_id"`
	Balance  float64 `db:"balance"`
	Currency string  `db:"currency"`
}

type DomainBankAccount struct {
	Id       string
	Balance  float64
	Currency string
}

type BankAccount struct {
	UserId   string
	Currency string
}
