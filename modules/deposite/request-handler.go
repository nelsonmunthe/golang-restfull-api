package deposit

import (
	"anteraja/backend/dto"
	"anteraja/backend/repository"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DepositHandler struct {
	db   *gorm.DB
	ctrl DepositController
}

func NewRequestDepositHandler(db *gorm.DB) DepositHandler {
	return DepositHandler{
		db: db,
	}
}

func (deposit DepositHandler) HandleDeposit(router *gin.Engine) {
	depositRepo := repository.NewDeposit(deposit.db)
	userDetal := repository.NewUserDetail(deposit.db)
	depositUsecase := DepositUsecase{
		depositRepo: depositRepo,
		userDetail:  userDetal,
	}
	deposit.ctrl = DepositController{
		usecase: depositUsecase,
	}

	depositRouter := router.Group("/deposit")

	depositRouter.GET("/list", deposit.GetList)
	depositRouter.GET("/filter", deposit.FilterDeposit)
	depositRouter.GET("/location", deposit.FilterUserByLocation)
}

func (deposit DepositHandler) GetList(context *gin.Context) {
	query := dto.RequestDeposit{}
	err := context.BindQuery(&query)

	if err != nil {
		context.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	var pagination dto.PaginationRequest

	if err := context.ShouldBindQuery(&pagination); err != nil {
		context.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	if query.Subsidiary_id == "" {
		context.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage("Subsidiary_id is mandatory"))
		return
	}

	response, err := deposit.ctrl.GetList(context, query, pagination)
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

func (deposit DepositHandler) FilterDeposit(context *gin.Context) {
	queryDeposit := dto.RequestFilterDeposit{}
	err := context.BindQuery(&queryDeposit)

	if err != nil {
		context.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	if queryDeposit.Subsidiary_id == "" {
		context.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage("Subsidiary_id can't be empty"))
		return
	}

	var filtering string

	values := reflect.ValueOf(queryDeposit)

	for index := 0; index < values.NumField(); index++ {
		field := values.Field(index)
		if field.Interface() == true {
			filtering = values.Type().Field(index).Tag.Get("json")
		}
	}

	response, err := deposit.ctrl.FilterDeposit(context, queryDeposit.Subsidiary_id, filtering)

	if err != nil {
		context.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	if response.Success == false {
		context.JSON(http.StatusNotFound, response)
		return
	}

	context.JSON(http.StatusOK, response)
}

func (deposit DepositHandler) FilterUserByLocation(context *gin.Context) {
	queryDepositLocation := dto.UserDetailLocationEntity{}
	err := context.BindQuery(&queryDepositLocation)

	if err != nil {
		context.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	type Keywords struct {
		Keyword string `form:"keyword" json:"keyword"`
	}

	keywords := Keywords{}

	if err := context.ShouldBindQuery(&keywords); err != nil {
		context.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	if keywords.Keyword == "" || queryDepositLocation.Subsidiary_id == "" {
		context.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage("subsidiary_id can't be empty"))
		return
	}

	response, err := deposit.ctrl.FilterUserByLocation(context, queryDepositLocation, keywords.Keyword)

	if err != nil {
		context.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	if response.Success == false {
		context.JSON(http.StatusNotFound, response)
		return
	}

	context.JSON(http.StatusOK, response)
}
