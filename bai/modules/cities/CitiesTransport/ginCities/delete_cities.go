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

func DeleteCities(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := citiesstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := citiesbiz.NewDeleteCitiestBiz(store)

		if err := biz.DeleteCitiest(c.Request.Context(), id); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
