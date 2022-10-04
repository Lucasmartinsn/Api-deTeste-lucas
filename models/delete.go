package models

import "github.com/Lucasmartinsn/ApiProjeto-da-tropa/db"

func Delete(id int64) ( int64, error){
	conn, err := db.OpenConnection()
	if err != nil{
		return 0, err
	}

	defer conn.Close()

	response, err := conn.Exec(` DELETE FROM cadastros WHERE id=$'1`, id)

	if err != nil {
		return 0, err
	}

	return response.RowsAffected()
}