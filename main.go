package main

import (
	"Food-delivery/component/appctx"
	"Food-delivery/component/uploadprovider"
	"Food-delivery/middleware"
	pblocal "Food-delivery/pubsub/localpubsub"
	ginroutes "Food-delivery/routes"
	"Food-delivery/skio"
	"Food-delivery/subscriber"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// test := Restaurant{
	// 	Id:   1,
	// 	Name: "200lab",
	// 	Addr: "somewhere",
	// }
	// jsByte, err := json.Marshal(test)
	// log.Println(string(jsByte), err)
	// json.Unmarshal([]byte("{\"id\":1,\"name\":\"200lab\",\"addr\":\"somewhere\"}"), &test)
	// log.Println(test)

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// s3BucketName := os.Getenv("S3BucketName")
	// s3Region := os.Getenv("S3Region")
	// s3APIKey := os.Getenv("S3APIKey")
	// s3SecretKey := os.Getenv("S3SecretKey")
	// s3Domain := os.Getenv("S3Domain")set s3BucketName=g05-images-food-delyvery
	s3BucketName := "g05-images-food-delyvery"
	s3Region := "ap-southeast-1"
	s3APIKey := "AKIA5OCFQCU32MTW4FII"
	s3SecretKey := "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"
	s3Domain := ""
	secretKey := "CristianoRonaldo"

	dsn := "food_delivery_g06:19e5a718a54a9fe0559dfbce6908@tcp(127.0.0.1:3306)/food_delivery_g06?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)
	ps := pblocal.NewPubsub()
	appContext := appctx.NewAppContext(db, s3Provider, secretKey, ps)
	// subscriber.Setup(appContext)
	if err := subscriber.NewEngine(appContext).Start(); err != nil {
		log.Println(err)
	}
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	//middleware
	r.Use(middleware.Recover(appContext))
	r.GET("/ping", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user.html", gin.H{
			"message": "pong",
		})
	})
	// static upload file
	r.Static("static", "./static")
	r.StaticFile("/demo/", "./demo.html")
	// viết hàm POST

	v1 := r.Group("/v1")
	ginroutes.SettupRoutesAdmin(appContext, v1)
	ginroutes.SetupRoutes(appContext, v1)
	ginroutes.SetupRoutesLike(appContext, v1)
	ginroutes.SetupRoutesFood(appContext, v1)
	ginroutes.SetupRoutesLikeFood(appContext, v1)
	ginroutes.SetupRoutesRatingRestaurant(appContext, v1)

	rtEngine := skio.NewEngine()
	appContext.SetRealtimeEngine(rtEngine)
	_ = rtEngine.Run(appContext, r)

	r.Run()
}
