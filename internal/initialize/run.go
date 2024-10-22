package initialize

import (
	"fmt"

	"github.com/a6slab/shorter_url/global"
	"github.com/a6slab/shorter_url/internal/handler"
	"github.com/a6slab/shorter_url/internal/repository"
)

func Run() {
	Configuration()
	session := InitCassandra()
	repo := repository.NewShorterURLRepo(session)
	handler := handler.NewShorterHandler(&repo)

	r := InitRouter(handler)
	serverAddr := fmt.Sprintf(":%v", global.Config.Server.Port)
	if global.Config.Server.Mode != "release" {
		fmt.Println(serverAddr)
	}

	r.GET("/ping", handler.HealthCheck)
	r.POST("/shorten", handler.ShortenURL)
	r.GET("/:short_url_key", handler.GetRedirectURL)

	r.Run(serverAddr)
}
