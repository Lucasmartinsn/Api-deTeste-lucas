package main

import (
	"fmt"
	"net/http"

	"github.com/Lucasmartinsn/ApiProjeto-da-tropa/configs"
	"github.com/Lucasmartinsn/ApiProjeto-da-tropa/handlers"
	"github.com/go-chi/chi/v5"
)

func main(){
	err := configs.Load()
	if err != nil{
		panic(err)
	}

	r := chi.NewRouter()
	r.Post("/cadastro", handlers.Create)
	r.Put("/cadastro/{id}", handlers.Update)
	r.Delete("/cadastro/{id}", handlers.Delete)
	r.Get("/cadastro", handlers.List)
	r.Get("/cadastro/{id}", handlers.Get)

	http.ListenAndServe(fmt.Sprintf(":%s",configs.GetServerPort()), r)
}