package route

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nauhalf/mobilcantik/controller/auth"
	"github.com/nauhalf/mobilcantik/controller/bankaccount"
	"github.com/nauhalf/mobilcantik/controller/city"
	"github.com/nauhalf/mobilcantik/controller/condition"
	"github.com/nauhalf/mobilcantik/controller/facility"
	"github.com/nauhalf/mobilcantik/controller/fuel"
	"github.com/nauhalf/mobilcantik/controller/province"
	"github.com/nauhalf/mobilcantik/controller/vehiclebrand"
	"github.com/nauhalf/mobilcantik/controller/vehicletype"
	custommiddleware "github.com/nauhalf/mobilcantik/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func InitRoutes(e *echo.Echo) {

	e.Validator = &CustomValidator{validator: validator.New()}

	v1 := e.Group("/v1")
	{
		v1.POST("/request-token", auth.GenerateToken)

		middleware.ErrJWTMissing = echo.NewHTTPError(http.StatusUnauthorized, "invalid or expired jwt")
		// --------------------------------------------------------
		// GROUP Vehicle Type
		// --------------------------------------------------------
		vehicletypes := v1.Group("/vehicletypes", custommiddleware.JWT())
		{
			vehicletypes.GET("", vehicletype.GetAll)
			vehicletypes.POST("", vehicletype.Create, custommiddleware.JWTUserCredential)
			vehicletypes.PATCH("", vehicletype.Update, custommiddleware.JWTUserCredential)
			vehicletypes.DELETE("/:id", vehicletype.Delete, custommiddleware.JWTUserCredential)
		}

		v1.GET("/vehicletype/:id", vehicletype.GetById, custommiddleware.JWT())

		// --------------------------------------------------------
		// GROUP Vehicle Brand
		// --------------------------------------------------------

		vehiclebrands := v1.Group("/vehiclebrands", custommiddleware.JWT())
		{
			vehiclebrands.GET("", vehiclebrand.GetAll)
			vehiclebrands.POST("", vehiclebrand.Create, custommiddleware.JWTUserCredential)
			vehiclebrands.PATCH("", vehiclebrand.Update, custommiddleware.JWTUserCredential)
			vehiclebrands.DELETE("/:id", vehiclebrand.Delete, custommiddleware.JWTUserCredential)
		}

		v1.GET("/vehiclebrand/:id", vehiclebrand.GetById, custommiddleware.JWT())

		// --------------------------------------------------------
		// GROUP Fuel
		// --------------------------------------------------------

		fuels := v1.Group("/fuels", custommiddleware.JWT())
		{
			fuels.GET("", fuel.GetAll)
			fuels.POST("", fuel.Create, custommiddleware.JWTUserCredential)
			fuels.PATCH("", fuel.Update, custommiddleware.JWTUserCredential)
			fuels.DELETE("/:id", fuel.Delete, custommiddleware.JWTUserCredential)
		}

		v1.GET("/fuel/:id", fuel.GetById, custommiddleware.JWT())

		fuel_mapping := v1.Group("/fuel_mapping", custommiddleware.JWT())
		{
			fuel_mapping.GET("", fuel.GetMapping)
			fuel_mapping.POST("", fuel.DoMapping, custommiddleware.JWTUserCredential)
		}

		// --------------------------------------------------------
		// GROUP Condition
		// --------------------------------------------------------

		conditions := v1.Group("/conditions", custommiddleware.JWT())
		{
			conditions.GET("", condition.GetAll)
			conditions.POST("", condition.Create, custommiddleware.JWTUserCredential)
			conditions.PATCH("", condition.Update, custommiddleware.JWTUserCredential)
			conditions.DELETE("/:id", condition.Delete, custommiddleware.JWTUserCredential)
		}

		v1.GET("/condition/:id", condition.GetById, custommiddleware.JWT())

		// --------------------------------------------------------
		// GROUP Facility
		// --------------------------------------------------------

		facilities := v1.Group("/facilities", custommiddleware.JWT())
		{
			facilities.GET("", facility.GetAll)
			facilities.POST("", facility.Create, custommiddleware.JWTUserCredential)
			facilities.PATCH("", facility.Update, custommiddleware.JWTUserCredential)
			facilities.DELETE("/:id", facility.Delete, custommiddleware.JWTUserCredential)
		}

		v1.GET("/facility/:id", facility.GetById, custommiddleware.JWT())

		// --------------------------------------------------------
		// GROUP Province
		// --------------------------------------------------------

		provinces := v1.Group("/provinces", custommiddleware.JWT())
		{
			provinces.GET("", province.GetAll)
		}

		v1.GET("/province/:id", province.GetById, custommiddleware.JWT())

		// --------------------------------------------------------
		// GROUP City
		// --------------------------------------------------------

		cities := v1.Group("/cities", custommiddleware.JWT())
		{
			cities.GET("", city.GetAll)
		}

		v1.GET("/city/:id", city.GetById, custommiddleware.JWT())

		bankaccounts := v1.Group("/bankaccounts", custommiddleware.JWT())
		{
			bankaccounts.GET("", bankaccount.GetAll)
			bankaccounts.POST("", bankaccount.Create, custommiddleware.JWTUserCredential)
			bankaccounts.PATCH("", bankaccount.Update, custommiddleware.JWTUserCredential)
			bankaccounts.DELETE("/:id", bankaccount.Delete, custommiddleware.JWTUserCredential)
		}

		v1.GET("/bankaccount/:id", bankaccount.GetById, custommiddleware.JWT())

	}
}
