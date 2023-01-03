package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"go-messanger/dto"
	"go-messanger/dto/request"
	"go-messanger/service/postgres/provider"
	"net/http"
)

type AccountHandler struct {
	provider *provider.AccountProvider
}

func NewAccountHandler(provider *provider.AccountProvider) *AccountHandler {
	return &AccountHandler{
		provider: provider,
	}
}

func (ah *AccountHandler) HandleRegistration(ctx *gin.Context) {
	var req request.RegisterRequest
	if err := getRequestBody(ctx, &req); err != nil {
		log.Errorf("Error parsing req body: %+v", err)
		sendResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if err := ah.provider.RegisterNewUser(ctx, dto.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}); err != nil {
		if errors.Is(err, provider.ErrEmailNotUnique) || errors.Is(err, provider.ErrUsernameNotUnique) {
			sendResponse(ctx, http.StatusBadRequest, err.Error())
		} else {
			log.Errorf("Error creating new account: %+v", err)
			sendResponse(ctx, http.StatusInternalServerError, err.Error())
		}
	}

	sendResponse(ctx, http.StatusOK, req)
}
