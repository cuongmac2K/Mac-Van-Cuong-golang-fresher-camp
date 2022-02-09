package qlbhtransport

import (
	"awesomeProject/modules/common"
	"awesomeProject/modules/compoment"
	"awesomeProject/modules/qlbh/qlbhbiz"
	"awesomeProject/modules/qlbh/qlbhmodel"
	"awesomeProject/modules/qlbh/qlbhstorage"
	_ "errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListQLBH(appCtx compoment.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter qlbhmodel.Filter
		fmt.Println("da chay ==========")
		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}
		paging.Fulfill()
		store := qlbhstorage.NewSQLStore(appCtx.GetmainDBConnection())
		biz := qlbhbiz.NewlistQLBHBiz(store)
		result, err := biz.ListQLBH(c.Request.Context(), &filter, &paging)

		if err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
