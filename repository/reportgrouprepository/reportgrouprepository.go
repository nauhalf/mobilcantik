package reportgrouprepository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/nauhalf/mobilcantik/repository/model"
	errorutils "github.com/nauhalf/mobilcantik/utils/error"
)

func GetAll(db *sql.DB) ([]model.ReportGroup, error) {

	var reportgroups []model.ReportGroup
	rows, err := db.Query("SELECT * FROM GEN_ReportGroup")

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		reportgroup := new(model.ReportGroup)
		err := rows.Scan(
			&reportgroup.ReportGroupId,
			&reportgroup.ReportGroupName,
			&reportgroup.Annotation,
			&reportgroup.CreatedDate,
			&reportgroup.ModifiedDate,
		)

		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		reportgroups = append(reportgroups, *reportgroup)

	}

	if err := rows.Err(); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return reportgroups, nil
}

func Create(db *sql.DB, reportgroup *model.ReportGroup) (*model.ReportGroup, error) {

	tx, err := db.Begin()

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	stmt, _ := tx.Prepare("INSERT INTO GEN_ReportGroup (szReportGroupName, szAnnotation) VALUES (?, ?)")
	insert, err := stmt.Exec(reportgroup.ReportGroupName, reportgroup.Annotation)

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

func Update(db *sql.DB, reportgroup *model.ReportGroup) error {

	tx, err := db.Begin()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	stmt, _ := tx.Prepare("UPDATE GEN_ReportGroup SET szReportGroupName=?, szAnnotation=?, dtmModifiedDate=NOW() WHERE intReportGroupId=?")
	update, err := stmt.Exec(reportgroup.ReportGroupName, reportgroup.Annotation, reportgroup.ReportGroupId)

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

	stmt, _ := tx.Prepare("DELETE FROM GEN_ReportGroup WHERE intReportGroupId=?")
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

func GetById(db *sql.DB, id uint64) (*model.ReportGroup, error) {
	reportgroup := new(model.ReportGroup)
	err := db.QueryRow("SELECT * FROM GEN_ReportGroup WHERE intReportGroupId = ?", id).
		Scan(&reportgroup.ReportGroupId,
			&reportgroup.ReportGroupName,
			&reportgroup.Annotation,
			&reportgroup.CreatedDate,
			&reportgroup.ModifiedDate)

	if err != nil {
		return nil, err
	}

	return reportgroup, nil
}
