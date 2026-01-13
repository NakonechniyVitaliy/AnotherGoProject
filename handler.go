package main

import (
	"github.com/gin-gonic/gin"
	"studyProject/service"
)

type Handler struct {
	EmployeeService *service.EmployeeService
}

func NewHandler(EmployeeService *service.EmployeeService) *Handler {
	return &Handler{
		EmployeeService: EmployeeService,
	}
}

func (h *Handler) CreateEmployee(c *gin.Context) {
	h.EmployeeService.NewEmployee(c)
}

func (h *Handler) UpdateEmployee(c *gin.Context) {
	h.EmployeeService.UpdateEmployee(c)
}

func (h *Handler) DeleteEmployee(c *gin.Context) {
	h.EmployeeService.DeleteEmployee(c)
}

func (h *Handler) GetEmployee(c *gin.Context) {
	h.EmployeeService.GetEmployee(c)
}

func (h *Handler) GetAllEmployee(c *gin.Context) {
	h.EmployeeService.GetAllEmployee(c)
}
