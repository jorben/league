package router

import (
	"embed"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"io/fs"
	"league/common/context"
	"league/common/errs"
	"league/middleware"
	"league/router/api"
	"net/http"
	"time"
)

// SetupRouter 设置路由
func SetupRouter(s *gin.Engine, feEmbed embed.FS) {
	// 健康检查
	s.GET("/health", func(ctx *gin.Context) {
		c := context.CustomContext{Context: ctx}
		c.CJSON(errs.Success)
	})

	backend := s.Group("/api")
	// TODO: 移除跨域支持
	backend.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "HEAD", "PATCH"},
		AllowHeaders: []string{"Origin", "X-Token", "X-Csrf-Token", "X-Request-Id"},
		//AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))
	backend.OPTIONS("/*any", func(ctx *gin.Context) {
		// 处理OPTIONS请求，例如返回204 No Content
		ctx.AbortWithStatus(http.StatusNoContent)
	})

	backend.Use(middleware.RequestId(), middleware.Logger(), middleware.Auth(), gin.Recovery())
	// 权限相关接口
	backend.GET("/auth/login", api.AuthLogin)
	backend.GET("/auth/callback", api.AuthCallback)
	backend.GET("/auth/renew", api.AuthRenew)
	backend.GET("/auth/logout", api.AuthLogout)

	// 菜单相关接口
	backend.GET("/menu", api.GetIndexMenus)

	// 用户相关接口
	backend.GET("/user/current", api.GetUserinfo)

	backendAdmin := backend.Group("/admin")
	backendAdmin.GET("/menu", api.GetAdminMenus)

	s.StaticFS("/static", getFileSystem(feEmbed, "web/build/static"))
	s.NoRoute(func(ctx *gin.Context) {
		// 注意：不能使用FileFromFS的方式获取index.html，底层存在某些Bug
		indexData, _ := feEmbed.ReadFile("web/build/index.html")
		ctx.Data(http.StatusOK, "text/html", indexData)
	})

}

func getFileSystem(embeddedFiles embed.FS, path string) http.FileSystem {
	fileSystem, err := fs.Sub(embeddedFiles, path)
	if err != nil {
		panic(err)
	}
	return http.FS(fileSystem)
}
