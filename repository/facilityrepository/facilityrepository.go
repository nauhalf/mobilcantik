package facilityrepository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/nauhalf/mobilcantik/repository/model"
	"github.com/nauhalf/mobilcantik/repository/vehicletyperepository"
	errorutils "github.com/nauhalf/mobilcantik/utils/error"
)

type FacilityResponse struct {
	model.VehicleType
	Facilities []model.Facility `json:"facilities"`
}

type FacilityCreateResponse struct {
	model.Facility
	VehicleType model.VehicleType `json:"vehicletype"`
}

func GetAll(db *sql.DB, vehicletype_id *uint64) ([]FacilityResponse, error) {

	respobj := []FacilityResponse{}

	if vehicletype_id == nil {
		types, _ := vehicletyperepository.GetAll(db)
		facilities, _ := GetAllFacilities(db)
		for _, vt := range types {
			tmp := FacilityResponse{Facilities: []model.Facility{}}
			tmp.VehicleType = vt
			for _, fc := range facilities {
				if fc.VehicleTypeId == vt.VehicleTypeId {
					tmp.Facilities = append(tmp.Facilities, fc)
				}
			}
			respobj = append(respobj, tmp)
		}
	} else {
		types, _ := vehicletyperepository.GetById(db, *vehicletype_id)

		if types != nil {
			tmp := FacilityResponse{Facilities: []model.Facility{}}
			facilities, _ := GetAllFacilitiesByVehicleType(db, *vehicletype_id)
			tmp.VehicleType = *types
			for _, fc := range facilities {
				if fc.VehicleTypeId == types.VehicleTypeId {
					tmp.Facilities = append(tmp.Facilities, fc)
				}
			}
			respobj = append(respobj, tmp)
		}
	}

	return respobj, nil
}

func GetAllFacilities(db *sql.DB) ([]model.Facility, error) {

	var facilities []model.Facility
	rows, err := db.Query("SELECT * FROM GEN_Facility")

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		facility := new(model.Facility)
		err := rows.Scan(
			&facility.FacilityId,
			&facility.VehicleTypeId,
			&facility.FacilityName,
			&facility.Annotation,
			&facility.CreatedDate,
			&facility.ModifiedDate,
		)

		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		facilities = append(facilities, *facility)

	}

	if err := rows.Err(); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return facilities, nil
}

func GetAllFacilitiesByVehicleType(db *sql.DB, vehicletype_id uint64) ([]model.Facility, error) {

	var facilities []model.Facility
	rows, err := db.Query("SELECT * FROM GEN_Facility WHERE intVehicleTypeId = ?", vehicletype_id)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		facility := new(model.Facility)
		err := rows.Scan(
			&facility.FacilityId,
			&facility.VehicleTypeId,
			&facility.FacilityName,
			&facility.Annotation,
			&facility.CreatedDate,
			&facility.ModifiedDate,
		)

		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		facilities = append(facilities, *facility)

	}

	if err := rows.Err(); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return facilities, nil
}

func Create(db *sql.DB, facility *model.Facility) (*FacilityCreateResponse, error) {

	tx, err := db.Begin()

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	stmt, _ := tx.Prepare("INSERT INTO GEN_Facility (intVehicleTypeId, szFacilityName, szAnnotation) VALUES (?, ?, ?)")
	insert, err := stmt.Exec(facility.VehicleTypeId, facility.FacilityName, facility.Annotation)

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
	respobj := FacilityCreateResponse{}
	types := new(model.VehicleType)
	err = tx.QueryRow("SELECT * FROM GEN_VehicleType WHERE intVehicleTypeId = ?", facility.VehicleTypeId).
		Scan(&types.VehicleTypeId,
			&types.VehicleTypeName,
			&types.Annotation,
			&types.CreatedDate,
			&types.ModifiedDate)

	if err != nil {
		tx.Rollback()
		fmt.Println(err.Error())
		return nil, err
	}

	respobj.VehicleType = *types
	newFacility := new(model.Facility)

	err = tx.QueryRow("SELECT * FROM GEN_Facility WHERE intFacilityId = ?", id).
		Scan(&newFacility.FacilityId,
			&newFacility.VehicleTypeId,
			&newFacility.FacilityName,
			&newFacility.Annotation,
			&newFacility.CreatedDate,
			&newFacility.ModifiedDate)

	if err != nil {
		tx.Rollback()
		fmt.Println(err.Error())
		return nil, err
	}

	respobj.Facility = *newFacility
	err = tx.Commit()

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &respobj, nil

}

func Update(db *sql.DB, facility *model.Facility) error {

	tx, err := db.Begin()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	stmt, _ := tx.Prepare("UPDATE GEN_Facility SET intVehicleTypeId=?, szFacilityName=?, szAnnotation=?, dtmModifiedDate=NOW() WHERE intFacilityId=?")
	update, err := stmt.Exec(facility.VehicleTypeId, facility.FacilityName, facility.Annotation, facility.FacilityId)

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

	stmt, _ := tx.Prepare("DELETE FROM GEN_Facility WHERE intFacilityId=?")
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

func GetById(db *sql.DB, id uint64) (*FacilityCreateResponse, error) {

	respobj := FacilityCreateResponse{}

	facility := new(model.Facility)

	err := db.QueryRow("SELECT * FROM GEN_Facility WHERE intFacilityId = ?", id).
		Scan(&facility.FacilityId,
			&facility.VehicleTypeId,
			&facility.FacilityName,
			&facility.Annotation,
			&facility.CreatedDate,
			&facility.ModifiedDate)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	respobj.Facility = *facility

	vtypes := new(model.VehicleType)
	err = db.QueryRow("SELECT * FROM GEN_VehicleType WHERE intVehicleTypeId = ?", facility.VehicleTypeId).
		Scan(&vtypes.VehicleTypeId,
			&vtypes.VehicleTypeName,
			&vtypes.Annotation,
			&vtypes.CreatedDate,
			&vtypes.ModifiedDate)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	respobj.VehicleType = *vtypes

	return &respobj, nil
}

func IsExists(db *sql.DB, vehicletype_id uint64, ids []uint64) (bool, error) {

	qry := `SELECT 
  COUNT(f.intFacilityId)
FROM
  GEN_Facility f
  WHERE f.intVehicleTypeId = ?
  AND f.intFacilityId IN (`
	binds := []interface{}{}
	binds = append(binds, vehicletype_id)
	for i, id := range ids {
		qry += "?"
		binds = append(binds, id)

		if i < len(ids)-1 {
			qry += ", "
		} else {
			qry += ")"
		}
	}

	var count int

	err := db.QueryRow(qry, binds...).
		Scan(&count)

	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	if count == len(ids) {
		return true, nil
	} else {
		return false, nil
	}
}
