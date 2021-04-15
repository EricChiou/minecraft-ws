package config

import (
	"database/sql"
	"minecraft-ws/db/mysql"
	"minecraft-ws/po/configpo"
)

// Dao config dao
type Dao struct{}

// GetConfig dao
func (dao *Dao) GetConfig(key string) GetConfigResVo {
	var g mysql.Gooq

	g.SQL.
		Select(configpo.Key, configpo.Value, configpo.Description).
		From(configpo.Table).
		Where(configpo.Key).Eq(key)

	var getConfigRes GetConfigResVo
	g.QueryRow(func(row *sql.Row) error {
		return row.Scan(&getConfigRes.Key, &getConfigRes.Value, &getConfigRes.Description)
	})

	return getConfigRes
}
