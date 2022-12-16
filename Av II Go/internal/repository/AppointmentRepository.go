package repository

import (
	"AV II GO/internal/domain"
	
)

type Repository interface {
	CreateById(appointment domain.Appointment, idPatient int, idDentist int) (domain.Appointment, error)
	CreateByRgRegistration(appointment domain.Appointment, rgPatient string, registrationDentist string) (domain.Appointment, error)
	GetById(id int) (domain.Appointment, error)
	GetByRg(rg string) ([]domain.Appointment, error)
	Update(id int, appointment domain.Appointment) (domain.Appointment, error)
	Patch(id int, appointment domain.Appointment) (domain.Appointment, error)
	Delete(id int) error
}

type repository struct {
	storage store.StoreInterfaceAppointment
}

func Repository(storage store.StoreInterfaceAppointment) Repository {
	return &repository{storage}
}

func (r *repository) CreateById(appointment domain.Appointment, idPatient int, idDentist int) (domain.Appointment, error) {
	createdAppointment, err := r.storage.CreateById(appointment, idPatient, idDentist)
	if err != nil {
		return domain.Appointment{}, err
	}
	return createdAppointment, nil
}

func (r *repository) CreateByRgRegistration(appointment domain.Appointment, rgPatient string, registrationDentist string) (domain.Appointment, error) {
	createdAppointment, err := r.storage.CreateByRgRegistration(appointment, rgPatient, registrationDentist)
	if err != nil {
		return domain.Appointment{}, err
	}
	return createdAppointment, nil
}

func (r *repository) GetById(id int) (domain.Appointment, error) {
	appointment, err := r.storage.GetById(id)
	if err != nil {
		return domain.Appointment{}, err
	}
	return appointment, nil
}

func (r *repository) GetByRg(rg string) ([]domain.Appointment, error) {
	appointments, err := r.storage.GetByRg(rg)
	if err != nil {
		return []domain.Appointment{}, err
	}
	return appointments, nil
}

func (r *repository) Update(id int, appointment domain.Appointment) (domain.Appointment, error) {
	updatedAppointment, err := r.storage.Update(id, appointment)
	if err != nil {
		return domain.Appointment{}, err
	}
	return updatedAppointment, nil
}

func (r *repository) Patch(id int, appointment domain.Appointment) (domain.Appointment, error) {
	updatedAppointment, err := r.storage.Patch(id, appointment)
	if err != nil {
		return domain.Appointment{}, err
	}
	return updatedAppointment, nil
}

func (r *repository) Delete(id int) error {
	err := r.storage.Delete(id)
	if err != nil {
		return err
	}
	return nil
}