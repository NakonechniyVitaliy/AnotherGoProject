package main

import (
	"context"
	"fmt"
	"studyProject/handler"
	mongo2 "studyProject/repository/mongo"
	routerPkg "studyProject/router"
	"studyProject/service"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client := createMongoDBclient(ctx)

	EmployeeDAO, err := mongo2.NewEmployeeDAO(ctx, client)
	if err != nil {
		return
	}
	DepartmentDAO, err := mongo2.NewDepartmentDAO(ctx, client)
	if err != nil {
		return
	}

	EmployeeService := service.NewEmployeeService(EmployeeDAO)
	DepartmentService := service.NewDepartmentService(DepartmentDAO)

	appHandler := handler.NewHandler(EmployeeService, DepartmentService)

	router := gin.Default()

	routerPkg.RegisterDepartmentRoutes(router, appHandler)
	routerPkg.RegisterEmployeeRoutes(router, appHandler)

	err = router.Run(":8001")
	if err != nil {
		return
	}

}

func createMongoDBclient(ctx context.Context) *mongo.Client {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		fmt.Printf("Error at create mongoDB client %s", err)
	}

	return client
}
