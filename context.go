package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Deparment struct {
	MaPB   string `json:"MaPB,omitempty" gorm:"column:MaPB;"`
	HoTen  string `json:"HoTen" gorm:"column:HoTen;"`
	DiaChi string `json:"DiaChi" gorm:"column:DiaChi;"`
	SDT    string `json:"SDT" gorm:"column:SDT;"`
}
type DeparmentUpdate struct {
	MaPB   string `json:"MaPB,omitempty" gorm:"column:MaPB;"`
	HoTen  string `json:"HoTen" gorm:"column:HoTen;"`
	DiaChi string `json:"DiaChi" gorm:"column:DiaChi;"`
	SDT    string `json:"SDT" gorm:"column:SDT;"`
}

func (Deparment) TableName() string {
	return "deparment"
}
func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/customer?charset=utf8mb4&parseTime=True&loc=Local"

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
		departments := api.Group("/departments")
		{

			departments.POST("", func(c *gin.Context) {
				var data Deparment

				if err := c.ShouldBind(&data); err != nil {
					c.JSON(401, gin.H{
						"error": err.Error(),
					})

					return
				}

				if err := db.Create(&data).Error; err != nil {
					c.JSON(401, gin.H{
						"error": err.Error(),
					})

					return
				}

				c.JSON(http.StatusOK, data)
			})

			departments.GET("/:MaPB", func(c *gin.Context) {
				MaPB, err := strconv.Atoi(c.Param("MaPB"))

				if err != nil {
					c.JSON(401, gin.H{
						"error": err.Error(),
					})

					return
				}

				var data Deparment

				if err := db.Where("MaPB = ?", MaPB).First(&data).Error; err != nil {
					c.JSON(401, gin.H{
						"error": err.Error(),
					})

					return
				}

				c.JSON(http.StatusOK, data)
			})

			departments.GET("", func(c *gin.Context) {
				var data []Deparment

				type Filter struct {
					DateOfBirth string `json:"dateOfBirth" gorm:"column:dateOfBirth;"`
				}

				var filter Filter

				c.ShouldBind(&filter)

				newDb := db

				if filter.DateOfBirth == "" {
					newDb = db.Where("status = ?", filter.DateOfBirth)
				}

				if err := newDb.Find(&data).Error; err != nil {
					c.JSON(401, gin.H{
						"error": err.Error(),
					})

					return
				}

				c.JSON(http.StatusOK, data)
			})

			departments.PATCH("/:MaPB", func(c *gin.Context) {
				MaPB, err := strconv.Atoi(c.Param("MaPB"))

				if err != nil {
					c.JSON(401, gin.H{
						"error": err.Error(),
					})
					return
				}
				var data Deparment

				if err := c.ShouldBind(&data); err != nil {
					c.JSON(401, map[string]interface{}{
						"error": err.Error(),
					})

					return
				}

				if err := db.Where("MaPB = ?", MaPB).Updates(data).Error; err != nil {
					c.JSON(401, map[string]interface{}{
						"error": "Update failed",
					})

					return
				}

				c.JSON(http.StatusOK, map[string]interface{}{
					"update successfully": 1,
				})
			})
			departments.DELETE("/:MapB", func(c *gin.Context) {
				MaPB, err := strconv.Atoi(c.Param("MaPB"))

				if err != nil {
					c.JSON(401, map[string]interface{}{
						"error": err.Error(),
					})

					return
				}

				if err := db.Table(Deparment{}.TableName()).Where("MaPB = ?", MaPB).Delete(nil).Error; err != nil {
					c.JSON(401, map[string]interface{}{
						"error": err.Error(),
					})

					return
				}

				c.JSON(200, map[string]interface{}{"ok": 1})
			})
		}

	}
	return router.Run(":8080")
}
