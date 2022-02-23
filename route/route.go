package route

import (
	"address-book-go/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.New()

	// gin 檔案上傳body限制
	r.MaxMultipartMemory = 64 << 20 // 8 MiB

	// Middleware
	//r.Use(middleware.LogRequest())
	//r.Use(middleware.ErrorResponse())

	corsConf := cors.DefaultConfig()
	corsConf.AllowCredentials = true
	corsConf.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"}
	corsConf.AllowHeaders = []string{"Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization", "Bearer", "Accept-Language"}
	corsConf.AllowOriginFunc = config.GetCorsRule
	r.Use(cors.New(corsConf))

	ContactV1(r)

	return r
}
