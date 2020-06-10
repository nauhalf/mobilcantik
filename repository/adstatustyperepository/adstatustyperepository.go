package adstatustyperepository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/nauhalf/mobilcantik/repository/model"
	errorutils "github.com/nauhalf/mobilcantik/utils/error"
)

func GetAll(db *sql.DB) ([]model.AdStatusType, error) {

	var adstatustypes []model.AdStatusType
	rows, err := db.Query("SELECT * FROM GEN_AdStatusType")

	defer rows.Close()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	for rows.Next() {
		adstatustype := new(model.AdStatusType)
		err := rows.Scan(
			&adstatustype.AdStatusTypeId,
			&adstatustype.AdStatusTypeName,
			&adstatustype.CreatedDate,
			&adstatustype.ModifiedDate,
		)

		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		adstatustypes = append(adstatustypes, *adstatustype)

	}

	if err := rows.Err(); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return adstatustypes, nil
}

func Create(db *sql.DB, adstatustype *model.AdStatusType) (*model.AdStatusType, error) {

	tx, err := db.Begin()

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	stmt, _ := tx.Prepare("INSERT INTO GEN_AdStatusType (szAdStatusTypeName) VALUES (?)")
	insert, err := stmt.Exec(adstatustype.AdStatusTypeName)

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

func Update(db *sql.DB, adstatustype *model.AdStatusType) error {

	tx, err := db.Begin()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	stmt, _ := tx.Prepare("UPDATE GEN_AdStatusType SET szAdStatusTypeName=?, dtmModifiedDate=NOW() WHERE intAdStatusTypeId=?")
	update, err := stmt.Exec(adstatustype.AdStatusTypeName, adstatustype.AdStatusTypeId)

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

	stmt, _ := tx.Prepare("DELETE FROM GEN_AdStatusType WHERE intAdStatusTypeId=?")
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

func GetById(db *sql.DB, id uint64) (*model.AdStatusType, error) {
	adstatustype := new(model.AdStatusType)
	err := db.QueryRow("SELECT * FROM GEN_AdStatusType WHERE intAdStatusTypeId = ?", id).
		Scan(&adstatustype.AdStatusTypeId,
			&adstatustype.AdStatusTypeName,
			&adstatustype.CreatedDate,
			&adstatustype.ModifiedDate)

	if err != nil {
		return nil, err
	}

	return adstatustype, nil
}
