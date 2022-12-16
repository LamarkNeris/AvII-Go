package repository

import (
	"AV II GO/internal/domain"
)

type Repository interface {
	Create(patient domain.Patient) (domain.Patient, error)
	GetById(id int) (domain.Patient, error)
	GetByRg(rg string) (domain.Patient, error)
	GetAll() ([]domain.Patient, error)
	Update(id int, patient domain.Patient) (domain.Patient, error)
	Patch(id int, patient domain.Patient) (domain.Patient, error)
	Delete(id int) error
}

type repository struct {
	storage store.StoreInterfacePatient
}

func Repository(storage store.StoreInterfacePatient) Repository {
	return &repository{storage}
}

func (r *repository) Create(patient domain.Patient) (domain.Patient, error) {
	createdPatient, err := r.storage.Create(patient)
	if err != nil {
		return domain.Patient{}, err
	}
	return createdPatient, nil
}

func (r *repository) GetById(id int) (domain.Patient, error) {
	patient, err := r.storage.GetById(id)
	if err != nil {
		return domain.Patient{}, err
	}
	return patient, nil
}

func (r *repository) GetByRg(rg string) (domain.Patient, error) {
	patient, err := r.storage.GetByRg(rg)
	if err != nil {
		return domain.Patient{}, err
	}
	return patient, nil
}

func (r *repository) GetAll() ([]domain.Patient, error) {
	patients, err := r.storage.GetAll()
	if err != nil {
		return []domain.Patient{}, err
	}
	return patients, nil
}

func (r *repository) Update(id int, patient domain.Patient) (domain.Patient, error) {
	updatedPatient, err := r.storage.Update(id, patient)
	if err != nil {
		return domain.Patient{}, err
	}
	return updatedPatient, nil
}

func (r *repository) Patch(id int, patient domain.Patient) (domain.Patient, error) {
	updatedPatient, err := r.storage.Patch(id, patient)
	if err != nil {
		return domain.Patient{}, err
	}
	return updatedPatient, nil
}

func (r *repository) Delete(id int) error {
	err := r.storage.Delete(id)
	if err != nil {
		return err
	}
	return nil
}