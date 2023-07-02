package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"test-go-api/response"
)

func List(ctx echo.Context) error {
	return response.ListRes(ctx, http.StatusOK, nil)
}
