package adrepository

import (
	"database/sql"
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/nauhalf/mobilcantik/repository/model"
	"github.com/nauhalf/mobilcantik/response"
	"github.com/nauhalf/mobilcantik/utils"
	errorutils "github.com/nauhalf/mobilcantik/utils/error"
)

type AdCreateResponse struct {
	model.Ad
	Vehicle    model.AdVehicle    `json:"vehicle"`
	Facilities []FacilityResponse `json:"facilities"`
	Images     []string           `json:"images"`
	BillId     *uint64            `json:"intBillId"`
}

type AdSingleResponse struct {
	model.Ad
	Vehicle    model.AdVehicle    `json:"vehicle"`
	Facilities []FacilityResponse `json:"facilities"`
	Images     []string           `json:"images"`
}

type FacilityResponse struct {
	FacilityId   uint64  `json:"intFacilityId" form:"intFacilityId"`
	FacilityName string  `json:"szFacilityName" form:"szFacilityName"`
	Annotation   *string `json:"szAnnotation" form:"szAnnotation"`
}

type AdPagingResponse struct {
	AdId            uint64     `json:"intAdId"`
	Email           string     `json:"szEmail"`
	Title           string     `json:"szTitle"`
	Sender          string     `json:"szSender"`
	ProvinceId      string     `json:"szProvinceId"`
	ProvinceName    string     `json:"szProvinceName"`
	CityId          string     `json:"szCityId"`
	CityName        string     `json:"szCityName"`
	PhoneNumber     string     `json:"szPhoneNumber"`
	Annotation      string     `json:"szAnnotation"`
	IsActive        bool       `json:"isActive"`
	AdTypeId        uint64     `json:"intAdTypeId"`
	DtmPosted       *time.Time `json:"dtmPosted"`
	DtmConfirmation *time.Time `json:"dtmConfirmation"`
	DtmExpired      *time.Time `json:"dtmExpired"`
	AdStatusTypeId  uint64     `json:"intAdStatusTypeId"`
	Hits            int        `json:"intHits"`
	IsExpired       bool       `json:"isExpired"`
	VehicleTypeId   uint64     `json:"intVehicleTypeId"`
	VehicleTypeName string     `json:"szVehicleTypeName"`
	VehicleBrandId  uint64     `json:"intVehicleBrandId"`
	BrandName       string     `json:"szBrandName"`
	ConditionId     uint64     `json:"intConditionId"`
	ConditionName   string     `json:"szConditionName"`
	ProductYear     uint64     `json:"intProductYear"`
	Type            string     `json:"szType"`
	Model           string     `json:"szModel"`
	Price           float64    `json:"decPrice"`
	CreatedDate     time.Time  `json:"dtmCreatedDate"`
	ModifiedDate    *time.Time `json:"dtmModifiedDate"`
	Thumbnail       string     `json:"thumbnail"`
}

type GetAllResponse struct {
	Metadata response.Metadata  `json:"metadata"`
	Ads      []AdPagingResponse `json:"ads"`
}

func GetAll(db *sql.DB, page int, arr_province []string, arr_city []string, arr_brand []string, vehicletype *uint64) (interface{}, error) {
	getAllResponse := new(GetAllResponse)
	metadata := new(response.Metadata)

	qryFrom := `FROM
	APP_Ad ad 
	LEFT JOIN APP_AdVehicle adv
	ON ad.intAdId = adv.intAdId
	LEFT JOIN GEN_VehicleBrand vb
	ON adv.intVehicleBrandId = vb.intVehicleBrandId
	LEFT JOIN GEN_VehicleType vt
	ON adv.intVehicleTypeId = vt.intVehicleTypeId
	LEFT JOIN GEN_Condition c
	ON adv.intConditionId = c.intConditionId
	LEFT JOIN GEN_City city
	ON ad.szCityId = city.szCityId
	LEFT JOIN GEN_Province prov
	ON city.szProvinceId = prov.szProvinceId 
	WHERE TRUE 
	AND ad.isActive = 1
	AND ad.isExpired = 0
	AND DATE(ad.dtmExpired) > NOW()`

	qrySelectTotal := `SELECT COUNT(*) AS total `

	qrySelectData := `SELECT ad.intAdId,
						ad.szEmail,
						ad.szTitle,
						ad.szSender,
						city.szProvinceId,
						prov.szProvinceName,
						ad.szCityId,
						city.szCityName,
						ad.szPhoneNumber,
						ad.szAnnotation,
						ad.isActive,
						ad.intAdTypeId,
						ad.dtmPosted,
						ad.dtmConfirmation,
						ad.dtmExpired,
						ad.intAdStatusTypeId,
						ad.intHits,
						ad.isExpired,
						adv.intVehicleTypeId,
						vt.szVehicleTypeName,
						adv.intVehicleBrandId,
						vb.szBrandName,
						adv.intConditionId,
						c.szConditionName,
						adv.intProductYear,
						adv.szType,
						adv.szModel,
  						adv.decPrice,
						ad.dtmCreatedDate,
						ad.dtmModifiedDate `

	binds := []interface{}{}

	whereProvince := ""
	if len(arr_province) > 0 {
		whereProvince = ` AND prov.szProvinceId IN (`
		for i, province_id := range arr_province {
			whereProvince += "?"
			binds = append(binds, province_id)
			if i < len(arr_province)-1 {
				whereProvince += ", "
			} else {
				whereProvince += ")"
			}
		}
	}

	whereCity := ""
	if len(arr_city) > 0 {
		whereCity = ` AND city.szCityId IN (`
		for i, city_id := range arr_city {
			whereCity += "?"
			binds = append(binds, city_id)
			if i < len(arr_city)-1 {
				whereCity += ", "
			} else {
				whereCity += ")"
			}
		}
	}

	whereBrand := ""
	if len(arr_brand) > 0 {
		whereBrand = ` AND vb.intVehicleBrandId IN (`
		for i, brand_id := range arr_brand {
			whereBrand += "?"
			binds = append(binds, brand_id)
			if i < len(arr_brand)-1 {
				whereBrand += ", "
			} else {
				whereBrand += ")"
			}
		}
	}

	whereVehicleType := ""

	if vehicletype != nil {
		whereVehicleType = ` AND vt.intVehicleTypeId=?`
		binds = append(binds, vehicletype)
	}

	orderBy := ` ORDER BY ad.dtmPosted DESC`
	joinQryTotal := qrySelectTotal + qryFrom + whereProvince + whereCity + whereBrand + whereVehicleType + orderBy

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

	joinQrySelect := qrySelectData + qryFrom + whereProvince + whereCity + whereBrand + whereVehicleType + orderBy + qryLimit

	binds = append(binds, limitbinds...)
	rows, err := db.Query(joinQrySelect, binds...)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		ad := new(AdPagingResponse)
		err := rows.Scan(
			&ad.AdId,
			&ad.Email,
			&ad.Title,
			&ad.Sender,
			&ad.ProvinceId,
			&ad.ProvinceName,
			&ad.CityId,
			&ad.CityName,
			&ad.PhoneNumber,
			&ad.Annotation,
			&ad.IsActive,
			&ad.AdTypeId,
			&ad.DtmPosted,
			&ad.DtmConfirmation,
			&ad.DtmExpired,
			&ad.AdStatusTypeId,
			&ad.Hits,
			&ad.IsExpired,
			&ad.VehicleTypeId,
			&ad.VehicleTypeName,
			&ad.VehicleBrandId,
			&ad.BrandName,
			&ad.ConditionId,
			&ad.ConditionName,
			&ad.ProductYear,
			&ad.Type,
			&ad.Model,
			&ad.Price,
			&ad.CreatedDate,
			&ad.ModifiedDate,
		)

		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		getAllResponse.Ads = append(getAllResponse.Ads, *ad)

	}

	if err := rows.Err(); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return getAllResponse, nil
}

func Create(db *sql.DB, ad *model.Ad, password string, advehicle *model.AdVehicle, facilityIds []uint64, filename []string, filelocation string, decAmount *float64) (*AdCreateResponse, error) {

	tx, err := db.Begin()

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	insertAd, err := CreateAd(tx, ad, password)

	if err != nil {
		tx.Rollback()
		fmt.Println(err.Error())
		return nil, err
	}

	insertAdVehicle, err := CreateVehicle(tx, advehicle, *insertAd)

	if err != nil {
		tx.Rollback()
		fmt.Println(err.Error())
		return nil, err
	}

	err = CreateFacility(tx, facilityIds, *insertAd)

	if err != nil {
		tx.Rollback()
		fmt.Println(err.Error())
		return nil, err
	}

	err = CreateImage(tx, filename, filelocation, *insertAd)

	if err != nil {
		tx.Rollback()
		fmt.Println(err.Error())
		return nil, err
	}

	var idBill *uint64
	idBill = nil
	if decAmount != nil {
		idBill, err = CreateBill(tx, *decAmount, *insertAd)

		if err != nil {
			tx.Rollback()
			fmt.Println(err.Error())
			return nil, err
		}
	}

	err = tx.Commit()

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	resp := new(AdCreateResponse)

	getAd, _ := GetAdById(db, *insertAd)
	getV, _ := GetAdVehicleById(db, *insertAdVehicle)
	getF, _ := GetFacilitiesByAdId(db, *insertAd)
	resp.Ad = *getAd
	resp.Vehicle = *getV
	resp.Facilities = getF
	resp.BillId = idBill

	for _, file := range filename {
		path := fmt.Sprintf("%s/%s", filelocation, file)
		resp.Images = append(resp.Images, path)
	}

	return resp, nil
}

func CreateAd(tx *sql.Tx, ad *model.Ad, password string) (*uint64, error) {
	qryAd := `INSERT INTO APP_Ad (
				szEmail,
				szPassword,
				szTitle,
				szSender,
				szCityId,
				szPhoneNumber,
				szAnnotation,
				intAdTypeId,
				dtmPosted,
				dtmConfirmation,
				dtmExpired,
				intAdStatusTypeId,
				intHits)
			VALUES (
				?,
				?,
				?,
				?,
				?,
				?,
				?,
				?,
				?,
				?,
				?,
				?,
				0
			)`
	stmtAd, _ := tx.Prepare(qryAd)
	insertAd, err := stmtAd.Exec(ad.Email,
		password,
		ad.Title,
		ad.Sender,
		ad.CityId,
		ad.PhoneNumber,
		ad.Annotation,
		ad.AdTypeId,
		ad.DtmPosted,
		ad.DtmConfirmation,
		ad.DtmExpired,
		ad.AdStatusTypeId,
	)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	stmtAd.Close()

	idAd, err := insertAd.LastInsertId()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	uintId := uint64(idAd)
	return &uintId, nil
}

func CreateVehicle(tx *sql.Tx, advehicle *model.AdVehicle, idAd uint64) (*uint64, error) {
	qryVehicle := `INSERT INTO APP_AdVehicle (
						intAdId,
						intVehicleTypeId,
						intVehicleBrandId,
						szType,
						szModel,
						intConditionId,
						intProductYear,
						intFuelId,
						szColor,
						intKilometer,
						decPrice,
						szAnnotation
						) VALUES (
							?,
							?,
							?,
							?,
							?,
							?,
							?,
							?,
							?,
							?,
							?,
							?
						)`
	stmtAdVehicle, _ := tx.Prepare(qryVehicle)
	insertAdVehicle, err := stmtAdVehicle.Exec(
		idAd,
		advehicle.VehicleTypeId,
		advehicle.VehicleBrandId,
		advehicle.Type,
		advehicle.Model,
		advehicle.ConditionId,
		advehicle.ProductYear,
		advehicle.FuelId,
		advehicle.Color,
		advehicle.Kilometer,
		advehicle.Price,
		advehicle.Annotation,
	)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	stmtAdVehicle.Close()

	idAdVehicle, err := insertAdVehicle.LastInsertId()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	uintId := uint64(idAdVehicle)
	return &uintId, nil
}

func CreateImage(tx *sql.Tx, filename []string, directory string, idAd uint64) error {
	qryImages := `INSERT INTO APP_AdImage (
						intAdId,
						szDirectory,
						szFilename
					) 
					VALUES (
						?,
						?,
						?
					)`
	stmtImage, _ := tx.Prepare(qryImages)

	for _, file := range filename {
		_, err := stmtImage.Exec(
			idAd,
			directory,
			file,
		)

		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}
	stmtImage.Close()

	return nil
}

func CreateFacility(tx *sql.Tx, facilityIds []uint64, idAd uint64) error {
	qryFacility := `INSERT INTO APP_AdFacility (
						intAdId,
						intFacilityId
					) 
					VALUES (
						?,
						?
					)`
	stmtFacility, _ := tx.Prepare(qryFacility)

	for _, facilityId := range facilityIds {
		_, err := stmtFacility.Exec(
			idAd,
			facilityId,
		)

		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}
	stmtFacility.Close()

	return nil
}

func CreateBill(tx *sql.Tx, decAmount float64, idAd uint64) (*uint64, error) {
	qryBill := `INSERT INTO APP_AdBill (
						intAdId,
						decAmount,
						decBalance
					) 
					VALUES (
						?,
						?,
						?
					)`

	stmtBill, err := tx.Prepare(qryBill)

	insertBill, err := stmtBill.Exec(
		idAd,
		decAmount,
		decAmount,
	)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	id, err := insertBill.LastInsertId()

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	stmtBill.Close()

	uintId := uint64(id)
	return &uintId, nil
}

func Update(db *sql.DB, adclosingreason *model.AdClosingReason) error {

	tx, err := db.Begin()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	stmt, _ := tx.Prepare("UPDATE GEN_AdClosingReason SET szReasonName=?, szAnnotation=?, dtmModifiedDate=NOW() WHERE intAdClosingReasonId=?")
	update, err := stmt.Exec(adclosingreason.ReasonName, adclosingreason.Annotation, adclosingreason.AdClosingReasonId)

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

	stmt, _ := tx.Prepare("DELETE FROM APP_Ad WHERE intAdId=?")
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

func GetAdById(db *sql.DB, id uint64) (*model.Ad, error) {
	ad := new(model.Ad)
	qryGetAd := `SELECT 
				intAdId,
				szEmail,
				szTitle,
				szSender,
				szCityId,
				szPhoneNumber,
				szAnnotation,
				isActive,
				intAdTypeId,
				dtmPosted,
				dtmConfirmation,
				dtmExpired,
				intAdStatusTypeId,
				intHits,
				isExpired,
				dtmCreatedDate,
				dtmModifiedDate 
				FROM
				APP_Ad
				WHERE intAdId = ?`
	err := db.QueryRow(qryGetAd, id).
		Scan(&ad.AdId,
			&ad.Email,
			&ad.Title,
			&ad.Sender,
			&ad.CityId,
			&ad.PhoneNumber,
			&ad.Annotation,
			&ad.IsActive,
			&ad.AdTypeId,
			&ad.DtmPosted,
			&ad.DtmConfirmation,
			&ad.DtmExpired,
			&ad.AdStatusTypeId,
			&ad.Hits,
			&ad.IsExpired,
			&ad.CreatedDate,
			&ad.ModifiedDate,
		)

	if err != nil {
		return nil, err
	}

	return ad, nil
}

func GetAdById2(db *sql.DB, id uint64) (*model.Ad, *string, error) {
	ad := new(model.Ad)
	var pw string
	qryGetAd := `SELECT 
				intAdId,
				szEmail,
				szPassword,
				szTitle,
				szSender,
				szCityId,
				szPhoneNumber,
				szAnnotation,
				isActive,
				intAdTypeId,
				dtmPosted,
				dtmConfirmation,
				dtmExpired,
				intAdStatusTypeId,
				intHits,
				isExpired,
				dtmCreatedDate,
				dtmModifiedDate 
				FROM
				APP_Ad
				WHERE intAdId = ?`
	err := db.QueryRow(qryGetAd, id).
		Scan(&ad.AdId,
			&ad.Email,
			&pw,
			&ad.Title,
			&ad.Sender,
			&ad.CityId,
			&ad.PhoneNumber,
			&ad.Annotation,
			&ad.IsActive,
			&ad.AdTypeId,
			&ad.DtmPosted,
			&ad.DtmConfirmation,
			&ad.DtmExpired,
			&ad.AdStatusTypeId,
			&ad.Hits,
			&ad.IsExpired,
			&ad.CreatedDate,
			&ad.ModifiedDate,
		)

	if err != nil {
		return nil, nil, err
	}

	return ad, &pw, nil
}

func GetAdVehicleById(db *sql.DB, id uint64) (*model.AdVehicle, error) {
	adVehicle := new(model.AdVehicle)
	qryGetAdVehicle := `SELECT 
			intAdVehicleId,
			intAdId,
			intVehicleTypeId,
			intVehicleBrandId,
			szType,
			szModel,
			intConditionId,
			intProductYear,
			intFuelId,
			szColor,
			intKilometer,
			decPrice,
			szAnnotation,
			dtmCreatedDate,
			dtmModifiedDate 
			FROM
			APP_AdVehicle 
			WHERE intAdVehicleId = ?`
	err := db.QueryRow(qryGetAdVehicle, id).
		Scan(
			&adVehicle.AdVehicleId,
			&adVehicle.AdId,
			&adVehicle.VehicleTypeId,
			&adVehicle.VehicleBrandId,
			&adVehicle.Type,
			&adVehicle.Model,
			&adVehicle.ConditionId,
			&adVehicle.ProductYear,
			&adVehicle.FuelId,
			&adVehicle.Color,
			&adVehicle.Kilometer,
			&adVehicle.Price,
			&adVehicle.Annotation,
			&adVehicle.CreatedDate,
			&adVehicle.ModifiedDate,
		)

	if err != nil {
		return nil, err
	}

	return adVehicle, nil
}

func GetAdVehicleByAdId(db *sql.DB, adId uint64) (*model.AdVehicle, error) {
	adVehicle := new(model.AdVehicle)
	qryGetAdVehicle := `SELECT 
			intAdVehicleId,
			intAdId,
			intVehicleTypeId,
			intVehicleBrandId,
			szType,
			szModel,
			intConditionId,
			intProductYear,
			intFuelId,
			szColor,
			intKilometer,
			decPrice,
			szAnnotation,
			dtmCreatedDate,
			dtmModifiedDate 
			FROM
			APP_AdVehicle 
			WHERE intAdId = ?`
	err := db.QueryRow(qryGetAdVehicle, adId).
		Scan(
			&adVehicle.AdVehicleId,
			&adVehicle.AdId,
			&adVehicle.VehicleTypeId,
			&adVehicle.VehicleBrandId,
			&adVehicle.Type,
			&adVehicle.Model,
			&adVehicle.ConditionId,
			&adVehicle.ProductYear,
			&adVehicle.FuelId,
			&adVehicle.Color,
			&adVehicle.Kilometer,
			&adVehicle.Price,
			&adVehicle.Annotation,
			&adVehicle.CreatedDate,
			&adVehicle.ModifiedDate,
		)

	if err != nil {
		return nil, err
	}

	return adVehicle, nil
}

func GetFacilitiesByAdId(db *sql.DB, adId uint64) ([]FacilityResponse, error) {
	var fs []FacilityResponse
	qryFacility := `SELECT 
			f.intFacilityId,
			f.szFacilityName,
			f.szAnnotation
			FROM
			APP_AdFacility af
			LEFT JOIN GEN_Facility f
			ON af.intFacilityId = f.intFacilityId
			WHERE af.intAdId = ?`
	rows, err := db.Query(qryFacility, adId)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		f := new(FacilityResponse)
		err := rows.Scan(
			&f.FacilityId,
			&f.FacilityName,
			&f.Annotation,
		)

		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		fs = append(fs, *f)

	}

	if err := rows.Err(); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return fs, nil
}

func Verification(db *sql.DB, adId uint64) error {
	qry := `UPDATE APP_Ad 
		SET 
			isActive = 1,
			intAdStatusTypeId = 3,
			dtmPosted=?, 
			dtmConfirmation=?, 
			dtmExpired=?,
			dtmModifiedDate=NOW()
		WHERE intAdId=?`

	tx, err := db.Begin()

	stmt, _ := tx.Prepare(qry)
	now := time.Now()
	exp := now.AddDate(0, 0, utils.GetExpiredInterval())
	update, err := stmt.Exec(now, now, exp, adId)

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

func Increment(db *sql.DB, intAdId uint64) error {
	qry := `UPDATE APP_Ad 
		SET 
			dtmModifiedDate=NOW(),
			intHits = intHits + 1
		WHERE intAdId=?`

	tx, err := db.Begin()

	stmt, _ := tx.Prepare(qry)
	update, err := stmt.Exec(intAdId)

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

func IsAdExists(db *sql.DB, intAdId uint64) (int, error) {
	var count int
	count = 0
	qryGetAd := `SELECT 
				COUNT(*)
				FROM
				APP_Ad
				WHERE intAdId = ?`
	err := db.QueryRow(qryGetAd, intAdId).
		Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil
}

func GetImagesByAdId(db *sql.DB, adId uint64) ([]string, error) {
	var paths []string

	qry := `SELECT 
				szDirectory, szFilename
				FROM
				APP_AdImage
				WHERE intAdId = ?`
	rows, err := db.Query(qry, adId)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var dir, file string
		err := rows.Scan(
			&dir,
			&file,
		)

		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		path := dir + "/" + file
		paths = append(paths, path)

	}
	if err != nil {
		return []string{}, err
	}

	return paths, nil
}
