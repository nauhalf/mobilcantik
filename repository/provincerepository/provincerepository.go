package provincerepository

import (
	"database/sql"
	"fmt"

	"github.com/nauhalf/mobilcantik/repository/model"
)

func GetAll(db *sql.DB) ([]model.Province, error) {

	var provinces []model.Province
	rows, err := db.Query("SELECT szProvinceId, szProvinceName, szAnnotation FROM GEN_Province")

	defer rows.Close()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	for rows.Next() {
		province := new(model.Province)
		err := rows.Scan(
			&province.ProvinceId,
			&province.ProvinceName,
			&province.Annotation,
		)

		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		provinces = append(provinces, *province)

	}

	if err := rows.Err(); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return provinces, nil
}

func GetById(db *sql.DB, id string) (*model.Province, error) {
	province := new(model.Province)
	err := db.QueryRow("SELECT szProvinceId, szProvinceName, szAnnotation FROM GEN_Province WHERE szProvinceId = ?", id).
		Scan(&province.ProvinceId,
			&province.ProvinceName,
			&province.Annotation)

	if err != nil {
		return nil, err
	}

	return province, nil
}
