package config

import (
	"minecraft-ws/constants/configkey"
	"minecraft-ws/service/apihandler"

	"github.com/EricChiou/httprouter"
)

// Controller config controller
type Controller struct {
	dao Dao
}

// GetStatus by GetConfig dao
func (con *Controller) GetStatus(ctx *httprouter.Context) apihandler.ResponseEntity {

	result := con.dao.GetConfig(configkey.Status)

	return responseEntity.OK(ctx.Ctx, result)
}
