package mys

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/lianhong2758/RosmBot-MUL/message"
	"github.com/lianhong2758/RosmBot-MUL/rosm"
	"github.com/lianhong2758/RosmBot-MUL/server/mys/mysmsg"
	"github.com/lianhong2758/RosmBot-MUL/tool"
	"github.com/lianhong2758/RosmBot-MUL/tool/web"
	log "github.com/sirupsen/logrus"
)

const (
	sendMessage = "https://bbs-api.miyoushe.com/vila/api/bot/platform/sendMessage"
)

func (c *Config) BotSend(ctx *rosm.CTX, msg ...message.MessageSegment) any {
	msgContentInfo, objectStr := mysmsg.MakeMsgContent(ctx, msg...)
	contentStr, _ := json.Marshal(msgContentInfo)
	data, _ := json.Marshal(H{"room_id": tool.Int64(ctx.Being.RoomID), "object_name": objectStr, "msg_content": tool.BytesToString(contentStr)})
	log.Debugln("[send]", tool.BytesToString(data))
	data, err := web.Web(&http.Client{}, sendMessage, http.MethodPost, makeHeard(ctx), bytes.NewReader(data))
	if err != nil {
		log.Errorln("[send]", err)
	}
	sendState := new(SendState)
	_ = json.Unmarshal(data, sendState)
	log.Infoln("[send]["+sendState.Message+"]", tool.BytesToString(contentStr))
	log.Debugln("[send]", tool.BytesToString(data))
	return sendState
}

func makeHeard(ctx *rosm.CTX) func(req *http.Request) {
	return func(req *http.Request) {
		req.Header.Add("x-rpc-bot_id", ctx.Bot.(*Config).BotToken.BotID)
		req.Header.Add("x-rpc-bot_secret", ctx.Bot.(*Config).BotToken.BotSecret)
		req.Header.Add("x-rpc-bot_villa_id", ctx.Being.RoomID2)
		req.Header.Add("Content-Type", "application/json")
	}
}
