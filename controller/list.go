package controller

import (
	"github.com/labstack/echo/v4"
	"go-render-api/response"
	"net/http"
)

func List(ctx echo.Context) error {
	return response.ListRes(ctx, http.StatusOK, nil)
}
