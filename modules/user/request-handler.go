package user

import (
	"anteraja/backend/dto"
	"anteraja/backend/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RequestHandler struct {
	evoDB *gorm.DB
	ctrl  Controller
}

func NewRequestHandler(evoDB *gorm.DB) RequestHandler {
	return RequestHandler{
		evoDB: evoDB,
	}
}

func (h RequestHandler) Handle(router *gin.Engine) {
	userRepo := repository.NewAnterajaUser(h.evoDB)
	UserUc := UseCase{
		userRepo: userRepo,
	}
	h.ctrl = Controller{
		useCase: UserUc,
	}
	router.GET("/user-detail", h.reparationProduct)
}

func (h RequestHandler) reparationProduct(c *gin.Context) {
	var request RequestBackfillReparation

	res, err := h.ctrl.backfillReparationProduct(c.Request.Context(), request, 1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorInvalidDataWithMessage(err.Error()))
		return
	}

	c.JSON(http.StatusOK, res)
}
