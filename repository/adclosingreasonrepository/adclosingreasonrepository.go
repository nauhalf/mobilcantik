package adclosingreasonrepository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/nauhalf/mobilcantik/repository/model"
	errorutils "github.com/nauhalf/mobilcantik/utils/error"
)

func GetAll(db *sql.DB) ([]model.AdClosingReason, error) {

	var adclosingreasons []model.AdClosingReason
	rows, err := db.Query("SELECT * FROM GEN_AdClosingReason")

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		adclosingreason := new(model.AdClosingReason)
		err := rows.Scan(
			&adclosingreason.AdClosingReasonId,
			&adclosingreason.ReasonName,
			&adclosingreason.Annotation,
			&adclosingreason.CreatedDate,
			&adclosingreason.ModifiedDate,
		)

		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		adclosingreasons = append(adclosingreasons, *adclosingreason)

	}

	if err := rows.Err(); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return adclosingreasons, nil
}

func Create(db *sql.DB, adclosingreason *model.AdClosingReason) (*model.AdClosingReason, error) {

	tx, err := db.Begin()

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	stmt, _ := tx.Prepare("INSERT INTO GEN_AdClosingReason (szReasonName, szAnnotation) VALUES (?, ?)")
	insert, err := stmt.Exec(adclosingreason.ReasonName, adclosingreason.Annotation)

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

func Update(db *sql.DB, adclosingreason *model.AdClosingReason) error {

	tx, err := db.Begin()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	stmt, _ := tx.Prepare("UPDATE GEN_AdClosingReason SET szReasonName=?, szAnnotation=?, dtmModifiedDate=NOW() WHERE intAdClosingReasonId=?")
	update, err := stmt.Exec(adclosingreason.ReasonName, adclosingreason.Annotation, adclosingreason.AdClosingReasonId)

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

	stmt, _ := tx.Prepare("DELETE FROM GEN_AdClosingReason WHERE intAdClosingReasonId=?")
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

func GetById(db *sql.DB, id uint64) (*model.AdClosingReason, error) {
	adclosingreason := new(model.AdClosingReason)
	err := db.QueryRow("SELECT * FROM GEN_AdClosingReason WHERE intAdClosingReasonId = ?", id).
		Scan(&adclosingreason.AdClosingReasonId,
			&adclosingreason.ReasonName,
			&adclosingreason.Annotation,
			&adclosingreason.CreatedDate,
			&adclosingreason.ModifiedDate)

	if err != nil {
		return nil, err
	}

	return adclosingreason, nil
}
