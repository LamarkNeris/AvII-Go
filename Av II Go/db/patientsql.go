package db

import (
	"AV II GO/internal/domain"
	"database/sql"
	"errors"
	"log"
)

type sqlStorePatient struct {
	db *sql.DB
}

func SQLStorePatient(db *sql.DB) StoreInterfacePatient {
	return &sqlStorePatient{
		db: db,
	}
}

func (s *sqlPatient) GetById(id int) (domain.Patient, error) {
	queryGetById := "SELECT id, name, last_name, rg, register_date FROM patient where id = ?"

	row := s.db.QueryRow(queryGetById, id)

	patient := domain.Patient{}

	err := row.Scan(
		&patient.Id,
		&patient.Name,
		&patient.LastName,
		&patient.Rg,
		&patient.RegisterDate,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return patient, errors.New("patient not found")
	}

	return patient, nil
}

func (s *sqlPatient) GetByRg(rg string) (domain.Patient, error) {
	queryGetByRg := "SELECT id, name, last_name, rg, register_date FROM patient where rg = ?"

	row := s.db.QueryRow(queryGetByRg, rg)

	patient := domain.Patient{}

	err := row.Scan(
		&patient.Id,
		&patient.Name,
		&patient.LastName,
		&patient.Rg,
		&patient.RegisterDate,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return patient, errors.New("patient not found")
	}

	return patient, nil
}

func (s *sqlPatient) GetAll() ([]domain.Patient, error) {

	queryGetAll := "SELECT id, name, last_name, rg, register_date FROM patient"

	var patients []domain.Patient
	rows, err := s.db.Query(queryGetAll)
	if err != nil {
		return []domain.Patient{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var patient domain.Patient

		if err := rows.Scan(
			&patient.Id,
			&patient.Name,
			&patient.LastName,
			&patient.Rg,
			&patient.RegisterDate,
		); err != nil {
			return patients, err
		}

		patients = append(patients, patient)
	}
	return patients, nil
}

func (s *sqlPatient) Create(patient domain.Patient) (domain.Patient, error) {
	queryInsert := "INSERT INTO patient (name, last_name, rg, register_date) VALUES (?, ?, ?, ?)"

	stmt, err := s.db.Prepare(queryInsert)

	if err != nil {
		return domain.Patient{}, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		patient.Name,
		patient.LastName,
		patient.Rg,
		patient.RegisterDate)
	if err != nil {
		return domain.Patient{}, err
	}

	RowsAffected, _ := res.RowsAffected()
	if RowsAffected == 0 {
		return domain.Patient{}, errors.New("fail to save patient")
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return domain.Patient{}, err
	}

	patient.Id = int(lastId)

	return patient, nil
}

func (s *sqlPatient) Update(id int, patient domain.Patient) (domain.Patient, error) {
	queryUpdate  := "UPDATE patient SET name = ?, last_name = ?, rg = ?, register_date = ? WHERE id = ?"

	persistedPatient, err := s.GetById(id)
	if err != nil {
		return domain.Patient{}, errors.New("patient not found")
	}

	persistedPatient.Name = patient.Name
	persistedPatient.LastName = patient.LastName
	persistedPatient.Rg = patient.Rg
	persistedPatient.RegisterDate = patient.RegisterDate

	result, err := s.db.Exec(
		queryUpdate,
		persistedPatient.Name,
		persistedPatient.LastName,
		persistedPatient.Rg,
		persistedPatient.RegisterDate,
		id,
	)

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return domain.Patient{}, err
	}
	log.Println(affectedRows)

	return persistedPatient, nil
}

func (s *sqlPatient) Patch(id int, patient domain.Patient) (domain.Patient, error) {
	queryUpdate  := "UPDATE patient SET name = ?, last_name = ?, rg = ?, register_date = ? WHERE id = ?"

	persistedPatient, err := s.ReadById(id)
	if err != nil {
		return domain.Patient{}, errors.New("patient not found")
	}

	if patient.Name != "" {
		persistedPatient.Name = patient.Name
	}
	if patient.LastName != "" {
		persistedPatient.LastName = patient.LastName
	}
	if patient.Rg != "" {
		persistedPatient.Rg = patient.Rg
	}
	if patient.RegisterDate != "" {
		persistedPatient.RegisterDate = patient.RegisterDate
	}

	result, err := s.db.Exec(
		queryUpdate,
		persistedPatient.Name,
		persistedPatient.LastName,
		persistedPatient.Rg,
		persistedPatient.RegisterDate,
		id,
	)

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return domain.Patient{}, err
	}
	log.Println(affectedRows)

	return persistedPatient, nil
}

func (s *sqlPatient) Delete(id int) error {
	queryDelete := "DELETE FROM patient WHERE id = ?"

	result, err := s.db.Exec(queryDelete, id)
	if err != nil {
		return err
	}

	affectedRows, err := result.RowsAffected()

	if affectedRows == 0 {
		return errors.New("patient not found")
	}

	return nil
}