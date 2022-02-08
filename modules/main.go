package main

import (
	"awesomeProject/modules/compoment"
	"awesomeProject/modules/qlbh/qlbhtransport/gin2"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"log"
	"net/http"
	"os"

	"gorm.io/gorm"
)

type Nhanvien struct {
	Id       int    `json:"id" gorm:"column:id;"`
	Name     string `json:"Name" gorm:"column:Name;"`
	Age      int    `json:"Age" gorm:"column:Age;"`
	Adress   string `json:"Adress" gorm:"column:Adress;"`
	ThamNien int    `json:"ThamNien" gorm:"column:ThamNien;"`
}

func (Nhanvien) TableName() string {
	return "qlbh"
}

type NhanvienUpdate struct {
	Name   *string `json:"Name" gorm:"column:Name;"`
	Adress *string `json:"Adress" gorm:"column:Adress;"`
}

func (NhanvienUpdate) TableName() string {
	return Nhanvien{}.TableName()
}

func main() {
	//dsn := "root:123456@tcp(127.0.0.1:3306)/qlbh?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := os.Getenv("DatabaseConntext")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("k ket noi dc", err)
	}
	if err := runServer(db); err != nil {
		log.Fatalln("k run dc ", err)
	}

}
func runServer(db *gorm.DB) error {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "test successful",
			})
		})
		//CURD
		appCtx := compoment.NewAppContext(db)
		departments := api.Group("/qlbh")
		{

			departments.POST("", qlbhtransport.CreateQlbh(appCtx))
			departments.GET("/:MaPB", qlbhtransport.GetQLBH(appCtx))
			departments.GET("", qlbhtransport.ListQLBH(appCtx))
			departments.PATCH("/:MaPB", qlbhtransport.UpdateQLBH(appCtx))
			departments.DELETE("/:MapB", qlbhtransport.DeleteQLBH(appCtx))

		}

	}
	return router.Run(":8080")
}
