package main

import (
	"fmt"

	"github.com/Sebas03446/Challenge-Go/libs"
	"github.com/Sebas03446/Challenge-Go/pkg/domain/models"
	"github.com/Sebas03446/Challenge-Go/pkg/http/rest"
	"github.com/Sebas03446/Challenge-Go/pkg/repository/mysql"
	"github.com/Sebas03446/Challenge-Go/pkg/service"
)

func main() {
	dbConfig := libs.Configure("./config", "mysql")
	libs.DB = dbConfig.InitMysqlDB()
	error := libs.DB.AutoMigrate(&models.Property{})
	if error != nil {
		fmt.Println(error)
	}
	propertyRepo := mysql.NewQueryBuilder(libs.DB)
	propertyService := service.NewService(propertyRepo)
	rest.UserService = propertyService
	rest.Controller()
}
