package initialize

import (
	"github.com/a6slab/shorter_url/global"
	"github.com/a6slab/shorter_url/internal/handler"
	"github.com/gin-gonic/gin"
)

func InitRouter(handler handler.ShorterHandler) *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	return r
}
