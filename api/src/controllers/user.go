package controllers

import (
	"api/src/database"
	"api/src/model"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//CreateUser insere um usuario no banco de dados
func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var user model.User
	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		log.Fatal(erro)
	}

	db, erro := database.Connect()
	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close()

	repo := repositories.NewRepositoryUser(db)
	userID, erro := repo.Create(user)
	if erro != nil {
		log.Fatal(erro)
	}

	w.Write([]byte(fmt.Sprintf("Id inserido: %d", userID)))

}

//FindUsers busca todos os usuarios do banco de dados
func FindUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os usuarios"))
}

//FindUser busca um usuario no banco de dados
func FindUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscar um usuario"))
}

//UpdateUser atualiza um usuario no banco de dados
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizar usuario"))
}

//DeleteUser deleta um usuario no banco de dados
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletar usuario"))
}
