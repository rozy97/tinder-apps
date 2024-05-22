package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rozy97/tinder-apps/user-service/request"
	"github.com/rozy97/tinder-apps/user-service/response"
)

type Handler struct {
	usecase Usecase
}

func NewHandler(usecase Usecase) *Handler {
	return &Handler{usecase: usecase}
}

type Usecase interface {
	Register(ctx context.Context, req request.Register) (*response.Register, error)
	Login(ctx context.Context, req request.Login) (*response.Login, error)
	Activation(ctx context.Context, userID uint64, profile request.Profile) error
}

func (h *Handler) Register(ctx *gin.Context) {
	var req request.Register
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	result, err := h.usecase.Register(ctx.Request.Context(), req)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if result.Code != 0 {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, result)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (h *Handler) Login(ctx *gin.Context) {
	var req request.Login
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	result, err := h.usecase.Login(ctx.Request.Context(), req)
	if err != nil {
		if err.Error() == response.ErrMessageIncorrectPassword {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
		}

		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (h *Handler) Activation(ctx *gin.Context) {
	var req request.Profile
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID, _ := strconv.ParseInt(ctx.Request.Header.Get("x-user-id"), 10, 64)

	if err := h.usecase.Activation(ctx.Request.Context(), uint64(userID), req); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}
