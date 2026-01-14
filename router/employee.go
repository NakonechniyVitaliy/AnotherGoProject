package router

import (
	"studyProject/handler"

	"github.com/gin-gonic/gin"
)

func RegisterEmployeeRoutes(r *gin.Engine, h *handler.Handler) {
	employees := r.Group("/employees")

	{
		employees.POST("", h.CreateEmployee)
		employees.GET("", h.GetAllEmployee)
		employees.GET("/:id", h.GetEmployee)
		employees.PUT("/:id", h.UpdateEmployee)
		employees.DELETE("/:id", h.DeleteEmployee)
		employees.GET("/by-department/:id", h.GetAllEmployeeByDepartment)
	}
}
