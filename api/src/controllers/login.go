package controllers

import (
	"api/src/database"
	"api/src/model"
	"api/src/repositories"
	"api/src/response"
	"api/src/secure"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Login é responsavel por autenticar o usuario na API
func Login(w http.ResponseWriter, r *http.Request) {
	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		response.Error(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user model.User
	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		response.Error(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repo := repositories.NewRepositoryUser(db)
	userSaveDB, erro := repo.FindByEmail(user.Email)
	if erro != nil {
		response.Error(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = secure.VerifyPass(userSaveDB.Pass, user.Pass); erro != nil {
		response.Error(w, http.StatusUnauthorized, erro)
		return
	}

}
