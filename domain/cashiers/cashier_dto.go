package cashiers

import (
	"strings"

	"github.com/federicoleon/bookstore_utils-go/rest_errors"
)


type Cashier struct {
	Id          int64  `json:"cashierId,omitempty"`
	Name   		string `json:"name,omitempty"`
	Passcode    string `json:"passcode,omitempty"`
}

type Cashiers []Cashier

func (cashier *Cashier) Validate() rest_errors.RestErr {
	cashier.Name = strings.TrimSpace(cashier.Name)
	cashier.Passcode = strings.TrimSpace(cashier.Passcode)
	if cashier.Passcode == "" {
		return rest_errors.NewBadRequestError("invalid password")
	}
	return nil
}
