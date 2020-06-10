package adbillrepository

import (
	"database/sql"
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/nauhalf/mobilcantik/repository/adconfirmationrepository"
	"github.com/nauhalf/mobilcantik/repository/model"
	"github.com/nauhalf/mobilcantik/response"
	"github.com/nauhalf/mobilcantik/utils"
	errorutils "github.com/nauhalf/mobilcantik/utils/error"
)

type AdBillSingleResponse struct {
	model.AdBill
	AdConfirmations []adconfirmationrepository.AdConfirmationResponse `json:"confirmations"`
}

type GetAllResponse struct {
	Metadata response.Metadata      `json:"metadata"`
	Bills    []AdBillSingleResponse `json:"bills"`
}

func GetAll(db *sql.DB, page int) (*GetAllResponse, error) {
	getAllResponse := new(GetAllResponse)
	metadata := new(response.Metadata)

	qryFrom := `FROM
	APP_AdBill adb `

	qrySelectTotal := `SELECT COUNT(*) AS total `

	qrySelectData := `SELECT adb.* `

	binds := []interface{}{}

	orderBy := ` ORDER BY adb.intAdBillId DESC`
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
		adbill := new(model.AdBill)
		adBillSingleResponse := new(AdBillSingleResponse)
		_ = rows.Scan(&adbill.AdBillId,
			&adbill.AdId,
			&adbill.Amount,
			&adbill.Balance,
			&adbill.CreatedDate,
			&adbill.ModifiedDate)
		adBillSingleResponse.AdBill = *adbill
		getAllResponse.Bills = append(getAllResponse.Bills, *adBillSingleResponse)
	}

	rows.Close()

	for index, g := range getAllResponse.Bills {
		adconfirmation, _ := adconfirmationrepository.GetByAdId(db, g.AdBill.AdId)
		getAllResponse.Bills[index].AdConfirmations = adconfirmation

	}

	return getAllResponse, nil
}

func ConfirmPaymentBill(db *sql.DB, adbill *model.AdBill) error {

	tx, err := db.Begin()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	stmtbill, _ := tx.Prepare("UPDATE APP_AdBill SET decBalance=?, dtmModifiedDate=NOW() WHERE intAdBillId=?")
	updatebill, err := stmtbill.Exec(adbill.Balance, adbill.AdBillId)

	if err != nil {
		tx.Rollback()
		fmt.Println(err.Error())
		return err
	}
	stmtbill.Close()

	rows, err := updatebill.RowsAffected()
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

	qry := `UPDATE APP_Ad 
		SET 
			isActive = 1,
			intAdStatusTypeId = 3,
			dtmPosted=?, 
			dtmConfirmation=?, 
			dtmExpired=?,
			dtmModifiedDate=NOW()
		WHERE intAdId=?`

	stmtad, _ := tx.Prepare(qry)
	now := time.Now()
	exp := now.AddDate(0, 0, utils.GetExpiredInterval())
	updatead, err := stmtad.Exec(now, now, exp, adbill.AdId)

	if err != nil {
		tx.Rollback()
		fmt.Println(err.Error())
		return err
	}
	stmtad.Close()

	rows, err = updatead.RowsAffected()
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

func GetById(db *sql.DB, id uint64) (*AdBillSingleResponse, error) {
	adBillSingleResponse := new(AdBillSingleResponse)
	adbill := new(model.AdBill)
	err := db.QueryRow("SELECT * FROM APP_AdBill WHERE intAdBillId = ?", id).
		Scan(&adbill.AdBillId,
			&adbill.AdId,
			&adbill.Amount,
			&adbill.Balance,
			&adbill.CreatedDate,
			&adbill.ModifiedDate)

	if err != nil {
		return nil, err
	}
	adBillSingleResponse.AdBill = *adbill
	return adBillSingleResponse, nil
}

func GetByAdId(db *sql.DB, adId uint64) (*model.AdBill, error) {
	adbill := new(model.AdBill)
	err := db.QueryRow("SELECT * FROM APP_AdBill WHERE intAdId = ?", adId).
		Scan(&adbill.AdBillId,
			&adbill.AdId,
			&adbill.Amount,
			&adbill.Balance,
			&adbill.CreatedDate,
			&adbill.ModifiedDate)

	if err != nil {
		return nil, err
	}

	return adbill, nil
}
