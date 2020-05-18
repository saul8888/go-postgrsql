package dbpostgrsql

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	//psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	//	"password=%s dbname=%s sslmode=disable",
	//	host, port, user, password, dbname)
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("HOST_DB"), os.Getenv("PORT_DB"), user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		fmt.Println("-------error1234-----------------")
		fmt.Println(err)

	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		fmt.Println("1234")
	}
	fmt.Println(os.Getenv("LOCAL_HOST"))
	return db
}

type Postgrsql interface {
	GetById(id string) (*Employee, error)
	GetTotal(params *GetEmployeesRequest) ([]*Employee, error)
	GetCount() (int, error)
	Insert(params *Employee) (string, error)
	Update(params *UpdateEmployeeRequest) (*Employee, error)
	DeleteById(id string) (*Employee, error)
	Search(params *DateValidate) (*Employee, error)
}

type postgrsql struct {
	db *sql.DB
}

func NewDataBase(dbconection *sql.DB) Postgrsql {
	return &postgrsql{db: dbconection}
}

func (repo postgrsql) GetById(id string) (*Employee, error) {
	rows, err := repo.db.Query(`SELECT * FROM persons WHERE ID = $1`, id)
	employee := &Employee{}
	for rows.Next() {
		err = rows.Scan(&employee.ID, &employee.Name, &employee.Email,
			&employee.Password, &employee.CreatedAt, &employee.UpdateAt)
	}
	return employee, err
}

func (repo postgrsql) GetTotal(params *GetEmployeesRequest) ([]*Employee, error) {
	rows, err := repo.db.Query("SELECT * FROM persons LIMIT $1 OFFSET $2", params.Limit, params.Offset)

	var employees []*Employee
	for rows.Next() {
		employee := &Employee{}
		err = rows.Scan(&employee.ID, &employee.Name, &employee.Email,
			&employee.Password, &employee.CreatedAt, &employee.UpdateAt)
		employees = append(employees, employee)
	}
	return employees, err

}

func (repo postgrsql) GetCount() (int, error) {
	var total int
	rows := repo.db.QueryRow("SELECT Count(*) FROM persons")
	err := rows.Scan(&total)
	return total, err
}

func (repo postgrsql) Insert(params *Employee) (string, error) {
	const sql = `INSERT INTO persons(id,firstname,email,pass,created_at,update_at) VALUES($1,$2,$3,$4,$5,$6)`
	Newid := RandomString(10)
	_, err := repo.db.Exec(sql, Newid, params.Name, params.Email,
		params.Password, time.Now(), time.Now())
	if err != nil {
		return "err", err
	}
	//id, _ := result.LastInsertId()
	return Newid, err
}

func (repo postgrsql) Update(params *UpdateEmployeeRequest) (*Employee, error) {
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

func (repo postgrsql) DeleteById(id string) (*Employee, error) {
	row, _ := repo.GetById(id)
	result, err := repo.db.Exec(`DELETE FROM persons WHERE ID = $1`, id)
	count, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	fmt.Println(count)
	return row, err
}

func (repo postgrsql) Search(params *DateValidate) (*Employee, error) {
	rows, err := repo.db.Query(`SELECT * FROM persons WHERE email = $1 AND pass = $2 `, params.Email, params.Password)
	if err != nil {
		return nil, err
	}

	employee := &Employee{}
	for rows.Next() {
		err = rows.Scan(&employee.ID, &employee.Name, &employee.Email,
			&employee.Password, &employee.CreatedAt, &employee.UpdateAt)
	}
	return employee, err
}
