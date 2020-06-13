package vehiclebrand

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/nauhalf/mobilcantik/db"
	"github.com/nauhalf/mobilcantik/repository/model"
	"github.com/nauhalf/mobilcantik/repository/vehiclebrandrepository"
	"github.com/nauhalf/mobilcantik/repository/vehicletyperepository"
	"github.com/nauhalf/mobilcantik/response"
	errorutils "github.com/nauhalf/mobilcantik/utils/error"
)

type RequestCreate struct {
	VehicleBrandName string  `json:"szBrandName" form:"szBrandName" query:"szBrandName" validate:"required"`
	Annotation       *string `json:"szAnnotation" form:"szAnnotation" query:"szAnnotation"`
	VehicleTypeId    uint64  `json:"intVehicleTypeId" form:"intVehicleTypeId" query:"intVehicleTypeId" validate:"required"`
}

type RequestUpdate struct {
	VehicleBrandId uint64 `json:"intVehicleBrandId" form:"intVehicleBrandId" query:"intVehicleBrandId" validate:"required"`
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

	vehiclebrands, err := vehiclebrandrepository.GetAll(db.DBCon, id)

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
		Message: fmt.Sprintf(errorutils.StatusErrorSuccessfullyRetrieved, "List Vehicle Brands"),
		Data:    vehiclebrands,
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

	vb := new(model.VehicleBrand)
	vb.VehicleBrandName = r.VehicleBrandName
	vb.VehicleTypeId = r.VehicleTypeId
	vb.Annotation = r.Annotation
	newVehicleBrand, err := vehiclebrandrepository.Create(db.DBCon, vb)
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
		Message: fmt.Sprintf(errorutils.StatusErrorSuccessfullyCreated, "Vehicle Brand"),
		Data:    newVehicleBrand,
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

	vBexists, _ := vehiclebrandrepository.GetById(db.DBCon, r.VehicleBrandId)

	if vBexists == nil {
		resp := response.ResponseError{
			Code:      http.StatusNotFound,
			Message:   fmt.Sprintf(errorutils.StatusErrorResourceNotExists, "Vehicle Brand"),
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

	v := new(model.VehicleBrand)
	v.VehicleBrandId = r.VehicleBrandId
	v.VehicleTypeId = r.VehicleTypeId
	v.VehicleBrandName = r.VehicleBrandName
	v.Annotation = r.Annotation

	err := vehiclebrandrepository.Update(db.DBCon, v)

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
		Message: fmt.Sprintf(errorutils.StatusErrorSuccessfullyUpdated, "Vehicle Brand"),
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
	err = vehiclebrandrepository.Delete(db.DBCon, intID)

	if err != nil {
		if err.Error() == errorutils.StatusZeroAffectedRows {
			resp := response.ResponseError{
				Code:      http.StatusNotFound,
				Message:   fmt.Sprintf(errorutils.StatusErrorResourceNotExists, "Vehicle Brand"),
				ErrorCode: 1,
			}
			return c.JSON(resp.Code, resp)
		} else if err.(*mysql.MySQLError).Number == errorutils.ErrorMySQLDeleteConstraintFK {
			resp := response.ResponseError{
				Code:      http.StatusConflict,
				Message:   fmt.Sprintf(errorutils.StatusErrorAlreadyInUsed, "Vehicle Brand"),
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
		Message: fmt.Sprintf(errorutils.StatusErrorSuccessfullyDeleted, "Vehicle Brand"),
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

	vehiclebrand, err := vehiclebrandrepository.GetById(db.DBCon, intID)

	if vehiclebrand == nil {
		resp := response.ResponseError{
			Code:      http.StatusNotFound,
			Message:   fmt.Sprintf(errorutils.StatusErrorResourceNotExists, "Vehicle Brand"),
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
		Message: fmt.Sprintf(errorutils.StatusErrorSuccessfullyRetrieved, "Vehicle Brand"),
		Data:    vehiclebrand,
	}
	return c.JSON(http.StatusOK, resp)
}
