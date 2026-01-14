package handler

import "studyProject/service"

type ErrorResponce struct {
	Message string `json:"message"`
}

type Handler struct {
	EmployeeService   *service.EmployeeService
	DepartmentService *service.DepartmentService
}

func NewHandler(EmployeeService *service.EmployeeService, DepartmentService *service.DepartmentService) *Handler {
	return &Handler{
		EmployeeService:   EmployeeService,
		DepartmentService: DepartmentService,
	}
}
