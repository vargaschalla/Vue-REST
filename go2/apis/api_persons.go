package apis

import (
	"net/http"

	"github.com/202lp2/go2/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//CRUD for items table
func PersonsIndex(c *gin.Context) {
	var lis []models.Person

	db, _ := c.Get("db")
	conn := db.(gorm.DB)

	// Migrate the schema
	conn.AutoMigrate(&models.Person{})

	conn.Find(&lis)
	c.JSON(http.StatusOK, gin.H{
		"msg": "thank you",
		"r":   lis,
	})

}

func PersonsCreate(c *gin.Context) {
	db, _ := c.Get("db")

	conn := db.(gorm.DB)

	//var d Person
	d := models.Person{Name: c.PostForm("name"), Age: c.PostForm("age"), ApPaterno: c.PostForm("apPaterno"),
		ApMaterno: c.PostForm("apMaterno"), EstadoCivil: c.PostForm("estadoCivil")}
	if err := c.BindJSON(&d); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Create(&d)
	c.JSON(http.StatusOK, &d)
}
