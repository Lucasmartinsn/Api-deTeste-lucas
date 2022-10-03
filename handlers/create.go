package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Lucasmartinsn/ApiProjeto-da-tropa/models"
)

func Create(w http.ResponseWriter, r *http.Request){
	var cadastro models.Cadastro

	err := json.NewDecoder(r.Body).Decode(&cadastro)
	if err != nil{
		log.Printf("erro ao fazer decode do json %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(cadastro)

	var resp map[string] any
	
	if  err != nil {
		resp = map[string]any{
			"Error": true,
			"menssage": fmt.Sprintf("ocorreu um erro ao tenta inserir: %v ", err),
		}
	}else {
		resp = map[string]any{
			"Error": false,
			"menssage": fmt.Sprintf("scadastrado com sucesso! Id : %v ", id),
		}
	}
	
	w.Header().Add("Content-Type", "appllication/json")
	json.NewEncoder(w).Encode(resp)
	
}