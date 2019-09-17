package public

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

func QueryParamToUint(c echo.Context, query_key string) uint {
	query_param_int, _ := strconv.Atoi(c.QueryParam(query_key))
	return uint(query_param_int)
}
