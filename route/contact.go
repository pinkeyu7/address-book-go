package route

import (
	apiV1 "address-book-go/api/v1"
	"github.com/gin-gonic/gin"
)

func ContactV1(r *gin.Engine) {
	v1Auth := r.Group("/v1/contacts")
	//v1Auth.Use(middleware.TokenAuth())

	v1Auth.GET("/", func(c *gin.Context) {
		apiV1.ListContact(c)
	})
}
