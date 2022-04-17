package routers

import (
	_ "blog-service/docs"
	"blog-service/global"
	"blog-service/internal/middleware"
	"blog-service/internal/routers/api"
	v1 "blog-service/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(),
		gin.Recovery(),
		middleware.Translations(),
	)

	article := v1.NewArticle()
	tag := v1.NewTag()

	// swagger 管理器的url
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 上传文件的路由
	upload := api.NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	//提供静态资源的访问
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	// auth验证路由
	r.GET("/auth", api.GetAuth)
	apiv1 := r.Group("/api/v1")
	{
		//针对标签管理的操作
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List)

		//针对文章管理的操作
		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)
	}
	return r
}
