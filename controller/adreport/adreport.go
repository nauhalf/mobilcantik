package adreport

import (
	"net/http"
	"strconv"

	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/nauhalf/mobilcantik/db"
	"github.com/nauhalf/mobilcantik/repository/adreportrepository"
	"github.com/nauhalf/mobilcantik/repository/adrepository"
	"github.com/nauhalf/mobilcantik/repository/model"
	"github.com/nauhalf/mobilcantik/repository/reporttyperepository"
	"github.com/nauhalf/mobilcantik/response"
	"github.com/nauhalf/mobilcantik/utils"
	errorutils "github.com/nauhalf/mobilcantik/utils/error"
)

type RequestCreate struct {
	AdId          uint64 `json:"intAdId" form:"intAdId" validate:"required"`
	ReportTypeId  uint64 `json:"intReportTypeId " form:"intReportTypeId" validate:"required"`
	Annotation    string `json:"szAnnotation" form:"szAnnotation" validate:"required"`
	ReporterName  string `json:"szReporterName" form:"szReporterName" validate:"required"`
	ReporterEmail string `json:"szReporterEmail" form:"szReporterEmail" validate:"required"`
}

func GetAll(c echo.Context) error {

	page := c.QueryParam("page")

	var p int
	tmpP, err := strconv.Atoi(page)
	if tmpP < 1 || err != nil {
		p = 1
	} else {
		p = tmpP
	}

	all, err := adreportrepository.GetAll(db.DBCon, p)

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
		Message: "List Ad Report successfully retrieved",
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
			Message:   http.StatusText(http.StatusUnprocessableEntity),
			ErrorCode: 1,
		}
		return c.JSON(resp.Code, resp)
	}

	validEmail := utils.ValidateEmail(r.ReporterEmail)

	if !validEmail {
		resp := response.ResponseError{
			Code:      http.StatusUnprocessableEntity,
			Message:   http.StatusText(http.StatusUnprocessableEntity),
			ErrorCode: 2,
		}
		return c.JSON(resp.Code, resp)
	}

	if len(r.Annotation) < 15 {
		resp := response.ResponseError{
			Code:      http.StatusUnprocessableEntity,
			Message:   http.StatusText(http.StatusUnprocessableEntity),
			ErrorCode: 3,
		}
		return c.JSON(resp.Code, resp)
	}

	adExists, err := adrepository.IsAdExists(db.DBCon, r.AdId)

	if err != nil {
		resp := response.ResponseError{
			Code:      http.StatusInternalServerError,
			Message:   http.StatusText(http.StatusInternalServerError),
			ErrorCode: nil,
		}
		return c.JSON(resp.Code, resp)
	}

	if adExists == 0 {
		resp := response.ResponseError{
			Code:      http.StatusNotFound,
			Message:   http.StatusText(http.StatusNotFound),
			ErrorCode: 1,
		}
		return c.JSON(resp.Code, resp)
	}

	reportTypeExists, _ := reporttyperepository.GetById(db.DBCon, r.ReportTypeId)

	if reportTypeExists == nil {
		resp := response.ResponseError{
			Code:      http.StatusNotFound,
			Message:   http.StatusText(http.StatusNotFound),
			ErrorCode: 2,
		}
		return c.JSON(resp.Code, resp)
	}

	adreport := new(model.AdReport)
	adreport.AdId = r.AdId
	adreport.ReportTypeId = r.ReportTypeId
	adreport.Annotation = r.Annotation
	adreport.ReporterName = r.ReporterName
	adreport.ReporterEmail = r.ReporterEmail

	newAdReport, err := adreportrepository.Create(db.DBCon, adreport)

	ad, err := adrepository.GetAdById(db.DBCon, adreport.AdId)
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

	newAdReport.AdSingleResponse = *adsingleresponse
	if err != nil {
		resp := response.ResponseError{
			Code:      http.StatusInternalServerError,
			Message:   http.StatusText(http.StatusInternalServerError),
			ErrorCode: nil,
		}
		return c.JSON(resp.Code, resp)
	}

	resp := response.ResponseSuccess{
		Code:    http.StatusCreated,
		Message: "Ad Report successfully created",
		Data:    newAdReport,
	}

	return c.JSON(resp.Code, resp)
}

func Delete(c echo.Context) error {

	cID := c.Param("id")

	if cID == "" {
		resp := response.ResponseError{
			Code:      http.StatusUnprocessableEntity,
			Message:   http.StatusText(http.StatusUnprocessableEntity),
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

	err = adreportrepository.Delete(db.DBCon, intID)

	if err != nil {
		if err.Error() == errorutils.StatusZeroAffectedRows {
			resp := response.ResponseError{
				Code:      http.StatusNotFound,
				Message:   http.StatusText(http.StatusNotFound),
				ErrorCode: 1,
			}
			return c.JSON(resp.Code, resp)
		} else if err.(*mysql.MySQLError).Number == errorutils.ErrorMySQLDeleteConstraintFK {
			resp := response.ResponseError{
				Code:      http.StatusConflict,
				Message:   "Ad Report is already in used, failed to delete it.",
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

	resp := response.ResponseSuccess{
		Code:    http.StatusOK,
		Message: "Ad Report successfully deleted",
		Data:    nil,
	}

	return c.JSON(resp.Code, resp)
}

func GetById(c echo.Context) error {

	cID := c.Param("id")

	if cID == "" {
		resp := response.ResponseError{
			Code:      http.StatusUnprocessableEntity,
			Message:   http.StatusText(http.StatusUnprocessableEntity),
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

	adreport, err := adreportrepository.GetById(db.DBCon, intID)

	if adreport == nil {
		resp := response.ResponseError{
			Code:      http.StatusNotFound,
			Message:   http.StatusText(http.StatusNotFound),
			ErrorCode: 1,
		}
		return c.JSON(resp.Code, resp)
	}

	if err != nil {
		resp := response.ResponseError{
			Code:      http.StatusInternalServerError,
			Message:   http.StatusText(http.StatusInternalServerError),
			ErrorCode: nil,
		}
		return c.JSON(http.StatusInternalServerError, resp)
	}

	ad, err := adrepository.GetAdById(db.DBCon, adreport.AdReport.AdId)
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

	adreport.AdSingleResponse = *adsingleresponse

	resp := response.ResponseSuccess{
		Code:    http.StatusOK,
		Message: "Ad Report successfully retrieved",
		Data:    adreport,
	}
	return c.JSON(http.StatusOK, resp)
}
