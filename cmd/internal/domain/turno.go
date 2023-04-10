package domain

type Turno struct {
	Id          int        `json:"id"`
	Descripcion string     `json:"descripcion"`
	Fecha       string     `json:"fecha"`
	Paciente    Paciente   `json:"paciente"`
	Odontologo  Odontologo `json:"odontologo"`
}
