package main

import (
	"github.com/th1nk-er/fast-get/server/routers"
)

func main() {
	r := routers.InitRouters()
	_ = r.Run(":8080")
}
