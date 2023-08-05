package api

import (
	"everything-is-twenty/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (server *Server) listProducts(c *gin.Context) {
	var product model.Product

	if result := server.db.Find(&product); result.Error != nil {
		log.Fatal(result.Error)
	}

	log.Printf("api/products.go - get AllProducts %s", product)

	c.JSON(http.StatusOK, gin.H{"body": product, "status": "ok"})
}

func (server *Server) getProduct(c *gin.Context) {
	var product model.Product
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Fatal(err)
	}

	if result := server.db.Find(&product, id); result.Error != nil {
		log.Fatal(result.Error)
	}

	log.Printf("api/products.go - get product by id = %d and product  %+v", id, product)

	c.JSON(http.StatusOK, gin.H{"body": product, "status": "ok"})
}

func (server *Server) createProduct(c *gin.Context) {
	var product model.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Can not bind body with %s", err.Error())})
		return
	}

	if result := server.db.Create(&product); result.Error != nil {
		log.Println(result.Error.Error())
		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Can't create a product with %s", result.Error.Error())})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"body": product, "status": "created"})
}
