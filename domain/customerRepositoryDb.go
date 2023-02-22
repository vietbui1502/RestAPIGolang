package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	db *sql.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {
	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"

	rows, err := d.db.Query(findAllSql)

	if err != nil {
		log.Println("Error when quering customer table" + err.Error())
		return nil, err
	}

	customers := make([]Customer, 0)

	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateofBirth, &c.Status)
		if err != nil {
			log.Println("Error when scaning row" + err.Error())
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func (d CustomerRepositoryDb) FindCustomerbyID(id string) (*Customer, error) {
	findSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	rows, err := d.db.Query(findSql, id)

	if err != nil {
		log.Println("Error when quering customer table" + err.Error())
		return nil, err
	}

	if rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateofBirth, &c.Status)
		if err != nil {
			log.Println("Error when scaning row" + err.Error())
			return nil, err
		}
		return &c, nil
	}

	return nil, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "root:codecamp@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{db: client}
}
