package service

import (
	"AV II GO/internal/domain"
)

type Service interface {
	CreateById(appointment domain.Appointment, idPatient int, idDentist int) (domain.Appointment, error)
	CreateByRgRegistration(appointment domain.Appointment, rgPatient string, registrationDentist string) (domain.Appointment, error)
	GetById(id int) (domain.Appointment, error)
	GetByRg(rg string) ([]domain.Appointment, error)
	Update(id int, appointment domain.Appointment) (domain.Appointment, error)
	Patch(id int, appointment domain.Appointment) (domain.Appointment, error)
	Delete(id int) error
}

type service struct {
	r Repository
}

func Service(r Repository) Service {
	return &service{r}
}

func (s *service) CreateById(appointment domain.Appointment, idPatient int, idDentist int) (domain.Appointment, error) {
	createdAppointment, err := s.r.CreateById(appointment, idPatient, idDentist)
	if err != nil {
		return domain.Appointment{}, err
	}
	return createdAppointment, nil
}

func (s *service) CreateByRgRegistration(appointment domain.Appointment, rgPatient string, registrationDentist string) (domain.Appointment, error) {
	createdAppointment, err := s.r.CreateByRgRegistration(appointment, rgPatient, registrationDentist)
	if err != nil {
		return domain.Appointment{}, err
	}
	return createdAppointment, nil
}

func (s *service) GetById(id int) (domain.Appointment, error) {
	appointment, err := s.r.GetById(id)
	if err != nil {
		return domain.Appointment{}, err
	}
	return appointment, nil
}

func (s *service) GetByRg(rg string) ([]domain.Appointment, error) {
	appointments, err := s.r.GetByRg(rg)
	if err != nil {
		return []domain.Appointment{}, err
	}
	return appointments, nil
}

func (s *service) Update(id int, appointment domain.Appointment) (domain.Appointment, error) {
	updatedAppointment, err := s.r.Update(id, appointment)
	if err != nil {
		return domain.Appointment{}, err
	}
	return updatedAppointment, nil
}

func (s *service) Patch(id int, appointment domain.Appointment) (domain.Appointment, error) {
	updatedAppointment, err := s.r.Patch(id, appointment)
	if err != nil {
		return domain.Appointment{}, err
	}
	return updatedAppointment, nil
}

func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}

	return nil
}