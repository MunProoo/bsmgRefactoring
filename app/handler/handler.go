package handler

import (
	"BsmgRefactoring/usecase"
	"fmt"

	"github.com/labstack/echo/v4"
)

type BsmgHandler struct {
	bmUsecase usecase.BsmgUsecase
}

func NewBsmgHandler(bu usecase.BsmgUsecase) *BsmgHandler {
	return &BsmgHandler{bu}
}

func (h *BsmgHandler) SelectUserList(c *echo.Context) {
	userList, err := h.bmUsecase.SelectUserList()
	fmt.Println(userList, err)
}
