package facility

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/nauhalf/mobilcantik/db"
	"github.com/nauhalf/mobilcantik/repository/facilityrepository"
	"github.com/nauhalf/mobilcantik/repository/model"
	"github.com/nauhalf/mobilcantik/repository/vehicletyperepository"
	"github.com/nauhalf/mobilcantik/response"
	errorutils "github.com/nauhalf/mobilcantik/utils/error"
)

type RequestCreate struct {
	FacilityName  string  `json:"szFacilityName" form:"szFacilityName" query:"szFacilityName" validate:"required"`
	Annotation    *string `json:"szAnnotation" form:"szAnnotation" query:"szAnnotation"`
	VehicleTypeId uint64  `json:"intVehicleTypeId" form:"intVehicleTypeId" query:"intVehicleTypeId" validate:"required"`
}

type RequestUpdate struct {
	FacilityId uint64 `json:"intFacilityId" form:"intFacilityId" query:"intFacilityId" validate:"required"`
	RequestCreate
}

func GetAll(c echo.Context) error {

	typeid := c.QueryParam("vehicletype_id")
	var id *uint64
	if typeid == "" {
		id = nil
	} else {

		intId, err := strconv.ParseUint(typeid, 10, 64)

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

	facilities, err := facilityrepository.GetAll(db.DBCon, id)

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
		Message: fmt.Sprintf(errorutils.StatusErrorSuccessfullyRetrieved, "List Facilities"),
		Data:    facilities,
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

	fc := new(model.Facility)
	fc.FacilityName = r.FacilityName
	fc.VehicleTypeId = r.VehicleTypeId
	fc.Annotation = r.Annotation
	newFacility, err := facilityrepository.Create(db.DBCon, fc)
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
		Message: fmt.Sprintf(errorutils.StatusErrorSuccessfullyCreated, "Facility"),
		Data:    newFacility,
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

	fExists, _ := facilityrepository.GetById(db.DBCon, r.FacilityId)

	if fExists == nil {
		resp := response.ResponseError{
			Code:      http.StatusNotFound,
			Message:   fmt.Sprintf(errorutils.StatusErrorResourceNotExists, "Facility"),
			ErrorCode: 1,
		}
		return c.JSON(resp.Code, resp)
	}

	vExists, _ := vehicletyperepository.GetById(db.DBCon, r.VehicleTypeId)

	if vExists == nil {
		resp := response.ResponseError{
			Code:      http.StatusNotFound,
			Message:   fmt.Sprintf(errorutils.StatusErrorResourceNotExists, "Vehicle Type"),
			ErrorCode: 2,
		}
		return c.JSON(resp.Code, resp)
	}

	fc := new(model.Facility)
	fc.FacilityId = r.FacilityId
	fc.VehicleTypeId = r.VehicleTypeId
	fc.FacilityName = r.FacilityName
	fc.Annotation = r.Annotation

	err := facilityrepository.Update(db.DBCon, fc)

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
		Message: fmt.Sprintf(errorutils.StatusErrorSuccessfullyUpdated, "Facility"),
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
	err = facilityrepository.Delete(db.DBCon, intID)

	if err != nil {
		if err.Error() == errorutils.StatusZeroAffectedRows {
			resp := response.ResponseError{
				Code:      http.StatusNotFound,
				Message:   fmt.Sprintf(errorutils.StatusErrorResourceNotExists, "Facility"),
				ErrorCode: 1,
			}
			return c.JSON(resp.Code, resp)
		} else if err.(*mysql.MySQLError).Number == errorutils.ErrorMySQLDeleteConstraintFK {
			resp := response.ResponseError{
				Code:      http.StatusConflict,
				Message:   fmt.Sprintf(errorutils.StatusErrorAlreadyInUsed, "Facility"),
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
		Message: fmt.Sprintf(errorutils.StatusErrorSuccessfullyDeleted, "Facility"),
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

	facility, err := facilityrepository.GetById(db.DBCon, intID)

	if facility == nil {
		resp := response.ResponseError{
			Code:      http.StatusNotFound,
			Message:   fmt.Sprintf(errorutils.StatusErrorResourceNotExists, "Facility"),
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
		Message: fmt.Sprintf(errorutils.StatusErrorSuccessfullyRetrieved, "Facility"),
		Data:    facility,
	}
	return c.JSON(http.StatusOK, resp)
}
