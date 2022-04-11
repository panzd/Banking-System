package domain

import (
	"database/sql"
	"log"
	"time"

	"github.com/Banking-System/errs"
	"github.com/Banking-System/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	// var rows *sql.Rows
	var err error
	customers := make([]Customer, 0)
	// 每次查询都会创建一个连接池
	if status == ""{
	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
	err := d.client.Select(&customers, findAllSql)
	// rows, err := d.client.Query(findAllSql) //查询？
	}else{
	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status=="?""
	err := d.client.Select(&customers, findAllSql, status)
	}


	if err != nil {
		logger.Error("Error while querying customer table" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	
	// err = sqlx.StructScan(rows, &customers)
	// if err != nil{
	// 	logger.Error("Error while scanning customers" + err.Error())
	// 	return nil, errs.NewUnexpectedError("Unexpected database error")
	// }

	// for rows.Next() {
	// 	var c Customer
	// 	err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
	// 	if err != nil {
	// 		logger.Error("Error while scanning customers" + err.Error())
	// 		return nil, errs.NewUnexpectedError("Unexpected database error")
	// 	}
	// 	customers = append(customers, c)
	// }
	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer,*errs.AppError){
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer-id ?"

	row := d.client.QueryRow(customerSql, id)

	var c Customer

	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while scanning customer" + err.Error())
			// log.Println("Error while scanning customer" + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected dataabase")
		}
	}
	return &c, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sqlx.Open("mysql", "root:pan@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client}
}
