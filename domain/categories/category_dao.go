package categories

import (
	"errors"
	"fmt"

	"github.com/DeVasu/gotoken-api/datasources/mysql/cashiers_db"
	"github.com/federicoleon/bookstore_utils-go/logger"
	"github.com/federicoleon/bookstore_utils-go/rest_errors"
)

const (
	queryInsertCategory         = "INSERT INTO categories(name) VALUES(?);"
	queryGetCategory = "SELECT * from categories WHERE id=?;"
	queryUpdateCategory = "UPDATE categories SET name = ? WHERE id=?;"
	queryDeleteCategory = "DELETE FROM categories WHERE id=?;"
	// queryGetCategories               = "SELECT Categoriesd, first_name, last_name, email, date_created, status FROM Categories WHERE id=?;"
	// queryFindByCategoriesd           = "SELECT Categoriesd, name, passcode FROM Categories WHERE Categoriesd=?;"

	// queryUpdateCategories            = "UPDATE Categories SET name=?, passcode=? WHERE Categoriesd=?;"
	// queryDeleteCategories            = "DELETE FROM Categories WHERE Categoriesd=?;"
	// // queryFindByStatus           = "SELECT id, first_name, last_name, email, date_created, status FROM Categories WHERE status=?;"
	// queryFindByEmailAndPassword = "SELECT id, first_name, last_name, email, date_created, status FROM Categories WHERE email=? AND password=? AND status=?"
	queryListCategories 			= "SELECT id, name from categories;"
)

func(category *Category) Delete() rest_errors.RestErr {
	stmt, err := cashiers_db.Client.Prepare(queryDeleteCategory)
	if err != nil {
		logger.Error("error when trying to prepare get cashier statement", err)
		return rest_errors.NewInternalServerError("error when tying to get cashier", errors.New("database error"))
	}
	defer stmt.Close()

	_, err = stmt.Exec(category.Id)
	if err != nil {
		logger.Error("error when trying to update user", err)
		return rest_errors.NewInternalServerError("error when tying to update user", errors.New("database error"))
	}
	return nil
}

func(category *Category) Update() rest_errors.RestErr {
	stmt, err := cashiers_db.Client.Prepare(queryUpdateCategory)
	if err != nil {
		logger.Error("error when trying to prepare get cashier statement", err)
		return rest_errors.NewInternalServerError("error when tying to get cashier", errors.New("database error"))
	}
	defer stmt.Close()

	_, err = stmt.Exec(category.Name, category.Id)
	if err != nil {
		logger.Error("error when trying to update user", err)
		return rest_errors.NewInternalServerError("error when tying to update user", errors.New("database error"))
	}
	return nil
}

func(category *Category) GetById() rest_errors.RestErr {
	stmt, err := cashiers_db.Client.Prepare(queryGetCategory)
	if err != nil {
		logger.Error("error when trying to prepare get cashier statement", err)
		return rest_errors.NewInternalServerError("error when tying to get cashier", errors.New("database error"))
	}
	defer stmt.Close()

	result := stmt.QueryRow(category.Id)

	if getErr := result.Scan(&category.Id, &category.Name); getErr != nil {
		logger.Error("error when trying to get user by id", getErr)
		return rest_errors.NewInternalServerError("error when tying to get user", errors.New("database error"))
	}

	return nil
}

func(category *Category) Create() rest_errors.RestErr {
	stmt, err := cashiers_db.Client.Prepare(queryInsertCategory)
	if err != nil {
		logger.Error("error when trying to prepare get category statement", err)
		return rest_errors.NewInternalServerError("error when tying to get category", errors.New("database error"))
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(category.Name)
	if saveErr != nil {
		logger.Error("error when trying to save user", saveErr)
		return rest_errors.NewInternalServerError("error when tying to save user", errors.New("database error"))
	}

	categoryId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert id after creating a new user", err)
		return rest_errors.NewInternalServerError("error when tying to save user", errors.New("database error"))
	}
	category.Id = categoryId

	return nil
}

func (c *Category) List() ([]Category, rest_errors.RestErr) {
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

	results := make([]Category, 0)
	for rows.Next() {
		var category Category
		if err := rows.Scan(&category.Id, &category.Name); err != nil {
			logger.Error("error when scan cashier row into cashier struct", err)
			return nil, rest_errors.NewInternalServerError("error when tying to gett cashier", errors.New("database error"))
		}
		results = append(results, category)
	}
	if len(results) == 0 {
		return nil, rest_errors.NewNotFoundError(fmt.Sprintf("no cashiers matching status %s", "ok"))
	}
	return results, nil	
}