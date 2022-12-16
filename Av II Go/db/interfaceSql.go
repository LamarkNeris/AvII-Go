package db

import "AV II GO/internal/domain"

type StoreInterfacePatient interface {

	Create(patient domain.Patient) (domain.Patient, error)
	GetById(id int) (domain.Patient, error)
	GetByRg(rg string) (domain.Patient, error)
	GetAll() ([]domain.Patient, error)
	Update(id int, patient domain.Patient) (domain.Patient, error)
	Patch(id int, patient domain.Patient) (domain.Patient, error)
	Delete(id int) error
}

type StoreInterfaceDentist interface {

	Create(dentist domain.Dentist) (domain.Dentist, error)
	getById(id int) (domain.Dentist, error)
	GetByRegistration(enrollment string) (domain.Dentist, error)
	GetAll() ([]domain.Dentist, error)
	Update(id int, dentist domain.Dentist) (domain.Dentist, error)
	Patch(id int, dentist domain.Dentist) (domain.Dentist, error)
	Delete(id int) error
}

type StoreInterfaceAppointment interface {

	CreateById(appointment domain.Appointment, idPatient int, idDentist int) (domain.Appointment, error)
	CreateByRgRegistration(appointment domain.Appointment, rgPatient string, enrollmentDentist string) (domain.Appointment, error)
	GetById(id int) (domain.Appointment, error)
	GetByRg(rg string) ([]domain.Appointment, error)
	Update(id int, appointment domain.Appointment) (domain.Appointment, error)
	Patch(id int, appointment domain.Appointment) (domain.Appointment, error)
	Delete(id int) error
}