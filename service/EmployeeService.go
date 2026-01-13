package service

import (
	"context"
	"studyProject/dao"
	"studyProject/model"
)

type EmployeeService struct {
	EmployeeDAO *dao.EmployeeDAO
}

func NewEmployeeService(EmployeeDAO *dao.EmployeeDAO) *EmployeeService {
	return &EmployeeService{
		EmployeeDAO: EmployeeDAO,
	}
}

func (service *EmployeeService) NewEmployee(ctx context.Context, employee *model.Employee) (*model.Employee, error) {

	err := service.EmployeeDAO.NewEmployee(ctx, employee)
	if err != nil {
		return nil, err
	}
	return employee, nil
}

func (service *EmployeeService) UpdateEmployee(ctx context.Context, employeeFromRequest *model.Employee, id int) (*model.Employee, error) {

	currentEmployee, err := service.EmployeeDAO.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	currentEmployee.Name = employeeFromRequest.Name
	currentEmployee.Sex = employeeFromRequest.Sex
	currentEmployee.Age = employeeFromRequest.Age
	currentEmployee.Salary = employeeFromRequest.Salary

	err = service.EmployeeDAO.Update(ctx, currentEmployee)
	if err != nil {
		return nil, err
	}
	return currentEmployee, nil
}

func (service *EmployeeService) DeleteEmployee(ctx context.Context, id int) error {

	err := service.EmployeeDAO.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (service *EmployeeService) GetEmployee(ctx context.Context, id int) (*model.Employee, error) {

	employee, err := service.EmployeeDAO.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return employee, nil

}

func (service *EmployeeService) GetAllEmployee(ctx context.Context) ([]*model.Employee, error) {

	employees, err := service.EmployeeDAO.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return employees, nil
}
