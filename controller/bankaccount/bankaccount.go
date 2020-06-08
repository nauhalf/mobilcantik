package bankaccount

import (
	"net/http"
	"strconv"

	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/nauhalf/mobilcantik/db"
	"github.com/nauhalf/mobilcantik/repository/bankaccountrepository"
	"github.com/nauhalf/mobilcantik/repository/model"
	"github.com/nauhalf/mobilcantik/response"
	errorutils "github.com/nauhalf/mobilcantik/utils/error"
)

type RequestCreate struct {
	BankAccountNumber string `json:"szBankAccountNumber" form:"szBankAccountNumber" query:"szBankAccountNumber" validate:"required"`
	BankCode          string `json:"szBankCode" form:"szBankCode" query:"szBankCode" validate:"required"`
	BankAccountName   string `json:"szBankAccountName" form:"szBankAccountName" query:"szBankAccountName" validate:"required"`
}

type RequestUpdate struct {
	BankAccountId uint64 `json:"intBankAccountId" form:"intBankAccountId" query:"intBankAccountId" validate:"required"`
	IsActive      string `json:"isActive" form:"isActive" query:"isActive" validate:"required"`
	RequestCreate
}

func GetAll(c echo.Context) error {

	is_active := c.QueryParam("isactive")
	var isActive *bool

	if is_active == "" {
		isActive = nil
	} else {
		tmpIsActive, err := strconv.ParseBool(is_active)
		if err != nil {
			resp := response.ResponseError{
				Code:      http.StatusInternalServerError,
				Message:   http.StatusText(http.StatusInternalServerError),
				ErrorCode: nil,
			}
			return c.JSON(resp.Code, resp)
		}
		isActive = &tmpIsActive
	}

	banks, err := bankaccountrepository.GetAll(db.DBCon, isActive)

	if err != nil {
		resp := response.ResponseError{
			Code:      http.StatusInternalServerError,
			Message:   http.StatusText(http.StatusInternalServerError),
			ErrorCode: nil,
		}
		return c.JSON(http.StatusInternalServerError, resp)
	}

	if banks == nil {
		banks = []model.BankAccount{}
	}

	resp := response.ResponseSuccess{
		Code:    http.StatusOK,
		Message: "List bank account successfully retrieved",
		Data:    banks,
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

	bankaccount := new(model.BankAccount)
	bankaccount.BankAccountNumber = r.BankAccountNumber
	bankaccount.BankCode = r.BankCode
	bankaccount.BankAccountName = r.BankAccountName

	newbank, err := bankaccountrepository.Create(db.DBCon, bankaccount)
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
		Message: "Bank Account successfully created",
		Data:    newbank,
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

	exists, _ := bankaccountrepository.GetById(db.DBCon, r.BankAccountId)

	if exists == nil {
		resp := response.ResponseError{
			Code:      http.StatusNotFound,
			Message:   http.StatusText(http.StatusNotFound),
			ErrorCode: 1,
		}
		return c.JSON(resp.Code, resp)
	}

	isActive, err := strconv.ParseBool(r.IsActive)

	if err != nil {
		resp := response.ResponseError{
			Code:      http.StatusInternalServerError,
			Message:   http.StatusText(http.StatusInternalServerError),
			ErrorCode: nil,
		}
		return c.JSON(resp.Code, resp)
	}

	bankaccount := new(model.BankAccount)
	bankaccount.BankAccountId = r.BankAccountId
	bankaccount.BankAccountNumber = r.BankAccountNumber
	bankaccount.BankCode = r.BankCode
	bankaccount.BankAccountName = r.BankAccountName

	bankaccount.IsActive = isActive

	err = bankaccountrepository.Update(db.DBCon, bankaccount)

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
		Message: "Bank Account successfully updated",
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

	err = bankaccountrepository.Delete(db.DBCon, intID)

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
				Message:   "Bank Account is already in used, failed to delete it.",
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
		Message: "Bank Account successfully deleted",
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

	bankaccount, err := bankaccountrepository.GetById(db.DBCon, intID)

	if bankaccount == nil {
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
		Message: "Bank Account successfully retrieved",
		Data:    bankaccount,
	}
	return c.JSON(http.StatusOK, resp)
}
