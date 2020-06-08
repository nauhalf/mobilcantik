package adstatustype

import (
	"net/http"
	"strconv"

	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/nauhalf/mobilcantik/db"
	"github.com/nauhalf/mobilcantik/repository/adstatustyperepository"
	"github.com/nauhalf/mobilcantik/repository/model"
	"github.com/nauhalf/mobilcantik/response"
	errorutils "github.com/nauhalf/mobilcantik/utils/error"
)

type RequestCreate struct {
	AdStatusTypeName string `json:"szAdStatusTypeName" form:"szAdStatusTypeName" query:"szAdType" validate:"required"`
}

type RequestUpdate struct {
	AdStatusTypeId uint64 `json:"intAdStatusTypeId" form:"intAdStatusTypeId" query:"intAdStatusTypeId" validate:"required"`
	RequestCreate
}

func GetAll(c echo.Context) error {

	adstatustypes, err := adstatustyperepository.GetAll(db.DBCon)

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
		Message: "List Ad Status Type successfully retrieved",
		Data:    adstatustypes,
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

	adstatustype := new(model.AdStatusType)
	adstatustype.AdStatusTypeName = r.AdStatusTypeName

	newAdStatusType, err := adstatustyperepository.Create(db.DBCon, adstatustype)
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
		Message: "Ad Status Type successfully created",
		Data:    newAdStatusType,
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
			Message:   http.StatusText(http.StatusUnprocessableEntity),
			ErrorCode: 1,
		}
		return c.JSON(resp.Code, resp)
	}

	exists, _ := adstatustyperepository.GetById(db.DBCon, r.AdStatusTypeId)

	if exists == nil {
		resp := response.ResponseError{
			Code:      http.StatusNotFound,
			Message:   http.StatusText(http.StatusNotFound),
			ErrorCode: 1,
		}
		return c.JSON(resp.Code, resp)
	}

	adstatustype := new(model.AdStatusType)
	adstatustype.AdStatusTypeId = r.AdStatusTypeId
	adstatustype.AdStatusTypeName = r.AdStatusTypeName

	err := adstatustyperepository.Update(db.DBCon, adstatustype)

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
		Message: "Ad Status Type successfully updated",
		Data:    nil,
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

	err = adstatustyperepository.Delete(db.DBCon, intID)

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
				Message:   "Ad Status Type is already in used, failed to delete it.",
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
		Code:    http.StatusCreated,
		Message: "Ad Status Type successfully deleted",
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

	adstatustype, err := adstatustyperepository.GetById(db.DBCon, intID)

	if adstatustype == nil {
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

	resp := response.ResponseSuccess{
		Code:    http.StatusOK,
		Message: "Ad Status Type successfully retrieved",
		Data:    adstatustype,
	}
	return c.JSON(http.StatusOK, resp)
}
