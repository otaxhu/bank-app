package entity

// Esto representa el user que debe devolver cualquier implementacion de UsersRepository
type RepositoryUser struct {
	Id       string
	Email    string
	Password string
	Roles    []string
}

// Esto representa la tabla user de Mysql
type MysqlEntityUser struct {
	Id       string `db:"id"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

// Esto representa la tabla user_roles de Mysql
type MysqlEntityUserRole struct {
	Id     string `db:"id"`
	Role   string `db:"role"`
	UserId string `db:"user_id"`
}

// Esto representa las credenciales del usuario para registrarse o iniciar sesion
type UserCredentials struct {
	Email    string
	Password string
}

// Esto representa el user en la aplicacion de mi dominio
type DomainUser struct {
	Id    string
	Roles []string
}
