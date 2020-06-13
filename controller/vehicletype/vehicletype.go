package vehicletype

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/nauhalf/mobilcantik/db"
	"github.com/nauhalf/mobilcantik/repository/model"
	"github.com/nauhalf/mobilcantik/repository/vehicletyperepository"
	"github.com/nauhalf/mobilcantik/response"
	errorutils "github.com/nauhalf/mobilcantik/utils/error"
)

func GetAll(c echo.Context) error {

	vehicletypes, err := vehicletyperepository.GetAll(db.DBCon)

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
		Message: "List vehicle type successfully retrieved",
		Data:    vehicletypes,
	}
	return c.JSON(http.StatusOK, resp)
}

func Create(c echo.Context) error {
	vehicleTypeName := c.FormValue("szVehicleTypeName")
	annotation := c.FormValue("szAnnotation")

	if vehicleTypeName == "" {
		resp := response.ResponseError{
			Code:      http.StatusUnprocessableEntity,
			Message:   errorutils.StatusErrorFillRequiredForm,
			ErrorCode: 1,
		}

		return c.JSON(resp.Code, resp)
	}

	vehicleType := new(model.VehicleType)
	vehicleType.VehicleTypeName = vehicleTypeName
	if annotation == "" {
		vehicleType.Annotation = nil
	} else {
		vehicleType.Annotation = &annotation
	}

	newVehicleType, err := vehicletyperepository.Create(db.DBCon, vehicleType)
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
		Message: fmt.Sprintf(errorutils.StatusErrorSuccessfullyCreated, "Vehicle Type"),
		Data:    newVehicleType,
	}

	return c.JSON(resp.Code, resp)
}

type Request struct {
	VehicleTypeId   uint64 `json:"intVehicleTypeId" form:"intVehicleTypeId" query:"intVehicleTypeId" validate:"required"`
	VehicleTypeName string `json:"szVehicleTypeName" form:"szVehicleTypeName" query:"szVehicleTypeName" validate:"required"`
	Annotation      string `json:"szAnnotation" form:"szAnnotation" query:"szAnnotation"`
}

func Update(c echo.Context) error {

	r := new(Request)

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

	vExists, _ := vehicletyperepository.GetById(db.DBCon, r.VehicleTypeId)

	if vExists == nil {
		resp := response.ResponseError{
			Code:      http.StatusNotFound,
			Message:   fmt.Sprintf(errorutils.StatusErrorResourceNotExists, "Vehicle Type"),
			ErrorCode: 1,
		}
		return c.JSON(resp.Code, resp)
	}

	v := new(model.VehicleType)
	v.VehicleTypeId = r.VehicleTypeId
	v.VehicleTypeName = r.VehicleTypeName
	if r.Annotation == "" {
		v.Annotation = nil
	} else {
		v.Annotation = &r.Annotation
	}
	err := vehicletyperepository.Update(db.DBCon, v)

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
		Message: fmt.Sprintf(errorutils.StatusErrorSuccessfullyUpdated, "Vehicle Type"),
		Data:    nil,
	}

	return c.JSON(resp.Code, resp)
}

func Delete(c echo.Context) error {

	vID := c.Param("id")

	if vID == "" {
		resp := response.ResponseError{
			Code:      http.StatusUnprocessableEntity,
			Message:   errorutils.StatusErrorFillRequiredForm,
			ErrorCode: 1,
		}
		return c.JSON(resp.Code, resp)
	}

	intID, err := strconv.ParseUint(vID, 10, 64)
	if err != nil {
		resp := response.ResponseError{
			Code:      http.StatusInternalServerError,
			Message:   http.StatusText(http.StatusInternalServerError),
			ErrorCode: nil,
		}
		return c.JSON(resp.Code, resp)
	}
	err = vehicletyperepository.Delete(db.DBCon, intID)

	if err != nil {
		if err.Error() == errorutils.StatusZeroAffectedRows {
			resp := response.ResponseError{
				Code:      http.StatusNotFound,
				Message:   fmt.Sprintf(errorutils.StatusErrorResourceNotExists, "Vehicle Type"),
				ErrorCode: 1,
			}
			return c.JSON(resp.Code, resp)
		} else if err.(*mysql.MySQLError).Number == errorutils.ErrorMySQLDeleteConstraintFK {
			resp := response.ResponseError{
				Code:      http.StatusConflict,
				Message:   fmt.Sprintf(errorutils.StatusErrorAlreadyInUsed, "Vehicle Type"),
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
		Message: fmt.Sprintf(errorutils.StatusErrorSuccessfullyDeleted, "Vehicle Type"),
		Data:    nil,
	}

	return c.JSON(resp.Code, resp)
}

func GetById(c echo.Context) error {

	vID := c.Param("id")

	if vID == "" {
		resp := response.ResponseError{
			Code:      http.StatusUnprocessableEntity,
			Message:   errorutils.StatusErrorFillRequiredForm,
			ErrorCode: 1,
		}
		return c.JSON(resp.Code, resp)
	}

	intID, err := strconv.ParseUint(vID, 10, 64)

	if err != nil {
		resp := response.ResponseError{
			Code:      http.StatusInternalServerError,
			Message:   http.StatusText(http.StatusInternalServerError),
			ErrorCode: nil,
		}
		return c.JSON(resp.Code, resp)
	}

	vehicletype, err := vehicletyperepository.GetById(db.DBCon, intID)

	if vehicletype == nil {
		resp := response.ResponseError{
			Code:      http.StatusNotFound,
			Message:   fmt.Sprintf(errorutils.StatusErrorResourceNotExists, "Vehicle Type"),
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
		Message: fmt.Sprintf(errorutils.StatusErrorSuccessfullyRetrieved, "Vehicle Type"),
		Data:    vehicletype,
	}
	return c.JSON(http.StatusOK, resp)
}
