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

func ListRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var fileter citiesmodel.Filter
		if err := c.ShouldBind(&fileter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.Fulfill()

		store := citiesstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := citiesbiz.NewListCities(store)

		result, err := biz.LisCities(c.Request.Context(), &fileter, &paging)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, fileter))
	}
}
