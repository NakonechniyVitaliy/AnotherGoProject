package mongo

import (
	"context"
	"studyProject/model"
	"studyProject/repository"

	"go.mongodb.org/mongo-driver/bson"
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

func (dao *EmployeeDAO) Create(ctx context.Context, e *model.Employee) error {
	_, err := dao.c.InsertOne(ctx, e)
	return err
}

func (dao *EmployeeDAO) FindByID(ctx context.Context, id int) (*model.Employee, error) {
	filter := bson.D{{"id", id}}

	var Employee model.Employee
	err := dao.c.FindOne(ctx, filter).Decode(&Employee)

	switch {
	case err == nil:
		return &Employee, nil
	case err == mongo.ErrNoDocuments:
		return nil, repository.ErrNotFound

	default:
		return nil, err
	}
}

func (dao *EmployeeDAO) Update(ctx context.Context, employee *model.Employee) error {
	filter := bson.D{{"id", employee.ID}}

	result, err := dao.c.ReplaceOne(ctx, filter, employee)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return repository.ErrNotFound
	}
	return nil
}

func (dao *EmployeeDAO) Delete(ctx context.Context, id int) error {
	filter := bson.D{{"id", id}}

	result, err := dao.c.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return repository.ErrNotFound
	}

	return nil
}

func (dao *EmployeeDAO) GetAll(ctx context.Context) ([]*model.Employee, error) {

	cursor, err := dao.c.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var employees []*model.Employee

	if err := cursor.All(ctx, &employees); err != nil {
		return nil, err
	}

	return employees, nil
}

func (dao *EmployeeDAO) GetAllByDepartment(ctx context.Context, departmentId int) ([]*model.Employee, error) {
	filter := bson.D{{"department_id", departmentId}}

	cursor, err := dao.c.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var employees []*model.Employee
	if err := cursor.All(ctx, &employees); err != nil {
		return nil, err
	}

	return employees, nil
}
