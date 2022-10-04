package models

import (

	"github.com/Lucasmartinsn/ApiProjeto-da-tropa/db"
)

func Insert(cadastro Cadastro) ( id int, err error ){
	conn, err := db.OpenConnection()
	if err != nil{
		return
	}

	defer conn.Close()

	sql := `INSERT INTO cadastro (foto, nome_usuario, email, senha) VALUES ($1,$2,$3,$4) RETURNING id`
	
	err = conn.QueryRow(sql, cadastro.Foto, cadastro.Nome_usuario, cadastro.Email, cadastro).Scan(&id)
	
	return
}