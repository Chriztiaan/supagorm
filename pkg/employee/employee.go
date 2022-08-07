package employee

import (
	"leave.gg/pkg"
	"leave.gg/pkg/crud"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name string `json:"name" gorm:"type:string; size:20; not null" binding:"required"`
	// ManagerID uint "This is how you set a relationship"
	Manager string `json:"Manager"`
}

type EmployeeController struct {
	router *gin.Engine
	*crud.CrudController
}

func NewEmployeeController(router *gin.Engine, crud *crud.CrudController) *EmployeeController {
	c := &EmployeeController{router, crud}
	c.AutoMigrate(&Employee{})
	c.registerRoutes()
	pkg.GenerateTypeScriptModel(Employee{})

	return c
}

func (c *EmployeeController) registerRoutes() {
	v1 := c.router.Group("api/v1")
	v1.GET("/employee/:id", c.getEmployee)
	v1.GET("/employees", c.getEmployees)
	v1.POST("/employee", c.createEmployee)
	v1.PUT("/employee/:id", c.updateEmployee)
	v1.DELETE("/employee/:id", c.deleteEmployee)
}

func (c *EmployeeController) getEmployee(ctx *gin.Context) {
	c.FirstBy(ctx, &Employee{}, "id")
}

func (c *EmployeeController) getEmployees(ctx *gin.Context) {
	c.All(ctx, &[]Employee{})
}

func (c *EmployeeController) createEmployee(ctx *gin.Context) {
	var request Employee

	if !c.Bind(ctx, &request) {
		return
	}

	e := Employee{Name: request.Name}

	c.Create(ctx, &e)
}

type employeeUpdateRequest struct {
	Name string `json:"name" binding:"required"`
}

func (c *EmployeeController) updateEmployee(ctx *gin.Context) {
	var updateRequest employeeUpdateRequest

	if !c.Bind(ctx, &updateRequest) {
		return
	}

	var e Employee
	if !c.FirstBySuppresed(ctx, &e, "id") {
		return
	}

	e.Name = updateRequest.Name

	c.Update(ctx, &e)
}

func (c *EmployeeController) deleteEmployee(ctx *gin.Context) {
	var e Employee
	if !c.FirstBySuppresed(ctx, &e, "id") {
		return
	}

	c.Delete(ctx, &e)
}
