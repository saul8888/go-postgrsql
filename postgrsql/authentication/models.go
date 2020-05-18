package authentication

import (
	"crypto/rsa"
	"io/ioutil"
	"log"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
)

/*
var datevalidate map[string]interface{} = map[string]interface{}{
	"email":    "jose@example.com",
	"password": "jose123",
}
*/

//Data that will be in the payload
type Customer struct {
	Name  string `bson:"name"`
	Email string `json:"email"`
}

//The token body
type Claim struct {
	Customer `json:"user"`
	//standar claim
	jwt.StandardClaims
}

//Token response
type Responsetoken struct {
	Token string `json:"token"`
}

//using method RS256
func init() {
	//the read archive in format bytes for save  the keys private anad public
	privateBytes, err := ioutil.ReadFile("./private.rsa")
	if err != nil {
		log.Fatal("private key was not read")
	}

	publicBytes, err := ioutil.ReadFile("./public.rsa.pub")
	if err != nil {
		log.Fatal("public key was not read")
	}

	//for load in the form of a key private and public
	PrivateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		log.Fatal("could not do the parse of private")
	}
	PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		log.Fatal("could not do the parse of public")
	}
}

//using method HS256
func Keys() []byte {
	privateBytes, err := ioutil.ReadFile("./private.rsa")
	if err != nil {
		log.Fatal("private key was not read")
	}
	return privateBytes
}
