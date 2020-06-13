package province

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nauhalf/mobilcantik/db"
	"github.com/nauhalf/mobilcantik/repository/provincerepository"
	"github.com/nauhalf/mobilcantik/response"
	errorutils "github.com/nauhalf/mobilcantik/utils/error"
)

func GetAll(c echo.Context) error {

	provinces, err := provincerepository.GetAll(db.DBCon)

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
		Message: fmt.Sprintf(errorutils.StatusErrorSuccessfullyRetrieved, "List Provinces"),
		Data:    provinces,
	}
	return c.JSON(http.StatusOK, resp)
}

func GetById(c echo.Context) error {

	pID := c.Param("id")

	if pID == "" {
		resp := response.ResponseError{
			Code:      http.StatusUnprocessableEntity,
			Message:   errorutils.StatusErrorFillRequiredForm,
			ErrorCode: 1,
		}
		return c.JSON(resp.Code, resp)
	}

	province, err := provincerepository.GetById(db.DBCon, pID)

	if province == nil {
		resp := response.ResponseError{
			Code:      http.StatusNotFound,
			Message:   fmt.Sprintf(errorutils.StatusErrorResourceNotExists, "Province"),
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
		Message: fmt.Sprintf(errorutils.StatusErrorSuccessfullyRetrieved, "Province"),
		Data:    province,
	}
	return c.JSON(http.StatusOK, resp)
}
