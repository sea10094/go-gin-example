package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/EDDYCJY/go-gin-example/docs"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	//"github.com/EDDYCJY/go-gin-example/middleware/jwt"
	"github.com/EDDYCJY/go-gin-example/pkg/export"
	"github.com/EDDYCJY/go-gin-example/pkg/qrcode"
	"github.com/EDDYCJY/go-gin-example/pkg/upload"
	//"github.com/EDDYCJY/go-gin-example/routers/api"
	"github.com/EDDYCJY/go-gin-example/routers/api/v1"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	//r.POST("/auth", api.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//r.POST("/upload", api.UploadImage)

	apiv1 := r.Group("/api/v1")
	//apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/getAgentList", v1.GetAgentList)
		apiv1.GET("/getTaskList", v1.GetTaskList)
		apiv1.GET("/getTaskAttr", v1.GetTaskAttr)
		//apiv1.GET("/startScan", v1.StartScan)
		apiv1.GET("/getAgentMonitor", v1.GetAgentMonitor)
		apiv1.GET("/rebootAgent", v1.RebootAgent)
		//apiv1.POST("/api/v1/StopScan", v1.StopScan)
		apiv1.GET("/getProgress", v1.GetProgress)
		apiv1.GET("/assetSearch", v1.AssetSearch)
		apiv1.GET("/getCommands", v1.GetCommands)
		//apiv1.POST("/api/v1/GetStatisticsData", v1.GetStatisticsData)


	}

	return r
}
