package adbill

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/nauhalf/mobilcantik/db"
	"github.com/nauhalf/mobilcantik/repository/adbillrepository"
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
		Message: fmt.Sprintf(errorutils.StatusErrorSuccessfullyRetrieved, "List ad bills"),
		Data:    all,
	}
	return c.JSON(http.StatusOK, resp)
}

func ConfirmPaymentBill(c echo.Context) error {

	id := c.FormValue("intAdId")
	password := c.FormValue("password")

	if id == "" || password == "" {

		resp := response.ResponseError{
			Code:      http.StatusUnprocessableEntity,
			Message:   errorutils.StatusErrorFillRequiredForm,
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
			Message:   errorutils.StatusErrorIncorrectPassword,
			ErrorCode: 1,
		}
		return c.JSON(resp.Code, resp)
	}

	if ad.AdStatusTypeId != 1 {
		resp := response.ResponseError{
			Code:      http.StatusUnprocessableEntity,
			Message:   errorutils.StatusErrorNotInWaitingFeeConfirmationState,
			ErrorCode: 2,
		}
		return c.JSON(resp.Code, resp)
	}

	adbill, _ := adbillrepository.GetByAdId(db.DBCon, intAdId)
	if ad == nil {
		resp := response.ResponseError{
			Code:      http.StatusNotFound,
			Message:   fmt.Sprintf(errorutils.StatusErrorResourceNotExists, "Ad Bill"),
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
		Message: fmt.Sprintf(errorutils.StatusErrorSuccessfullyConfirmed, "Ad Bill"),
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

	adbill, _ := adbillrepository.GetById(db.DBCon, intID)

	if adbill == nil {
		resp := response.ResponseError{
			Code:      http.StatusNotFound,
			Message:   fmt.Sprintf(errorutils.StatusErrorResourceNotExists, "Ad Bill"),
			ErrorCode: 1,
		}
		return c.JSON(resp.Code, resp)
	}

	adconfirmation, _ := adconfirmationrepository.GetByAdId(db.DBCon, adbill.AdBill.AdId)

	adbill.AdConfirmations = adconfirmation

	resp := response.ResponseSuccess{
		Code:    http.StatusOK,
		Message: fmt.Sprintf(errorutils.StatusErrorSuccessfullyRetrieved, "Ad Bill"),
		Data:    adbill,
	}
	return c.JSON(http.StatusOK, resp)
}
