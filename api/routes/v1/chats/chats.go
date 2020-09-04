// Copyright © 2020. Drew Lee. All rights reserved.

package chats

import (
	"KarlMalone/internal/service"
	"KarlMalone/pkg/logger"
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"gopkg.in/fatih/set.v0"
)

var clientMap = make(map[int64]*Node, 0)
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
	Id        int64     `json:"id,omitempty"`         // 消息ID
	MemberId  int64     `json:"member_id,omitempty"`  // 谁发的
	Cmd       int       `json:"cmd,omitempty"`        // 群聊还是私聊
	DstId     int64     `json:"dst_id,omitempty"`     // 对端用户ID/群ID
	Media     int       `json:"media,omitempty"`      // 消息按照什么格式展示
	Content   string    `json:"content,omitempty"`    // 消息内容
	Pic       string    `json:"pic,omitempty"`        // 预览图片
	Url       string    `json:"url,omitempty"`        // 服务的URL
	Memo      string    `json:"memo,omitempty"`       // 简单描述
	Amount    int       `json:"amount,omitempty"`     // 其他和数字相关的
	CreatedAt time.Time `json:"created_at,omitempty"` // 发送时间
}

func chat(c *gin.Context) {
	memberId, _ := strconv.ParseInt(c.Query("member_id"), 10, 64)
	token := c.Query("token")
	isLegal := checkToken(memberId, token)

	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return isLegal
		},
	}).Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.Error("chat controller websocket conn error", zap.String("reason", err.Error()))
		return
	}

	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}

	contact := service.Contact{
		OwnerId: memberId,
	}
	groupIds := contact.FindGroupIds()
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

func checkToken(memberId int64, token string) bool {
	m := service.Member{
		ID: memberId,
	}
	member, err := m.Find()
	if err != nil {
		return false
	}
	if member.ID == 0 {
		return false
	}
	return member.Token == token
}

func sendMsg(dstId int64, msg []byte) {
	mu.Lock()
	defer mu.Unlock()
	node, ok := clientMap[dstId]
	if ok {
		node.DataQueue <- msg
	}
}

// add new groupId to member's clientMap
func AddGroupId(memberId, groupId int64) {
	mu.Lock()
	defer mu.Unlock()
	if node, ok := clientMap[memberId]; ok {
		node.GroupSets.Add(groupId)
	}
}
