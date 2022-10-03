package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"github.com/go-chi/chi/v5"

	"github.com/Lucasmartinsn/ApiProjeto-da-tropa/models"
)


func Update(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	
	if err != nil{
		log.Printf("erro ao fazer parse do id %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var cadastro models.Cadastro

	err = json.NewDecoder(r.Body).Decode(&cadastro)
	if err != nil{
		log.Printf("erro ao fazer decode do json: %v",err)
		http.Error(w, http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
		return
	}

	rows, err := models.Update(int(id), cadastro)
	if err != nil{
		log.Printf("erro ao atualizar register: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("erro: foram atualizador %d registros", rows)
	}

	resp := map[string]any{
		"Error": false,
		"mensage": "dados atualizados com sucesso",
	}

	w.Header().Add("content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}