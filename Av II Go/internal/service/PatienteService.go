package service

import (
	"errors"
	"AV II GO/internal/domain"
)

type Service interface {
	Create(patient domain.Patient) (domain.Patient, error)
	GetById(id int) (domain.Patient, error)
	GetByRg(rg string) (domain.Patient, error)
	GetAll() ([]domain.Patient, error)
	Update(id int, patient domain.Patient) (domain.Patient, error)
	Patch(id int, patient domain.Patient) (domain.Patient, error)
	Delete(id int) error
}

type service struct {
	r Repository
}

func Service(r Repository) Service {
	return &service{r}
}

func (s *service) GetById(id int) (domain.Patient, error) {
	patient, err := s.r.GetById(id)
	if err != nil {
		return domain.Patient{}, err
	}
	return patient, nil
}

func (s *service) GetByRg(rg string) (domain.Patient, error) {
	patient, err := s.r.GetByRg(rg)
	if err != nil {
		return domain.Patient{}, err
	}
	return patient, nil
}

func (s *service) GetAll() ([]domain.Patient, error) {
	patients, err := s.r.GetAll()
	if err != nil {
		return []domain.Patient{}, err
	}
	return patients, nil
}

func (s *service) Create(patient domain.Patient) (domain.Patient, error) {
	patients, err := s.GetAll()
	if err != nil {
		return domain.Patient{}, err
	}

	for i := range patients {
		if patients[i].Rg == patient.Rg {
			return domain.Patient{}, errors.New("rg already exists")
		}
	}

	createdPatient, err := s.r.Create(patient)
	if err != nil {
		return domain.Patient{}, err
	}
	return createdPatient, nil
}

func (s *service) Update(id int, patient domain.Patient) (domain.Patient, error) {
	patients, err := s.GetAll()
	if err != nil {
		return domain.Patient{}, err
	}

	for i := range patients {
		if patients[i].Rg == patient.Rg {
			return domain.Patient{}, errors.New("rg already exists")
		}
	}

	updatedPatient, err := s.r.Update(id, patient)
	if err != nil {
		return domain.Patient{}, err
	}
	return updatedPatient, nil
}

func (s *service) Patch(id int, patient domain.Patient) (domain.Patient, error) {
	patients, err := s.GetAll()
	if err != nil {
		return domain.Patient{}, err
	}

	for i := range patients {
		if patients[i].Rg == patient.Rg {
			return domain.Patient{}, errors.New("rg already exists")
		}
	}

	updatedPatient, err := s.r.Patch(id, patient)
	if err != nil {
		return domain.Patient{}, err
	}
	return updatedPatient, nil
}

func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}