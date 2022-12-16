package repository

import (
	"AV II GO/internal/domain"

)

type Repository interface {

	Create(dentist domain.Dentist) (domain.Dentist, error)
	GetById(id int) (domain.Dentist, error)
	GetByRegistration(enrollment string) (domain.Dentist, error)
	GetAll() ([]domain.Dentist, error)
	Update(id int, dentist domain.Dentist) (domain.Dentist, error)
	Patch(id int, dentist domain.Dentist) (domain.Dentist, error)
	Delete(id int) error
}

type repository struct {
	storage store.StoreInterfaceDentist
}

func Repository(storage store.StoreInterfaceDentist) Repository {
	return &repository{storage}
}

func (r *repository) Create(dentist domain.Dentist) (domain.Dentist, error) {
	createdDentist, err := r.storage.Create(dentist)
	if err != nil {
		return domain.Dentist{}, err
	}
	return createdDentist, nil
}

func (r *repository) GetAll() ([]domain.Dentist, error) {
	dentists, err := r.storage.GetAll()
	if err != nil {
		return []domain.Dentist{}, err
	}
	return dentists, nil
}

func (r *repository) GetById(id int) (domain.Dentist, error) {
	dentist, err := r.storage.getById(id)
	if err != nil {
		return domain.Dentist{}, err
	}
	return dentist, nil
}

func (r *repository) GetByRegistration(registration string) (domain.Dentist, error) {
	dentist, err := r.storage.GetByRegistration(registration)
	if err != nil {
		return domain.Dentist{}, err
	}
	return dentist, nil
}

func (r *repository) Update(id int, dentist domain.Dentist) (domain.Dentist, error) {
	updatedDentist, err := r.storage.Update(id, dentist)
	if err != nil {
		return domain.Dentist{}, err
	}
	return updatedDentist, nil
}

func (r *repository) Patch(id int, dentist domain.Dentist) (domain.Dentist, error) {
	updatedDentist, err := r.storage.Patch(id, dentist)
	if err != nil {
		return domain.Dentist{}, err
	}
	return updatedDentist, nil
}

func (r *repository) Delete(id int) error {
	err := r.storage.Delete(id)
	if err != nil {
		return err
	}
	return nil
}