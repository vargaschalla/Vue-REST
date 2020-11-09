package routers

import (
	"github.com/202lp2/go2/apis"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupRouter() *gin.Engine {

	conn, err := connectDB()
	if err != nil {
		panic("failed to connect database: " + err.Error())
		//return
	}

	r := gin.Default()
	r.Use(cors.Default())
	//config := cors.DefaultConfig()
	//config.AllowOrigins = []string{"http://localhost", "http://localhost:8082"}

	//r.Use(cors.New(config))

	r.Use(dbMiddleware(*conn))

	//r.Run("localhost:8085")

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", apis.ItemsIndex)
		v1.GET("/persons", apis.PersonsIndex)
		v1.POST("/persons", apis.PersonsCreate)
	}
	return r
}

func connectDB() (c *gorm.DB, err error) {

	dsn := "root:aracelybriguit@tcp(localhost:3306)/integrator2?charset=utf8mb4&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	return conn, err
}

func dbMiddleware(conn gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", conn)
		c.Next()
	}
}
