package repository

import (
	"database/sql"
	"mini-project/structs"
)

func GetAllPerson(db *sql.DB) (result []structs.Person, err error) {
	sql := "SELECT * FROM person"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var person structs.Person

		err = rows.Scan(&person.ID, &person.FirstName, &person.LastName)
		if err != nil {
			return
		}

		result = append(result, person)
	}

	return
}

func InsertPerson(db *sql.DB, person structs.Person) (err error) {
	sql := "INSERT INTO person(id, first_name, last_name) VALUES($1, $2, $3)"

	errs := db.QueryRow(sql, person.ID, person.FirstName, person.LastName)

	return errs.Err()
}

func UpdatePerson(db *sql.DB, person structs.Person) (err error) {
	sql := "update person set first_name = $1, last_name = $2 where id = $3"

	errs := db.QueryRow(sql, person.FirstName, person.LastName, person.ID)

	return errs.Err()
}

func DeletePerson(db *sql.DB, person structs.Person) (err error) {
	sql := "delete from person where id = $1"

	errs := db.QueryRow(sql, person.ID)

	return errs.Err()
}
