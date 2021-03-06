package route

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nauhalf/mobilcantik/controller/ad"
	"github.com/nauhalf/mobilcantik/controller/adbill"
	"github.com/nauhalf/mobilcantik/controller/adclosingreason"
	"github.com/nauhalf/mobilcantik/controller/adconfirmation"
	"github.com/nauhalf/mobilcantik/controller/adreport"
	"github.com/nauhalf/mobilcantik/controller/adstatustype"
	"github.com/nauhalf/mobilcantik/controller/adtype"
	"github.com/nauhalf/mobilcantik/controller/auth"
	"github.com/nauhalf/mobilcantik/controller/bankaccount"
	"github.com/nauhalf/mobilcantik/controller/city"
	"github.com/nauhalf/mobilcantik/controller/condition"
	"github.com/nauhalf/mobilcantik/controller/facility"
	"github.com/nauhalf/mobilcantik/controller/fuel"
	"github.com/nauhalf/mobilcantik/controller/image"
	"github.com/nauhalf/mobilcantik/controller/partner"
	"github.com/nauhalf/mobilcantik/controller/province"
	"github.com/nauhalf/mobilcantik/controller/reportgroup"
	"github.com/nauhalf/mobilcantik/controller/reporttype"
	"github.com/nauhalf/mobilcantik/controller/testimonial"
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

	e.GET("/images/*", image.StreamImage)

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

		// --------------------------------------------------------
		// GROUP Bank Account
		// --------------------------------------------------------

		bankaccounts := v1.Group("/bankaccounts", custommiddleware.JWT())
		{
			bankaccounts.GET("", bankaccount.GetAll)
			bankaccounts.POST("", bankaccount.Create, custommiddleware.JWTUserCredential)
			bankaccounts.PATCH("", bankaccount.Update, custommiddleware.JWTUserCredential)
			bankaccounts.DELETE("/:id", bankaccount.Delete, custommiddleware.JWTUserCredential)
		}

		v1.GET("/bankaccount/:id", bankaccount.GetById, custommiddleware.JWT())

		// --------------------------------------------------------
		// GROUP Ad Type
		// --------------------------------------------------------

		adtypes := v1.Group("/adtypes", custommiddleware.JWT())
		{
			adtypes.GET("", adtype.GetAll)
			adtypes.POST("", adtype.Create, custommiddleware.JWTUserCredential)
			adtypes.PATCH("", adtype.Update, custommiddleware.JWTUserCredential)
			adtypes.DELETE("/:id", adtype.Delete, custommiddleware.JWTUserCredential)
		}

		v1.GET("/adtype/:id", adtype.GetById, custommiddleware.JWT())

		// --------------------------------------------------------
		// GROUP Ad Status Type
		// --------------------------------------------------------

		adstatustypes := v1.Group("/adstatustypes", custommiddleware.JWT())
		{
			adstatustypes.GET("", adstatustype.GetAll)
			adstatustypes.POST("", adstatustype.Create, custommiddleware.JWTUserCredential)
			adstatustypes.PATCH("", adstatustype.Update, custommiddleware.JWTUserCredential)
			adstatustypes.DELETE("/:id", adstatustype.Delete, custommiddleware.JWTUserCredential)
		}

		v1.GET("/adstatustype/:id", adstatustype.GetById, custommiddleware.JWT())

		// --------------------------------------------------------
		// GROUP Ad Closing Reason
		// --------------------------------------------------------

		adclosingreasons := v1.Group("/adclosingreasons", custommiddleware.JWT())
		{
			adclosingreasons.GET("", adclosingreason.GetAll)
			adclosingreasons.POST("", adclosingreason.Create, custommiddleware.JWTUserCredential)
			adclosingreasons.PATCH("", adclosingreason.Update, custommiddleware.JWTUserCredential)
			adclosingreasons.DELETE("/:id", adclosingreason.Delete, custommiddleware.JWTUserCredential)
		}

		v1.GET("/adclosingreason/:id", adclosingreason.GetById, custommiddleware.JWT())

		// --------------------------------------------------------
		// GROUP Partner
		// --------------------------------------------------------

		partners := v1.Group("/partners", custommiddleware.JWT())
		{
			partners.GET("", partner.GetAll)
			partners.POST("", partner.Create, custommiddleware.JWTUserCredential)
			partners.PATCH("", partner.Update, custommiddleware.JWTUserCredential)
			partners.DELETE("/:id", partner.Delete, custommiddleware.JWTUserCredential)
		}

		v1.GET("/partner/:id", partner.GetById, custommiddleware.JWT())

		// --------------------------------------------------------
		// GROUP Report Group
		// --------------------------------------------------------

		reportgroups := v1.Group("/reportgroups", custommiddleware.JWT())
		{
			reportgroups.GET("", reportgroup.GetAll)
			reportgroups.POST("", reportgroup.Create, custommiddleware.JWTUserCredential)
			reportgroups.PATCH("", reportgroup.Update, custommiddleware.JWTUserCredential)
			reportgroups.DELETE("/:id", reportgroup.Delete, custommiddleware.JWTUserCredential)
		}

		v1.GET("/reportgroup/:id", reportgroup.GetById, custommiddleware.JWT())

		// --------------------------------------------------------
		// GROUP Report Type
		// --------------------------------------------------------

		reporttypes := v1.Group("/reporttypes", custommiddleware.JWT())
		{
			reporttypes.GET("", reporttype.GetAll)
			reporttypes.POST("", reporttype.Create, custommiddleware.JWTUserCredential)
			reporttypes.PATCH("", reporttype.Update, custommiddleware.JWTUserCredential)
			reporttypes.DELETE("/:id", reporttype.Delete, custommiddleware.JWTUserCredential)
		}

		v1.GET("/reporttype/:id", reporttype.GetById, custommiddleware.JWT())

		// --------------------------------------------------------
		// GROUP Ad
		// --------------------------------------------------------

		ads := v1.Group("/ads", custommiddleware.JWT())
		{
			ads.GET("", ad.GetAll)
			ads.POST("", ad.Create)
			ads.POST("/close", testimonial.Create)
			ads.DELETE("/:id/:password", ad.Delete)
			ads.DELETE("/forcedelete/:id", ad.ForceDelete, custommiddleware.JWTUserCredential)
		}

		v1.GET("/ad/:id", ad.GetById, custommiddleware.JWT())
		v1.POST("/ad-verification", ad.Verification)

		// --------------------------------------------------------
		// GROUP Bill
		// --------------------------------------------------------

		bills := v1.Group("/bills", custommiddleware.JWT())
		{
			bills.GET("", adbill.GetAll)
			bills.POST("/confirmation-payment", adconfirmation.Create)
			bills.POST("/verify-payment", adbill.ConfirmPaymentBill, custommiddleware.JWTUserCredential)

		}

		v1.GET("/bill/:id", adbill.GetById, custommiddleware.JWT(), custommiddleware.JWTUserCredential)

		// --------------------------------------------------------
		// GROUP Ad Report
		// --------------------------------------------------------

		adreports := v1.Group("/adreports", custommiddleware.JWT())
		{
			adreports.GET("", adreport.GetAll, custommiddleware.JWTUserCredential)
			adreports.POST("", adreport.Create)
			adreports.DELETE("/:id", adreport.Delete, custommiddleware.JWTUserCredential)

		}

		v1.GET("/adreport/:id", adreport.GetById, custommiddleware.JWT(), custommiddleware.JWTUserCredential)

		// --------------------------------------------------------
		// GROUP Testimonial
		// --------------------------------------------------------

		testimonials := v1.Group("/testimonials", custommiddleware.JWT())
		{
			testimonials.GET("", testimonial.GetAll)
		}

		v1.GET("/testimonial/:id", testimonial.GetById, custommiddleware.JWT())

	}
}
