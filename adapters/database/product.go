package database

import (
	"database/sql"
	"fmt"

	"github.com/johnldev/go-hexagonal/app"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db: db}
}

func (p ProductDb) Get(id string) (app.ProductInterface, error) {
	var product app.Product

	stmt, err := p.db.Prepare("select id, name, status, price from products where id = ?")

	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Status, &product.Price)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p ProductDb) create(product app.ProductInterface) (app.ProductInterface, error) {
	stmt, err := p.db.Prepare("insert into products(id, name, status, price) values(?, ?, ?, ?)")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(product.GetID(), product.GetName(), product.GetStatus(), product.GetPrice())
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p ProductDb) update(product app.ProductInterface) (app.ProductInterface, error) {
	stmt, err := p.db.Prepare("update products set name = ?, status = ?, price = ? where id = ?")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(product.GetName(), product.GetStatus(), product.GetPrice(), product.GetID())
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p ProductDb) Save(product app.ProductInterface) (app.ProductInterface, error) {
	smt, err := p.db.Prepare("select id from products where id = ?")
	if err != nil {
		return nil, err
	}

	defer smt.Close()
	var rows string
	smt.QueryRow(product.GetID()).Scan(&rows)
	fmt.Printf("rows: %s\n", rows)
	if rows == "" {
		return p.create(product)
	} else {
		return p.update(product)
	}

}
