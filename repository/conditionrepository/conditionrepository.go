package conditionrepository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/nauhalf/mobilcantik/repository/model"
	errorutils "github.com/nauhalf/mobilcantik/utils/error"
)

func GetAll(db *sql.DB) ([]model.Condition, error) {

	var conditions []model.Condition
	rows, err := db.Query("SELECT * FROM GEN_Condition")

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		condition := new(model.Condition)
		err := rows.Scan(
			&condition.ConditionId,
			&condition.ConditionName,
			&condition.Annotation,
			&condition.CreatedDate,
			&condition.ModifiedDate,
		)

		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		conditions = append(conditions, *condition)

	}

	if err := rows.Err(); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return conditions, nil
}

func Create(db *sql.DB, condition *model.Condition) (*model.Condition, error) {

	tx, err := db.Begin()

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	stmt, _ := tx.Prepare("INSERT INTO GEN_Condition (szConditionName, szAnnotation) VALUES (?, ?)")
	insert, err := stmt.Exec(condition.ConditionName, condition.Annotation)

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

func Update(db *sql.DB, condition *model.Condition) error {

	tx, err := db.Begin()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	stmt, _ := tx.Prepare("UPDATE GEN_Condition SET szConditionName=?, szAnnotation=?, dtmModifiedDate=NOW() WHERE intConditionId=?")
	update, err := stmt.Exec(condition.ConditionName, condition.Annotation, condition.ConditionId)

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

	stmt, _ := tx.Prepare("DELETE FROM GEN_Condition WHERE intConditionId=?")
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

func GetById(db *sql.DB, id uint64) (*model.Condition, error) {
	condition := new(model.Condition)
	err := db.QueryRow("SELECT * FROM GEN_Condition WHERE intConditionId = ?", id).
		Scan(&condition.ConditionId,
			&condition.ConditionName,
			&condition.Annotation,
			&condition.CreatedDate,
			&condition.ModifiedDate)

	if err != nil {
		return nil, err
	}

	return condition, nil
}
