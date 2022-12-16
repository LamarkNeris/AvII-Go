package db

import (
	"AV II GO/internal/domain"
	"database/sql"
	"errors"
	"log"
)

type sqlDentist struct {
	db *sql.DB
}

func SQLDentist(db *sql.DB) InterfaceDentist {
	return &sqlDentist{
		db: db,
	}
}

func (s *sqlDentist) GetById(id int) (domain.Dentist, error) {
	queryGetById := "SELECT id, name, last_name, enrollment FROM dentist where id = ?"

	row := s.db.QueryRow(queryGetById, id)

	dentist := domain.Dentist{}

	err := row.Scan(
		&dentist.Id,
		&dentist.Name,
		&dentist.LastName,
		&dentist.Enrollment,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return dentist, errors.New("dentist not found")
	}

	return dentist, nil
}

func (s *sqlDentist) GetByRegistration(registration string) (domain.Dentist, error) {
	queryGetByRegistration := "SELECT id, name, last_name, registration FROM dentist where registration = ?"

	row := s.db.QueryRow(queryGetByRegistration, registration)

	dentist := domain.Dentist{}

	err := row.Scan(
		&dentist.Id,
		&dentist.Name,
		&dentist.LastName,
		&dentist.Enrollment,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return dentist, errors.New("dentist not found")
	}

	return dentist, nil
}

func (s *sqlDentist) GetAll() ([]domain.Dentist, error) {

	queryGetAll := "SELECT id, name, last_name, registration FROM dentist"

	var dentists []domain.Dentist
	rows, err := s.db.Query(queryGetAll)
	if err != nil {
		return []domain.Dentist{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var dentist domain.Dentist

		if err := rows.Scan(
			&dentist.Id,
			&dentist.Name,
			&dentist.LastName,
			&dentist.Enrollment,
		); err != nil {
			return dentists, err
		}

		dentists = append(dentists, dentist)
	}
	return dentists, nil
}

func (s *sqlDentist) Create(dentist domain.Dentist) (domain.Dentist, error) {
	queryInsert := "INSERT INTO dentist (name, last_name, enrollment) VALUES (?, ?, ?)"

	stmt, err := s.db.Prepare(queryInsert)

	if err != nil {
		return domain.Dentist{}, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		dentist.Name,
		dentist.LastName,
		dentist.Enrollment)
	if err != nil {
		return domain.Dentist{}, err
	}

	RowsAffected, _ := res.RowsAffected()
	if RowsAffected == 0 {
		return domain.Dentist{}, errors.New("fail to save dentist")
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return domain.Dentist{}, err
	}

	dentist.Id = int(lastId)

	return dentist, nil
}

func (s *sqlDentist) Update(id int, dentist domain.Dentist) (domain.Dentist, error) {
	queryUpdate  := "UPDATE dentist SET name = ?, last_name = ?, registration = ? WHERE id = ?"

	persistedDentist, err := s.ReadById(id)
	if err != nil {
		return domain.Dentist{}, errors.New("dentist not found")
	}

	persistedDentist.Name = dentist.Name
	persistedDentist.LastName = dentist.LastName
	persistedDentist.Registration = dentist.Registration

	result, err := s.db.Exec(
		queryUpdate,
		persistedDentist.Name,
		persistedDentist.LastName,
		persistedDentist.Registration,
		id,
	)

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return domain.Dentist{}, err
	}
	log.Println(affectedRows)

	return persistedDentist, nil
}

func (s *sqlStoreDentist) Patch(id int, dentist domain.Dentist) (domain.Dentist, error) {
	queryUpdate  := "UPDATE dentist SET name = ?, last_name = ?, enrollment = ? WHERE id = ?"

	persistedDentist, err := s.ReadById(id)
	if err != nil {
		return domain.Dentist{}, errors.New("dentist not found")
	}

	if dentist.Name != "" {
		persistedDentist.Name = dentist.Name
	}
	if dentist.LastName != "" {
		persistedDentist.LastName = dentist.LastName
	}
	if dentist.Enrollment != "" {
		persistedDentist.Enrollment = dentist.Enrollment
	}

	result, err := s.db.Exec(
		queryUpdate,
		persistedDentist.Name,
		persistedDentist.LastName,
		persistedDentist.Enrollment,
		id,
	)

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return domain.Dentist{}, err
	}
	log.Println(affectedRows)

	return persistedDentist, nil
}

func (s *sqlStoreDentist) Delete(id int) error {
	queryDelete := "DELETE FROM dentist WHERE id = ?"

	result, err := s.db.Exec(queryDelete, id)
	if err != nil {
		return err
	}

	affectedRows, err := result.RowsAffected()

	if affectedRows == 0 {
		return errors.New("dentist not found")
	}

	return nil
}