package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"studyProject/dao"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := createMongoDBclient(ctx)

	EmployeeDAO, err := dao.NewEmployeeDAO(ctx, client)

	memoryStorage := newMemoryStorage()
	handler := NewHandler(memoryStorage, EmployeeDAO)

	router := gin.Default()
	router.POST("/employee", handler.CreateEmployee)
	router.GET("/employee/:id", handler.GetEmployee)
	router.GET("/employees", handler.GetAllEmployee)
	router.PUT("/employee/:id", handler.UpdateEmployee)
	router.DELETE("/employee/:id", handler.DeleteEmployee)

	err = router.Run(":8000")
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
