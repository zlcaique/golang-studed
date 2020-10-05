package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//StringConnectDB é a string de conexao com o Banco
	StringConnectDB = ""

	//Port onde a Api vai estar rodando
	Port = 0
)

//Load vai inicializar as variaveis de ambiente
func Load() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Port, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Port = 9000
	}

	StringConnectDB = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_HOST_PORT"),
		os.Getenv("DB_NAME"),
	)

}
