package response

import (
	"github.com/labstack/echo/v4"
)

func ListRes(c echo.Context, code int, recipes interface{}) error {
	return c.JSON(
		code,
		struct {
			Recipes interface{} `json:"recipes"`
		}{
			Recipes: recipes,
		},
	)
}
