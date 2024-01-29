package storage

import (
	"errors"
	"fmt"
	"rest_project/internal/model"
)

func (s *Storage) PersonGetAll() ([]model.Person, error) {
	var persons []model.Person

	rows, err := s.db.Query("SELECT * FROM persons")
	if err != nil {
		return persons, err
	}
	defer rows.Close()

	for rows.Next() {
		person := model.Person{}
		err := rows.Scan(
			&person.Id,
			&person.Name,
			&person.Surname,
			&person.Patronymic,
			&person.Age,
			&person.Gender,
			&person.Nationality,
		)
		if err != nil {
			return persons, err
		}
		persons = append(persons, person)
	}

	return persons, nil
}

func (s *Storage) PersonSet(person model.Person) (int, error) {
	if  err := s.db.QueryRow(
		"INSERT INTO persons (Name, Surname, Patronymic, Age, Gender, Nationality) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		person.Name,
		person.Surname,
		person.Patronymic,
		person.Age,
		person.Gender,
		person.Nationality,
		).Scan(&person.Id); err != nil {
		return -1, err
	}

	return person.Id, nil
}

func (s *Storage) PersonGet(id int) (model.Person, error) {
	var person model.Person

	row := s.db.QueryRow("SELECT * FROM persons WHERE id = $1", id)
	if err := row.Scan(
		&person.Id,
		&person.Name,
		&person.Surname,
		&person.Patronymic,
		&person.Age,
		&person.Gender,
		&person.Nationality,
	); err != nil {
		return person, err
	}

	return person, nil
}

func (s *Storage) PersonDelete(id int) (error) {
	res, err := s.db.Exec("DELETE FROM persons WHERE id = $1", id)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New(fmt.Sprintf("No person with id %d exists to delete", id))
	}

	return nil
}

func (s *Storage) PersonUpdate(id int, person model.Person) (model.Person, error) {
	var u_person = person

	if  err := s.db.QueryRow(
		"UPDATE persons SET (Name, Surname, Patronymic, Age, Gender, Nationality) = ($1, $2, $3, $4, $5, $6) WHERE id = $7 RETURNING *",
		person.Name,
		person.Surname,
		person.Patronymic,
		person.Age,
		person.Gender,
		person.Nationality,
		id,
		).Scan(
			&person.Id,
			&person.Name,
			&person.Surname,
			&person.Patronymic,
			&person.Age,
			&person.Gender,
			&person.Nationality,); err != nil {
		return u_person, err
	}

	return u_person, nil
}
