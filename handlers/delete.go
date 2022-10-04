package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"github.com/go-chi/chi/v5"
	"github.com/Lucasmartinsn/ApiProjeto-da-tropa/models"
)


func Delete(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	
	if err != nil{
		log.Printf("erro ao fazer parse do id %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := models.Delete(int64(id))
	if err != nil{
		log.Printf("erro ao remover register: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("erro: foram removidos %d registros", rows)
	}

	resp := map[string]any{
		"Error": false,
		"mensage": "registro removido com sucesso",
	}

	w.Header().Add("content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}