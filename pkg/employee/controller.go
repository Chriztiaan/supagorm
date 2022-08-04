package employee

import (
	"fmt"
	"net/http"
	"strconv"

	"leave.gg/pkg"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type EmployeeController struct {
	router *gin.Engine
	*gorm.DB
}

func NewEmployeeController(router *gin.Engine, db *gorm.DB) *EmployeeController {
	c := &EmployeeController{router, db}
	c.prepareDb()
	c.prepareRouter()

	return c
}

func (c *EmployeeController) prepareDb() {
	fmt.Printf("Auto Migrating %T...\n", &Employee{})
	c.AutoMigrate(&Employee{})
}

func (c *EmployeeController) prepareRouter() {
	v1 := c.router.Group("api/v1")
	v1.GET("/employee/:id", c.getEmployee)
	v1.GET("/employees", c.getEmployees)
	v1.POST("/employee", c.createEmployee)
	v1.PUT("/employee/:id", c.updateEmployee)
	v1.DELETE("/employee/:id", c.deleteEmployee)
}

func (c *EmployeeController) getEmployee(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var e Employee

	if err := c.First(&e, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	ctx.JSON(http.StatusOK, e)
}

func (c *EmployeeController) getEmployees(ctx *gin.Context) {
	var employees []Employee

	if err := c.Find(&employees).Error; err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, employees)
}

func (c *EmployeeController) createEmployee(ctx *gin.Context) {
	var request Employee

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ParseBindJSONError(err))
		return
	}

	e := Employee{Name: request.Name}

	if err := c.Create(&e).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, e)
}

type employeeUpdateRequest struct {
	Name string `json:"name" binding:"required"`
}

func (c *EmployeeController) updateEmployee(ctx *gin.Context) {
	var updateRequest employeeUpdateRequest

	if err := ctx.ShouldBindJSON(&updateRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ParseBindJSONError(err))
		return
	}

	var e Employee
	if err := c.First(&e, ctx.Param("id")).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Couldn't update, employee not found"})
		return
	}

	e.Name = updateRequest.Name

	if err := c.Model(&e).Updates(e).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, e)
}

func (c *EmployeeController) deleteEmployee(ctx *gin.Context) {
	var grocery Employee

	if err := c.Where("id = ?", ctx.Param("id")).First(&grocery).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Employee not found!"})
		return
	}

	if err := c.Delete(&grocery).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Employee deleted"})
}
