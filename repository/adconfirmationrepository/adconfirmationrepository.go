package adconfirmationrepository

import (
	"database/sql"
	"fmt"

	"github.com/nauhalf/mobilcantik/repository/model"
)

type AdConfirmationResponse struct {
	model.AdConfirmation
	Path string `json:"szPath" form:"szPath"`
}

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

func Create(db *sql.DB, adconfirmation *model.AdConfirmation, filename string, filelocation string) error {

	tx, err := db.Begin()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	qryInsertAdConfirmation := `INSERT INTO APP_AdConfirmation (
									intAdId,
									intAdBillId,
									intBankAccountId,
									szBankAccountNumber,
									szBankName,
									szBankBeneficiaryName,
									szAnnotation,
									dtmDate
								) VALUES (
									?,
									?,
									?,
									?,
									?,
									?,
									?,
									NOW()
								)`

	stmt, _ := tx.Prepare(qryInsertAdConfirmation)
	insert, err := stmt.Exec(adconfirmation.AdId,
		adconfirmation.AdBillId,
		adconfirmation.BankAccountId,
		adconfirmation.BankAccountNumber,
		adconfirmation.BankName,
		adconfirmation.BankBeneficiaryName,
		adconfirmation.Annotation)

	if err != nil {
		tx.Rollback()
		fmt.Println(err.Error())
		return err
	}
	stmt.Close()

	if err != nil {
		tx.Rollback()
		fmt.Println(err.Error())
		return err
	}

	id, err := insert.LastInsertId()
	if err != nil {
		tx.Rollback()
		fmt.Println(err.Error())
		return err
	}

	qryImages := `INSERT INTO APP_AdConfirmationImage (
						intAdConfirmationId,
						szDirectory,
						szFilename
					) 
					VALUES (
						?,
						?,
						?
					)`
	stmtImage, _ := tx.Prepare(qryImages)
	_, err = stmtImage.Exec(id, filelocation, filename)
	if err != nil {
		tx.Rollback()
		fmt.Println(err.Error())
		return err
	}

	stmtImage.Close()

	err = tx.Commit()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil

}

func GetByAdId(db *sql.DB, adId uint64) ([]AdConfirmationResponse, error) {
	var adconfirmations []AdConfirmationResponse
	rows, err := db.Query("SELECT APP_AdConfirmation.*, APP_AdConfirmationImage.szDirectory, APP_AdConfirmationImage.szFilename FROM APP_AdConfirmation LEFT JOIN APP_AdConfirmationImage ON APP_AdConfirmationImage.intAdConfirmationId = APP_AdConfirmation.intAdConfirmationId WHERE intAdId = ? ", adId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		adconfirmationres := new(AdConfirmationResponse)
		adconfirmation := new(model.AdConfirmation)
		var dir, filename string
		err := rows.Scan(
			&adconfirmation.AdConfirmationId,
			&adconfirmation.AdId,
			&adconfirmation.AdBillId,
			&adconfirmation.BankAccountId,
			&adconfirmation.BankAccountNumber,
			&adconfirmation.BankName,
			&adconfirmation.BankBeneficiaryName,
			&adconfirmation.Annotation,
			&adconfirmation.Date,
			&adconfirmation.CreatedDate,
			&adconfirmation.ModifiedDate,
			&dir,
			&filename,
		)

		path := dir + filename
		adconfirmationres.Path = path
		adconfirmationres.AdConfirmation = *adconfirmation
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		adconfirmations = append(adconfirmations, *adconfirmationres)
	}

	if err != nil {
		return nil, err
	}

	return adconfirmations, nil
}
