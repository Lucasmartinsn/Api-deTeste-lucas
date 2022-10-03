package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"github.com/go-chi/chi/v5"

	"github.com/Lucasmartinsn/ApiProjeto-da-tropa/models"
)


func Get(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	
	if err != nil{
		log.Printf("erro ao fazer parse do id %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	cadastro, err := models.Get(int(id))
	if err != nil{
		log.Printf("erro ao atualizar register: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Add("content-Type", "application/json")
	json.NewEncoder(w).Encode(cadastro)
}