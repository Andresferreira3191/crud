package main

import (
	"time"
	"errors"
	"database/sql"
	"github.com/lib/pq"
)

// Estudiante estructura del estudiante
type Estudiante struct {
	ID        int
	Name      string
	Age       int16
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Crear registra un estudiante en la BD
func Crear(e Estudiante) error {
	q := `INSERT INTO
		estudiantes (name, age, active)
		VALUES ($1, $2, $3)`

	intNull := sql.NullInt64{}
	strNull := sql.NullString{}

	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if e.Age == 0 {
		intNull.Valid = false
	} else {
		intNull.Valid = true
		intNull.Int64 = int64(e.Age)
	}

	if e.Name == "" {
		strNull.Valid = false
	} else {
		strNull.Valid = true
		strNull.String = e.Name
	}

	r, err := stmt.Exec(strNull, intNull, e.Active)
	if err != nil {
		return err
	}

	i, _ := r.RowsAffected()
	if i != 1 {
		return errors.New("Error: Se esperaba 1 fila afectada.")
	}

	return nil
}

// Consultar busca la información de los estudiantes (todos)
func Consultar() (estudiantes []Estudiante, err error) {
	q := `SELECT id, name, age, active, created_at, updated_at
		FROM estudiantes`

	timeNull := pq.NullTime{}
	intNull := sql.NullInt64{}
	strNull := sql.NullString{}
	boolNull := sql.NullBool{}

	db := getConnection()
	defer db.Close()

	rows, err := db.Query(q)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		e := Estudiante{}
		err = rows.Scan(
			&e.ID,
			&strNull,
			&intNull,
			&boolNull,
			&e.CreatedAt,
			&timeNull,
		)
		if err != nil {
			return
		}

		e.UpdatedAt = timeNull.Time
		e.Name = strNull.String
		e.Age = int16(intNull.Int64)
		e.Active = boolNull.Bool

		estudiantes = append(estudiantes, e)
	}

	return estudiantes, nil
}