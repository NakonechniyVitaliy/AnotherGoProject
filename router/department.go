package router

import (
	"studyProject/handler"

	"github.com/gin-gonic/gin"
)

func RegisterDepartmentRoutes(r *gin.Engine, h *handler.Handler) {
	departments := r.Group("/departments")

	{
		departments.POST("", h.CreateDepartment)
		departments.GET("", h.GetAllDepartment)
		departments.GET("/:id", h.GetDepartment)
		departments.PUT("/:id", h.UpdateDepartment)
		departments.DELETE("/:id", h.DeleteDepartment)
	}
}
