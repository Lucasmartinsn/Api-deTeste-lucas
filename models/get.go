package models

import "github.com/Lucasmartinsn/ApiProjeto-da-tropa/db"

func Get(id int) (cadastro Cadastro,err error){
	conn, err := db.OpenConnection()

	if err != nil{
		return
	}

	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM cadastro WHERE id=$1`, id)

	err = row.Scan(&cadastro.ID, &cadastro.Foto, & cadastro.Nome_usuario, &cadastro.Email, &cadastro.Senha)

	return
}