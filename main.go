package main

import (
	_ "v-lottery/boot"
	_ "v-lottery/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
