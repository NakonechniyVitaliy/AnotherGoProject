package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"studyProject/dao"
	"studyProject/model"
	"studyProject/service"

	"github.com/gin-gonic/gin"
)

type ErrorResponce struct {
	Message string `json:"message"`
}

type Handler struct {
	EmployeeService *service.EmployeeService
}

func NewHandler(EmployeeService *service.EmployeeService) *Handler {
	return &Handler{
		EmployeeService: EmployeeService,
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

	createdEmployee, err := h.EmployeeService.NewEmployee(ctx, &employee)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponce{
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, createdEmployee)
}

func (h *Handler) UpdateEmployee(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("failed to convert id param to int: %s", err)
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

	var updatedEmployee *model.Employee
	updatedEmployee, err = h.EmployeeService.UpdateEmployee(ctx, &employeeFromRequest, id)

	if err != nil {
		switch {
		case errors.Is(err, dao.ErrNotFound):
			c.JSON(http.StatusNotFound, ErrorResponce{Message: err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, ErrorResponce{Message: "internal error"})
		}
		return
	}

	c.JSON(http.StatusOK, updatedEmployee)

}

func (h *Handler) DeleteEmployee(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("failed to convert id param to int: %s", err)
		c.JSON(http.StatusBadRequest, ErrorResponce{
			Message: err.Error(),
		})
		return
	}

	err = h.EmployeeService.DeleteEmployee(ctx, id)
	if err != nil {
		fmt.Printf("failed to delete employee: %s", err)
		c.JSON(http.StatusBadRequest, ErrorResponce{
			Message: err.Error(),
		})
		return
	}
	c.String(http.StatusOK, "employee deleted")

}

func (h *Handler) GetEmployee(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("failed to convert id param to int: %s", err)
		c.JSON(http.StatusBadRequest, ErrorResponce{
			Message: err.Error(),
		})
		return
	}

	var employee *model.Employee
	employee, err = h.EmployeeService.GetEmployee(ctx, id)

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
	ctx := c.Request.Context()

	var employees []*model.Employee
	employees, err := h.EmployeeService.GetAllEmployee(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponce{
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, employees)
}
