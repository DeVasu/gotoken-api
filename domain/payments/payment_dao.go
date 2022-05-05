package payments

import (
	"errors"
	"fmt"
	"time"

	"github.com/DeVasu/gotoken-api/datasources/mysql/cashiers_db"
	"github.com/federicoleon/bookstore_utils-go/logger"
	"github.com/federicoleon/bookstore_utils-go/rest_errors"
)



const (
	queryInsertPayment         = "INSERT INTO payments(name, type, logo, updatedAt, createdAt) VALUES(?, ?, ?, ?, ?);"
	queryListPayments 			= "SELECT * from payments;"
	queryById = "SELECT * from payments where id=?;"
	queryUpdatePayment = "UPDATE payments SET name=?, type=?, logo=? WHERE id = ?;"
	queryDeletePayment = "DELETE FROM payments WHERE id=?;"
	
)


func(p *Payment) Delete() rest_errors.RestErr {
	stmt, err := cashiers_db.Client.Prepare(queryDeletePayment)
	if err != nil {
		logger.Error("error when trying to prepare get cashier statement", err)
		return rest_errors.NewInternalServerError("error when tying to get cashier", errors.New("database error"))
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.Id)
	if err != nil {
		logger.Error("error when trying to update user", err)
		return rest_errors.NewInternalServerError("error when tying to update user", errors.New("database error"))
	}
	return nil
}

func(p *Payment) Update() rest_errors.RestErr {

	temp := &Payment{
		Id : p.Id,
	}
	temp.GetById()

	if p.Id != 0 {
		temp.Id = p.Id
	}
	if len(p.Name) != 0 {
		temp.Name = p.Name
	}
	if len(p.Type) != 0 {
		temp.Type = p.Type
	}


	stmt, err := cashiers_db.Client.Prepare(queryUpdatePayment)
	if err != nil {
		logger.Error("error when trying to prepare get cashier statement", err)
		return rest_errors.NewInternalServerError("error when tying to get cashier", errors.New("database error"))
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		temp.Name,
		temp.Type,
		temp.Logo,
		temp.Id,
	)
	if err != nil {
		logger.Error("error when trying to update user", err)
		return rest_errors.NewInternalServerError("error when tying to update user", errors.New("database error"))
	}
	return nil
}

func(p *Payment) GetById() rest_errors.RestErr {
	stmt, err := cashiers_db.Client.Prepare(queryById)
	if err != nil {
		logger.Error("error when trying to prepare get cashier statement", err)
		return rest_errors.NewInternalServerError("error when tying to get cashier", errors.New("database error"))
	}
	defer stmt.Close()

	result := stmt.QueryRow(p.Id)

	if err := result.Scan(
		&p.Id,
		&p.Name,
		&p.Type,
		&p.Logo,
		); err != nil {
		logger.Error("error when scan cashier row into cashier struct", err)
		return rest_errors.NewInternalServerError("error when tying to gett cashier", errors.New("database error"))
	}

	return nil
}

func(p *Payment) Create() rest_errors.RestErr {

	p.CreatedAt = time.Now().Format("2006-01-02T15:04:05Z")
	p.UpdatedAt = p.CreatedAt

	stmt, err := cashiers_db.Client.Prepare(queryInsertPayment)
	if err != nil {
		logger.Error("error when trying to prepare prepare Payment create statement", err)
		return rest_errors.NewInternalServerError("error when trying to get category", errors.New("database error"))
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(p.Name, p.Type, p.Logo, p.UpdatedAt, p.CreatedAt)
	if saveErr != nil {
		logger.Error("error when trying to save Payment", saveErr)
		return rest_errors.NewInternalServerError("error when tying to save user", errors.New("database error"))
	}

	paymentId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert id after creating a new user", err)
		return rest_errors.NewInternalServerError("error when tying to save user", errors.New("database error"))
	}
	p.Id = paymentId

	return nil
}

func (p *Payment) List() ([]Payment, rest_errors.RestErr) {
	stmt, err := cashiers_db.Client.Prepare(queryListPayments)
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

	results := make([]Payment, 0)
	for rows.Next() {
		var temp Payment
		if err := rows.Scan(
			&p.Id,
			&p.Name,
			&p.Type,
			&p.Logo,
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