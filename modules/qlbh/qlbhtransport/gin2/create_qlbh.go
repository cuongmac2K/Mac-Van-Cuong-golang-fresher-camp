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

func CreateQlbh(appCtx compoment.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data qlbhmodel.QlbhCreate
		fmt.Println("da chay ==========")
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := qlbhstorage.NewSQLStore(appCtx.GetmainDBConnection())
		biz := qlbhbiz.NewCreateqlbhbiz(store)
		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {

		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
