package ob11

import (
	"context"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lianhong2758/RosmBot-MUL/rosm"
	"github.com/lianhong2758/RosmBot-MUL/server/ob11/ent"
	"github.com/wdvxdr1123/ZeroBot/driver"
)

func (config *Config) Run() {
	config.mul()
	switch config.Types {
	case "WS":
		config.Driver = driver.NewWebSocketClient(config.URL, config.Token)
	case "WSS":
		config.Driver = driver.NewWebSocketServer(16, config.URL, config.Token)
	}
	botMap[config.BotID] = config
	config.Driver.Connect()
	db, err := ent.Open(config.DbDriver, config.Dsn)
	if err != nil {
		log.Fatalf("[ob11]连接数据库失败: %v", err)
		return
	}
	db = db.Debug()
	if err := db.Schema.Create(context.Background()); err != nil {
		log.Fatalf("create db Schema failed!err:=%v", err)
		return
	}
	config.Db = db

	config.Driver.Listen(config.processEvent())

}

func (c *Config) mul() {
	rosm.MULChan <- rosm.MUL{Types: "ob11", Name: c.BotName, BotID: c.BotID}
}
