package reporttype

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/nauhalf/mobilcantik/db"
	"github.com/nauhalf/mobilcantik/repository/model"
	"github.com/nauhalf/mobilcantik/repository/reportgrouprepository"
	"github.com/nauhalf/mobilcantik/repository/reporttyperepository"
	"github.com/nauhalf/mobilcantik/response"
	errorutils "github.com/nauhalf/mobilcantik/utils/error"
)

type RequestCreate struct {
	ReportTypeName string  `json:"szReportTypeName" form:"szReportTypeName" query:"szReportTypeName" validate:"required"`
	Annotation     *string `json:"szAnnotation" form:"szAnnotation" query:"szAnnotation"`
	ReportGroupId  uint64  `json:"intReportGroupId" form:"intReportGroupId" query:"intReportGroupId" validate:"required"`
}

type RequestUpdate struct {
	ReportTypeId uint64 `json:"intReportTypeId" form:"intReportTypeId" query:"intReportTypeId" validate:"required"`
	RequestCreate
}

func GetAll(c echo.Context) error {

	groupid := c.QueryParam("reportgroup_id")
	var id *uint64
	if groupid == "" {
		id = nil
	} else {

		intId, err := strconv.ParseUint(groupid, 10, 64)

		if err != nil {
			resp := response.ResponseError{
				Code:      http.StatusInternalServerError,
				Message:   http.StatusText(http.StatusInternalServerError),
				ErrorCode: nil,
			}
			return c.JSON(resp.Code, resp)
		}
		id = &intId
	}

	reporttypes, err := reporttyperepository.GetAll(db.DBCon, id)

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
		Message: fmt.Sprintf(errorutils.StatusErrorSuccessfullyRetrieved, "List Report Types"),
		Data:    reporttypes,
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

	rt := new(model.ReportType)
	rt.ReportTypeName = r.ReportTypeName
	rt.ReportGroupId = r.ReportGroupId
	rt.Annotation = r.Annotation
	newReportType, err := reporttyperepository.Create(db.DBCon, rt)
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
		Message: fmt.Sprintf(errorutils.StatusErrorSuccessfullyCreated, "Report Type"),
		Data:    newReportType,
	}

	return c.JSON(resp.Code, resp)
}

func Update(c echo.Context) error {

	r := new(RequestUpdate)

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

	rtExists, _ := reporttyperepository.GetById(db.DBCon, r.ReportTypeId)

	if rtExists == nil {
		resp := response.ResponseError{
			Code:      http.StatusNotFound,
			Message:   fmt.Sprintf(errorutils.StatusErrorResourceNotExists, "Report Type"),
			ErrorCode: 1,
		}
		return c.JSON(resp.Code, resp)
	}

	rgExists, _ := reportgrouprepository.GetById(db.DBCon, r.ReportGroupId)

	if rgExists == nil {
		resp := response.ResponseError{
			Code:      http.StatusNotFound,
			Message:   fmt.Sprintf(errorutils.StatusErrorResourceNotExists, "Report Group"),
			ErrorCode: 2,
		}
		return c.JSON(resp.Code, resp)
	}

	rt := new(model.ReportType)
	rt.ReportTypeId = r.ReportTypeId
	rt.ReportGroupId = r.ReportGroupId
	rt.ReportTypeName = r.ReportTypeName
	rt.Annotation = r.Annotation

	err := reporttyperepository.Update(db.DBCon, rt)

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
		Message: fmt.Sprintf(errorutils.StatusErrorSuccessfullyUpdated, "Report Type"),
		Data:    nil,
	}

	return c.JSON(resp.Code, resp)
}

func Delete(c echo.Context) error {

	fID := c.Param("id")

	if fID == "" {
		resp := response.ResponseError{
			Code:      http.StatusUnprocessableEntity,
			Message:   errorutils.StatusErrorFillRequiredForm,
			ErrorCode: 1,
		}
		return c.JSON(resp.Code, resp)
	}

	intID, err := strconv.ParseUint(fID, 10, 64)
	if err != nil {
		resp := response.ResponseError{
			Code:      http.StatusInternalServerError,
			Message:   http.StatusText(http.StatusInternalServerError),
			ErrorCode: nil,
		}
		return c.JSON(resp.Code, resp)
	}
	err = reporttyperepository.Delete(db.DBCon, intID)

	if err != nil {
		if err.Error() == errorutils.StatusZeroAffectedRows {
			resp := response.ResponseError{
				Code:      http.StatusNotFound,
				Message:   fmt.Sprintf(errorutils.StatusErrorResourceNotExists, "Report Type"),
				ErrorCode: 1,
			}
			return c.JSON(resp.Code, resp)
		} else if err.(*mysql.MySQLError).Number == errorutils.ErrorMySQLDeleteConstraintFK {
			resp := response.ResponseError{
				Code:      http.StatusConflict,
				Message:   fmt.Sprintf(errorutils.StatusErrorAlreadyInUsed, "Report Type"),
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
		Message: fmt.Sprintf(errorutils.StatusErrorSuccessfullyDeleted, "Report Type"),
		Data:    nil,
	}

	return c.JSON(resp.Code, resp)
}

func GetById(c echo.Context) error {

	fID := c.Param("id")

	if fID == "" {
		resp := response.ResponseError{
			Code:      http.StatusUnprocessableEntity,
			Message:   errorutils.StatusErrorFillRequiredForm,
			ErrorCode: 1,
		}
		return c.JSON(resp.Code, resp)
	}

	intID, err := strconv.ParseUint(fID, 10, 64)

	if err != nil {
		resp := response.ResponseError{
			Code:      http.StatusInternalServerError,
			Message:   http.StatusText(http.StatusInternalServerError),
			ErrorCode: nil,
		}
		return c.JSON(resp.Code, resp)
	}

	reporttype, err := reporttyperepository.GetById(db.DBCon, intID)

	if reporttype == nil {
		resp := response.ResponseError{
			Code:      http.StatusNotFound,
			Message:   fmt.Sprintf(errorutils.StatusErrorResourceNotExists, "Report Type"),
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

	resp := response.ResponseSuccess{
		Code:    http.StatusOK,
		Message: fmt.Sprintf(errorutils.StatusErrorSuccessfullyRetrieved, "Report Type"),
		Data:    reporttype,
	}
	return c.JSON(http.StatusOK, resp)
}
