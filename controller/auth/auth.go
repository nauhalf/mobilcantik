package auth

import (
	"log"
	"net/http"
	"os"

	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/nauhalf/mobilcantik/db"
	"github.com/nauhalf/mobilcantik/repository/apikeyrepository"
	"github.com/nauhalf/mobilcantik/repository/userrepository"
	"github.com/nauhalf/mobilcantik/response"
	errorutils "github.com/nauhalf/mobilcantik/utils/error"
	"golang.org/x/crypto/bcrypt"
)

func GenerateToken(c echo.Context) (err error) {

	szApiKey := c.FormValue("szApiKey")
	szUsername := c.FormValue("szUsername")
	szPassword := c.FormValue("szPassword")

	if szApiKey == "" {
		resp := response.ResponseError{
			Code:      http.StatusUnprocessableEntity,
			Message:   errorutils.StatusErrorFillRequiredForm,
			ErrorCode: 1,
		}
		return c.JSON(http.StatusBadRequest, resp)
	}

	if err != nil {
		resp := response.ResponseError{
			Code:      http.StatusInternalServerError,
			Message:   err.Error(),
			ErrorCode: nil,
		}
		return c.JSON(http.StatusInternalServerError, resp)
	}

	apiKey, err := apikeyrepository.GetApiKeyById(db.DBCon, szApiKey)
	if err != nil {
		resp := response.ResponseError{
			Code:      http.StatusInternalServerError,
			Message:   err.Error(),
			ErrorCode: nil,
		}
		return c.JSON(http.StatusInternalServerError, resp)
	}
	if apiKey == "" {
		resp := response.ResponseError{
			Code:      http.StatusUnauthorized,
			Message:   http.StatusText(http.StatusUnauthorized),
			ErrorCode: 1,
		}

		return c.JSON(http.StatusUnauthorized, resp)
	}

	if szUsername == "" {
		token, err := createJwt(apiKey, "")
		if err != nil {
			resp := response.ResponseError{
				Code:      http.StatusInternalServerError,
				Message:   err.Error(),
				ErrorCode: nil,
			}
			return c.JSON(http.StatusInternalServerError, resp)
		}

		resp := response.ResponseSuccess{
			Code:    http.StatusOK,
			Message: "Authorized",
			Data:    token,
		}

		return c.JSON(http.StatusOK, resp)
	}

	user, err := userrepository.GetUserByUsername(db.DBCon, szUsername)

	if user == nil {
		resp := response.ResponseError{
			Code:      http.StatusUnauthorized,
			Message:   http.StatusText(http.StatusUnauthorized),
			ErrorCode: 2,
		}

		return c.JSON(http.StatusUnauthorized, resp)
	}

	if compare := comparePasswords(user.Password, []byte(szPassword)); !compare {
		resp := response.ResponseError{
			Code:      http.StatusUnauthorized,
			Message:   http.StatusText(http.StatusUnauthorized),
			ErrorCode: 2,
		}

		return c.JSON(http.StatusUnauthorized, resp)
	}

	token, err := createJwt(apiKey, user.Username)
	if err != nil {
		resp := response.ResponseError{
			Code:      http.StatusInternalServerError,
			Message:   err.Error(),
			ErrorCode: 2,
		}
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp := response.ResponseSuccess{
		Code:    http.StatusOK,
		Message: "Authorized",
		Data:    token,
	}

	return c.JSON(http.StatusOK, resp)
}

func createJwt(apiKey string, username string) (map[string]string, error) {
	//create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set Claims
	// This is the information which frontend can use
	// The backend can also decode the token and get admin etc.

	claims := token.Claims.(jwt.MapClaims)
	claims["key_id"] = apiKey
	if username == "" {
		claims["username"] = nil
	} else {
		claims["username"] = username
	}
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	t, err := token.SignedString([]byte(os.Getenv("API_SECRET")))

	if err != nil {
		return nil, err
	}

	return map[string]string{
		"token": t,
	}, nil
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
