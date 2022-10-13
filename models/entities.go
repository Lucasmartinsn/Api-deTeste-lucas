package models

/// Tabela com os typos e os campos da tabela
type Cadastro struct {
	ID           int    `json:"id"`
	Foto         string `json:"foto"`
	Nome_usuario string `json:"nome_usuario"`
	Email        string `json:"Email"`
	Senha        int    `json:"Senha"`
}
