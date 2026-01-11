package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"studyProject/dao"
	"studyProject/model"
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

	//h.storage.Insert(&employee)
	//c.JSON(http.StatusOK, map[string]interface{}{
	//	"id": employee.ID,
	//})

}

func (h *Handler) UpdateEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		fmt.Printf("failed to convert id param to int: %s", err)
		c.JSON(http.StatusBadRequest, ErrorResponce{
			Message: err.Error(),
		})
		return
	}

	var employee model.Employee

	if err := c.BindJSON(&employee); err != nil {
		fmt.Printf("failed to bind employee: %v", err)
		c.JSON(http.StatusBadRequest, ErrorResponce{
			Message: err.Error(),
		})
		return

	}

	h.storage.Update(id, employee)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": employee.ID,
	})
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
