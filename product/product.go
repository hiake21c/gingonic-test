package product

import (
	"gingonic-test/database"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

type Product struct {
	gorm.Model
	Code  string `gorm:"unique"`
	Price uint64
}

func GetAllProduct(c *gin.Context) {

	db := database.Db
	var products []Product
	db.Find(&products)

	c.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}

func GetProduct(c *gin.Context) {

	id := c.Param("id")

	db := database.Db

	var product Product
	db.Find(&product, id)

	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})
}

func SaveProduct(c *gin.Context) {
	db := database.Db

	uuid := uuid.New()

	var product Product
	product.Code = uuid.String()
	product.Price = 10

	db.Save(&product)

	c.JSON(http.StatusOK, gin.H{
		"product": &product,
	})
}

func UpdateProduct(c *gin.Context) {
	var param = Product{}
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")

	db := database.Db
	var product Product
	db.Find(&product, id)

	//product.Price, _ = strconv.ParseUint(price, 10 , 32)
	product.Price = param.Price
	db.Save(&product)

	c.JSON(http.StatusOK, gin.H{
		"product": &product,
	})
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	db := database.Db

	var product Product
	db.First(&product, id)
	if product.ID == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   http.StatusInternalServerError,
			"message": "product is not exist",
		})
	} else {
		db.Delete(&product)
		c.JSON(http.StatusOK, gin.H{
			"delete": "true",
		})
	}

}
