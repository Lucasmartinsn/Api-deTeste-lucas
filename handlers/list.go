package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Lucasmartinsn/ApiProjeto-da-tropa/models"
)

func List(w http.ResponseWriter, r *http.Request){
	cadastros, err := models.GetAll()
	if err != nil{
		log.Printf("Erro ao obter registros: %v", err)
	}

	w.Header().Add("content-Type", "application/json")
	json.NewEncoder(w).Encode(cadastros)
}