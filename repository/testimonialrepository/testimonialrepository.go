package testimonialrepository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/nauhalf/mobilcantik/repository/model"
	errorutils "github.com/nauhalf/mobilcantik/utils/error"
)

func GetAll(db *sql.DB) ([]model.Testimonial, error) {

	var testimonials []model.Testimonial
	rows, err := db.Query("SELECT * FROM APP_Testimonial WHERE isActive = 1 ORDER BY intTestimonialId DESC LIMIT 10")

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		testimonial := new(model.Testimonial)
		err := rows.Scan(
			&testimonial.TestimonialId,
			&testimonial.AdId,
			&testimonial.Comment,
			&testimonial.AdClosingReasonId,
			&testimonial.DtmDate,
			&testimonial.IsActive,
			&testimonial.CreatedDate,
			&testimonial.ModifiedDate,
		)

		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		testimonials = append(testimonials, *testimonial)

	}

	if err := rows.Err(); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return testimonials, nil
}

func Create(db *sql.DB, testimonial *model.Testimonial) (*model.Testimonial, error) {

	tx, err := db.Begin()

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	stmt, _ := tx.Prepare("INSERT INTO APP_Testimonial (intAdId, szComment, intAdClosingReasonId, isActive) VALUES (?, ?, ?, ?)")
	insert, err := stmt.Exec(testimonial.AdId, testimonial.Comment, testimonial.AdClosingReasonId, testimonial.IsActive)

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

	stmtClose, _ := tx.Prepare("UPDATE APP_Ad SET isActive=0, dtmModifiedDate=NOW() WHERE intAdId=?")
	update, err := stmtClose.Exec(testimonial.AdId)

	if err != nil {
		tx.Rollback()
		fmt.Println(err.Error())
		return nil, err
	}
	rows, err := update.RowsAffected()
	if err != nil {
		tx.Rollback()
		fmt.Println(err.Error())
		return nil, err
	}

	stmtClose.Close()
	if rows == 0 {
		tx.Rollback()
		e := errors.New(errorutils.StatusZeroAffectedRows)
		return nil, e
	}

	err = tx.Commit()

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	c, _ := GetById(db, uint64(id))
	return c, nil

}

func GetById(db *sql.DB, id uint64) (*model.Testimonial, error) {
	testimonial := new(model.Testimonial)
	err := db.QueryRow("SELECT * FROM APP_Testimonial WHERE intTestimonialId = ?", id).
		Scan(
			&testimonial.TestimonialId,
			&testimonial.AdId,
			&testimonial.Comment,
			&testimonial.AdClosingReasonId,
			&testimonial.DtmDate,
			&testimonial.IsActive,
			&testimonial.CreatedDate,
			&testimonial.ModifiedDate,
		)

	if err != nil {
		return nil, err
	}

	return testimonial, nil
}

func GetByAdId(db *sql.DB, id uint64) (*model.Testimonial, error) {
	testimonial := new(model.Testimonial)
	err := db.QueryRow("SELECT * FROM APP_Testimonial WHERE intAdId = ?", id).
		Scan(
			&testimonial.TestimonialId,
			&testimonial.AdId,
			&testimonial.Comment,
			&testimonial.AdClosingReasonId,
			&testimonial.DtmDate,
			&testimonial.IsActive,
			&testimonial.CreatedDate,
			&testimonial.ModifiedDate,
		)

	if err != nil {
		return nil, err
	}

	return testimonial, nil
}
