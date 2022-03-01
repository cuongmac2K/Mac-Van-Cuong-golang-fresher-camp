package ginCities

import (
	"demo/common"
	"demo/component"
	"demo/modules/cities/citiesbiz"
	"demo/modules/cities/citiesmodel"
	"demo/modules/cities/citiesstorage"

	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UpdateCities(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data citiesmodel.CitiesUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := citiesstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := citiesbiz.NewUpdateCitiesBiz(store)

		if err := biz.UpdateCities(c.Request.Context(), id, &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
