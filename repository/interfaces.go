package repository

import (
	"context"
	"studyProject/model"
)

type EmployeeRepository interface {
	Create(ctx context.Context, employee *model.Employee) error
	FindByID(ctx context.Context, id int) (*model.Employee, error)
	Update(ctx context.Context, employee *model.Employee) error
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]*model.Employee, error)
	GetAllByDepartment(ctx context.Context, departmentId int) ([]*model.Employee, error)
}

type DepartmentRepository interface {
	Create(ctx context.Context, department *model.Department) error
	FindByID(ctx context.Context, id int) (*model.Department, error)
	Update(ctx context.Context, department *model.Department) error
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]*model.Department, error)
}
