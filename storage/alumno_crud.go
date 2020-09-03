package storage

import (
	"database/sql"
	"fmt"

	"github.com/Leonardo-Antonio/rest/model"
)

const (
	sqlCreate  = "INSERT INTO tb_alumnos VALUES (NULL, ?, ? ,?)"
	sqlGetAll  = "SELECT id, fullname, dni, age FROM tb_alumnos"
	sqlGetByID = "SELECT id, fullname, dni, age FROM tb_alumnos WHERE id = ?"
	sqlDelete  = "DELETE FROM tb_alumnos WHERE id = ?"
	sqlUpdate  = "UPDATE tb_alumnos SET fullname = ?, dni = ?, age = ?  WHERE id = ?"
)

// Alumn is class
type Alumn struct {
	db *sql.DB
}

// NewAlumn is the method constructor of class
func NewAlumn(db *sql.DB) *Alumn {
	return &Alumn{db}
}

// Create is a method implement of handler.storage
func (a *Alumn) Create(alumn model.Alumn) error {
	stmt, err := a.db.Prepare(sqlCreate)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rs, err := stmt.Exec(
		StringNull(alumn.Fullname),
		alumn.Dni,
		IntNull(int32(alumn.Age)),
	)
	if err != nil {
		return err
	}
	rAff, err := rs.RowsAffected()
	if err != nil {
		return err
	}
	if rAff != 1 {
		return ErrorRowAffected
	}
	fmt.Println("rows affected", rAff)
	return nil
}

// GetAll is a method implement of handler.storage
func (a *Alumn) GetAll() (alumns []model.Alumn, err error) {
	stmt, err := a.db.Prepare(sqlGetAll)
	if err != nil {
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return
	}

	// valid nulls
	fullnameNull := sql.NullString{}
	ageNull := sql.NullInt32{}

	for rows.Next() {
		alumn := model.Alumn{}
		err := rows.Scan(
			&alumn.ID,
			&fullnameNull,
			&alumn.Dni,
			&ageNull,
		)
		if err != nil {
			return nil, err
		}
		alumn.Fullname = fullnameNull.String
		alumn.Age = uint8(ageNull.Int32)
		alumns = append(alumns, alumn)
	}
	return
}

// GetByID is a method implement of handler.storage
func (a *Alumn) GetByID(id int) (alumn model.Alumn, err error) {
	stmt, err := a.db.Prepare(sqlGetByID)
	if err != nil {
		return
	}
	defer stmt.Close()

	// valid null
	fullnameNull := sql.NullString{}
	ageNull := sql.NullInt32{}

	err = stmt.QueryRow(id).Scan(
		&alumn.ID,
		&fullnameNull,
		&alumn.Dni,
		&ageNull,
	)
	if err != nil {
		return
	}
	alumn.Fullname = fullnameNull.String
	alumn.Age = uint8(ageNull.Int32)

	return
}

// Delete is a method implement of handler.storage
func (a *Alumn) Delete(id int) error {
	stmt, err := a.db.Prepare(sqlDelete)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rs, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	if rA, _ := rs.RowsAffected(); rA != 1 {
		return model.ErrorRowAffected
	}
	return nil
}

// Update is a method implement of handler.storage
func (a *Alumn) Update(id int, alumn model.Alumn) error {
	stmt, err := a.db.Prepare(sqlUpdate)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rs, err := stmt.Exec(
		StringNull(alumn.Fullname),
		alumn.Dni,
		IntNull(int32(alumn.Age)),
		id,
	)
	if err != nil {
		return err
	}
	if rA, _ := rs.RowsAffected(); rA != 1 {
		return model.ErrorRowAffected
	}
	return nil
}
