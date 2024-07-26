package router

import (
	"embed"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
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
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "HEAD", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "X-Token", "X-Csrf-Token", "X-Request-Id"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	backend.OPTIONS("/*any", func(ctx *gin.Context) {
		// 处理OPTIONS请求，例如返回204 No Content
		ctx.AbortWithStatus(http.StatusNoContent)
	})

	backend.Use(gin.Recovery(), middleware.RequestId(), middleware.Logger(), middleware.Auth())
	// 权限相关接口
	backend.GET("/auth/login", api.AuthLogin)
	backend.GET("/auth/callback", api.AuthCallback)
	backend.GET("/auth/renew", api.AuthRenew)
	backend.GET("/auth/logout", api.AuthLogout)

	// 菜单相关接口
	backend.GET("/menu", api.MenuIndex)

	// 用户相关接口
	backend.GET("/user/current", api.UserCurrent)

	// 管理后台相关接口
	backendAdmin := backend.Group("/admin")
	backendAdmin.GET("/menu", api.MenuAdmin)
	backendAdmin.GET("/user/list", api.UserList)
	backendAdmin.GET("/user/detail", api.UserDetail)
	backendAdmin.POST("/user/status", api.UserStatus)
	backendAdmin.POST("/user/unbind", api.UserUnbind)
	backendAdmin.POST("/user/delete", api.UserDelete)
	backendAdmin.POST("/user/join_group", api.UserJoinGroup)
	backendAdmin.POST("/user/exit_group", api.UserExitGroup)
	backendAdmin.GET("/group/list", api.UserGroupList)
	backendAdmin.GET("/setting/apilist", api.SettingApiList)
	backendAdmin.POST("/setting/api", api.SettingUpdateApi)
	backendAdmin.POST("/setting/api/delete", api.SettingDeleteApi)
	backendAdmin.GET("/setting/menulist", api.SettingMenuList)
	backendAdmin.GET("/auth/policylist", api.AuthPolicyList)
	backendAdmin.POST("/auth/policy", api.AuthUpdatePolicy)
	backendAdmin.POST("/auth/policy/delete", api.AuthDeletePolicy)

	s.Use(gzip.Gzip(gzip.DefaultCompression)).StaticFS("/static", getFileSystem(feEmbed, "web/build/static"))
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
