package db

import (
	"AV II GO/internal/domain"
	"database/sql"
	"errors"
	"log"
)

type sqlStoreAppointment struct {
	db *sql.DB
}

func QLStoreAppointment(db *sql.DB) StoreInterfaceAppointment {
	return &sqlStoreAppointment{
		db: db,
	}
}

func (s *sqlAppointment) GetById(id int) (domain.Appointment, error) {
	queryGetById := `SELECT appointment.id, patient.id, patient.name, patient.last_name, patient.rg, patient.register_date, 
							dentist.id, dentist.name, dentist.last_name, dentist.registration, 
							appointment.date, appointment.description 
							FROM appointment 
							INNER JOIN patient 
							ON patient.id = appointment.patient_id 
							INNER JOIN dentist 
							ON dentist.id = appointment.dentist_id 
							WHERE appointment.id = ?`

	row := s.db.QueryRow(queryGetById, id)

	appointment := domain.Appointment{}

	err := row.Scan(
		&appointment.Id,
		&appointment.Patient.Id,
		&appointment.Patient.Name,
		&appointment.Patient.LastName,
		&appointment.Patient.Rg,
		&appointment.Patient.RegisterDate,
		&appointment.Dentist.Id,
		&appointment.Dentist.Name,
		&appointment.Dentist.LastName,
		&appointment.Dentist.Enrollment,
		&appointment.Date,
		&appointment.Description,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return appointment, errors.New("appointment not found")
	}

	return appointment, nil
}

func (s *sqlStoreAppointment) GetByRg(rg string) ([]domain.Appointment, error) {
	queryGetByRg := `SELECT appointment.id, patient.id, patient.name, patient.last_name, patient.rg, patient.register_date, 
							dentist.id, dentist.name, dentist.last_name, dentist.registration, 
							appointment.date, appointment.description 
							FROM appointment 
							INNER JOIN patient 
							ON patient.id = appointment.patient_id 
							INNER JOIN dentist 
							ON dentist.id = appointment.dentist_id 
							WHERE patient.rg = ?`

	var appointments []domain.Appointment
	rows, err := s.db.Query(queryGetByRg, rg)

	defer rows.Close()

	for rows.Next() {
		appointment := domain.Appointment{}

		if err := rows.Scan(
			&appointment.Id,
			&appointment.Patient.Id,
			&appointment.Patient.Name,
			&appointment.Patient.LastName,
			&appointment.Patient.Rg,
			&appointment.Patient.RegisterDate,
			&appointment.Dentist.Id,
			&appointment.Dentist.Name,
			&appointment.Dentist.LastName,
			&appointment.Dentist.Enrollment,
			&appointment.Date,
			&appointment.Description,
		); err != nil {
			return appointments, err
		}
		appointments = append(appointments, appointment)
	}

	return appointments, nil
}

func (s *sqlStoreAppointment) CreateById(appointment domain.Appointment, idPatient int, idDentist int) (domain.Appointment, error) {
	queryInsert := "INSERT INTO appointment (patient_id, dentist_id, date, description) VALUES (?, ?, ?, ?)"

	stmt, err := s.db.Prepare(queryInsert)

	}
	defer stmt.Close()

	res, err := stmt.Exec(
		idPatient,
		idDentist,
		appointment.Date,
		appointment.Description)

	RowsAffected, _ := res.RowsAffected()
	if RowsAffected == 0 {
		return domain.Appointment{}, errors.New("fail to save appointment")
	}


	return appointment, nil
}

func (s *sqlAppointment) CreateByRgEnrollment(appointment domain.Appointment, rgPatient string, enrollmentDentist string) (domain.Appointment, error) {
	queryInsert := `INSERT INTO appointment (patient_id, dentist_id, date, description)
					VALUES ((SELECT patient.id FROM patient WHERE patient.rg = ?), 
					(SELECT dentist.id FROM dentist WHERE dentist.enrollment = ?), 
					?, ?)`

	stmt, err := s.db.Prepare(queryInsert)

	defer stmt.Close()

	res, err := stmt.Exec(
		rgPatient,
		enrollmentDentist,
		appointment.Date,
		appointment.Description
	)
	}

	RowsAffected, _ := res.RowsAffected()
	if RowsAffected == 0 {
		return domain.Appointment{}, errors.New("fail to save appointment")
	}

	return appointment, nil
}

func (s *sqlStoreAppointment) Update(id int, appointment domain.Appointment) (domain.Appointment, error) {
	queryUpdate  := "UPDATE appointment SET patient_id = ?, dentist_id = ?, date = ?, description = ? WHERE id = ?"

	persistedAppointment, err := s.ReadById(id)
	if err != nil {
		return domain.Appointment{}, errors.New("appointment not found")
	}

	persistedAppointment.Patient.Id = appointment.Patient.Id
	persistedAppointment.Dentist.Id = appointment.Dentist.Id
	persistedAppointment.Date = appointment.Date
	persistedAppointment.Description = appointment.Description

	result, err := s.db.Exec(
		queryUpdate,
		persistedAppointment.Patient.Id,
		persistedAppointment.Dentist.Id,
		persistedAppointment.Date,
		persistedAppointment.Description,
		id,
	)

	return appointment, nil
}

func (s *sqlStoreAppointment) Patch(id int, appointment domain.Appointment) (domain.Appointment, error) {
	queryUpdate  := "UPDATE appointment SET patient_id = ?, dentist_id = ?, date = ?, description = ? WHERE id = ?"

	persistedAppointment, err := s.ReadById(id)
	if err != nil {
		return domain.Appointment{}, errors.New("appointment not found")
	}

	if appointment.Patient.Id != 0 {
		persistedAppointment.Patient.Id = appointment.Patient.Id
	}
	if appointment.Dentist.Id != 0 {
		persistedAppointment.Dentist.Id = appointment.Dentist.Id
	}
	if appointment.Date != "" {
		persistedAppointment.Date = appointment.Date
	}
	if appointment.Description != "" {
		persistedAppointment.Description = appointment.Description
	}

	result, err := s.db.Exec(
		queryUpdate,
		persistedAppointment.Patient.Id,
		persistedAppointment.Dentist.Id,
		persistedAppointment.Date,
		persistedAppointment.Description,
		id,
	)

	return appointment, nil
}

func (s *sqlStoreAppointment) Delete(id int) error {
	queryDelete := "DELETE FROM appointment WHERE id = ?"

	result, err := s.db.Exec(queryDelete, id)
	if err != nil {
		return err
	}

	affectedRows, err := result.RowsAffected()

	if affectedRows == 0 {
		return errors.New("appointment not found")
	}
	return nil
}