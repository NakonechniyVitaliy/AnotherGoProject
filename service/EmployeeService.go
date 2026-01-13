package service

import (
	"fmt"
	"net/http"
	"strconv"
	"studyProject/dao"
	"studyProject/model"

	"github.com/gin-gonic/gin"
)

type ErrorResponce struct {
	Message string `json:"message"`
}

type EmployeeService struct {
	EmployeeDAO *dao.EmployeeDAO
}

func NewEmployeeService(EmployeeDAO *dao.EmployeeDAO) *EmployeeService {
	return &EmployeeService{
		EmployeeDAO: EmployeeDAO,
	}
}

func (service *EmployeeService) NewEmployee(c *gin.Context) {
	ctx := c.Request.Context()

	var employee model.Employee

	if err := c.ShouldBindJSON(&employee); err != nil {
		fmt.Printf("failed to bind employee: %v", err)
		c.JSON(http.StatusBadRequest, ErrorResponce{
			Message: err.Error(),
		})
		return
	}

	err := service.EmployeeDAO.NewEmployee(ctx, &employee)
	if err != nil {
		return
	} else {
		c.JSON(http.StatusOK, employee)
	}
}

func (service *EmployeeService) UpdateEmployee(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("failed to convert id param to int: %s", err)
		c.JSON(http.StatusBadRequest, ErrorResponce{
			Message: err.Error(),
		})
		return
	}
	currentEmployee, err := service.EmployeeDAO.FindByID(ctx, id)
	if err != nil {
		fmt.Printf("failed to find employee: %s", err)
		c.JSON(http.StatusBadRequest, ErrorResponce{
			Message: err.Error(),
		})
		return
	}

	var employeeFromRequest model.Employee
	if err := c.BindJSON(&employeeFromRequest); err != nil {
		fmt.Printf("failed to bind employee: %v", err)
		c.JSON(http.StatusBadRequest, ErrorResponce{
			Message: err.Error(),
		})
		return
	}
	currentEmployee.Name = employeeFromRequest.Name
	currentEmployee.Sex = employeeFromRequest.Sex
	currentEmployee.Age = employeeFromRequest.Age
	currentEmployee.Salary = employeeFromRequest.Salary

	err = service.EmployeeDAO.Update(ctx, currentEmployee)
	if err == nil {
		c.JSON(http.StatusOK, currentEmployee)
	} else {
		fmt.Printf("failed to update employee: %s", err)
		c.JSON(http.StatusBadRequest, ErrorResponce{
			Message: err.Error(),
		})
		return
	}
}

func (service *EmployeeService) DeleteEmployee(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("failed to convert id param to int: %s", err)
		c.JSON(http.StatusBadRequest, ErrorResponce{
			Message: err.Error(),
		})
		return
	}
	err = service.EmployeeDAO.Delete(ctx, id)
	if err != nil {
		fmt.Printf("failed to delete employee: %s", err)
		c.JSON(http.StatusBadRequest, ErrorResponce{
			Message: err.Error(),
		})
		return
	} else {
		c.String(http.StatusOK, "employee deleted")
	}

}

func (service *EmployeeService) GetEmployee(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("failed to convert id param to int: %s", err)
		c.JSON(http.StatusBadRequest, ErrorResponce{
			Message: err.Error(),
		})
		return
	}

	employee, err := service.EmployeeDAO.FindByID(ctx, id)
	if err != nil {
		fmt.Printf("employee not found: %s", err)
		c.JSON(http.StatusBadRequest, ErrorResponce{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, employee)
}

func (service *EmployeeService) GetAllEmployee(c *gin.Context) {
	ctx := c.Request.Context()

	employees, err := service.EmployeeDAO.GetAll(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponce{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, employees)
}
