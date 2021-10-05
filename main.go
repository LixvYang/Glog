package main

import (
	"glog/model"
	"glog/routes"
)

func main() {
	model.InitDb()

	routes.InitRouter()
}
