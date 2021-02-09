package routers

import (
	"gin-blog/pkg/setting"
	"gin-blog/routers/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	api_group := r.Group("/api")
	//获取标签列表
	api_group.GET("/tags", api.GetTags)
	//新建标签
	api_group.POST("/tags", api.AddTag)
	//更新指定标签
	api_group.PUT("/tags/:id", api.EditTag)
	//删除指定标签
	api_group.DELETE("/tags/:id", api.DeleteTag)

	//获取文章列表
	api_group.GET("/articles", api.GetArticles)
	//获取指定文章
	api_group.GET("/articles/:id", api.GetArticle)
	//新建文章
	api_group.POST("/articles", api.AddArticle)
	//更新指定文章
	api_group.PUT("/articles/:id", api.EditArticle)
	//删除指定文章
	api_group.DELETE("/articles/:id", api.DeleteArticle)

	return r

}
