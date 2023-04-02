package entity

type MysqlEntityUser struct {
	Id       string `db:"id"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

type UserCredentials struct {
	Email    string
	Password string
}

type DomainUser struct {
	Id       string
	Email    string
	Password string
}
