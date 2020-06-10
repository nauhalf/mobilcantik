package adtyperepository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/nauhalf/mobilcantik/repository/model"
	errorutils "github.com/nauhalf/mobilcantik/utils/error"
)

func GetAll(db *sql.DB) ([]model.AdType, error) {

	var adtypes []model.AdType
	rows, err := db.Query("SELECT * FROM GEN_AdType")

	defer rows.Close()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	for rows.Next() {
		adtype := new(model.AdType)
		err := rows.Scan(
			&adtype.AdTypeId,
			&adtype.AdTypeName,
			&adtype.Annotation,
			&adtype.CreatedDate,
			&adtype.ModifiedDate,
		)

		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		adtypes = append(adtypes, *adtype)

	}

	if err := rows.Err(); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return adtypes, nil
}

func Create(db *sql.DB, adtype *model.AdType) (*model.AdType, error) {

	tx, err := db.Begin()

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	stmt, _ := tx.Prepare("INSERT INTO GEN_AdType (szAdTypeName, szAnnotation) VALUES (?, ?)")
	insert, err := stmt.Exec(adtype.AdTypeName, adtype.Annotation)

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

	err = tx.Commit()

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	c, _ := GetById(db, uint64(id))
	return c, nil

}

func Update(db *sql.DB, adtype *model.AdType) error {

	tx, err := db.Begin()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	stmt, _ := tx.Prepare("UPDATE GEN_AdType SET szAdTypeName=?, szAnnotation=?, dtmModifiedDate=NOW() WHERE intAdTypeId=?")
	update, err := stmt.Exec(adtype.AdTypeName, adtype.Annotation, adtype.AdTypeId)

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

	stmt, _ := tx.Prepare("DELETE FROM GEN_AdType WHERE intAdTypeId=?")
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

func GetById(db *sql.DB, id uint64) (*model.AdType, error) {
	adtype := new(model.AdType)
	err := db.QueryRow("SELECT * FROM GEN_AdType WHERE intAdTypeId = ?", id).
		Scan(&adtype.AdTypeId,
			&adtype.AdTypeName,
			&adtype.Annotation,
			&adtype.CreatedDate,
			&adtype.ModifiedDate)

	if err != nil {
		return nil, err
	}

	return adtype, nil
}
