package handler

import (
	"service-transaction/auth"
	"service-transaction/core"
)

type transactionHandler struct {
	userService core.Service
	authService auth.Service
}

func NewTransactionHandler(transactionService core.Service, authService auth.Service) *transactionHandler {
	return &transactionHandler{transactionService, authService}
}

// func (h *userInvestorHandler) CreateTransaction(c *gin.Context) {
