package odontologo

import (
	"Projecto_integrador/cmd/pkg/store"
	"errors"

	"../domain"
)

type OdonRepository interface {
	GetById(id int) (domain.Odontologo, error)
	Insert(odontologo domain.Odontologo) (domain.Odontologo, error)
	Update(id int, odontologo domain.Odontologo) error
	UpdatePatch(campo string, valor string) error
	Delete(id int) error
}

type repository struct {
	storage store.Storage
}

func newRepository(storage store.Storage) OdonRepository{
	return &repository{storage}
}


func (r *OdonRepository) GetByID(id int) (domain.Odontologo, error) {
	odontologo, err := r.storage.GetByID(id)
	if err != nil {
		return domain.Odontologo{}, errors.New("product not found")
	}
	return odontologo, nil

}
