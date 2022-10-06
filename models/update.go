package models

import "github.com/Lucasmartinsn/ApiProjeto-da-tropa/db"

func Update(id int64, cadastro Cadastro) ( int64, error){
	conn, err := db.OpenConnection()
	if err != nil{
		return 0, err
	}

	defer conn.Close()

	response, err := conn.Exec(`UPDATE cadastro SET foto=$2, nome_usuario=$3, email=$4, senha=$5 WHERE id=$1`,
	id, cadastro.Foto,cadastro.Nome_usuario, cadastro.Email, cadastro.Senha)

	if err != nil {
		return 0, err
	}

	return response.RowsAffected()
}