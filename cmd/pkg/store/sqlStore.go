package store

import (
	"_/C_/Users/juanm/Google_Drive/Carrera_certified_Tech_Developer/Back_End_III/Projecto_integrador/cmd/internal/domain"
	"database/sql"
)

type Storage interface {
	GetByID()
	Post()
	Put()
	Patch()
	Delete()
}

type sqlStore struct {
	db *sql.DB
}

func NewSqlStore(db *sql.DB) Storage {
	return &sqlStore{
		db: db,
	}
}

func (s *sqlStore) GetByID(id int) (domain.Odontologo, error) {
	return domain.Odontologo{}, nil
}

func (s *sqlStore) Post(odontologo domain.Odontologo) error {
	return nil
}

func (s *sqlStore) Put(odontologo domain.Odontologo) error {
	return nil
}

func (s *sqlStore) Delete(id int) error {
	return nil
}

func (s *sqlStore) Patch(codeValue string) bool {
	return false
}
