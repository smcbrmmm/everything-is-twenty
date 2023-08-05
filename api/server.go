package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

type Server struct {
	router *gin.Engine
	db     *gorm.DB
}

func NewSever(db *gorm.DB) *Server {
	server := &Server{db: db}

	router := gin.Default()
	router.ForwardedByClientIP = true
	err := router.SetTrustedProxies([]string{"127.0.0.1"})

	if err != nil {
		log.Fatal("Can't trust proxies at address 127.0.0.1")
	}

	router.GET("/products", server.listProducts)
	router.GET("/products/:id", server.getProduct)
	router.POST("/product", server.createProduct)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
