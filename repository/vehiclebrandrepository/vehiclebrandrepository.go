package vehiclebrandrepository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/nauhalf/mobilcantik/repository/model"
	"github.com/nauhalf/mobilcantik/repository/vehicletyperepository"
	errorutils "github.com/nauhalf/mobilcantik/utils/error"
)

type VehicleBrandResponse struct {
	model.VehicleType
	VehicleBrands []model.VehicleBrand `json:"vehiclebrands"`
}

type VehicleBrandCreateResponse struct {
	model.VehicleBrand
	VehicleType model.VehicleType `json:"vehicletype"`
}

func GetAll(db *sql.DB, vehicletype_id *uint64) ([]VehicleBrandResponse, error) {

	respobj := []VehicleBrandResponse{}

	if vehicletype_id == nil {
		types, _ := vehicletyperepository.GetAll(db)
		brands, _ := GetAllBrands(db)
		for _, vt := range types {
			tmp := VehicleBrandResponse{VehicleBrands: []model.VehicleBrand{}}
			tmp.VehicleType = vt
			for _, vb := range brands {
				if vb.VehicleTypeId == vt.VehicleTypeId {
					tmp.VehicleBrands = append(tmp.VehicleBrands, vb)
				}
			}
			respobj = append(respobj, tmp)
		}
	} else {
		types, _ := vehicletyperepository.GetById(db, *vehicletype_id)

		if types != nil {
			tmp := VehicleBrandResponse{VehicleBrands: []model.VehicleBrand{}}
			brands, _ := GetAllBrandsByVehicleType(db, *vehicletype_id)
			tmp.VehicleType = *types
			for _, vb := range brands {
				if vb.VehicleTypeId == types.VehicleTypeId {
					tmp.VehicleBrands = append(tmp.VehicleBrands, vb)
				}
			}
			respobj = append(respobj, tmp)
		}
	}

	return respobj, nil
}

func GetAllBrands(db *sql.DB) ([]model.VehicleBrand, error) {

	var brands []model.VehicleBrand
	rows, err := db.Query("SELECT * FROM GEN_VehicleBrand")

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		vBrand := new(model.VehicleBrand)
		err := rows.Scan(
			&vBrand.VehicleBrandId,
			&vBrand.VehicleTypeId,
			&vBrand.VehicleBrandName,
			&vBrand.Count,
			&vBrand.Annotation,
			&vBrand.CreatedDate,
			&vBrand.ModifiedDate,
		)

		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		brands = append(brands, *vBrand)

	}

	if err := rows.Err(); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return brands, nil
}

func GetAllBrandsByVehicleType(db *sql.DB, vehicletype_id uint64) ([]model.VehicleBrand, error) {

	var brands []model.VehicleBrand
	rows, err := db.Query("SELECT * FROM GEN_VehicleBrand WHERE intVehicleTypeId = ?", vehicletype_id)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		vBrand := new(model.VehicleBrand)
		err := rows.Scan(
			&vBrand.VehicleBrandId,
			&vBrand.VehicleTypeId,
			&vBrand.VehicleBrandName,
			&vBrand.Count,
			&vBrand.Annotation,
			&vBrand.CreatedDate,
			&vBrand.ModifiedDate,
		)

		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		brands = append(brands, *vBrand)

		// var id uint64
		// var name, annotation string
		// var c_date time.Time
		// var m_date *time.Time
		// err := rows.Scan(
		// 	id,
		// 	name,
		// 	annotation,
		// 	c_date,
		// 	m_date,
		// )

		// if err != nil {
		// 	fmt.Println(err.Error())
		// 	return nil, err
		// }

		// brands = append(brands, model.VehicleType{VehicleTypeId: id, VehicleTypeName: name, Annotation: annotation, CreatedDate: c_date, ModifiedDate: m_date})

	}

	if err := rows.Err(); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return brands, nil
}

func Create(db *sql.DB, vehicleBrand *model.VehicleBrand) (*VehicleBrandCreateResponse, error) {

	tx, err := db.Begin()

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	stmt, _ := tx.Prepare("INSERT INTO GEN_VehicleBrand (intVehicleTypeId, szBrandName, szAnnotation) VALUES (?, ?, ?)")
	insert, err := stmt.Exec(vehicleBrand.VehicleTypeId, vehicleBrand.VehicleBrandName, vehicleBrand.Annotation)

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
	respobj := VehicleBrandCreateResponse{}
	types := new(model.VehicleType)
	err = tx.QueryRow("SELECT * FROM GEN_VehicleType WHERE intVehicleTypeId = ?", vehicleBrand.VehicleTypeId).
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
	newBrand := new(model.VehicleBrand)

	err = tx.QueryRow("SELECT * FROM GEN_VehicleBrand WHERE intVehicleBrandId = ?", id).
		Scan(&newBrand.VehicleBrandId,
			&newBrand.VehicleTypeId,
			&newBrand.VehicleBrandName,
			&newBrand.Count,
			&newBrand.Annotation,
			&newBrand.CreatedDate,
			&newBrand.ModifiedDate)

	if err != nil {
		tx.Rollback()
		fmt.Println(err.Error())
		return nil, err
	}

	respobj.VehicleBrand = *newBrand
	err = tx.Commit()

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &respobj, nil

}

func Update(db *sql.DB, vehicleType *model.VehicleBrand) error {

	tx, err := db.Begin()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	stmt, _ := tx.Prepare("UPDATE GEN_VehicleBrand SET intVehicleTypeId=?, szBrandName=?, szAnnotation=?, dtmModifiedDate=NOW() WHERE intVehicleBrandId=?")
	update, err := stmt.Exec(vehicleType.VehicleTypeId, vehicleType.VehicleBrandName, vehicleType.Annotation, vehicleType.VehicleBrandId)

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

	stmt, _ := tx.Prepare("DELETE FROM GEN_VehicleBrand WHERE intVehicleBrandId=?")
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

func GetById(db *sql.DB, id uint64) (*VehicleBrandCreateResponse, error) {

	respobj := VehicleBrandCreateResponse{}

	brand := new(model.VehicleBrand)

	err := db.QueryRow("SELECT * FROM GEN_VehicleBrand WHERE intVehicleBrandId = ?", id).
		Scan(&brand.VehicleBrandId,
			&brand.VehicleTypeId,
			&brand.VehicleBrandName,
			&brand.Count,
			&brand.Annotation,
			&brand.CreatedDate,
			&brand.ModifiedDate)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	respobj.VehicleBrand = *brand

	vtypes := new(model.VehicleType)
	err = db.QueryRow("SELECT * FROM GEN_VehicleType WHERE intVehicleTypeId = ?", brand.VehicleTypeId).
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
