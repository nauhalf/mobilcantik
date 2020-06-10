package adconfirmation

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/nauhalf/mobilcantik/db"
	"github.com/nauhalf/mobilcantik/repository/adbillrepository"
	"github.com/nauhalf/mobilcantik/repository/adconfirmationrepository"
	"github.com/nauhalf/mobilcantik/repository/adrepository"
	"github.com/nauhalf/mobilcantik/repository/bankaccountrepository"
	"github.com/nauhalf/mobilcantik/repository/model"
	"github.com/nauhalf/mobilcantik/response"
	"github.com/nauhalf/mobilcantik/utils"
)

type RequestCreate struct {
	AdId                uint64 `json:"intAdId" form:"intAdId" query:"intAdId" validate:"required"`
	Password            string `json:"szPassword" form:"szPassword" query:"szPassword" validate:"required"`
	AdBillId            uint64 `json:"intAdBillId" form:"intAdBillId" query:"intAdBillId" validate:"required"`
	BankAccountId       uint64 `json:"intBankAccountId" form:"intBankAccountId" query:"intBankAccountId" validate:"required"`
	BankAccountNumber   string `json:"szBankAccountNumber" form:"szBankAccountNumber" query:"szBankAccountNumber" validate:"required"`
	BankName            string `json:"szBankName" form:"szBankName" query:"szBankName" validate:"required"`
	BankBeneficiaryName string `json:"szBankBeneficiaryName" form:"szBankBeneficiaryName" query:"szBankBeneficiaryName" validate:"required"`
	Annotation          string `json:"szAnnotation" form:"szAnnotation" query:"szAnnotation" validate:"required"`
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
			Message:   http.StatusText(http.StatusUnprocessableEntity),
			ErrorCode: 1,
		}
		return c.JSON(resp.Code, resp)
	}

	//validate exists data

	if validate1 := ValidateCreateExists(*r); validate1 != nil {
		return c.JSON(validate1.Code, validate1)
	}

	form, _ := c.MultipartForm()

	images := form.File["image"]
	validate2 := ValidateFile(images)
	if validate2 != nil {
		return c.JSON(validate2.Code, validate2)
	}

	now := time.Now()
	year := now.Year()
	month := now.Month()
	day := now.Day()

	newfile, err := HandleUpload(images)
	dir := fmt.Sprintf("/%s/%s/%s", strconv.FormatInt(int64(year), 10), strconv.FormatInt(int64(month), 10), strconv.FormatInt(int64(day), 10))
	if err != nil {
		resp := response.ResponseError{
			Code:      http.StatusInternalServerError,
			Message:   err.Error(),
			ErrorCode: nil,
		}
		return c.JSON(resp.Code, resp)
	}

	adconfirmation := new(model.AdConfirmation)
	adconfirmation.AdId = r.AdId
	adconfirmation.AdBillId = r.AdBillId
	adconfirmation.BankAccountId = r.BankAccountId
	adconfirmation.BankAccountNumber = r.BankAccountNumber
	adconfirmation.BankName = r.BankName
	adconfirmation.BankBeneficiaryName = r.BankBeneficiaryName
	adconfirmation.Annotation = r.Annotation

	err = adconfirmationrepository.Create(db.DBCon, adconfirmation, newfile[0], dir)

	if err != nil {
		var deleteFile []string
		for _, file := range newfile {
			deleteFile = append(deleteFile, dir+"/"+file)
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
		Message: "Ad confirmation bill successfully created",
		Data:    nil,
	}

	return c.JSON(resp.Code, resp)
}

func ValidateCreateExists(r RequestCreate) *response.ResponseError {
	adExists, pw, _ := adrepository.GetAdById2(db.DBCon, r.AdId)
	if adExists == nil {
		resp := response.ResponseError{
			Code:      http.StatusNotFound,
			Message:   http.StatusText(http.StatusNotFound),
			ErrorCode: 1,
		}

		return &resp
	}

	if r.Password != *pw {
		resp := response.ResponseError{
			Code:      http.StatusUnprocessableEntity,
			Message:   http.StatusText(http.StatusUnprocessableEntity),
			ErrorCode: 2,
		}

		return &resp
	}

	adbillexists, _ := adbillrepository.GetByAdId(db.DBCon, r.AdId)
	if adbillexists == nil {
		resp := response.ResponseError{
			Code:      http.StatusNotFound,
			Message:   http.StatusText(http.StatusNotFound),
			ErrorCode: 2,
		}

		return &resp
	}

	if adbillexists.AdBillId != r.AdBillId {
		resp := response.ResponseError{
			Code:      http.StatusUnprocessableEntity,
			Message:   http.StatusText(http.StatusUnprocessableEntity),
			ErrorCode: 3,
		}

		return &resp
	}

	if adbillexists.Balance == 0 {
		resp := response.ResponseError{
			Code:      http.StatusUnprocessableEntity,
			Message:   http.StatusText(http.StatusUnprocessableEntity),
			ErrorCode: 4,
		}

		return &resp
	}

	bankAccountExists, _ := bankaccountrepository.GetById(db.DBCon, r.BankAccountId)
	if bankAccountExists == nil {
		resp := response.ResponseError{
			Code:      http.StatusNotFound,
			Message:   http.StatusText(http.StatusNotFound),
			ErrorCode: 3,
		}

		return &resp
	}

	return nil
}

func ValidateFile(file []*multipart.FileHeader) *response.ResponseError {

	if file == nil || len(file) < 1 {
		resp := response.ResponseError{
			Code:      http.StatusUnprocessableEntity,
			Message:   http.StatusText(http.StatusUnprocessableEntity),
			ErrorCode: 1,
		}
		return &resp

	}

	validateType := utils.IsFiletypeAccepted(file)
	if !validateType {
		resp := response.ResponseError{
			Code:      http.StatusUnprocessableEntity,
			Message:   http.StatusText(http.StatusUnprocessableEntity),
			ErrorCode: 4,
		}

		return &resp
	}

	for _, f := range file {
		validatesize := utils.IsFileSizeAccepted(f.Size)
		if !validatesize {
			resp := response.ResponseError{
				Code:      http.StatusUnprocessableEntity,
				Message:   http.StatusText(http.StatusUnprocessableEntity),
				ErrorCode: 5,
			}

			return &resp
		}
	}
	return nil
}

// func GetById(c echo.Context) error {

// 	cID := c.Param("id")

// 	if cID == "" {
// 		resp := response.ResponseError{
// 			Code:      http.StatusUnprocessableEntity,
// 			Message:   http.StatusText(http.StatusUnprocessableEntity),
// 			ErrorCode: 1,
// 		}
// 		return c.JSON(resp.Code, resp)
// 	}

// 	intID, err := strconv.ParseUint(cID, 10, 64)

// 	if err != nil {
// 		resp := response.ResponseError{
// 			Code:      http.StatusInternalServerError,
// 			Message:   http.StatusText(http.StatusInternalServerError),
// 			ErrorCode: nil,
// 		}
// 		return c.JSON(resp.Code, resp)
// 	}

// 	adx, _ := adrepository.GetAdById(db.DBCon, intID)

// 	if adx == nil {
// 		resp := response.ResponseError{
// 			Code:      http.StatusNotFound,
// 			Message:   http.StatusText(http.StatusNotFound),
// 			ErrorCode: 1,
// 		}
// 		return c.JSON(resp.Code, resp)
// 	}

// 	if adx.IsActive == false {
// 		resp := response.ResponseError{
// 			Code:      http.StatusUnprocessableEntity,
// 			Message:   http.StatusText(http.StatusUnprocessableEntity),
// 			ErrorCode: 1,
// 		}
// 		return c.JSON(resp.Code, resp)
// 	}

// 	err = adrepository.Increment(db.DBCon, intID)

// 	if err != nil {
// 		resp := response.ResponseError{
// 			Code:      http.StatusInternalServerError,
// 			Message:   http.StatusText(http.StatusInternalServerError),
// 			ErrorCode: nil,
// 		}
// 		return c.JSON(http.StatusInternalServerError, resp)
// 	}

// 	ad, err := adrepository.GetAdById(db.DBCon, intID)
// 	if err != nil {
// 		resp := response.ResponseError{
// 			Code:      http.StatusInternalServerError,
// 			Message:   http.StatusText(http.StatusInternalServerError),
// 			ErrorCode: nil,
// 		}
// 		return c.JSON(http.StatusInternalServerError, resp)
// 	}

// 	adsingleresponse := new(adrepository.AdSingleResponse)
// 	getV, _ := adrepository.GetAdVehicleByAdId(db.DBCon, ad.AdId)
// 	getF, _ := adrepository.GetFacilitiesByAdId(db.DBCon, ad.AdId)
// 	getI, _ := adrepository.GetImagesByAdId(db.DBCon, ad.AdId)
// 	adsingleresponse.Vehicle = *getV
// 	adsingleresponse.Facilities = getF
// 	adsingleresponse.Ad = *ad
// 	adsingleresponse.Images = getI

// 	resp := response.ResponseSuccess{
// 		Code:    http.StatusOK,
// 		Message: "Ad successfully retrieved",
// 		Data:    adsingleresponse,
// 	}
// 	return c.JSON(http.StatusOK, resp)
// }

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
		newFile := pathdir + "/" + fileName
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
