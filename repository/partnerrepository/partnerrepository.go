package partnerrepository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/nauhalf/mobilcantik/repository/model"
	errorutils "github.com/nauhalf/mobilcantik/utils/error"
)

func GetAll(db *sql.DB) ([]model.Partner, error) {

	var partners []model.Partner
	rows, err := db.Query("SELECT * FROM GEN_Partner")

	defer rows.Close()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	for rows.Next() {
		partner := new(model.Partner)
		err := rows.Scan(
			&partner.PartnerId,
			&partner.Title,
			&partner.Url,
			&partner.Annotation,
			&partner.CreatedDate,
			&partner.ModifiedDate,
		)

		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		partners = append(partners, *partner)

	}

	if err := rows.Err(); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return partners, nil
}

func Create(db *sql.DB, partner *model.Partner) (*model.Partner, error) {

	tx, err := db.Begin()

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	stmt, _ := tx.Prepare("INSERT INTO GEN_Partner (szTitle, szUrl, szAnnotation) VALUES (?, ?, ?)")
	insert, err := stmt.Exec(partner.Title, partner.Url, partner.Annotation)

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

func Update(db *sql.DB, partner *model.Partner) error {

	tx, err := db.Begin()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	stmt, _ := tx.Prepare("UPDATE GEN_Partner SET szTitle=?, szUrl=?, szAnnotation=?, dtmModifiedDate=NOW() WHERE intPartnerId=?")
	update, err := stmt.Exec(partner.Title, partner.Url, partner.Annotation, partner.PartnerId)

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

	stmt, _ := tx.Prepare("DELETE FROM GEN_Partner WHERE intPartnerId=?")
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

func GetById(db *sql.DB, id uint64) (*model.Partner, error) {
	partner := new(model.Partner)
	err := db.QueryRow("SELECT * FROM GEN_Partner WHERE intPartnerId = ?", id).
		Scan(&partner.PartnerId,
			&partner.Title,
			&partner.Url,
			&partner.Annotation,
			&partner.CreatedDate,
			&partner.ModifiedDate)

	if err != nil {
		return nil, err
	}

	return partner, nil
}
