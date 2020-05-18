package middelware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/middleware"
	"github.com/saul8888/postgrsql/postgrsql/authentication"
)

//authentication
var autoConfig = middleware.JWTConfig{
	Claims:        &authentication.Claim{},
	SigningMethod: jwt.SigningMethodHS256.Name,
	SigningKey:    authentication.Keys(),
	//SigningMethod: jwt.SigningMethodRS256.Name,
	//SigningKey:    authentication.PublicKey,
	//TokenLookup: "header:" + echo.HeaderAuthorization,
}

//Bearer {token}
