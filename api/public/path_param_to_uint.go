package public

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

func PathParamToUint(c echo.Context, path_key string) uint {
	path_param_int, _ := strconv.Atoi(c.Param(path_key))
	return uint(path_param_int)
}
