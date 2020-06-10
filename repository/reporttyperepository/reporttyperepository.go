package reporttyperepository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/nauhalf/mobilcantik/repository/model"
	"github.com/nauhalf/mobilcantik/repository/reportgrouprepository"
	errorutils "github.com/nauhalf/mobilcantik/utils/error"
)

type ReportTypeResponse struct {
	model.ReportGroup
	ReportTypes []model.ReportType `json:"reporttypes"`
}

type ReportTypeCreateResponse struct {
	model.ReportType
	ReportGroup model.ReportGroup `json:"reportgroup"`
}

func GetAll(db *sql.DB, reportgroup_id *uint64) ([]ReportTypeResponse, error) {

	respobj := []ReportTypeResponse{}

	if reportgroup_id == nil {
		groups, _ := reportgrouprepository.GetAll(db)
		reporttypes, _ := GetAllReportTypes(db)
		for _, rg := range groups {
			tmp := ReportTypeResponse{ReportTypes: []model.ReportType{}}
			tmp.ReportGroup = rg
			for _, fc := range reporttypes {
				if fc.ReportGroupId == rg.ReportGroupId {
					tmp.ReportTypes = append(tmp.ReportTypes, fc)
				}
			}
			respobj = append(respobj, tmp)
		}
	} else {
		group, _ := reportgrouprepository.GetById(db, *reportgroup_id)

		if group != nil {
			tmp := ReportTypeResponse{ReportTypes: []model.ReportType{}}
			reporttypes, _ := GetAllReportTypesByReportGroup(db, *reportgroup_id)
			tmp.ReportGroup = *group
			for _, fc := range reporttypes {
				if fc.ReportGroupId == group.ReportGroupId {
					tmp.ReportTypes = append(tmp.ReportTypes, fc)
				}
			}
			respobj = append(respobj, tmp)
		}
	}

	return respobj, nil
}

func GetAllReportTypes(db *sql.DB) ([]model.ReportType, error) {

	var reporttypes []model.ReportType
	rows, err := db.Query("SELECT * FROM GEN_ReportType")

	defer rows.Close()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	for rows.Next() {
		reporttype := new(model.ReportType)
		err := rows.Scan(
			&reporttype.ReportTypeId,
			&reporttype.ReportGroupId,
			&reporttype.ReportTypeName,
			&reporttype.Annotation,
			&reporttype.CreatedDate,
			&reporttype.ModifiedDate,
		)

		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		reporttypes = append(reporttypes, *reporttype)

	}

	if err := rows.Err(); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return reporttypes, nil
}

func GetAllReportTypesByReportGroup(db *sql.DB, reportgroup_id uint64) ([]model.ReportType, error) {

	var reporttypes []model.ReportType
	rows, err := db.Query("SELECT * FROM GEN_ReportType WHERE intReportGroupId = ?", reportgroup_id)

	defer rows.Close()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	for rows.Next() {
		reporttype := new(model.ReportType)
		err := rows.Scan(
			&reporttype.ReportTypeId,
			&reporttype.ReportGroupId,
			&reporttype.ReportTypeName,
			&reporttype.Annotation,
			&reporttype.CreatedDate,
			&reporttype.ModifiedDate,
		)

		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		reporttypes = append(reporttypes, *reporttype)

	}

	if err := rows.Err(); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return reporttypes, nil
}

func Create(db *sql.DB, reportgroup *model.ReportType) (*ReportTypeCreateResponse, error) {

	tx, err := db.Begin()

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	stmt, _ := tx.Prepare("INSERT INTO GEN_ReportType (intReportGroupId, szReportTypeName, szAnnotation) VALUES (?, ?, ?)")
	insert, err := stmt.Exec(reportgroup.ReportGroupId, reportgroup.ReportTypeName, reportgroup.Annotation)

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
	respobj := ReportTypeCreateResponse{}
	groups := new(model.ReportGroup)
	err = tx.QueryRow("SELECT * FROM GEN_ReportGroup WHERE intReportGroupId = ?", reportgroup.ReportGroupId).
		Scan(&groups.ReportGroupId,
			&groups.ReportGroupName,
			&groups.Annotation,
			&groups.CreatedDate,
			&groups.ModifiedDate)

	if err != nil {
		tx.Rollback()
		fmt.Println(err.Error())
		return nil, err
	}

	respobj.ReportGroup = *groups
	newReportType := new(model.ReportType)

	err = tx.QueryRow("SELECT * FROM GEN_ReportType WHERE intReportTypeId = ?", id).
		Scan(&newReportType.ReportTypeId,
			&newReportType.ReportGroupId,
			&newReportType.ReportTypeName,
			&newReportType.Annotation,
			&newReportType.CreatedDate,
			&newReportType.ModifiedDate)

	if err != nil {
		tx.Rollback()
		fmt.Println(err.Error())
		return nil, err
	}

	respobj.ReportType = *newReportType
	err = tx.Commit()

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &respobj, nil

}

func Update(db *sql.DB, reportgroup *model.ReportType) error {

	tx, err := db.Begin()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	stmt, _ := tx.Prepare("UPDATE GEN_ReportType SET intReportGroupId=?, szReportTypeName=?, szAnnotation=?, dtmModifiedDate=NOW() WHERE intReportTypeId=?")
	update, err := stmt.Exec(reportgroup.ReportGroupId, reportgroup.ReportTypeName, reportgroup.Annotation, reportgroup.ReportTypeId)

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

	stmt, _ := tx.Prepare("DELETE FROM GEN_ReportType WHERE intReportTypeId=?")
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

func GetById(db *sql.DB, id uint64) (*ReportTypeCreateResponse, error) {

	respobj := ReportTypeCreateResponse{}

	reportType := new(model.ReportType)

	err := db.QueryRow("SELECT * FROM GEN_ReportType WHERE intReportTypeId = ?", id).
		Scan(&reportType.ReportTypeId,
			&reportType.ReportGroupId,
			&reportType.ReportTypeName,
			&reportType.Annotation,
			&reportType.CreatedDate,
			&reportType.ModifiedDate)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	respobj.ReportType = *reportType

	rgroup := new(model.ReportGroup)
	err = db.QueryRow("SELECT * FROM GEN_ReportGroup WHERE intReportGroupId = ?", reportType.ReportGroupId).
		Scan(&rgroup.ReportGroupId,
			&rgroup.ReportGroupName,
			&rgroup.Annotation,
			&rgroup.CreatedDate,
			&rgroup.ModifiedDate)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	respobj.ReportGroup = *rgroup

	return &respobj, nil
}
