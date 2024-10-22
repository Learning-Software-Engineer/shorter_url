package initialize

import (
	"fmt"

	"github.com/a6slab/shorter_url/global"
)

func Run() {
	Configuration()
	r := InitRouter()
	serverAddr := fmt.Sprintf(":%v", global.Config.Server.Port)
	if global.Config.Server.Mode != "release" {
		fmt.Println(serverAddr)
	}
	r.Run(serverAddr)
}
