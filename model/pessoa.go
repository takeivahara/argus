package model

import "time"

type Pessoa struct
{
	Id int32 `json:"id"`
	Nome string `json:"name"`
	Matricula int32 `json:"matricula"`
	DataNascimento time.Time `json:"dataNascimento"`
}
