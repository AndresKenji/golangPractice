package models

type Cliente struct {
	Id       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Correo   string `json:"correo"`
	Telefono string `json:"telefono"`
}

type Clientes []Cliente
