package role

import (
	"anteraja/backend/dto"
	"anteraja/backend/repository"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	roleRouter.GET("/:roleId", role.findById)
	roleRouter.POST("/upload", role.uploadFile)

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

func (role RoleRequestHandler) findById(context *gin.Context) {
	roleId, err := strconv.ParseUint(context.Param("roleId"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}
	response, err := role.ctrl.FindById(context, int(roleId))

	if err != nil {
		context.JSON(http.StatusBadRequest, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	context.JSON(http.StatusOK, response)
}

func (role RoleRequestHandler) uploadFile(context *gin.Context) {
	file, err := context.FormFile("file")
	// The file cannot be received.
	if err != nil {
		context.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	extension := filepath.Ext(file.Filename)
	newFileName := uuid.New().String() + extension

	// newDir, err := filepath.Abs(filepath.Dir(os.Args[0]))

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		context.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	// The file is received, so let's save it
	path := dir + "/assets/" + newFileName
	err = context.SaveUploadedFile(file, path)

	if err != nil {
		context.JSON(http.StatusBadRequest, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	context.JSON(http.StatusOK, path)

}
