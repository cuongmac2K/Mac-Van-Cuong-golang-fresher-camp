package ginCities

import (
	"demo/common"
	"demo/component"
	"demo/modules/cities/citiesbiz"
	"demo/modules/cities/citiesstorage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetCities(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := citiesstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := citiesbiz.NewgGetCities(store)

		data, err := biz.GetCities(c.Request.Context(), id)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
