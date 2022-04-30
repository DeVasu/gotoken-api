package cashiers

import (
	"encoding/json"
)

type Publiccashier struct {
	Id          int64  `json:"id"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

type Privatecashier struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

func (cashiers Cashiers) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(cashiers))
	for index, cashier := range cashiers {
		result[index] = cashier.Marshall(isPublic)
	}
	return result
}

func (cashier *Cashier) Marshall(isPublic bool) interface{} {
	if isPublic {
		return Publiccashier{
			Id:          cashier.Id,
		}
	}

	cashierJson, _ := json.Marshal(cashier)
	var privatecashier Privatecashier
	json.Unmarshal(cashierJson, &privatecashier)
	return privatecashier
}
