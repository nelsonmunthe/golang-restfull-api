package userv2

import (
	"anteraja/backend/dto"
	"anteraja/backend/middleware"
	"anteraja/backend/repository"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRequestHandlerV2 struct {
	evoDB *gorm.DB
	ctrl  UserController
}

func NewUserRequestHandler(evoDB *gorm.DB) UserRequestHandlerV2 {
	return UserRequestHandlerV2{
		evoDB: evoDB,
	}
}

func (h UserRequestHandlerV2) HandleUserV2(router *gin.Engine) {
	userRepoV2 := repository.NewAnterajaUserV2(h.evoDB)
	userDetail := repository.NewUserDetail(h.evoDB)
	roleDetail := repository.NewController(h.evoDB)

	userV2Usecase := UserV2Usecase{
		userRepoV2: userRepoV2,
		userDetail: userDetail,
		roleDetail: roleDetail,
	}
	h.ctrl = UserController{
		useCase: userV2Usecase,
	}
	userV2Router := router.Group("/v2/users")

	userV2Router.GET("/detail/:userId", middleware.Authenticate(), h.FindByIdV2)
	userV2Router.GET("/list", middleware.Authenticate(), h.GetListUser)
	userV2Router.DELETE("/delete/:userId", middleware.Authenticate(), h.DeleteUser)
	userV2Router.PUT("/update/:userId", middleware.Authenticate(), h.UpdateUser)
	userV2Router.POST("/create", middleware.Authenticate(), h.CreateUser)
	userV2Router.PUT("/change-status/:userId", middleware.Authenticate(), h.changeStatus)
	userV2Router.POST("/login", h.login)
	userV2Router.POST("/set-position", h.setPosition)
	userV2Router.GET("/location", h.FindByLocation)
}

func (h UserRequestHandlerV2) FindByLocation(context *gin.Context) {
	query := dto.QUeryRequest{}
	err := context.BindQuery(&query)
	fmt.Println("query", query, err)
}

func (h UserRequestHandlerV2) FindByIdV2(context *gin.Context) {
	userId, err := strconv.ParseUint(context.Param("userId"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}
	response, err := h.ctrl.FindByIdV2(context.Request.Context(), uint(userId))
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

func (h UserRequestHandlerV2) GetListUser(context *gin.Context) {
	response, err := h.ctrl.GetListUser(context.Request.Context())
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

func (h UserRequestHandlerV2) DeleteUser(context *gin.Context) {
	userId, err := strconv.ParseUint(context.Param("userId"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	_, err = h.ctrl.FindByIdV2(context.Request.Context(), uint(userId))
	if err != nil {
		context.JSON(http.StatusBadRequest, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	response, err := h.ctrl.DeleteUser(context.Request.Context(), uint(userId))
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

func (h UserRequestHandlerV2) UpdateUser(context *gin.Context) {
	var request RequestUserUpdateUser
	err := context.BindJSON(&request)

	if err != nil {
		context.JSON(http.StatusBadRequest, dto.DefaultErrorResponse())
		return
	}

	userId, err := strconv.ParseUint(context.Param("userId"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	_, err = h.ctrl.FindByIdV2(context.Request.Context(), uint(userId))
	if err != nil {
		context.JSON(http.StatusBadRequest, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	response, err := h.ctrl.UpdateUser(context.Request.Context(), uint(userId), request)
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

func (h UserRequestHandlerV2) CreateUser(context *gin.Context) {
	var request RequestUser

	err := context.BindJSON(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, dto.DefaultErrorResponse())
		return
	}

	response, err := h.ctrl.CreateUser(context.Request.Context(), request)

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

func (h UserRequestHandlerV2) changeStatus(context *gin.Context) {
	var request RequestUserUpdateStatus
	err := context.BindJSON(&request)

	if err != nil {
		context.JSON(http.StatusBadRequest, dto.DefaultErrorResponse())
		return
	}

	userId, err := strconv.ParseUint(context.Param("userId"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	response, err := h.ctrl.changeStatus(context.Request.Context(), uint(userId), request)
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

func (h UserRequestHandlerV2) login(context *gin.Context) {
	var request RequestUserLogin
	err := context.BindJSON(&request)

	if err != nil {
		context.JSON(http.StatusBadRequest, dto.DefaultErrorResponse())
		return
	}

	response, err := h.ctrl.login(context.Request.Context(), request)
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

func (h UserRequestHandlerV2) setPosition(context *gin.Context) {
	var request RequestSetPosition
	err := context.BindJSON(&request)

	if err != nil {
		context.JSON(http.StatusBadRequest, dto.DefaultErrorResponse())
		return
	}

	response, err := h.ctrl.SetPosition(context.Request.Context(), request)
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
