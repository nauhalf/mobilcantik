package bankaccountrepository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/nauhalf/mobilcantik/repository/model"
	errorutils "github.com/nauhalf/mobilcantik/utils/error"
)

func GetAll(db *sql.DB, is_active *bool) ([]model.BankAccount, error) {

	var bankaccounts []model.BankAccount
	if is_active == nil {
		banks, _ := GetAllBankAccount(db)

		bankaccounts = append(bankaccounts, banks...)

	} else {
		banks, _ := GetAllBankAccountByStatusActive(db, *is_active)
		bankaccounts = append(bankaccounts, banks...)
	}

	return bankaccounts, nil
}

func GetAllBankAccount(db *sql.DB) ([]model.BankAccount, error) {

	var banks []model.BankAccount
	rows, err := db.Query("SELECT * FROM GEN_BankAccount")

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		bankAccount := new(model.BankAccount)
		err := rows.Scan(
			&bankAccount.BankAccountId,
			&bankAccount.BankAccountNumber,
			&bankAccount.BankCode,
			&bankAccount.BankAccountName,
			&bankAccount.IsActive,
			&bankAccount.CreatedDate,
			&bankAccount.ModifiedDate,
		)

		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		banks = append(banks, *bankAccount)

	}

	if err := rows.Err(); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return banks, nil
}

func GetAllBankAccountByStatusActive(db *sql.DB, is_active bool) ([]model.BankAccount, error) {

	var banks []model.BankAccount
	rows, err := db.Query("SELECT * FROM GEN_BankAccount WHERE isActive = ?", is_active)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		bankAccount := new(model.BankAccount)
		err := rows.Scan(
			&bankAccount.BankAccountId,
			&bankAccount.BankAccountNumber,
			&bankAccount.BankCode,
			&bankAccount.BankAccountName,
			&bankAccount.IsActive,
			&bankAccount.CreatedDate,
			&bankAccount.ModifiedDate,
		)

		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		banks = append(banks, *bankAccount)

	}

	if err := rows.Err(); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return banks, nil
}

func Create(db *sql.DB, bankAccount *model.BankAccount) (*model.BankAccount, error) {

	tx, err := db.Begin()

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	stmt, _ := tx.Prepare("INSERT INTO GEN_BankAccount (szBankAccountNumber, szBankCode, szBankAccountName, isActive) VALUES (?, ?, ?, ?)")
	insert, err := stmt.Exec(bankAccount.BankAccountNumber, bankAccount.BankCode, bankAccount.BankAccountName, 1)

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

	b, _ := GetById(db, uint64(id))
	return b, nil

}

func Update(db *sql.DB, bankAccount *model.BankAccount) error {

	tx, err := db.Begin()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	stmt, _ := tx.Prepare("UPDATE GEN_BankAccount SET szBankAccountNumber=?, szBankCode=?, szBankAccountName=?, isActive=?, dtmModifiedDate=NOW() WHERE intBankAccountId=?")
	update, err := stmt.Exec(bankAccount.BankAccountNumber, bankAccount.BankCode, bankAccount.BankAccountName, bankAccount.IsActive, bankAccount.BankAccountId)

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

	stmt, _ := tx.Prepare("DELETE FROM GEN_BankAccount WHERE intBankAccountId=?")
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

func GetById(db *sql.DB, id uint64) (*model.BankAccount, error) {
	bankAccount := new(model.BankAccount)
	err := db.QueryRow("SELECT intBankAccountId, szBankAccountNumber, szBankCode, szBankAccountName, isActive, dtmCreatedDate, dtmModifiedDate FROM GEN_BankAccount WHERE intBankAccountId = ?", id).
		Scan(
			&bankAccount.BankAccountId,
			&bankAccount.BankAccountNumber,
			&bankAccount.BankCode,
			&bankAccount.BankAccountName,
			&bankAccount.IsActive,
			&bankAccount.CreatedDate,
			&bankAccount.ModifiedDate,
		)
	if err != nil {
		return nil, err
	}

	return bankAccount, nil
}
