package vehicletyperepository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/nauhalf/mobilcantik/repository/model"
	errorutils "github.com/nauhalf/mobilcantik/utils/error"
)

func GetAll(db *sql.DB) ([]model.VehicleType, error) {

	var types []model.VehicleType
	rows, err := db.Query("SELECT * FROM GEN_VehicleType")

	defer rows.Close()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	for rows.Next() {
		vtype := new(model.VehicleType)
		err := rows.Scan(
			&vtype.VehicleTypeId,
			&vtype.VehicleTypeName,
			&vtype.Annotation,
			&vtype.CreatedDate,
			&vtype.ModifiedDate,
		)

		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		types = append(types, *vtype)

		// var id uint64
		// var name, annotation string
		// var c_date time.Time
		// var m_date *time.Time
		// err := rows.Scan(
		// 	id,
		// 	name,
		// 	annotation,
		// 	c_date,
		// 	m_date,
		// )

		// if err != nil {
		// 	fmt.Println(err.Error())
		// 	return nil, err
		// }

		// types = append(types, model.VehicleType{VehicleTypeId: id, VehicleTypeName: name, Annotation: annotation, CreatedDate: c_date, ModifiedDate: m_date})

	}

	if err := rows.Err(); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return types, nil
}

func Create(db *sql.DB, vehicleType *model.VehicleType) (*model.VehicleType, error) {

	tx, err := db.Begin()

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	stmt, _ := tx.Prepare("INSERT INTO GEN_VehicleType (szVehicleTypeName, szAnnotation) VALUES (?, ?)")
	insert, err := stmt.Exec(vehicleType.VehicleTypeName, vehicleType.Annotation)

	if err != nil {
		tx.Rollback()
		fmt.Println(err.Error())
		return nil, err
	}
	stmt.Close()

	if err != nil {
		tx.Rollback()
		fmt.Println(err.Error())
		return nil, err
	}

	id, err := insert.LastInsertId()
	if err != nil {
		tx.Rollback()
		fmt.Println(err.Error())
		return nil, err
	}
	err = tx.QueryRow("SELECT * FROM GEN_VehicleType WHERE intVehicleTypeId = ? LIMIT 1", id).
		Scan(&vehicleType.VehicleTypeId,
			&vehicleType.VehicleTypeName,
			&vehicleType.Annotation,
			&vehicleType.CreatedDate,
			&vehicleType.ModifiedDate)

	if err != nil {
		tx.Rollback()
		fmt.Println(err.Error())
		return nil, err
	}

	err = tx.Commit()

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return vehicleType, nil

}

func Update(db *sql.DB, vehicleType *model.VehicleType) error {

	tx, err := db.Begin()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	stmt, _ := tx.Prepare("UPDATE GEN_VehicleType SET szVehicleTypeName=?, szAnnotation=?, dtmModifiedDate=NOW() WHERE intVehicleTypeId=?")
	update, err := stmt.Exec(vehicleType.VehicleTypeName, vehicleType.Annotation, vehicleType.VehicleTypeId)

	if err != nil {
		tx.Rollback()
		fmt.Println(err.Error())
		return err
	}
	stmt.Close()

	rows, err := update.RowsAffected()
	if err != nil {
		tx.Rollback()
		fmt.Println(err.Error())
		return err
	}

	if rows == 0 {
		tx.Rollback()
		e := errors.New(errorutils.StatusZeroAffectedRows)
		return e
	}
	err = tx.Commit()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func Delete(db *sql.DB, id uint64) error {
	tx, err := db.Begin()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	stmt, _ := tx.Prepare("DELETE FROM GEN_VehicleType WHERE intVehicleTypeId=?")
	delete, err := stmt.Exec(id)

	if err != nil {
		tx.Rollback()
		fmt.Println(err.Error())
		return err
	}
	stmt.Close()

	rows, err := delete.RowsAffected()
	if err != nil {
		tx.Rollback()
		fmt.Println(err.Error())
		return err
	}

	if rows == 0 {
		tx.Rollback()
		e := errors.New(errorutils.StatusZeroAffectedRows)
		return e
	}

	err = tx.Commit()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil

}

func GetById(db *sql.DB, id uint64) (*model.VehicleType, error) {
	vehicleType := new(model.VehicleType)
	err := db.QueryRow("SELECT * FROM GEN_VehicleType WHERE intVehicleTypeId = ?", id).
		Scan(&vehicleType.VehicleTypeId,
			&vehicleType.VehicleTypeName,
			&vehicleType.Annotation,
			&vehicleType.CreatedDate,
			&vehicleType.ModifiedDate)

	if err != nil {
		return nil, err
	}

	return vehicleType, nil
}
