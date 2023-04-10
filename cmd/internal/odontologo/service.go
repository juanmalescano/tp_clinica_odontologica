package odontologo

import "_/C_/Users/juanm/Google_Drive/Carrera_certified_Tech_Developer/Back_End_III/Projecto_integrador/cmd/internal/domain"

type OdoService interface {
	GetById(id int) (domain.Odontologo, error)
	Insert(odontologo domain.Odontologo) (domain.Odontologo, error)
	Update(id int, odontologo domain.Odontologo) error
	UpdatePatch(campo string, valor string) error
	Delete(id int) error
}

type service struct {
	r domain.OdoRepository
}

func NewService(r OdoService) OdoService {
	return &domain.OdoService{r}
}

func (s *OdoService) GetByID(id int) (domain.Odontologo, error) {
	p, err := s.r.GetByID(id)
	
	if err != nil {
		return domain.Product{}, err
	}
	return p, nil
}

