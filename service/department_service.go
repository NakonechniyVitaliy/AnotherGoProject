package service

import (
	"context"
	"studyProject/model"
	"studyProject/repository"
)

type DepartmentService struct {
	repo repository.DepartmentRepository
}

func NewDepartmentService(repo repository.DepartmentRepository) *DepartmentService {
	return &DepartmentService{
		repo: repo,
	}
}

func (service *DepartmentService) Create(ctx context.Context, department *model.Department) (*model.Department, error) {

	err := service.repo.Create(ctx, department)
	if err != nil {
		return nil, err
	}
	return department, nil
}

func (service *DepartmentService) Update(ctx context.Context, departmentFromRequest *model.Department, id int) (*model.Department, error) {

	currentDepartment, err := service.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	currentDepartment.Title = departmentFromRequest.Title
	currentDepartment.Description = departmentFromRequest.Description

	err = service.repo.Update(ctx, currentDepartment)
	if err != nil {
		return nil, err
	}
	return currentDepartment, nil
}

func (service *DepartmentService) Delete(ctx context.Context, id int) error {

	err := service.repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (service *DepartmentService) FindByID(ctx context.Context, id int) (*model.Department, error) {

	department, err := service.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return department, nil

}

func (service *DepartmentService) GetAll(ctx context.Context) ([]*model.Department, error) {

	departments, err := service.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return departments, nil
}
