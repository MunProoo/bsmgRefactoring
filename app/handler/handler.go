package handler

import (
	"BsmgRefactoring/usecase"
)

type BsmgHandler struct {
	uc usecase.BsmgUsecase
}

func NewBsmgHandler(bu usecase.BsmgUsecase) *BsmgHandler {
	return &BsmgHandler{bu}
}
