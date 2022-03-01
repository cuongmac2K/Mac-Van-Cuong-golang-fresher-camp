package ginrestaurantlike

import (
	"demo/common"
	"demo/component"
	rslikebiz "demo/modules/restaurantlike/biz"
	restaurantlikemodel "demo/modules/restaurantlike/model"
	rsstorage "demo/modules/restaurantlike/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

//Get /v1/restaurant/:id/liked-users
func ListUser(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter restaurantlikemodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := rsstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := rslikebiz.NewListUserRestaurantBiz(store)

		result, err := biz.ListUser(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}
		for i := range result {
			result[i].Mask(false)
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
