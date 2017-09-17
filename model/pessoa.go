package model

import "time"

type Pessoa struct
{
	Id int `json:"id"`
	Nome string `json:"name"`
	Matricula int `json:"matricula"`
	DataNascimento time.Time `json:"dataNascimento"`
}

type PessoaBD []struct{
	IDPESSOA string
	NOME string
	MATRICULA string
	DATANASCIMENTO string
}