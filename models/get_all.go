package models

import "github.com/Lucasmartinsn/ApiProjeto-da-tropa/db"

func GetAll() (cadastros []Cadastro,err error){
	conn, err := db.OpenConnection()

	if err != nil{
		return
	}

	defer conn.Close()

	rows := conn.Query(`SELECT * FROM cadastro`)
	if err != nil{
		return
	}

	for  rows.Next(){
		var cadastro Cadastro

		err = rows.Scan(&cadastro.ID, &cadastro.Foto, & cadastro.Nome_usuario, &cadastro.Email, &cadastro.Senha)
		if err != nil{
			continue
		}

		cadastros = append(cadastros, cadastro)
	}
	return
}