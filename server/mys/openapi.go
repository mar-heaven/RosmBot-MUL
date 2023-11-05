package mys

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/lianhong2758/RosmBot-MUL/rosm"
	"github.com/lianhong2758/RosmBot-MUL/tool"
	"github.com/lianhong2758/RosmBot-MUL/tool/web"
	log "github.com/sirupsen/logrus"
)

const (
	host           = "https://bbs-api.miyoushe.com"
	getRoomListURL = "/vila/api/bot/platform/getVillaGroupRoomList"
	getUserDataURL = "/vila/api/bot/platform/getMember"
	recallURL      = "/vila/api/bot/platform/recallMessage"
	deleteUserURL  = "/vila/api/bot/platform/deleteVillaMember"
	getVillaURL    = "/vila/api/bot/platform/getVilla"
)

// 获取房间列表
func GetRoomList(ctx *rosm.CTX) (r *RoomList, err error) {
	data, err := web.Web(&http.Client{}, host+getRoomListURL, http.MethodGet, makeHeard(ctx), nil)
	log.Debugln("[GetRoomList]", string(data))
	if err != nil {
		return nil, err
	}
	r = new(RoomList)
	err = json.Unmarshal(data, r)
	return
}

// 获取用户信息
func GetUserData(ctx *rosm.CTX, uid string) (r *UserData, err error) {
	data, _ := json.Marshal(H{"uid": uid})
	data, err = web.Web(&http.Client{}, host+getUserDataURL, http.MethodGet, makeHeard(ctx), bytes.NewReader(data))
	log.Debugln("[GetUserData]", string(data))
	if err != nil {
		return nil, err
	}
	r = new(UserData)
	err = json.Unmarshal(data, r)
	return
}

// 踢人
func DeleteUser(ctx *rosm.CTX, uid string) (err error) {
	data, _ := json.Marshal(H{"uid": uid})
	data, err = web.Web(&http.Client{}, host+deleteUserURL, http.MethodPost, makeHeard(ctx), bytes.NewReader(data))
	log.Debugln("[DeleteUser]", string(data))
	var r ApiCode
	_ = json.Unmarshal(data, &r)
	if r.Retcode != 0 {
		return errors.New(r.Message)
	}
	return
}

// 撤回消息,消息id,房间id,发送时间
func Recall(ctx *rosm.CTX, msgid string, msgtime int64, roomid string) (err error) {
	data, _ := json.Marshal(H{"msg_uid": msgid, "room_id": tool.Int64(roomid), "msg_time": msgtime})
	data, err = web.Web(&http.Client{}, host+recallURL, http.MethodPost, makeHeard(ctx), bytes.NewReader(data))
	log.Debugln("[Recall]", string(data))
	var r ApiCode
	_ = json.Unmarshal(data, &r)
	if r.Retcode != 0 {
		return errors.New(r.Message)
	}
	return
}

// 获取别野信息
func GetVillaData(ctx *rosm.CTX) (r *VillaData, err error) {
	data, err := web.Web(&http.Client{}, host+getVillaURL, http.MethodGet, makeHeard(ctx), nil)
	log.Debugln("[GetVillaData]", string(data))
	if err != nil {
		return nil, err
	}
	r = new(VillaData)
	err = json.Unmarshal(data, r)
	return
}

type RoomList struct {
	Retcode int    `json:"retcode"`
	Message string `json:"message"`
	Data    struct {
		List []struct {
			GroupID   string `json:"group_id"`
			GroupName string `json:"group_name"`
			RoomList  []struct {
				RoomID   string `json:"room_id"`
				RoomName string `json:"room_name"`
				RoomType string `json:"room_type"`
				GroupID  string `json:"group_id"`
			} `json:"room_list"`
		} `json:"list"`
	} `json:"data"`
}

type UserData struct {
	Retcode int    `json:"retcode"`
	Message string `json:"message"`
	Data    struct {
		Member struct {
			Basic struct {
				UID       string `json:"uid"`
				Nickname  string `json:"nickname"`
				Introduce string `json:"introduce"`
				Avatar    string `json:"avatar"`
				AvatarURL string `json:"avatar_url"`
			} `json:"basic"`
			RoleIDList []string `json:"role_id_list"`
			JoinedAt   string   `json:"joined_at"`
			RoleList   []struct {
				ID       string `json:"id"`
				Name     string `json:"name"`
				Color    string `json:"color"`
				RoleType string `json:"role_type"`
				VillaID  string `json:"villa_id"`
				WebColor string `json:"web_color"`
			} `json:"role_list"`
		} `json:"member"`
	} `json:"data"`
}

type VillaData struct {
	Retcode int    `json:"retcode"`
	Message string `json:"message"`
	Data    struct {
		Villa struct {
			VillaID        string   `json:"villa_id"`
			Name           string   `json:"name"`
			VillaAvatarURL string   `json:"villa_avatar_url"`
			OwnerUID       string   `json:"owner_uid"`
			IsOfficial     bool     `json:"is_official"`
			Introduce      string   `json:"introduce"`
			CategoryID     int      `json:"category_id"`
			Tags           []string `json:"tags"`
		} `json:"villa"`
	} `json:"data"`
}