package main

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

type Handler struct {
	storage     Storage
	EmployeeDAO *dao.EmployeeDAO
}

func NewHandler(storage Storage, EmployeeDAO *dao.EmployeeDAO) *Handler {
	return &Handler{
		storage:     storage,
		EmployeeDAO: EmployeeDAO,
	}
}

func (h *Handler) CreateEmployee(c *gin.Context) {
	ctx := c.Request.Context()

	var employee model.Employee

	if err := c.ShouldBindJSON(&employee); err != nil {
		fmt.Printf("failed to bind employee: %v", err)
		c.JSON(http.StatusBadRequest, ErrorResponce{
			Message: err.Error(),
		})
		return
	}

	err := h.EmployeeDAO.NewEmployee(ctx, &employee)
	if err != nil {
		return
	}

}

func (h *Handler) UpdateEmployee(c *gin.Context) (*model.Employee, error) {
	ctx := c.Request.Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("failed to convert id param to int: %s", err)
		c.JSON(http.StatusBadRequest, ErrorResponce{
			Message: err.Error(),
		})
		return nil, err
	}
	currentEmployee, err := h.EmployeeDAO.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	var employeeFromRequest model.Employee
	if err := c.BindJSON(&employeeFromRequest); err != nil {
		fmt.Printf("failed to bind employee: %v", err)
		c.JSON(http.StatusBadRequest, ErrorResponce{
			Message: err.Error(),
		})
		return nil, err
	}
	currentEmployee.Name = employeeFromRequest.Name
	currentEmployee.Sex = employeeFromRequest.Sex
	currentEmployee.Age = employeeFromRequest.Age
	currentEmployee.Salary = employeeFromRequest.Salary

	return currentEmployee, h.EmployeeDAO.Update(ctx, currentEmployee)
}

func (h *Handler) DeleteEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("failed to convert id param to int: %s", err)
		c.JSON(http.StatusBadRequest, ErrorResponce{
			Message: err.Error(),
		})
		return
	}
	h.storage.Delete(id)

	c.String(http.StatusOK, "employee deleted")
}

func (h *Handler) GetEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("failed to convert id param to int: %s", err)
		c.JSON(http.StatusBadRequest, ErrorResponce{
			Message: err.Error(),
		})
	}

	employee, err := h.storage.Get(id)

	if err != nil {
		fmt.Printf("employee not found: %s", err)
		c.JSON(http.StatusBadRequest, ErrorResponce{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, employee)

}

func (h *Handler) GetAllEmployee(c *gin.Context) {
	c.JSON(http.StatusOK, h.storage.GetAll())
}
