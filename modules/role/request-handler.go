package role

import (
	"anteraja/backend/dto"
	"anteraja/backend/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RoleRequestHandler struct {
	db   *gorm.DB
	ctrl RoleController
}

func NewRoleRequestHandler(db *gorm.DB) RoleRequestHandler {
	return RoleRequestHandler{
		db: db,
	}
}

func (role RoleRequestHandler) HandleRole(router *gin.Engine) {
	roleRepo := repository.NewRole(role.db)
	roleUsecase := RoleUsecase{
		roleRepo: roleRepo,
	}
	role.ctrl = RoleController{
		usecase: roleUsecase,
	}

	roleRouter := router.Group("/role")
	roleRouter.GET("/list", role.getList)

}

func (role RoleRequestHandler) getList(context *gin.Context) {
	query := dto.QUeryRequest{}
	err := context.BindQuery(&query)

	if err != nil {
		context.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	response, err := role.ctrl.GetList(context.Request.Context(), query)
	if response.Success == false {
		context.JSON(http.StatusNotFound, response)
		return
	}
	if err != nil {
		context.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}
	context.JSON(http.StatusOK, response)
}
