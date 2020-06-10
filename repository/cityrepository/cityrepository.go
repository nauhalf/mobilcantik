package cityrepository

import (
	"database/sql"
	"fmt"

	"github.com/nauhalf/mobilcantik/repository/model"
	"github.com/nauhalf/mobilcantik/repository/provincerepository"
)

type CityResponse struct {
	model.Province
	Cities []model.City `json:"cities"`
}

type CitySingleResponse struct {
	model.City
	Province model.Province `json:"province"`
}

func GetAll(db *sql.DB, province_id string) (CityResponse, error) {

	respobj := CityResponse{}

	province, _ := provincerepository.GetById(db, province_id)

	if province != nil {
		cities, _ := GetAllCitiesByProvince(db, province_id)
		respobj.Province = *province
		for _, city := range cities {
			if city.ProvinceId == province.ProvinceId {
				respobj.Cities = append(respobj.Cities, city)
			}
		}
	}

	return respobj, nil
}

func GetById(db *sql.DB, id string) (*CitySingleResponse, error) {

	respobj := CitySingleResponse{}

	city := new(model.City)

	err := db.QueryRow("SELECT szCityId, szProvinceId, szCityName, szAnnotation FROM GEN_City WHERE szCityId = ?", id).
		Scan(&city.CityId,
			&city.ProvinceId,
			&city.CityName,
			&city.Annotation)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	respobj.City = *city

	province := new(model.Province)
	err = db.QueryRow("SELECT szProvinceId, szProvinceName, szAnnotation FROM GEN_Province WHERE szProvinceId = ?", city.ProvinceId).
		Scan(&province.ProvinceId,
			&province.ProvinceName,
			&province.Annotation)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	respobj.Province = *province

	return &respobj, nil
}

func GetAllCitiesByProvince(db *sql.DB, province_id string) ([]model.City, error) {

	var cities []model.City
	rows, err := db.Query("SELECT szCityId, szProvinceId, szCityName, szAnnotation FROM GEN_City WHERE szProvinceId = ?", province_id)

	defer rows.Close()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	for rows.Next() {
		city := new(model.City)
		err := rows.Scan(
			&city.CityId,
			&city.ProvinceId,
			&city.CityName,
			&city.Annotation,
		)

		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		cities = append(cities, *city)

	}

	if err := rows.Err(); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return cities, nil
}
