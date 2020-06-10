package adreportrepository

import (
	"database/sql"
	"errors"
	"fmt"
	"math"

	"github.com/nauhalf/mobilcantik/repository/adrepository"
	"github.com/nauhalf/mobilcantik/repository/model"
	"github.com/nauhalf/mobilcantik/response"
	"github.com/nauhalf/mobilcantik/utils"
	errorutils "github.com/nauhalf/mobilcantik/utils/error"
)

type AdReportResponse struct {
	model.AdReport
	ReportTypeName string `json:"szReportTypeName"`
}

type AdReportSingleResponse struct {
	model.AdReport
	ReportTypeName   string                        `json:"szReportTypeName"`
	AdSingleResponse adrepository.AdSingleResponse `json:"ad"`
}

type GetAllResponse struct {
	Metadata  response.Metadata  `json:"metadata"`
	AdReports []AdReportResponse `json:"adreports"`
}

func GetAll(db *sql.DB, page int) (*GetAllResponse, error) {
	getAllResponse := new(GetAllResponse)
	metadata := new(response.Metadata)

	qryFrom := `FROM APP_AdReport adr 
  LEFT JOIN GEN_ReportType rt
  ON adr.intReportTypeId = rt.intReportTypeId `

	qrySelectTotal := `SELECT COUNT(*) AS total `

	qrySelectData := `SELECT adr.*, rt.szReportTypeName `

	binds := []interface{}{}

	orderBy := ` ORDER BY adr.intAdReportId DESC`
	joinQryTotal := qrySelectTotal + qryFrom + orderBy

	var total = 0

	limitbinds := []interface{}{}
	l, o := utils.GetLimitOffset(page, utils.GetPerPage())
	limitbinds = append(limitbinds, l, o)
	err := db.QueryRow(joinQryTotal, binds...).Scan(&total)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	metadata.Page = page
	metadata.TotalItems = total
	metadata.ItemPerPage = o
	metadata.TotalPage = int(math.Ceil(float64(total) / float64(o)))

	getAllResponse.Metadata = *metadata
	qryLimit := ` LIMIT ?, ?`

	joinQrySelect := qrySelectData + qryFrom + orderBy + qryLimit

	binds = append(binds, limitbinds...)
	rows, err := db.Query(joinQrySelect, binds...)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	for rows.Next() {
		adreportresponse := new(AdReportResponse)
		_ = rows.Scan(
			&adreportresponse.AdReportId,
			&adreportresponse.AdId,
			&adreportresponse.ReportTypeId,
			&adreportresponse.Annotation,
			&adreportresponse.ReportTime,
			&adreportresponse.ReporterName,
			&adreportresponse.ReporterEmail,
			&adreportresponse.CreatedDate,
			&adreportresponse.ModifiedDate,
			&adreportresponse.ReportTypeName,
		)
		getAllResponse.AdReports = append(getAllResponse.AdReports, *adreportresponse)
	}

	rows.Close()

	return getAllResponse, nil
}

func Create(db *sql.DB, adreport *model.AdReport) (*AdReportSingleResponse, error) {
	tx, err := db.Begin()

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	qry := `INSERT INTO APP_AdReport (
  intAdId,
  intReportTypeId,
  szAnnotation,
  dtmReportTime,
  szReporterName,
  szReporterEmail
) VALUES (
  ?,
  ?,
  ?,
  NOW(),
  ?,
  ?)`

	stmt, _ := tx.Prepare(qry)
	insert, err := stmt.Exec(adreport.AdId, adreport.ReportTypeId, adreport.Annotation, adreport.ReporterName, adreport.ReporterEmail)

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

func Delete(db *sql.DB, id uint64) error {
	tx, err := db.Begin()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	stmt, _ := tx.Prepare("DELETE FROM APP_AdReport WHERE intAdReportId=?")
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

func GetById(db *sql.DB, id uint64) (*AdReportSingleResponse, error) {
	adreportresponse := new(AdReportSingleResponse)
	err := db.QueryRow("SELECT adr.*, rt.szReportTypeName FROM APP_AdReport adr LEFT JOIN GEN_ReportType rt ON adr.intReportTypeId = rt.intReportTypeId  WHERE adr.intAdReportId = ?", id).
		Scan(&adreportresponse.AdReportId,
			&adreportresponse.AdId,
			&adreportresponse.ReportTypeId,
			&adreportresponse.Annotation,
			&adreportresponse.ReportTime,
			&adreportresponse.ReporterName,
			&adreportresponse.ReporterEmail,
			&adreportresponse.CreatedDate,
			&adreportresponse.ModifiedDate,
			&adreportresponse.ReportTypeName)

	if err != nil {
		return nil, err
	}

	return adreportresponse, nil
}
