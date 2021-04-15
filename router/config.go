package router

import (
	"minecraft-ws/api/config"
)

func initConfigAPI() {
	groupName := "/config"
	var controller config.Controller

	// Get
	get(groupName+"/status", controller.GetStatus)

}
