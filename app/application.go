package app

import (
	"fmt"

	"github.com/DeVasu/gotoken-api/datasources/mysql/cashiers_db"
	"github.com/federicoleon/bookstore_utils-go/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {

	fmt.Println("You need to add information about MYSQL server to start")

	var username, password, host, db_name string
	host = "localhost:3306"
	username = "root"
	password = ""
	db_name = "users_db"

	cashiers_db.Init(username, password, host, db_name)


	mapUrls()


	logger.Info("about to start the application...")
	router.Run(":3030")
}
