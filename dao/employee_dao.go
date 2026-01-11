package dao

import (
	"context"
	"studyProject/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeDAO struct {
	c *mongo.Collection
}

func NewEmployeeDAO(ctx context.Context, client *mongo.Client) (*EmployeeDAO, error) {
	return &EmployeeDAO{
		c: client.Database("core").Collection("employees"),
	}, nil
}

func (dao *EmployeeDAO) NewEmployee(ctx context.Context, e *model.Employee) error {
	_, err := dao.c.InsertOne(ctx, e)
	return err
}
