package adbill

import (
	"net/http"
	"strconv"

	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/nauhalf/mobilcantik/db"
	"github.com/nauhalf/mobilcantik/repository/adbillrepository"
	"github.com/nauhalf/mobilcantik/repository/adclosingreasonrepository"
	"github.com/nauhalf/mobilcantik/repository/adconfirmationrepository"
	"github.com/nauhalf/mobilcantik/repository/adrepository"
	"github.com/nauhalf/mobilcantik/repository/model"
	"github.com/nauhalf/mobilcantik/response"
	errorutils "github.com/nauhalf/mobilcantik/utils/error"
)

type RequestCreate struct {
	ReasonName string `json:"szReasonName" form:"szReasonName" query:"szReasonName" validate:"required"`
	Annotation string `json:"szAnnotation" form:"szAnnotation" query:"szAnnotation"`
}

type RequestUpdate struct {
	AdClosingReasonId uint64 `json:"intAdClosingReasonId" form:"intAdClosingReasonId" query:"intAdClosingReasonId" validate:"required"`
	RequestCreate
}

type AdBillResponse struct {
	model.AdBill
	confirmations []adconfirmationrepository.AdConfirmationResponse `json:"confirmations"`
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

	all, err := adbillrepository.GetAll(db.DBCon, p)

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
		Message: "List Ad Bill successfully retrieved",
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

	adclosingreason := new(model.AdClosingReason)
	adclosingreason.ReasonName = r.ReasonName
	adclosingreason.Annotation = &r.Annotation

	newAdClosingReason, err := adclosingreasonrepository.Create(db.DBCon, adclosingreason)
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
		Message: "Ad Closing Reason successfully created",
		Data:    newAdClosingReason,
	}

	return c.JSON(resp.Code, resp)
}

func ConfirmPaymentBill(c echo.Context) error {

	id := c.FormValue("intAdId")
	password := c.FormValue("password")

	if id == "" || password == "" {

		resp := response.ResponseError{
			Code:      http.StatusUnprocessableEntity,
			Message:   http.StatusText(http.StatusUnprocessableEntity),
			ErrorCode: 1,
		}
		return c.JSON(resp.Code, resp)
	}

	intAdId, err := strconv.ParseUint(id, 10, 64)
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
			Message:   http.StatusText(http.StatusNotFound),
			ErrorCode: 1,
		}
		return c.JSON(resp.Code, resp)
	}

	if password != *pw {
		resp := response.ResponseError{
			Code:      http.StatusUnprocessableEntity,
			Message:   http.StatusText(http.StatusUnprocessableEntity),
			ErrorCode: 1,
		}
		return c.JSON(resp.Code, resp)
	}

	if ad.AdStatusTypeId != 1 {
		resp := response.ResponseError{
			Code:      http.StatusUnprocessableEntity,
			Message:   http.StatusText(http.StatusUnprocessableEntity),
			ErrorCode: 2,
		}
		return c.JSON(resp.Code, resp)
	}

	adbill, _ := adbillrepository.GetByAdId(db.DBCon, intAdId)
	if ad == nil {
		resp := response.ResponseError{
			Code:      http.StatusNotFound,
			Message:   http.StatusText(http.StatusNotFound),
			ErrorCode: 2,
		}
		return c.JSON(resp.Code, resp)
	}

	adbill.Balance = 0

	err = adbillrepository.ConfirmPaymentBill(db.DBCon, adbill)

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
		Message: "Ad Bill Payment successfully confirmed",
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

	err = adclosingreasonrepository.Delete(db.DBCon, intID)

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
				Message:   "Ad Closing Reason is already in used, failed to delete it.",
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
		Message: "Ad Closing Reason successfully deleted",
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

	adbill, err := adbillrepository.GetById(db.DBCon, intID)

	if adbill == nil {
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

	adconfirmation, _ := adconfirmationrepository.GetByAdId(db.DBCon, adbill.AdBill.AdId)

	adbill.AdConfirmations = adconfirmation

	resp := response.ResponseSuccess{
		Code:    http.StatusOK,
		Message: "Ad Bill successfully retrieved",
		Data:    adbill,
	}
	return c.JSON(http.StatusOK, resp)
}
