// Copyright © 2020. Drew Lee. All rights reserved.

package controller

import (
	"KarlMalone/pkg/logger"
	"encoding/json"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"gopkg.in/fatih/set.v0"
)

var clientMap = make(map[uint]*Node, 0)
var mu sync.RWMutex

// cmd category
const (
	CmdSingleMsg = 10
	CmdRoomMsg   = 11
	CmdHeart     = 0
)

type Node struct {
	Conn      *websocket.Conn
	DataQueue chan []byte
	GroupSets set.Interface
}

type Message struct {
	Id       uint   `json:"id,omitempty"`        //消息ID
	MemberId uint   `json:"member_id,omitempty"` //谁发的
	Cmd      int    `json:"cmd,omitempty"`       //群聊还是私聊
	DstId    uint   `json:"dst_id,omitempty"`    //对端用户ID/群ID
	Media    int    `json:"media,omitempty"`     //消息按照什么格式展示
	Content  string `json:"content,omitempty"`   //消息内容
	Pic      string `json:"pic,omitempty"`       //预览图片
	Url      string `json:"url,omitempty"`       //服务的URL
	Memo     string `json:"memo,omitempty"`      //简单描述
	Amount   int    `json:"amount,omitempty"`    //其他和数字相关的
}

func Chat(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id, _ := strconv.ParseUint(query.Get("id"), 10, 64)
	memberId := uint(id)
	token := query.Get("token")
	isLegal := checkToken(memberId, token)

	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return isLegal
		},
	}).Upgrade(w, r, nil)
	if err != nil {
		logger.Error("chat controller websocket conn error", zap.String("reason", err.Error()))
		return
	}

	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}

	groupIds := contactService.SearchGroupIds(memberId)
	for _, groupId := range groupIds {
		node.GroupSets.Add(groupId)
	}

	mu.Lock()
	clientMap[memberId] = node
	mu.Unlock()

	go sendProc(node)

	go recvProc(node)

	sendMsg(memberId, []byte("hello."))
}

func sendProc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			if err := node.Conn.WriteMessage(websocket.TextMessage, data); err != nil {
				logger.Error("chat controller conn write error", zap.String("reason", err.Error()))
				return
			}

		}
	}
}

func recvProc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			logger.Error("chat controller conn read error", zap.String("reason", err.Error()))
			return
		}

		dispatch(data)
		logger.Info("chat controller receive data", zap.ByteString("data", data))
	}
}

func dispatch(data []byte) {
	msg := Message{}
	if err := json.Unmarshal(data, &msg); err != nil {
		logger.Error("chat controller json unmarshal error", zap.String("reason", err.Error()))
		return
	}
	switch msg.Cmd {
	case CmdSingleMsg:
		sendMsg(msg.DstId, data)
	case CmdRoomMsg:
		for _, v := range clientMap {
			if v.GroupSets.Has(msg.DstId) {
				v.DataQueue <- data
			}
		}
	case CmdHeart:
		// 检测客户端心跳
	}
}

func checkToken(memberId uint, token string) bool {
	member, err := MemberService.Find(memberId)
	if err != nil {
		return false
	}
	return member.Token == token
}

func sendMsg(dstId uint, msg []byte) {
	mu.RLock()
	node, ok := clientMap[dstId]
	mu.RUnlock()
	if ok {
		node.DataQueue <- msg
	}
}

// add new groupId to member's clientMap
func AddGroupId(memberId, groupId uint) {
	mu.Lock()
	defer mu.Unlock()
	if node, ok := clientMap[memberId]; ok {
		node.GroupSets.Add(groupId)
	}
}
