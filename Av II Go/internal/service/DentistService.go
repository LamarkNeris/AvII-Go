package service

import (
	"AV II GO/internal/domain"
	"errors"
)

type Service interface {

	Create(dentist domain.Dentist) (domain.Dentist, error)
	GetById(id int) (domain.Dentist, error)
	GetByRegistration(enrollment string) (domain.Dentist, error)
	GetAll() ([]domain.Dentist, error)
	Update(id int, dentist domain.Dentist) (domain.Dentist, error)
	Patch(id int, dentist domain.Dentist) (domain.Dentist, error)
	Delete(id int) error
}

type service struct {
	r Repository
}

func Service(r Repository) Service {
	return &service{r}
}

func (s *service) Create(dentist domain.Dentist) (domain.Dentist, error) {
	dentists, err := s.GetAll()
	if err != nil {
		return domain.Dentist{}, err
	}

	for i := range dentists {
		if dentists[i].Registration == dentist.Registration {
			return domain.Dentist{}, errors.New("registration already exists")
		}
	}

	createdDentist, err := s.r.Create(dentist)
	if err != nil {
		return domain.Dentist{}, err
	}
	return createdDentist, nil
}

func (s *service) GetAll() ([]domain.Dentist, error) {
	dentists, err := s.r.GetAll()
	if err != nil {
		return []domain.Dentist{}, err
	}
	return dentists, nil
}

func (s *service) GetById(id int) (domain.Dentist, error) {
	dentist, err := s.r.GetById(id)
	if err != nil {
		return domain.Dentist{}, err
	}
	return dentist, nil
}

func (s *service) GetByRegistration(registration string) (domain.Dentist, error) {
	dentist, err := s.r.GetByRegistration(registration)
	if err != nil {
		return domain.Dentist{}, err
	}
	return dentist, nil
}

func (s *service) Update(id int, dentist domain.Dentist) (domain.Dentist, error) {
	dentists, err := s.GetAll()
	if err != nil {
		return domain.Dentist{}, err
	}

	for i := range dentists {
		if dentists[i].Enrollment == dentist.Registration {
			return domain.Dentist{}, errors.New("registration already exists")
		}
	}

	updatedDentist, err := s.r.Update(id, dentist)
	if err != nil {
		return domain.Dentist{}, err
	}
	return updatedDentist, nil
}

func (s *service) Patch(id int, dentist domain.Dentist) (domain.Dentist, error) {
	dentists, err := s.GetAll()
	if err != nil {
		return domain.Dentist{}, err
	}

	for i := range dentists {
		if dentists[i].Enrollment == dentist.Registration {
			return domain.Dentist{}, errors.New("registration already exists")
		}
	}

	updatedDentist, err := s.r.Patch(id, dentist)
	if err != nil {
		return domain.Dentist{}, err
	}
	return updatedDentist, nil
}

func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}