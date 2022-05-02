package cashiers

import (
	// "net/http"
	// "strconv"

	// "github.com/federicoleon/bookstore_oauth-go/oauth"
	// "github.com/DeVasu/gotoken-api/domain/users"

	// "github.com/federicoleon/bookstore_utils-go/rest_errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/DeVasu/gotoken-api/domain/cashiers"
	"github.com/DeVasu/gotoken-api/utils"
	"github.com/federicoleon/bookstore_utils-go/rest_errors"
	"github.com/gin-gonic/gin"
)



func getUserId(userIdParam string) (int64, rest_errors.RestErr) {
	if userIdParam == "" {
		return 0, nil
	}
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, rest_errors.NewBadRequestError("user id should be a number")
	}
	return userId, nil
}

func Common(c *gin.Context) {

	if strings.Contains(c.Request.URL.String(), "passcode") {
		GetPasscode(c)
		fmt.Println("in passcode")
	} else if strings.Contains(c.Request.URL.String(), "login") {
		Login(c)
		fmt.Println("in login")
	} else if  strings.Contains(c.Request.URL.String(), "logout") {
		Logout(c)
	} else {
		GetById(c)
		fmt.Println("in getbyid")
	}

}


func Logout(c *gin.Context) {
	
	cashierId, idErr := getUserId(c.Param("cashierId"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}
	var result cashiers.Cashier
	if err := c.ShouldBindJSON(&result); err != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	temp := &cashiers.Cashier{Id: cashierId,}
	if err := temp.GetById(); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	msg := fmt.Sprintf("{\"%s\"}:{\"%s\"}", "message", "successfully logged out")

	if temp.Passcode == result.Passcode {
		c.JSON(http.StatusOK, msg)
	} else {
		c.JSON(http.StatusBadRequest, "{\"token\":\"bad request\"}")
	}


}
func Login(c *gin.Context) {
	
	cashierId, idErr := getUserId(c.Param("cashierId"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}
	var result cashiers.Cashier
	if err := c.ShouldBindJSON(&result); err != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	temp := &cashiers.Cashier{Id: cashierId,}
	if err := temp.GetById(); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	token := 
		"{\"token\":\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2NDYzNjk2NDQsInN1YiI6IjEifQ.OXOV-TjfCbCCJ7z1w1osQ1lz99rK89V_Ert_Y1JUfCM\"}"

	if temp.Passcode == result.Passcode {
		c.JSON(http.StatusOK, token)
	} else {
		c.JSON(http.StatusBadRequest, rest_errors.NewBadRequestError("passcode not correct"))
	}


}

func GetPasscode(c *gin.Context) {
	
	cashierId, idErr := getUserId(c.Param("cashierId"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}
	result := &cashiers.Cashier{
		Id: cashierId,
	}
	err := result.GetById()
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	result.Id=0
	result.Name=""
	c.JSON(http.StatusOK, result)

}


func Delete(c *gin.Context) {

	cashierId, idErr := getUserId(c.Param("cashierId"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	var temp cashiers.Cashier
	// if err := c.ShouldBindJSON(&temp); err != nil {
	// 	c.JSON(http.StatusBadGateway, err)
	// 	return
	// }

	temp.Id = cashierId
	fmt.Println(temp)
	err := temp.Delete()
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	
	c.JSON(http.StatusOK, temp)

}

func Update(c *gin.Context) {

	cashierId, idErr := getUserId(c.Param("cashierId"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	var temp cashiers.Cashier
	if err := c.ShouldBindJSON(&temp); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}

	temp.Id = cashierId

	err := temp.Update()
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	
	c.JSON(http.StatusOK, temp)

}

func GetById(c *gin.Context) {
	
	cashierId, idErr := getUserId(c.Param("cashierId"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}
	result := &cashiers.Cashier{
		Id: cashierId,
	}
	err := result.GetById()
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	result.Passcode=""
	c.JSON(http.StatusOK, result)

}

//i want to list cashiers
func List(c *gin.Context) {
	
	temp := &cashiers.Cashier{}

	limit, err := getUserId(c.Query("limit"))
	if err != nil {
		res := utils.NewFailureResponse("limit value was not parsed", utils.Meta{})
		c.JSON(http.StatusBadRequest, res)
		return
	}
	skip, err := getUserId(c.Query("skip"))
	if err != nil {
		res := utils.NewFailureResponse("skip value was not parsed", utils.Meta{})
		c.JSON(http.StatusBadRequest, res)
		return
	}
	listOf, err := temp.GetList(limit, skip)

	fmt.Println("this is limit, skip", limit, skip)

	meta := utils.Meta{
		Limit: limit,
		Skip: skip,
	}

	if err != nil {
		fmt.Println(err.Error())
		res := utils.NewFailureResponse("something wrong while executing query", meta)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	meta.Total = int64(len(listOf))
	res := utils.NewSuccessResponse(listOf, meta)

	c.JSON(http.StatusOK, res)

}

func Create(c *gin.Context) {

	res := &cashiers.Cashier{};
	
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