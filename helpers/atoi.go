package helpers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func ConvertStringToInt(params *gin.Params, val string) (int, error) {
	paramsString := params.ByName(val)
	valInt, err := strconv.Atoi(paramsString)
	if err != nil {
		return 0, err
	}
	return valInt, nil
}
