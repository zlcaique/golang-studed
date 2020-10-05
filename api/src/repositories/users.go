package repositories

import (
	"api/src/model"
	"database/sql"
)

//Users representa um repositorio de usuarios
type Users struct {
	db *sql.DB
}

//NewRepositoryUser cria um repositorio de usuarios
func NewRepositoryUser(db *sql.DB) *Users {
	return &Users{db}
}

// Create insere um usuario no banco de dados
func (repo Users) Create(user model.User) (uint64, error) {
	statement, erro := repo.db.Prepare("insert into users (name,nick,email,pass) values (?, ?, ?, ?)")
	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	result, erro := statement.Exec(user.Name, user.Nick, user.Email, user.Pass)
	if erro != nil {
		return 0, erro
	}

	lastID, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(lastID), nil

}
