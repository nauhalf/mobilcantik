package fuel

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/nauhalf/mobilcantik/db"
	"github.com/nauhalf/mobilcantik/repository/fuelrepository"
	"github.com/nauhalf/mobilcantik/repository/model"
	"github.com/nauhalf/mobilcantik/repository/vehicletyperepository"
	"github.com/nauhalf/mobilcantik/response"
	errorutils "github.com/nauhalf/mobilcantik/utils/error"
)

type RequestCreate struct {
	FuelName   string  `json:"szFuelName" form:"szFuelName" query:"szFuelName" validate:"required"`
	Annotation *string `json:"szAnnotation" form:"szAnnotation" query:"szAnnotation"`
}

type RequestUpdate struct {
	FuelId uint64 `json:"intFuelId" form:"intFuelId" query:"intFuelId" validate:"required"`
	RequestCreate
}

type RequestMapping struct {
	VehicleTypeId uint64   `json:"intVehicleTypeId" form:"intVehicleTypeId" validate:"required"`
	FuelIds       []uint64 `json:"intFuelIds" form:"intFuelIds" validate:"required"`
}

func GetAll(c echo.Context) error {

	fuels, err := fuelrepository.GetAll(db.DBCon)

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
		Message: fmt.Sprintf(errorutils.StatusErrorSuccessfullyRetrieved, "List Fuels"),
		Data:    fuels,
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

	f := new(model.Fuel)
	f.FuelName = r.FuelName
	f.Annotation = r.Annotation

	newFuel, err := fuelrepository.Create(db.DBCon, f)
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
		Message: fmt.Sprintf(errorutils.StatusErrorSuccessfullyCreated, "Fuel"),
		Data:    newFuel,
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

	vExists, _ := fuelrepository.GetById(db.DBCon, r.FuelId)

	if vExists == nil {
		resp := response.ResponseError{
			Code:      http.StatusNotFound,
			Message:   fmt.Sprintf(errorutils.StatusErrorResourceNotExists, "Fuel"),
			ErrorCode: 1,
		}
		return c.JSON(resp.Code, resp)
	}

	f := new(model.Fuel)
	f.FuelId = r.FuelId
	f.FuelName = r.FuelName
	f.Annotation = r.Annotation

	err := fuelrepository.Update(db.DBCon, f)

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
		Message: fmt.Sprintf(errorutils.StatusErrorSuccessfullyUpdated, "Fuel"),
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
	err = fuelrepository.Delete(db.DBCon, intID)

	if err != nil {
		if err.Error() == errorutils.StatusZeroAffectedRows {
			resp := response.ResponseError{
				Code:      http.StatusNotFound,
				Message:   fmt.Sprintf(errorutils.StatusErrorResourceNotExists, "Fuel"),
				ErrorCode: 1,
			}
			return c.JSON(resp.Code, resp)
		} else if err.(*mysql.MySQLError).Number == errorutils.ErrorMySQLDeleteConstraintFK {
			resp := response.ResponseError{
				Code:      http.StatusConflict,
				Message:   fmt.Sprintf(errorutils.StatusErrorAlreadyInUsed, "Fuel"),
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
		Message: fmt.Sprintf(errorutils.StatusErrorSuccessfullyDeleted, "Fuel"),
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

	fuel, err := fuelrepository.GetById(db.DBCon, intID)

	if fuel == nil {
		resp := response.ResponseError{
			Code:      http.StatusNotFound,
			Message:   fmt.Sprintf(errorutils.StatusErrorResourceNotExists, "Fuel"),
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
		Message: "Fuel successfully retrieved",
		Data:    fuel,
	}
	return c.JSON(http.StatusOK, resp)
}

func GetMapping(c echo.Context) error {

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

	fuels, err := fuelrepository.GetAllMapping(db.DBCon, id)

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
		Message: fmt.Sprintf(errorutils.StatusErrorSuccessfullyRetrieved, "Fuel Mapping"),
		Data:    fuels,
	}
	return c.JSON(http.StatusOK, resp)
}

func DoMapping(c echo.Context) error {
	r := new(RequestMapping)

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

	if r.FuelIds[0] == 0 || len(r.FuelIds) < 1 {
		resp := response.ResponseError{
			Code:      http.StatusUnprocessableEntity,
			Message:   "Fuel array must be greated than zero",
			ErrorCode: 2,
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

	mapping, err := fuelrepository.DoMap(db.DBCon, r.VehicleTypeId, r.FuelIds)
	if err != nil {

		if err.(*mysql.MySQLError).Number == errorutils.ErrorMySQLInsertUpdateConstraintFK {
			resp := response.ResponseError{
				Code:      http.StatusNotFound,
				Message:   fmt.Sprintf(errorutils.StatusErrorResourceNotExists, "Facility / Vehicle Type"),
				ErrorCode: 2,
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
		Code:    http.StatusCreated,
		Message: "Fuel with vehicle type successfully mapped",
		Data:    mapping,
	}

	return c.JSON(resp.Code, resp)
}
