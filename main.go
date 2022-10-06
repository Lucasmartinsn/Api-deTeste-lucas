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
	r.Post("/cadastros", handlers.Create)
	r.Put("/cadastros/{id}", handlers.Update)
	r.Delete("/cadastros/{id}", handlers.Delete)
	r.Get("/cadastros", handlers.List)
	r.Get("/cadastros/{id}", handlers.Get)

	http.ListenAndServe(fmt.Sprintf(":%s",configs.GetServerPort()), r)
}