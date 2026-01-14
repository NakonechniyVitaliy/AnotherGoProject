package repository

import (
	"context"
	"studyProject/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DepartmentDAO struct {
	c *mongo.Collection
}

func NewDepartmentDAO(ctx context.Context, client *mongo.Client) (*DepartmentDAO, error) {
	return &DepartmentDAO{
		c: client.Database("core").Collection("departments"),
	}, nil
}

func (dao *DepartmentDAO) NewDepartment(ctx context.Context, e *model.Department) error {
	_, err := dao.c.InsertOne(ctx, e)
	return err
}

func (dao *DepartmentDAO) FindByID(ctx context.Context, id int) (*model.Department, error) {
	filter := bson.D{{"id", id}}

	var Department model.Department
	err := dao.c.FindOne(ctx, filter).Decode(&Department)

	switch {
	case err == nil:
		return &Department, nil
	case err == mongo.ErrNoDocuments:
		return nil, ErrNotFound

	default:
		return nil, err
	}
}

func (dao *DepartmentDAO) Update(ctx context.Context, department *model.Department) error {
	filter := bson.D{{"id", department.ID}}

	result, err := dao.c.ReplaceOne(ctx, filter, department)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return ErrNotFound
	}
	return nil
}

func (dao *DepartmentDAO) Delete(ctx context.Context, id int) error {
	filter := bson.D{{"id", id}}

	result, err := dao.c.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return ErrNotFound
	}

	return nil
}

func (dao *DepartmentDAO) GetAll(ctx context.Context) ([]*model.Department, error) {

	cursor, err := dao.c.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var departments []*model.Department

	if err := cursor.All(ctx, &departments); err != nil {
		return nil, err
	}

	return departments, nil
}
