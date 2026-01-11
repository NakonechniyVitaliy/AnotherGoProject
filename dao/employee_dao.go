package dao

import (
	"context"
	"errors"
	"studyProject/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeDAO struct {
	c *mongo.Collection
}

var ErrNotFound = errors.New("employee not found")

func NewEmployeeDAO(ctx context.Context, client *mongo.Client) (*EmployeeDAO, error) {
	return &EmployeeDAO{
		c: client.Database("core").Collection("employees"),
	}, nil
}

func (dao *EmployeeDAO) NewEmployee(ctx context.Context, e *model.Employee) error {
	_, err := dao.c.InsertOne(ctx, e)
	return err
}

func (dao *EmployeeDAO) FindByID(ctx context.Context, id int) (*model.Employee, error) {
	filter := bson.D{{"_id", id}}

	var Employee model.Employee
	err := dao.c.FindOne(ctx, filter).Decode(&Employee)

	switch {
	case err == nil:
		return &Employee, nil
	case err == mongo.ErrNoDocuments:
		return nil, ErrNotFound

	default:
		return nil, err
	}

}
func (dao *EmployeeDAO) Update(ctx context.Context, employee *model.Employee) error {
	filter := bson.D{{"_id", employee.ID}}

	result, err := dao.c.ReplaceOne(ctx, filter, employee)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return ErrNotFound
	}
	return nil

}
