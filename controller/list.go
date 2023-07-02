package controller

import (
	"github.com/Poul-george/go-render-ap/go-render-api/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

func List(ctx echo.Context) error {
	return response.ListRes(ctx, http.StatusOK, nil)
}
