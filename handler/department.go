package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"studyProject/model"
	"studyProject/repository"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateDepartment(c *gin.Context) {
	ctx := c.Request.Context()

	var department model.Department
	if err := c.ShouldBindJSON(&department); err != nil {
		fmt.Printf("failed to bind department: %v", err)
		c.JSON(http.StatusBadRequest, ErrorResponce{
			Message: err.Error(),
		})
		return
	}

	createdDepartment, err := h.DepartmentService.NewDepartment(ctx, &department)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponce{
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, createdDepartment)
}

func (h *Handler) UpdateDepartment(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("failed to convert id param to int: %s", err)
		c.JSON(http.StatusBadRequest, ErrorResponce{
			Message: err.Error(),
		})
		return
	}

	var departmentFromRequest model.Department
	if err := c.BindJSON(&departmentFromRequest); err != nil {
		fmt.Printf("failed to bind department: %v", err)
		c.JSON(http.StatusBadRequest, ErrorResponce{
			Message: err.Error(),
		})
		return
	}

	var updatedDepartment *model.Department
	updatedDepartment, err = h.DepartmentService.UpdateDepartment(ctx, &departmentFromRequest, id)

	if err != nil {
		switch {
		case errors.Is(err, repository.ErrNotFound):
			c.JSON(http.StatusNotFound, ErrorResponce{Message: err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, ErrorResponce{Message: "internal error"})
		}
		return
	}

	c.JSON(http.StatusOK, updatedDepartment)

}

func (h *Handler) DeleteDepartment(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("failed to convert id param to int: %s", err)
		c.JSON(http.StatusBadRequest, ErrorResponce{
			Message: err.Error(),
		})
		return
	}

	err = h.DepartmentService.DeleteDepartment(ctx, id)
	if err != nil {
		fmt.Printf("failed to delete department: %s", err)
		c.JSON(http.StatusBadRequest, ErrorResponce{
			Message: err.Error(),
		})
		return
	}
	c.String(http.StatusOK, "department deleted")

}

func (h *Handler) GetDepartment(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("failed to convert id param to int: %s", err)
		c.JSON(http.StatusBadRequest, ErrorResponce{
			Message: err.Error(),
		})
		return
	}

	var department *model.Department
	department, err = h.DepartmentService.GetDepartment(ctx, id)

	if err != nil {
		fmt.Printf("department not found: %s", err)
		c.JSON(http.StatusBadRequest, ErrorResponce{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, department)
}

func (h *Handler) GetAllDepartment(c *gin.Context) {
	ctx := c.Request.Context()

	var departments []*model.Department
	departments, err := h.DepartmentService.GetAllDepartment(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponce{
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, departments)
}
