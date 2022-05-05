package payments

import (
	// "net/http"
	// "strconv"

	// "github.com/federicoleon/bookstore_oauth-go/oauth"
	// "github.com/DeVasu/gotoken-api/domain/users"

	// "github.com/federicoleon/bookstore_utils-go/rest_errors"

	"fmt"
	"net/http"
	"strconv"

	"github.com/DeVasu/gotoken-api/domain/payments"
	"github.com/DeVasu/gotoken-api/utils"
	"github.com/federicoleon/bookstore_utils-go/rest_errors"
	"github.com/gin-gonic/gin"
)



func getUserId(userIdParam string) (int64, rest_errors.RestErr) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, rest_errors.NewBadRequestError("user id should be a number")
	}
	return userId, nil
}

func Delete(c *gin.Context) {

	paymentId, idErr := getUserId(c.Param("paymentId"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	var temp payments.Payment

	temp.Id = paymentId
	err := temp.Delete()
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	res := &utils.Response{
		Success: true,
		Message: "Success",
	}
	
	c.JSON(http.StatusOK, res)

}

func Update(c *gin.Context) {

	paymentId, idErr := getUserId(c.Param("paymentId"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	var temp payments.Payment
	if err := c.ShouldBindJSON(&temp); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}

	temp.Id = paymentId

	err := temp.Update()
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	res := &utils.Response{
		Success: true,
		Message: "Success",
	}
	
	c.JSON(http.StatusOK, res)

}

func GetById(c *gin.Context) {
	
	paymentId, idErr := getUserId(c.Param("paymentId"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}
	result := &payments.Payment{
		Id: paymentId,
	}
	err := result.GetById()
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, result)
}

// //i want to list cashiers
func List(c *gin.Context) {

	result := &payments.Payment{}
	

	listOf, err := result.List()
	fmt.Println(listOf[:2])


	if err != nil {
		c.JSON(http.StatusBadRequest, "\"{\"err\":\"wrong\"}")
		return
	}

	c.JSON(http.StatusOK, listOf)

}

func Create(c *gin.Context) {

	res := &payments.Payment{}
	
	if err := c.ShouldBindJSON(&res); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	
	err := res.Create()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, res)

}