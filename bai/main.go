package main

import (
	"demo/component"
	"demo/component/uploadprovider"
	"demo/middleware"
	"demo/modules/cities/CitiesTransport/ginCities"
	"demo/modules/restaurant/restauranttransport/ginrestaurant"
	"demo/upload/transport/ginupload"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	//dsn := "root:123456@tcp(127.0.0.1:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := os.Getenv("DBConnection")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//S3
	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3APIKey := os.Getenv("S3APIKey")
	s3SecretKey := os.Getenv("S3SecretKey")
	s3Domain := os.Getenv("S3Domain")
	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	if err != nil {
		log.Fatalln(err)
	}

	if err := runService(db, s3Provider); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB, upProvider uploadprovider.UploadProvider) error {

	appCtx := component.NewAppContext(db, upProvider)
	r := gin.Default()

	r.Use(middleware.Recover(appCtx))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//upload file
	r.POST("/upload", ginupload.Upload(appCtx))
	// CRUD
	cities := r.Group("/cities")
	{
		cities.POST("", ginCities.CreaCities(appCtx))
		cities.GET("/:id", ginCities.GetCities(appCtx))
		cities.GET("", ginCities.ListRestaurant(appCtx))
		cities.PATCH("/:id", ginCities.UpdateCities(appCtx))
		cities.DELETE("/:id", ginCities.DeleteCities(appCtx))
	}
	restaurants := r.Group("/restaurants")
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))
		restaurants.GET("/:id", ginrestaurant.GetRestaurant(appCtx))
		restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))
		restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurant(appCtx))
		restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))
	}

	return r.Run(":8080")
}

//CREATE TABLE `restaurants` (
//	`id` int(11) NOT NULL AUTO_INCREMENT,
//	`owner_id` int(11) NOT NULL,
//	`name` varchar(50) NOT NULL,
//	`addr` varchar(255) NOT NULL,
//	`city_id` int(11) DEFAULT NULL,
//	`lat` double DEFAULT NULL,
//	`lng` double DEFAULT NULL,rủ
//	`cover` json NOT NULL,
//	`logo` json NOT NULL,
//	`shipping_fee_per_km` double DEFAULT '0',
//	`status` int(11) NOT NULL DEFAULT '1',
//	`created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
//	`updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
//	PRIMARY KEY (`id`),
//	KEY `owner_id` (`owner_id`) USING BTREE,
//	KEY `city_id` (`city_id`) USING BTREE,
//	KEY `status` (`status`) USING BTREE
//) ENGINE=InnoDB DEFAULT CHARSET=utf8;
//
//type Restaurant struct {
//	Id   int    `json:"id" gorm:"column:id;"`
//	Name string `json:"name" gorm:"column:name;"`
//	Addr string `json:"address" gorm:"column:addr;"`
//}
//
//func (Restaurant) TableName() string {
//	return "restaurants"
//}
//
//type RestaurantUpdate struct {
//	Name *string `json:"name" gorm:"column:name;"`
//	Addr *string `json:"address" gorm:"column:addr;"`
//}
//
//func (RestaurantUpdate) TableName() string {
//	return Restaurant{}.TableName()
