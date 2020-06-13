package testimonial

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/nauhalf/mobilcantik/db"
	"github.com/nauhalf/mobilcantik/repository/adrepository"
	"github.com/nauhalf/mobilcantik/repository/model"
	"github.com/nauhalf/mobilcantik/repository/testimonialrepository"
	"github.com/nauhalf/mobilcantik/response"
	errorutils "github.com/nauhalf/mobilcantik/utils/error"
)

type RequestCreate struct {
	AdId              uint64 `json:"intAdId" form:"intAdId" validate:"required"`
	Comment           string `json:"szComment" form:"szComment" validate:"required"`
	Password          string `json:"szPassword" form:"szPassword" validate:"required"`
	AdClosingReasonId uint64 `json:"intAdClosingReasonId" form:"intAdClosingReasonId"`
}

func GetAll(c echo.Context) error {

	testimonials, err := testimonialrepository.GetAll(db.DBCon)

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
		Message: fmt.Sprintf(errorutils.StatusErrorSuccessfullyRetrieved, "List Testimonials"),
		Data:    testimonials,
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

	ad, pw, _ := adrepository.GetAdById2(db.DBCon, r.AdId)

	if ad == nil {
		resp := response.ResponseError{
			Code:      http.StatusNotFound,
			Message:   fmt.Sprintf(errorutils.StatusErrorResourceNotExists, "Ad"),
			ErrorCode: 1,
		}
		return c.JSON(resp.Code, resp)
	}

	if r.Password != *pw {
		resp := response.ResponseError{
			Code:      http.StatusUnprocessableEntity,
			Message:   errorutils.StatusErrorIncorrectPassword,
			ErrorCode: 2,
		}
		return c.JSON(resp.Code, resp)
	}

	testimonialExists, _ := testimonialrepository.GetByAdId(db.DBCon, r.AdId)

	if testimonialExists != nil {
		resp := response.ResponseError{
			Code:      http.StatusUnprocessableEntity,
			Message:   "Testimonial already exists",
			ErrorCode: 3,
		}
		return c.JSON(resp.Code, resp)
	}

	testimonial := new(model.Testimonial)
	testimonial.AdId = r.AdId
	testimonial.Comment = r.Comment
	testimonial.AdClosingReasonId = r.AdClosingReasonId
	testimonial.IsActive = ad.IsActive

	newTesti, err := testimonialrepository.Create(db.DBCon, testimonial)
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
		Message: fmt.Sprintf(errorutils.StatusErrorSuccessfullyCreated, "Testimonial"),
		Data:    newTesti,
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

	testi, err := testimonialrepository.GetById(db.DBCon, intID)

	if testi == nil {
		resp := response.ResponseError{
			Code:      http.StatusNotFound,
			Message:   fmt.Sprintf(errorutils.StatusErrorResourceNotExists, "Testimonial"),
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
		Message: fmt.Sprintf(errorutils.StatusErrorSuccessfullyRetrieved, "Testimonial"),
		Data:    testi,
	}
	return c.JSON(http.StatusOK, resp)
}
