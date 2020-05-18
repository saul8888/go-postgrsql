package authentication

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/labstack/echo"
	"github.com/saul8888/postgrsql/postgrsql/dbpostgrsql"
)

type Service interface {
	GenerateCustomer(context echo.Context) error
	ValidateToken(context echo.Context) error
}

type service struct {
	repo dbpostgrsql.Postgrsql
}

func NewService(repo dbpostgrsql.Postgrsql) Service {
	return &service{repo: repo}
}

func (s *service) GenerateCustomer(c echo.Context) (err error) {
	jsonResult := new(Responsetoken)
	var answer = http.StatusOK
	date := new(dbpostgrsql.DateValidate)
	if err = c.Bind(date); err != nil {
		answer = http.StatusForbidden
		return c.JSON(answer, err)
	}
	autho, err := s.repo.Search(date)
	if err != nil {
		answer = http.StatusForbidden
		return c.JSON(answer, err)
	}

	customer := &Customer{}
	customer.Name = autho.Name
	customer.Email = autho.Email

	if customer.Name != "" && customer.Email != "" {
		//create a struct of my Claim
		claims := Claim{
			Customer: *customer,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
				Issuer:    "token test", //object of token
			},
		}

		//--------------------encode to base64-----------------//
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokeS, err := token.SignedString(Keys())
		//token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
		//tokeS, err := token.SignedString(PrivateKey)
		if err != nil {
			tokeS = "could not sign private token"
		}
		answer = http.StatusOK
		jsonResult.Token = tokeS
	} else {
		answer = http.StatusForbidden
		jsonResult.Token = "usser or password invalid"
	}
	return c.JSON(answer, jsonResult)
}

func (s *service) ValidateToken(context echo.Context) error {
	jsonResult := new(Responsetoken)
	var answer = http.StatusOK
	token, err := request.ParseFromRequestWithClaims(context.Request(), request.OAuth2Extractor, &Claim{},
		func(token *jwt.Token) (interface{}, error) {
			return PublicKey, nil
		})

	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)
			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				answer = http.StatusUnauthorized
				jsonResult.Token = "your token expired"
			case jwt.ValidationErrorSignatureInvalid:
				answer = http.StatusUnauthorized
				jsonResult.Token = "the signature does not match"
			default:
				answer = http.StatusUnauthorized
				jsonResult.Token = "the signature does not match"
			}
		default:
			answer = http.StatusUnauthorized
			jsonResult.Token = "your token is not valid"
		}
	}
	if token.Valid {
		answer = http.StatusAccepted
		jsonResult.Token = "welcome to the system"
	} else {
		answer = http.StatusUnauthorized
		jsonResult.Token = "your token is not valid"
	}
	return context.JSON(answer, jsonResult)
}
