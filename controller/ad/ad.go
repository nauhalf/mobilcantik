package ad

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/nauhalf/mobilcantik/db"
	"github.com/nauhalf/mobilcantik/repository/adrepository"
	"github.com/nauhalf/mobilcantik/repository/adtyperepository"
	"github.com/nauhalf/mobilcantik/repository/cityrepository"
	"github.com/nauhalf/mobilcantik/repository/conditionrepository"
	"github.com/nauhalf/mobilcantik/repository/facilityrepository"
	"github.com/nauhalf/mobilcantik/repository/fuelrepository"
	"github.com/nauhalf/mobilcantik/repository/model"
	"github.com/nauhalf/mobilcantik/repository/vehiclebrandrepository"
	"github.com/nauhalf/mobilcantik/repository/vehicletyperepository"
	"github.com/nauhalf/mobilcantik/response"
	"github.com/nauhalf/mobilcantik/utils"
	errorutils "github.com/nauhalf/mobilcantik/utils/error"
)

type RequestCreate struct {
	Email          string   `json:"szEmail" form:"szEmail" query:"szEmail" validate:"required"`
	Title          string   `json:"szTitle" form:"szTitle" query:"szTitle" validate:"required"`
	Sender         string   `json:"szSender" form:"szSender" query:"szSender" validate:"required"`
	CityId         string   `json:"szCityId" form:"szCityId" query:"szCityId" validate:"required"`
	PhoneNumber    string   `json:"szPhoneNumber" form:"szPhoneNumber" query:"szPhoneNumber" validate:"required"`
	AdTypeId       uint64   `json:"intAdTypeId" form:"intAdTypeId" query:"intAdTypeId" validate:"required"`
	VehicleTypeId  uint64   `json:"intVehicleTypeId" form:"intVehicleTypeId" query:"intVehicleTypeId" validate:"required"`
	VehicleBrandId uint64   `json:"intVehicleBrandId" form:"intVehicleBrandId" query:"intVehicleBrandId" validate:"required"`
	Type           string   `json:"szType" form:"szType" query:"szType" validate:"required"`
	Model          string   `json:"szModel" form:"szModel" query:"szModel" validate:"required"`
	ConditionId    uint64   `json:"intConditionId" form:"intConditionId" query:"intConditionId" validate:"required"`
	ProductYear    int      `json:"intProductYear" form:"intProductYear" query:"intProductYear" validate:"required"`
	FuelId         uint64   `json:"intFuelId" form:"intFuelId" query:"intFuelId" validate:"required"`
	Color          string   `json:"szColor" form:"szColor" query:"szColor" validate:"required"`
	Kilometer      int      `json:"intKilometer" form:"intKilometer" query:"intKilometer" validate:"required"`
	Price          float64  `json:"decPrice" form:"decPrice" query:"decPrice" validate:"required"`
	Annotation     string   `json:"szAnnotation" form:"szAnnotation" query:"szAnnotation" validate:"required"`
	FacilityIds    []uint64 `json:"intFacilityId" form:"intFacilityId" query:"intFacilityId"`
}

type RequestUpdate struct {
	AdClosingReasonId uint64 `json:"intAdClosingReasonId" form:"intAdClosingReasonId" query:"intAdClosingReasonId" validate:"required"`
	RequestCreate
}

type RequestVerification struct {
	AdId     uint64 `query:"id" validate:"required"`
	Password string `query:"password" validate:"required"`
}

func GetAll(c echo.Context) error {

	page := c.QueryParam("page")
	provinces := c.QueryParam("provinces")

	cities := c.QueryParam("cities")
	brands := c.QueryParam("brands")
	vehicletype := c.QueryParam("vehicletype")

	var p int
	tmpP, err := strconv.Atoi(page)
	if tmpP < 1 || err != nil {
		p = 1
	} else {
		p = tmpP
	}

	var vtx *uint64
	vt, err := strconv.ParseUint(vehicletype, 10, 64)
	if vehicletype == "" || err != nil {
		vtx = nil
	} else {
		vtx = &vt
	}

	var arr_province, arr_city, arr_brand = []string{}, []string{}, []string{}

	if provinces != "" {
		arr_province = strings.Split(provinces, ",")
	}

	if cities != "" {
		arr_city = strings.Split(cities, ",")
	}

	if brands != "" {
		arr_brand = strings.Split(brands, ",")
	}

	all, err := adrepository.GetAll(db.DBCon, p, arr_province, arr_city, arr_brand, vtx)

	if err != nil {
		resp := response.ResponseError{
			Code:      http.StatusInternalServerError,
			Message:   http.StatusText(http.StatusInternalServerError),
			ErrorCode: nil,
		}
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp := response.ResponseSuccess{
		Code:    http.StatusOK,
		Message: fmt.Sprintf(errorutils.StatusErrorSuccessfullyRetrieved, "List ads"),
		Data:    all,
	}
	return c.JSON(http.StatusOK, resp)
}

func Create(c echo.Context) error {

	r := new(RequestCreate)

	if err := c.Bind(r); err != nil {
		resp := response.ResponseError{
			Code:      http.StatusInternalServerError,
			Message:   http.StatusText(http.StatusInternalServerError),
			ErrorCode: nil,
		}
		return c.JSON(resp.Code, resp)
	}

	if err := c.Validate(r); err != nil {
		resp := response.ResponseError{
			Code:      http.StatusUnprocessableEntity,
			Message:   errorutils.StatusErrorFillRequiredForm,
			ErrorCode: 1,
		}
		return c.JSON(resp.Code, resp)
	}

	//validate exists data

	if validate1 := ValidateCreateExists(*r); validate1 != nil {
		return c.JSON(validate1.Code, validate1)
	}

	form, _ := c.MultipartForm()

	images := form.File["images"]
	validate2 := ValidateFile(images)
	if validate2 != nil {
		return c.JSON(validate2.Code, validate2)
	}

	if validate3 := ValidateEmail(r.Email); validate3 != nil {
		return c.JSON(validate3.Code, validate3)
	}

	now := time.Now()
	year := now.Year()
	month := now.Month()
	day := now.Day()

	newfile, err := HandleUpload(images)
	dir := fmt.Sprintf("%s/%s/%s/", strconv.FormatInt(int64(year), 10), strconv.FormatInt(int64(month), 10), strconv.FormatInt(int64(day), 10))
	if err != nil {
		resp := response.ResponseError{
			Code:      http.StatusInternalServerError,
			Message:   err.Error(),
			ErrorCode: nil,
		}
		return c.JSON(resp.Code, resp)
	}

	ad := new(model.Ad)
	advehicle := new(model.AdVehicle)
	ad.Email = r.Email
	ad.Title = r.Title
	ad.Sender = r.Sender
	ad.CityId = r.CityId
	ad.PhoneNumber = r.PhoneNumber
	ad.AdTypeId = r.AdTypeId
	ad.Annotation = r.Annotation

	password, _ := utils.RandomFile16Char()
	var decAmount *float64
	if utils.GetMode() == 1 {

		decAmount = nil
		ad.AdStatusTypeId = 2
	} else {
		ad.AdStatusTypeId = 1
		fee := utils.GetAmountFee()
		decAmount = &fee
	}

	advehicle.VehicleTypeId = r.VehicleTypeId
	advehicle.VehicleBrandId = r.VehicleBrandId
	advehicle.Type = r.Type
	advehicle.Model = r.Model
	advehicle.ConditionId = r.ConditionId
	advehicle.ProductYear = r.ProductYear
	advehicle.FuelId = r.FuelId
	advehicle.Color = r.Color
	advehicle.Kilometer = r.Kilometer
	advehicle.Price = r.Price
	advehicle.Annotation = &r.Annotation

	newAd, err := adrepository.Create(db.DBCon, ad, password, advehicle, r.FacilityIds, newfile, dir, decAmount)

	if err != nil {
		var deleteFile []string
		for _, file := range newfile {
			deleteFile = append(deleteFile, dir+file)
		}

		_ = HandleDelete(deleteFile)
		resp := response.ResponseError{
			Code:      http.StatusInternalServerError,
			Message:   http.StatusText(http.StatusInternalServerError),
			ErrorCode: nil,
		}
		return c.JSON(resp.Code, resp)
	}

	resp := response.ResponseSuccess{
		Code:    http.StatusCreated,
		Message: fmt.Sprintf(errorutils.StatusErrorSuccessfullyCreated, "Ad"),
		Data:    newAd,
	}

	return c.JSON(resp.Code, resp)
}

func ValidateCreateExists(r RequestCreate) *response.ResponseError {
	cityExists, _ := cityrepository.GetById(db.DBCon, r.CityId)
	if cityExists == nil {
		resp := response.ResponseError{
			Code:      http.StatusNotFound,
			Message:   fmt.Sprintf(errorutils.StatusErrorResourceNotExists, "City"),
			ErrorCode: 1,
		}

		return &resp
	}

	adTypeExists, _ := adtyperepository.GetById(db.DBCon, r.AdTypeId)
	if adTypeExists == nil {
		resp := response.ResponseError{
			Code:      http.StatusNotFound,
			Message:   fmt.Sprintf(errorutils.StatusErrorResourceNotExists, "Ad Type"),
			ErrorCode: 2,
		}

		return &resp
	}

	vehicleTypeExists, _ := vehicletyperepository.GetById(db.DBCon, r.VehicleTypeId)
	if vehicleTypeExists == nil {
		resp := response.ResponseError{
			Code:      http.StatusNotFound,
			Message:   fmt.Sprintf(errorutils.StatusErrorResourceNotExists, "Vehicle Type"),
			ErrorCode: 3,
		}

		return &resp
	}

	vehicleBrandExists, _ := vehiclebrandrepository.GetById(db.DBCon, r.VehicleBrandId)
	if vehicleBrandExists == nil {
		resp := response.ResponseError{
			Code:      http.StatusNotFound,
			Message:   fmt.Sprintf(errorutils.StatusErrorResourceNotExists, "Brand"),
			ErrorCode: 4,
		}

		return &resp
	}

	conditionExists, _ := conditionrepository.GetById(db.DBCon, r.ConditionId)
	if conditionExists == nil {
		resp := response.ResponseError{
			Code:      http.StatusNotFound,
			Message:   fmt.Sprintf(errorutils.StatusErrorResourceNotExists, "Condition"),
			ErrorCode: 5,
		}

		return &resp
	}

	fuelExists, _ := fuelrepository.GetById(db.DBCon, r.FuelId)
	if fuelExists == nil {
		resp := response.ResponseError{
			Code:      http.StatusNotFound,
			Message:   fmt.Sprintf(errorutils.StatusErrorResourceNotExists, "Fuel"),
			ErrorCode: 6,
		}

		return &resp
	}

	if len(r.FacilityIds) < 1 {
		resp := response.ResponseError{
			Code:      http.StatusUnprocessableEntity,
			Message:   errorutils.StatusErrorFillRequiredForm,
			ErrorCode: 1,
		}

		return &resp
	}

	facilityExists, _ := facilityrepository.IsExists(db.DBCon, r.VehicleTypeId, r.FacilityIds)
	if !facilityExists {
		resp := response.ResponseError{
			Code:      http.StatusNotFound,
			Message:   fmt.Sprintf(errorutils.StatusErrorResourceNotExists, "Facility"),
			ErrorCode: 7,
		}

		return &resp
	}

	return nil
}

func ValidateFile(file []*multipart.FileHeader) *response.ResponseError {

	if len(file) < 1 {
		resp := response.ResponseError{
			Code:      http.StatusUnprocessableEntity,
			Message:   errorutils.StatusErrorFillRequiredForm,
			ErrorCode: 1,
		}
		return &resp

	}

	validateType := utils.IsFiletypeAccepted(file)
	if !validateType {
		resp := response.ResponseError{
			Code:      http.StatusUnprocessableEntity,
			Message:   errorutils.StatusErrorImageType,
			ErrorCode: 4,
		}

		return &resp
	}

	for _, f := range file {
		validatesize := utils.IsFileSizeAccepted(f.Size)
		if !validatesize {
			resp := response.ResponseError{
				Code:      http.StatusUnprocessableEntity,
				Message:   fmt.Sprintf(errorutils.StatusErrorImageSize, "1 MB"),
				ErrorCode: 5,
			}

			return &resp
		}
	}
	return nil
}

func ValidateYear(c echo.Context, year int, now int) *response.ResponseError {
	validyear := (year > 1800 && year <= now)
	if !validyear {
		resp := response.ResponseError{
			Code:      http.StatusUnprocessableEntity,
			Message:   errorutils.StatusErrorInvalidYear,
			ErrorCode: 3,
		}

		return &resp
	}
	return nil
}

func ValidateEmail(email string) *response.ResponseError {
	if !utils.ValidateEmail(email) {
		resp := response.ResponseError{
			Code:      http.StatusUnprocessableEntity,
			Message:   errorutils.StatusErrorInvalidEmail,
			ErrorCode: 2,
		}
		return &resp
	}
	return nil
}

func Delete(c echo.Context) error {

	adId := c.Param("id")
	password := c.Param("password")

	if adId == "" || password == "" {
		resp := response.ResponseError{
			Code:      http.StatusUnprocessableEntity,
			Message:   errorutils.StatusErrorFillRequiredForm,
			ErrorCode: 1,
		}
		return c.JSON(resp.Code, resp)
	}

	intAdId, err := strconv.ParseUint(adId, 10, 64)

	if err != nil {
		resp := response.ResponseError{
			Code:      http.StatusInternalServerError,
			Message:   http.StatusText(http.StatusInternalServerError),
			ErrorCode: nil,
		}
		return c.JSON(resp.Code, resp)
	}

	ad, pw, _ := adrepository.GetAdById2(db.DBCon, intAdId)

	if ad == nil {
		resp := response.ResponseError{
			Code:      http.StatusNotFound,
			Message:   fmt.Sprintf(errorutils.StatusErrorResourceNotExists, "Ad"),
			ErrorCode: 1,
		}
		return c.JSON(resp.Code, resp)
	}

	if password != *pw {
		resp := response.ResponseError{
			Code:      http.StatusUnprocessableEntity,
			Message:   http.StatusText(http.StatusInternalServerError),
			ErrorCode: 1,
		}
		return c.JSON(resp.Code, resp)
	}

	images, _ := adrepository.GetImagesByAdId(db.DBCon, intAdId)

	err = adrepository.Delete(db.DBCon, intAdId)

	if err != nil {
		if err.Error() == errorutils.StatusZeroAffectedRows {
			resp := response.ResponseError{
				Code:      http.StatusNotFound,
				Message:   fmt.Sprintf(errorutils.StatusErrorResourceNotExists, "Ad"),
				ErrorCode: 1,
			}
			return c.JSON(resp.Code, resp)
		} else if err.(*mysql.MySQLError).Number == errorutils.ErrorMySQLDeleteConstraintFK {
			resp := response.ResponseError{
				Code:      http.StatusConflict,
				Message:   fmt.Sprintf(errorutils.StatusErrorAlreadyInUsed, "Ad"),
				ErrorCode: nil,
			}
			return c.JSON(resp.Code, resp)
		}

		resp := response.ResponseError{
			Code:      http.StatusInternalServerError,
			Message:   http.StatusText(http.StatusInternalServerError),
			ErrorCode: nil,
		}
		return c.JSON(resp.Code, resp)
	}

	_ = HandleDelete(images)
	resp := response.ResponseSuccess{
		Code:    http.StatusOK,
		Message: fmt.Sprintf(errorutils.StatusErrorSuccessfullyDeleted, "Ad"),
		Data:    nil,
	}

	return c.JSON(resp.Code, resp)
}

func ForceDelete(c echo.Context) error {

	adId := c.Param("id")

	if adId == "" {
		resp := response.ResponseError{
			Code:      http.StatusUnprocessableEntity,
			Message:   http.StatusText(http.StatusUnprocessableEntity),
			ErrorCode: 1,
		}
		return c.JSON(resp.Code, resp)
	}

	intAdId, err := strconv.ParseUint(adId, 10, 64)

	if err != nil {
		resp := response.ResponseError{
			Code:      http.StatusInternalServerError,
			Message:   http.StatusText(http.StatusInternalServerError),
			ErrorCode: nil,
		}
		return c.JSON(resp.Code, resp)
	}

	ad, _, _ := adrepository.GetAdById2(db.DBCon, intAdId)

	if ad == nil {
		resp := response.ResponseError{
			Code:      http.StatusNotFound,
			Message:   fmt.Sprintf(errorutils.StatusErrorResourceNotExists, "Ad"),
			ErrorCode: 1,
		}
		return c.JSON(resp.Code, resp)
	}

	images, _ := adrepository.GetImagesByAdId(db.DBCon, intAdId)

	err = adrepository.Delete(db.DBCon, intAdId)

	if err != nil {
		if err.Error() == errorutils.StatusZeroAffectedRows {
			resp := response.ResponseError{
				Code:      http.StatusNotFound,
				Message:   fmt.Sprintf(errorutils.StatusErrorResourceNotExists, "Ad"),
				ErrorCode: 1,
			}
			return c.JSON(resp.Code, resp)
		} else if err.(*mysql.MySQLError).Number == errorutils.ErrorMySQLDeleteConstraintFK {
			resp := response.ResponseError{
				Code:      http.StatusConflict,
				Message:   fmt.Sprintf(errorutils.StatusErrorSuccessfullyDeleted, "Ad"),
				ErrorCode: nil,
			}
			return c.JSON(resp.Code, resp)
		}

		resp := response.ResponseError{
			Code:      http.StatusInternalServerError,
			Message:   http.StatusText(http.StatusInternalServerError),
			ErrorCode: nil,
		}
		return c.JSON(resp.Code, resp)
	}

	_ = HandleDelete(images)
	resp := response.ResponseSuccess{
		Code:    http.StatusOK,
		Message: fmt.Sprintf(errorutils.StatusErrorSuccessfullyDeleted, "Ad"),
		Data:    nil,
	}

	return c.JSON(resp.Code, resp)
}

func GetById(c echo.Context) error {

	cID := c.Param("id")

	if cID == "" {
		resp := response.ResponseError{
			Code:      http.StatusUnprocessableEntity,
			Message:   errorutils.StatusErrorFillRequiredForm,
			ErrorCode: 1,
		}
		return c.JSON(resp.Code, resp)
	}

	intID, err := strconv.ParseUint(cID, 10, 64)

	if err != nil {
		resp := response.ResponseError{
			Code:      http.StatusInternalServerError,
			Message:   http.StatusText(http.StatusInternalServerError),
			ErrorCode: nil,
		}
		return c.JSON(resp.Code, resp)
	}

	adx, _ := adrepository.GetAdById(db.DBCon, intID)

	if adx == nil {
		resp := response.ResponseError{
			Code:      http.StatusNotFound,
			Message:   http.StatusText(http.StatusNotFound),
			ErrorCode: 1,
		}
		return c.JSON(resp.Code, resp)
	}

	if !adx.IsActive {
		resp := response.ResponseError{
			Code:      http.StatusUnprocessableEntity,
			Message:   "Ad still not active",
			ErrorCode: 1,
		}
		return c.JSON(resp.Code, resp)
	}

	err = adrepository.Increment(db.DBCon, intID)

	if err != nil {
		resp := response.ResponseError{
			Code:      http.StatusInternalServerError,
			Message:   http.StatusText(http.StatusInternalServerError),
			ErrorCode: nil,
		}
		return c.JSON(http.StatusInternalServerError, resp)
	}

	ad, err := adrepository.GetAdById(db.DBCon, intID)
	if err != nil {
		resp := response.ResponseError{
			Code:      http.StatusInternalServerError,
			Message:   http.StatusText(http.StatusInternalServerError),
			ErrorCode: nil,
		}
		return c.JSON(http.StatusInternalServerError, resp)
	}

	adsingleresponse := new(adrepository.AdSingleResponse)
	getV, _ := adrepository.GetAdVehicleByAdId(db.DBCon, ad.AdId)
	getF, _ := adrepository.GetFacilitiesByAdId(db.DBCon, ad.AdId)
	getI, _ := adrepository.GetImagesByAdId(db.DBCon, ad.AdId)
	adsingleresponse.Vehicle = *getV
	adsingleresponse.Facilities = getF
	adsingleresponse.Ad = *ad
	adsingleresponse.Images = getI

	resp := response.ResponseSuccess{
		Code:    http.StatusOK,
		Message: fmt.Sprint(errorutils.StatusErrorSuccessfullyRetrieved, "Ad"),
		Data:    adsingleresponse,
	}
	return c.JSON(http.StatusOK, resp)
}

func HandleUpload(files []*multipart.FileHeader) ([]string, error) {

	pathdir, err := utils.GenerateFolder()

	if err != nil {
		return nil, err
	}
	var newFileName []string

	for _, file := range files {

		// Source
		src, err := file.Open()
		if err != nil {
			return nil, err
		}
		defer src.Close()

		// Destination
		fileName := utils.GenerateNewFile(file.Filename)
		newFile := pathdir + fileName
		dst, err := os.Create(newFile)
		if err != nil {
			return nil, err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return nil, err
		}

		newFileName = append(newFileName, fileName)
	}

	return newFileName, nil
}

func HandleDelete(files []string) error {

	pathimg := utils.GetImageDir()
	for _, file := range files {

		var err = os.Remove(pathimg + file)
		if err != nil {
			return err
		}
	}

	return nil
}

func Verification(c echo.Context) error {
	r := new(RequestVerification)

	if err := c.Bind(r); err != nil {
		resp := response.ResponseError{
			Code:      http.StatusInternalServerError,
			Message:   http.StatusText(http.StatusInternalServerError),
			ErrorCode: nil,
		}
		return c.JSON(resp.Code, resp)
	}

	if err := c.Validate(r); err != nil {
		resp := response.ResponseError{
			Code:      http.StatusUnprocessableEntity,
			Message:   errorutils.StatusErrorFillRequiredForm,
			ErrorCode: 1,
		}
		return c.JSON(resp.Code, resp)
	}

	ad, pw, _ := adrepository.GetAdById2(db.DBCon, r.AdId)

	if ad == nil {
		resp := response.ResponseError{
			Code:      http.StatusNotFound,
			Message:   fmt.Sprintf(errorutils.StatusErrorResourceNotExists, "Ad"),
			ErrorCode: 1,
		}
		return c.JSON(resp.Code, resp)
	}

	if r.Password != *pw {
		resp := response.ResponseError{
			Code:      http.StatusUnprocessableEntity,
			Message:   errorutils.StatusErrorIncorrectPassword,
			ErrorCode: 2,
		}
		return c.JSON(resp.Code, resp)
	}

	if ad.AdStatusTypeId == 1 {
		resp := response.ResponseError{
			Code:      http.StatusUnprocessableEntity,
			Message:   errorutils.StatusErrorPaidToActivateAd,
			ErrorCode: 3,
		}
		return c.JSON(resp.Code, resp)
	}

	if ad.IsActive {
		resp := response.ResponseError{
			Code:      http.StatusUnprocessableEntity,
			Message:   fmt.Sprintf(errorutils.StatusErrorAlreadyActive, "Ad"),
			ErrorCode: 4,
		}
		return c.JSON(resp.Code, resp)
	}

	err := adrepository.Verification(db.DBCon, r.AdId)

	if err != nil {
		resp := response.ResponseError{
			Code:      http.StatusInternalServerError,
			Message:   http.StatusText(http.StatusInternalServerError),
			ErrorCode: nil,
		}
		return c.JSON(resp.Code, resp)
	}

	resp := response.ResponseSuccess{
		Code:    http.StatusOK,
		Message: "Ad successfully verified",
		Data:    nil,
	}

	return c.JSON(resp.Code, resp)
}
