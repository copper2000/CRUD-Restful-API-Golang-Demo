package entities

import (
	"../models"
	"database/sql"
)

type ProductModel struct {
	Db *sql.DB
}

func (productModel ProductModel) FindAll() (product []models.Product, err error) {
	rows, err := productModel.Db.Query("SELECT * FROM product")

	if err != nil {
		return nil, err
	} else {
		var products []models.Product

		// map data from database to model
		for rows.Next() {
			var id int64
			var name string
			var price float64
			var quantity int64

			// assign value to model
			err2 := rows.Scan(&id, &name, &price, &quantity)

			if err2 != nil {
				return nil, err2
			} else {
				product := models.Product{
					Id:       id,
					Name:     name,
					Price:    price,
					Quantity: quantity,
				}
				products = append(products, product)
			}
		}
		return products, nil
	}
}

func (productModel ProductModel) FindAllWithPaging(pageIndex int, pageSize int) (product []models.Product, err error) { // ()th pass to f() | ()nd receive by f()
	rows, err := productModel.Db.Query("SELECT * FROM product LIMIT ? OFFSET ?", pageSize, pageIndex-1)

	if err != nil {
		return nil, err
	} else {
		var products []models.Product

		// map data from database to model
		for rows.Next() {
			var id int64
			var name string
			var price float64
			var quantity int64

			// assign value to model
			err2 := rows.Scan(&id, &name, &price, &quantity)

			if err2 != nil {
				return nil, err2
			} else {
				product := models.Product{
					Id:       id,
					Name:     name,
					Price:    price,
					Quantity: quantity,
				}
				products = append(products, product)
			}
		}
		return products, nil
	}
}

func (productModel ProductModel) SearchByName(keyword string) (product []models.Product, err error) {
	rows, err := productModel.Db.Query("SELECT * FROM product WHERE name LIKE ?", "%"+keyword+"%")

	if err != nil {
		return nil, err
	} else {
		var products []models.Product

		// map data from database to model
		for rows.Next() {
			var id int64
			var name string
			var price float64
			var quantity int64

			// assign value to model
			err2 := rows.Scan(&id, &name, &price, &quantity)

			if err2 != nil {
				return nil, err2
			} else {
				product := models.Product{
					Id:       id,
					Name:     name,
					Price:    price,
					Quantity: quantity,
				}
				products = append(products, product)
			}
		}
		return products, nil
	}
}

func (productModel ProductModel) SearchByPricesRange(min, max float64) (product []models.Product, err error) {
	rows, err := productModel.Db.Query("SELECT * FROM product WHERE price >= ? AND price <= ?", min, max)

	if err != nil {
		return nil, err
	} else {
		var products []models.Product

		// map data from database to model
		for rows.Next() {
			var id int64
			var name string
			var price float64
			var quantity int64

			// assign value to model
			err2 := rows.Scan(&id, &name, &price, &quantity)

			if err2 != nil {
				return nil, err2
			} else {
				product := models.Product{
					Id:       id,
					Name:     name,
					Price:    price,
					Quantity: quantity,
				}
				products = append(products, product)
			}
		}
		return products, nil
	}
}

func (productModel ProductModel) Create(product *models.Product) (err error) {

	result, err := productModel.Db.Exec("INSERT INTO product (name, price, quantity) VALUES (?, ?, ?)",
		product.Name, product.Price, product.Quantity)

	if err != nil {
		return err
	} else {
		product.Id, _ = result.LastInsertId()
		return nil
	}
}

func (productModel ProductModel) Update(product *models.Product) (err error) {

	result, err := productModel.Db.Exec("UPDATE product SET name = ?, price = ?, quantity = ? WHERE id = ?",
		product.Name, product.Price, product.Quantity, product.Id)

	if err != nil {
		return err
	} else {
		_, _ = result.RowsAffected()
		return nil
	}
}

func (productModel ProductModel) Delete(product *models.Product) (err error) {

	result, err := productModel.Db.Exec("DELETE FROM product WHERE id = ?", product.Id)

	if err != nil {
		return err
	} else {
		_, _ = result.RowsAffected()
		return nil
	}
}
