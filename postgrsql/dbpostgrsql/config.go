package dbpostgrsql

import (
	"math/rand"
	"time"
)

const (
	//host = "172.18.0.2"
	host = "localhost"
	//host     = os.Getenv("LOCAL_HOST")
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "testdb"
)

type GetCustomersRequest struct {
	Limit  int `json:"limit" form:"limit" query:"limit"`
	Offset int `json:"offset" form:"offset" query:"offset"`
}

type Customer struct {
	ID        string    `bson:"id" json:"id,omitempty"`
	Name      string    `bson:"name"`
	Email     string    `json:"email"`
	Password  string    `bson:"password"`
	CreatedAt time.Time `bson:"created_at"`
	UpdateAt  time.Time `bson:"update_at"`
}

type CustomerList struct {
	Data         []*Customer `json:"data"`
	TotalRecords int         `json:"totalRecords"`
}

type UpdateCustomerRequest struct {
	ID       string `json:"id" form:"id" query:"id"`
	Name     string `json:"name" form:"name" query:"name"`
	Email    string `json:"email" form:"email" query:"email"`
	Password string `json:"password" form:"password" query:"password"`
}

//Data that you will use to obtain the token
type DateValidate struct {
	Email    string `json:"email" form:"email" query:"email"`
	Password string `json:"password" form:"password" query:"password"`
}

// Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// Generate a random string of A-Z chars with len = l
func RandomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(65, 90))
	}
	return string(bytes)
}
