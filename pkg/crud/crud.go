package crud

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"leave.gg/pkg"
)

type CrudController struct {
	DB *gorm.DB
}

func (c *CrudController) AutoMigrate(entity interface{}) {
	fmt.Printf("Auto Migrating %s...\n", deriveEntityName(entity))
	c.DB.AutoMigrate(entity)
}

func (c *CrudController) Bind(ctx *gin.Context, request interface{}) (ok bool) {
	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ParseBindJSONError(err))
		return false
	}

	return true
}

func (c *CrudController) FirstBy(ctx *gin.Context, entity interface{}, param string) (ok bool) {
	return c.firstBy(ctx, entity, param, true)
}

// FirstBySuprresed doesn't write success status, can then be used by something like update recipe
func (c *CrudController) FirstBySuppresed(ctx *gin.Context, entity interface{}, param string) (ok bool) {
	return c.firstBy(ctx, entity, param, false)
}

func (c *CrudController) firstBy(ctx *gin.Context, entity interface{}, param string, log bool) (ok bool) {
	id, _ := strconv.Atoi(ctx.Param(param))

	if err := c.DB.First(entity, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": deriveEntityName(entity) + " not found."})
		return false
	}
	if log {
		ctx.JSON(http.StatusOK, entity)
	}
	return true
}

// All expects the "entities" to a slice
func (c *CrudController) All(ctx *gin.Context, entities interface{}) (ok bool) {
	if err := c.DB.Find(entities).Error; err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return false
	}

	ctx.JSON(http.StatusOK, entities)
	return true
}

func (c *CrudController) Create(ctx *gin.Context, entity interface{}) (ok bool) {
	if err := c.DB.Create(entity).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return false
	}

	ctx.JSON(http.StatusOK, entity)
	return true
}

func (c *CrudController) Update(ctx *gin.Context, entity interface{}) (ok bool) {
	if err := c.DB.Model(entity).Updates(entity).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return false
	}

	ctx.JSON(http.StatusOK, entity)
	return true
}

func (c *CrudController) Delete(ctx *gin.Context, entity interface{}) (ok bool) {
	if err := c.DB.Delete(entity).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return false
	}

	ctx.JSON(http.StatusOK, gin.H{"message": deriveEntityName(entity) + " deleted."})

	return true
}

func deriveEntityName(entity interface{}) string {
	return reflect.TypeOf(entity).Elem().Name()
}
