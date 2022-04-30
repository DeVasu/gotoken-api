package cashiers

import (
	"errors"
	"fmt"

	"github.com/DeVasu/gotoken-api/datasources/mysql/cashiers_db"
	"github.com/federicoleon/bookstore_utils-go/logger"
	"github.com/federicoleon/bookstore_utils-go/rest_errors"
)

const (
	queryInsertcashier             = "INSERT INTO cashiers(name, passcode) VALUES(?, ?);"
	queryGetcashier                = "SELECT cashierId, first_name, last_name, email, date_created, status FROM cashiers WHERE id=?;"
	queryFindByCashierId           = "SELECT cashierId, name, passcode FROM cashiers WHERE cashierId=?;"

	queryUpdatecashier             = "UPDATE cashiers SET name=?, passcode=? WHERE cashierId=?;"
	queryDeletecashier             = "DELETE FROM cashiers WHERE cashierId=?;"
	// queryFindByStatus           = "SELECT id, first_name, last_name, email, date_created, status FROM cashiers WHERE status=?;"
	queryFindByEmailAndPassword = "SELECT id, first_name, last_name, email, date_created, status FROM cashiers WHERE email=? AND password=? AND status=?"
	queryListCashiers 			= "SELECT cashierId, name from cashiers;"
)



func(cashier *Cashier) Delete() rest_errors.RestErr {
	stmt, err := cashiers_db.Client.Prepare(queryDeletecashier)
	if err != nil {
		logger.Error("error when trying to prepare get cashier statement", err)
		return rest_errors.NewInternalServerError("error when tying to get cashier", errors.New("database error"))
	}
	defer stmt.Close()

	_, err = stmt.Exec(cashier.Id)
	if err != nil {
		logger.Error("error when trying to update user", err)
		return rest_errors.NewInternalServerError("error when tying to update user", errors.New("database error"))
	}
	return nil
}
func(cashier *Cashier) Update() rest_errors.RestErr {
	stmt, err := cashiers_db.Client.Prepare(queryUpdatecashier)
	if err != nil {
		logger.Error("error when trying to prepare get cashier statement", err)
		return rest_errors.NewInternalServerError("error when tying to get cashier", errors.New("database error"))
	}
	defer stmt.Close()

	_, err = stmt.Exec(cashier.Name, cashier.Passcode, cashier.Id)
	if err != nil {
		logger.Error("error when trying to update user", err)
		return rest_errors.NewInternalServerError("error when tying to update user", errors.New("database error"))
	}
	return nil
}


func(cashier *Cashier) GetById() rest_errors.RestErr {
	stmt, err := cashiers_db.Client.Prepare(queryFindByCashierId)
	if err != nil {
		logger.Error("error when trying to prepare get cashier statement", err)
		return rest_errors.NewInternalServerError("error when tying to get cashier", errors.New("database error"))
	}
	defer stmt.Close()

	result := stmt.QueryRow(cashier.Id)

	if getErr := result.Scan(&cashier.Id, &cashier.Name, &cashier.Passcode); getErr != nil {
		logger.Error("error when trying to get user by id", getErr)
		return rest_errors.NewInternalServerError("error when tying to get user", errors.New("database error"))
	}

	return nil
}

func(cashier *Cashier) Create() rest_errors.RestErr {
	stmt, err := cashiers_db.Client.Prepare(queryInsertcashier)
	if err != nil {
		logger.Error("error when trying to prepare get cashier statement", err)
		return rest_errors.NewInternalServerError("error when tying to get cashier", errors.New("database error"))
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(cashier.Name, cashier.Passcode)
	if saveErr != nil {
		logger.Error("error when trying to save user", saveErr)
		return rest_errors.NewInternalServerError("error when tying to save user", errors.New("database error"))
	}

	cashierId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert id after creating a new user", err)
		return rest_errors.NewInternalServerError("error when tying to save user", errors.New("database error"))
	}
	cashier.Id = cashierId

	return nil
}

func (cashier *Cashier) GetList(limit, offset int ) ([]Cashier, rest_errors.RestErr) {
	stmt, err := cashiers_db.Client.Prepare(queryListCashiers)
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

	results := make([]Cashier, 0)
	for rows.Next() {
		var cashier Cashier
		if err := rows.Scan(&cashier.Id, &cashier.Name); err != nil {
			logger.Error("error when scan cashier row into cashier struct", err)
			return nil, rest_errors.NewInternalServerError("error when tying to gett cashier", errors.New("database error"))
		}
		results = append(results, cashier)
	}
	if len(results) == 0 {
		return nil, rest_errors.NewNotFoundError(fmt.Sprintf("no cashiers matching status %s", "ok"))
	}
	return results, nil	
}
