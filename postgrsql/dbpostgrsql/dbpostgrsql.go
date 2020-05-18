package dbpostgrsql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	//psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	//	"password=%s dbname=%s sslmode=disable",
	//	os.Getenv("LOCAL_HOST"), port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		fmt.Println("-------error1234-----------------")
		fmt.Println(err)
		//panic(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		//panic(err)
	}
	return db
}

type Postgrsql interface {
	GetById(id string) (*Customer, error)
	GetTotal(params *GetCustomersRequest) ([]*Customer, error)
	GetCount() (int, error)
	Insert(params *Customer) (string, error)
	Update(params *UpdateCustomerRequest) (*Customer, error)
	DeleteById(id string) (*Customer, error)
	Search(params *DateValidate) (*Customer, error)
}

type postgrsql struct {
	db *sql.DB
}

func NewDataBase(dbconection *sql.DB) Postgrsql {
	return &postgrsql{db: dbconection}
}

func (repo postgrsql) GetById(id string) (*Customer, error) {
	rows, err := repo.db.Query(`SELECT * FROM persons WHERE ID = $1`, id)
	//customer := &employee.Customer{}
	customer := &Customer{}
	for rows.Next() {
		err = rows.Scan(&customer.ID, &customer.Name, &customer.Email,
			&customer.Password, &customer.CreatedAt, &customer.UpdateAt)
	}
	return customer, err
}

func (repo postgrsql) GetTotal(params *GetCustomersRequest) ([]*Customer, error) {
	rows, err := repo.db.Query("SELECT * FROM persons LIMIT $1 OFFSET $2", params.Limit, params.Offset)

	var customers []*Customer
	for rows.Next() {
		customer := &Customer{}
		err = rows.Scan(&customer.ID, &customer.Name, &customer.Email,
			&customer.Password, &customer.CreatedAt, &customer.UpdateAt)
		customers = append(customers, customer)
	}
	return customers, err

}

func (repo postgrsql) GetCount() (int, error) {
	var total int
	rows := repo.db.QueryRow("SELECT Count(*) FROM persons")
	err := rows.Scan(&total)
	return total, err
}

func (repo postgrsql) Insert(params *Customer) (string, error) {
	const sql = `INSERT INTO persons(id,firstname,email,pass,created_at,update_at) VALUES($1,$2,$3,$4,$5,$6)`
	//result, err := repo.db.Exec(sql, "1235", "akemi", "akemi@example.com", "1234", time.Now(), time.Now())
	Newid := RandomString(10)
	_, err := repo.db.Exec(sql, Newid, params.Name, params.Email,
		params.Password, time.Now(), time.Now())
	if err != nil {
		return "err", err
	}
	//id, _ := result.LastInsertId()
	return Newid, err
}

func (repo postgrsql) Update(params *UpdateCustomerRequest) (*Customer, error) {
	const sql = `
			UPDATE persons SET
    		firstname= $2,
			email= $3,
			pass= $4,
			update_at= $5
			WHERE ID = $1`
	result, err := repo.db.Exec(sql, params.ID, params.Name, params.Email, params.Password, time.Now())
	if err != nil {
		return nil, err
	}
	idi, _ := result.LastInsertId()
	fmt.Println(idi)
	row, _ := repo.GetById(params.ID)
	return row, nil
}

func (repo postgrsql) DeleteById(id string) (*Customer, error) {
	row, _ := repo.GetById(id)
	result, err := repo.db.Exec(`DELETE FROM persons WHERE ID = $1`, id)
	count, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	fmt.Println(count)
	return row, err
}

func (repo postgrsql) Search(params *DateValidate) (*Customer, error) {
	rows, err := repo.db.Query(`SELECT * FROM persons WHERE email = $1 AND pass = $2 `, params.Email, params.Password)
	//rows, err := repo.db.Query(`SELECT * FROM persons WHERE email = $1 AND password = $2 `, params.Email, params.Password)
	if err != nil {
		return nil, err
	}
	//customer := &employee.Customer{}
	customer := &Customer{}
	for rows.Next() {
		err = rows.Scan(&customer.ID, &customer.Name, &customer.Email,
			&customer.Password, &customer.CreatedAt, &customer.UpdateAt)
	}
	return customer, err
}
