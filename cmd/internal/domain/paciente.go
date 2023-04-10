package domain

type Paciente struct {
	Id        int    `json:"id"`
	DNI       int    `json:"dni"`
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	domicilio string `json:"domicilio"`
	F_alta    string `json:"f_alta"`
}
