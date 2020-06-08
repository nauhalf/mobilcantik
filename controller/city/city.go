package city

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nauhalf/mobilcantik/db"
	"github.com/nauhalf/mobilcantik/repository/cityrepository"
	"github.com/nauhalf/mobilcantik/repository/provincerepository"
	"github.com/nauhalf/mobilcantik/response"
)

func GetAll(c echo.Context) error {

	province_id := c.QueryParam("province_id")

	if province_id == "" {
		resp := response.ResponseError{
			Code:      http.StatusNotFound,
			Message:   http.StatusText(http.StatusNotFound),
			ErrorCode: 1,
		}

		return c.JSON(resp.Code, resp)
	}

	pExists, _ := provincerepository.GetById(db.DBCon, province_id)

	if pExists == nil {
		resp := response.ResponseError{
			Code:      http.StatusNotFound,
			Message:   http.StatusText(http.StatusNotFound),
			ErrorCode: 1,
		}
		return c.JSON(resp.Code, resp)
	}

	cities, err := cityrepository.GetAll(db.DBCon, province_id)

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
		Message: "List city successfully retrieved",
		Data:    cities,
	}
	return c.JSON(http.StatusOK, resp)
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

	city, err := cityrepository.GetById(db.DBCon, cID)

	if city == nil {
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
		Message: "City successfully retrieved",
		Data:    city,
	}
	return c.JSON(http.StatusOK, resp)
}
