package deposit

import (
	"anteraja/backend/dto"
	"anteraja/backend/repository"
	"fmt"
	"net/http"

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
	depositUsecase := DepositUsecase{
		depositRepo: depositRepo,
	}
	deposit.ctrl = DepositController{
		usecase: depositUsecase,
	}

	depositRouter := router.Group("/deposit")

	depositRouter.GET("/list", deposit.GetList)
}

func (deposit DepositHandler) GetList(context *gin.Context) {
	query := dto.RequestDeposit{}
	err := context.BindQuery(&query)

	if err != nil {
		context.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	pagination := dto.PaginationRequest{}

	if pagination.Page <= 0 {
		pagination.Page = int(1)
	}

	if pagination.PerPage <= 0 {
		pagination.PerPage = int(10)
	}
	fmt.Println("pagination", pagination)
	// // err := context.BindQuery(&pagination)

	if err != nil {
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
