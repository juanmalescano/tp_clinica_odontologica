package domain

type Odontologo struct {
	Id        int    `json:"id"`
	Matricula string `json:"matricula"`
	Apellido  string `json:"apellido"`
	Nombre    string `json:"nombre"`
}
