package ginCities

import (
	"demo/common"
	"demo/component"
	"demo/modules/cities/citiesbiz"
	"demo/modules/cities/citiesmodel"
	"demo/modules/cities/citiesstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreaCities(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data citiesmodel.CitiesCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := citiesstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := citiesbiz.NewCreateCitiesBiz(store)

		if err := biz.CreateCities(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
