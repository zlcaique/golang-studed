package database

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver Mysql
)

//Connect abre a aconexao com o banco de dados e retorna
func Connect() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConnectDB)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}
