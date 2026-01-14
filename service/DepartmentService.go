package service

import (
	"context"
	"studyProject/model"
	"studyProject/repository"
)

type DepartmentService struct {
	DepartmentDAO *repository.DepartmentDAO
}

func NewDepartmentService(DepartmentDAO *repository.DepartmentDAO) *DepartmentService {
	return &DepartmentService{
		DepartmentDAO: DepartmentDAO,
	}
}

func (service *DepartmentService) NewDepartment(ctx context.Context, department *model.Department) (*model.Department, error) {

	err := service.DepartmentDAO.NewDepartment(ctx, department)
	if err != nil {
		return nil, err
	}
	return department, nil
}

func (service *DepartmentService) UpdateDepartment(ctx context.Context, departmentFromRequest *model.Department, id int) (*model.Department, error) {

	currentDepartment, err := service.DepartmentDAO.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	currentDepartment.Title = departmentFromRequest.Title
	currentDepartment.Description = departmentFromRequest.Description

	err = service.DepartmentDAO.Update(ctx, currentDepartment)
	if err != nil {
		return nil, err
	}
	return currentDepartment, nil
}

func (service *DepartmentService) DeleteDepartment(ctx context.Context, id int) error {

	err := service.DepartmentDAO.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (service *DepartmentService) GetDepartment(ctx context.Context, id int) (*model.Department, error) {

	department, err := service.DepartmentDAO.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return department, nil

}

func (service *DepartmentService) GetAllDepartment(ctx context.Context) ([]*model.Department, error) {

	departments, err := service.DepartmentDAO.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return departments, nil
}
