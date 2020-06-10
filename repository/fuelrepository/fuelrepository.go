package fuelrepository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/nauhalf/mobilcantik/repository/model"
	"github.com/nauhalf/mobilcantik/repository/vehicletyperepository"
	errorutils "github.com/nauhalf/mobilcantik/utils/error"
)

type FuelMappingResponse struct {
	VehicleTypeId   uint64         `json:"intVehicleTypeId" form:"intVehicleTypeId"`
	VehicleTypeName string         `json:"szVehicleTypeName" form:"szVehicleTypeName"`
	Annotation      *string        `json:"szAnnotation" form:"szAnnotation"`
	Fuels           []FuelResponse `json:"fuels"`
}

type FuelResponse struct {
	FuelId     uint64  `json:"intFuelId" form:"intFuelId"`
	FuelName   string  `json:"szFuelName" form:"szFuelName"`
	Annotation *string `json:"szAnnotation" form:"szAnnotation"`
}

type VehicleTypeFuelMapping struct {
	FuelResponse
	VehicleTypeId uint64 `json:"intVehicleTypeId"`
}

func GetAll(db *sql.DB) ([]model.Fuel, error) {

	var types []model.Fuel
	rows, err := db.Query("SELECT * FROM GEN_Fuel")

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		vtype := new(model.Fuel)
		err := rows.Scan(
			&vtype.FuelId,
			&vtype.FuelName,
			&vtype.Annotation,
			&vtype.CreatedDate,
			&vtype.ModifiedDate,
		)

		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		types = append(types, *vtype)

	}

	if err := rows.Err(); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return types, nil
}

func Create(db *sql.DB, fuel *model.Fuel) (*model.Fuel, error) {

	tx, err := db.Begin()

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	stmt, _ := tx.Prepare("INSERT INTO GEN_Fuel (szFuelName, szAnnotation) VALUES (?, ?)")
	insert, err := stmt.Exec(fuel.FuelName, fuel.Annotation)

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
	err = tx.QueryRow("SELECT * FROM GEN_Fuel WHERE intFuelId = ? LIMIT 1", id).
		Scan(&fuel.FuelId,
			&fuel.FuelName,
			&fuel.Annotation,
			&fuel.CreatedDate,
			&fuel.ModifiedDate)

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

	return fuel, nil

}

func Update(db *sql.DB, fuel *model.Fuel) error {

	tx, err := db.Begin()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	stmt, _ := tx.Prepare("UPDATE GEN_Fuel SET szFuelName=?, szAnnotation=?, dtmModifiedDate=NOW() WHERE intFuelId=?")
	update, err := stmt.Exec(fuel.FuelName, fuel.Annotation, fuel.FuelId)

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

	stmt, _ := tx.Prepare("DELETE FROM GEN_Fuel WHERE intFuelId=?")
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

func GetById(db *sql.DB, id uint64) (*model.Fuel, error) {
	fuel := new(model.Fuel)
	err := db.QueryRow("SELECT * FROM GEN_Fuel WHERE intFuelId = ?", id).
		Scan(&fuel.FuelId,
			&fuel.FuelName,
			&fuel.Annotation,
			&fuel.CreatedDate,
			&fuel.ModifiedDate)

	if err != nil {
		return nil, err
	}

	return fuel, nil
}

func GetAllMapping(db *sql.DB, vehicletype_id *uint64) ([]FuelMappingResponse, error) {

	respobj := []FuelMappingResponse{}

	if vehicletype_id == nil {
		types, _ := vehicletyperepository.GetAll(db)
		fuels, _ := GetVehicleFuelMapping(db)
		for _, vt := range types {
			tmp := FuelMappingResponse{
				VehicleTypeId:   vt.VehicleTypeId,
				VehicleTypeName: vt.VehicleTypeName,
				Annotation:      vt.Annotation,
				Fuels:           []FuelResponse{},
			}
			for _, f := range fuels {
				if f.VehicleTypeId == vt.VehicleTypeId {
					tmp.Fuels = append(tmp.Fuels, f.FuelResponse)
				}
			}
			respobj = append(respobj, tmp)
		}
	} else {
		types, _ := vehicletyperepository.GetById(db, *vehicletype_id)

		if types != nil {
			tmp := FuelMappingResponse{
				VehicleTypeId:   types.VehicleTypeId,
				VehicleTypeName: types.VehicleTypeName,
				Annotation:      types.Annotation,
				Fuels:           []FuelResponse{},
			}
			fuels, _ := GetVehicleFuelMappingById(db, types.VehicleTypeId)

			for _, f := range fuels {
				if f.VehicleTypeId == types.VehicleTypeId {
					tmp.Fuels = append(tmp.Fuels, f.FuelResponse)
				}
			}
			respobj = append(respobj, tmp)
		}
	}

	return respobj, nil
}

func GetVehicleFuelMapping(db *sql.DB) ([]VehicleTypeFuelMapping, error) {

	var vf []VehicleTypeFuelMapping
	rows, err := db.Query("SELECT f.intFuelId, f.szFuelName, f.szAnnotation, vf.intVehicleTypeId FROM GEN_Fuel f LEFT JOIN GEN_VehicleFuel vf ON vf.intFuelId = f.intFuelId")

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		fr := new(FuelResponse)
		var vId uint64
		err := rows.Scan(
			&fr.FuelId,
			&fr.FuelName,
			&fr.Annotation,
			&vId,
		)

		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		vf = append(vf, VehicleTypeFuelMapping{FuelResponse: *fr, VehicleTypeId: vId})

	}

	if err := rows.Err(); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return vf, nil
}

func GetVehicleFuelMappingById(db *sql.DB, vehicletype_id uint64) ([]VehicleTypeFuelMapping, error) {

	var vf []VehicleTypeFuelMapping
	rows, err := db.Query("SELECT f.intFuelId, f.szFuelName, f.szAnnotation, vf.intVehicleTypeId FROM GEN_Fuel f LEFT JOIN GEN_VehicleFuel vf ON vf.intFuelId = f.intFuelId WHERE vf.intVehicleTypeId = ?", vehicletype_id)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		fr := new(FuelResponse)
		var vId uint64
		err := rows.Scan(
			&fr.FuelId,
			&fr.FuelName,
			&fr.Annotation,
			&vId,
		)

		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		vf = append(vf, VehicleTypeFuelMapping{FuelResponse: *fr, VehicleTypeId: vId})

	}

	if err := rows.Err(); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return vf, nil
}

func DoMap(db *sql.DB, vehicletype_id uint64, fuel_ids []uint64) (*FuelMappingResponse, error) {

	tx, err := db.Begin()

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	deletestmt, _ := tx.Prepare("DELETE FROM GEN_VehicleFuel WHERE intVehicleTypeId=?")
	_, err = deletestmt.Exec(vehicletype_id)

	if err != nil {
		tx.Rollback()
		fmt.Println(err.Error())
		return nil, err
	}
	deletestmt.Close()

	qry := "INSERT INTO GEN_VehicleFuel (intVehicleTypeId, intFuelId) VALUES "
	binds := []interface{}{}

	for i, fuel_id := range fuel_ids {
		qry += "(?, ?)"
		binds = append(binds, vehicletype_id, fuel_id)

		if i < len(fuel_ids)-1 {
			qry += ","
		}
	}

	stmt, _ := tx.Prepare(qry)
	_, err = stmt.Exec(binds...)

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

	err = tx.Commit()

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	respobj := new(FuelMappingResponse)
	types, _ := vehicletyperepository.GetById(db, vehicletype_id)
	respobj.VehicleTypeId = types.VehicleTypeId
	respobj.VehicleTypeName = types.VehicleTypeName
	respobj.Annotation = types.Annotation
	if types != nil {

		fuels, _ := GetVehicleFuelMappingById(db, types.VehicleTypeId)

		for _, f := range fuels {
			if f.VehicleTypeId == types.VehicleTypeId {
				respobj.Fuels = append(respobj.Fuels, f.FuelResponse)
			}
		}
	}

	return respobj, nil

}
