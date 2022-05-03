package products

import (
	"errors"
	"fmt"
	"time"

	"github.com/DeVasu/gotoken-api/datasources/mysql/cashiers_db"
	"github.com/federicoleon/bookstore_utils-go/logger"
	"github.com/federicoleon/bookstore_utils-go/rest_errors"
)



const (
	queryInsertProduct         = "INSERT INTO products(categoryId, name, image, price, stock, discountQty, discountType, discountResult, discountExpiredAt, createdAt, updatedAt) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
	queryListCategories 			= "SELECT * from products;"
	queryById = "SELECT * from products where id=?;"
	
)
// 	queryGetProduct = "SELECT * from products WHERE id=?;"
// 	queryUpdateProduct = "UPDATE products SET name = ? WHERE id=?;"
// 	queryDeleteCategory = "DELETE FROM categories WHERE id=?;"
// 	// queryGetCategories               = "SELECT Categoriesd, first_name, last_name, email, date_created, status FROM Categories WHERE id=?;"
// 	// queryFindByCategoriesd           = "SELECT Categoriesd, name, passcode FROM Categories WHERE Categoriesd=?;"

// 	// queryUpdateCategories            = "UPDATE Categories SET name=?, passcode=? WHERE Categoriesd=?;"
// 	// queryDeleteCategories            = "DELETE FROM Categories WHERE Categoriesd=?;"
// 	// // queryFindByStatus           = "SELECT id, first_name, last_name, email, date_created, status FROM Categories WHERE status=?;"
// 	// queryFindByEmailAndPassword = "SELECT id, first_name, last_name, email, date_created, status FROM Categories WHERE email=? AND password=? AND status=?"
// 	queryListCategories 			= "SELECT id, name from categories;"
// )

// func(category *Category) Delete() rest_errors.RestErr {
// 	stmt, err := cashiers_db.Client.Prepare(queryDeleteCategory)
// 	if err != nil {
// 		logger.Error("error when trying to prepare get cashier statement", err)
// 		return rest_errors.NewInternalServerError("error when tying to get cashier", errors.New("database error"))
// 	}
// 	defer stmt.Close()

// 	_, err = stmt.Exec(category.Id)
// 	if err != nil {
// 		logger.Error("error when trying to update user", err)
// 		return rest_errors.NewInternalServerError("error when tying to update user", errors.New("database error"))
// 	}
// 	return nil
// }

// func(category *Category) Update() rest_errors.RestErr {
// 	stmt, err := cashiers_db.Client.Prepare(queryUpdateCategory)
// 	if err != nil {
// 		logger.Error("error when trying to prepare get cashier statement", err)
// 		return rest_errors.NewInternalServerError("error when tying to get cashier", errors.New("database error"))
// 	}
// 	defer stmt.Close()

// 	_, err = stmt.Exec(category.Name, category.Id)
// 	if err != nil {
// 		logger.Error("error when trying to update user", err)
// 		return rest_errors.NewInternalServerError("error when tying to update user", errors.New("database error"))
// 	}
// 	return nil
// }

func(p *Product) GetById() rest_errors.RestErr {
	stmt, err := cashiers_db.Client.Prepare(queryById)
	if err != nil {
		logger.Error("error when trying to prepare get cashier statement", err)
		return rest_errors.NewInternalServerError("error when tying to get cashier", errors.New("database error"))
	}
	defer stmt.Close()

	result := stmt.QueryRow(p.Id)

	if err := result.Scan(&p.Id,
		&p.CategoryId,
		&p.Name,
		&p.Image,
		&p.Price,
		&p.Stock,
		&p.UpdatedAt,
		&p.CreatedAt,
		&p.Discount.Qty, 
		&p.Discount.Type,
		&p.Discount.Result,
		&p.Discount.ExpiredAt,
		); err != nil {
		logger.Error("error when scan cashier row into cashier struct", err)
		return rest_errors.NewInternalServerError("error when tying to gett cashier", errors.New("database error"))
	}

	return nil
}

func(product *Product) Create() rest_errors.RestErr {

	product.CreatedAt = time.Now().Format("2006-01-02T15:04:05Z")
	product.UpdatedAt = product.CreatedAt

	stmt, err := cashiers_db.Client.Prepare(queryInsertProduct)
	if err != nil {
		logger.Error("error when trying to prepare prepare product create statement", err)
		return rest_errors.NewInternalServerError("error when trying to get category", errors.New("database error"))
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(product.CategoryId, product.Name, product.Image, product.Price, product.Stock, product.Discount.Qty, product.Discount.Type, product.Discount.Result, product.Discount.ExpiredAt, product.CreatedAt, product.UpdatedAt)
	if saveErr != nil {
		logger.Error("error when trying to save product", saveErr)
		return rest_errors.NewInternalServerError("error when tying to save user", errors.New("database error"))
	}

	productId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert id after creating a new user", err)
		return rest_errors.NewInternalServerError("error when tying to save user", errors.New("database error"))
	}
	product.Id = productId

	return nil
}

func (p *Product) List() ([]Product, rest_errors.RestErr) {
	stmt, err := cashiers_db.Client.Prepare(queryListCategories)
	if err != nil {
		logger.Error("error when trying to prepare get cashier statement", err)
		return nil, rest_errors.NewInternalServerError("error when tying to get cashier", errors.New("database error"))
	}
	defer stmt.Close()
	rows, err := stmt.Query() //update with limit and skip
	if err != nil {
		logger.Error("error when trying list cahisers", err)
		return nil, rest_errors.NewInternalServerError("error when tying to get cashier", errors.New("database error"))
	}
	defer rows.Close()

	results := make([]Product, 0)
	for rows.Next() {
		var temp Product
		if err := rows.Scan(&temp.Id,
			&temp.CategoryId,
			&temp.Name,
			&temp.Image,
			&temp.Price,
			&temp.Stock,
			&temp.UpdatedAt,
			&temp.CreatedAt,
			&temp.Discount.Qty, 
			&temp.Discount.Type,
			&temp.Discount.Result,
			&temp.Discount.ExpiredAt,
			); err != nil {
			logger.Error("error when scan cashier row into cashier struct", err)
			return nil, rest_errors.NewInternalServerError("error when tying to gett cashier", errors.New("database error"))
		}
		results = append(results, temp)
	}
	if len(results) == 0 {
		return nil, rest_errors.NewNotFoundError(fmt.Sprintf("no cashiers matching status %s", "ok"))
	}
	return results, nil	
}