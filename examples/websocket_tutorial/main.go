package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// 配置ws连接 在此只配置了读写缓存大小
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	setupRoutes()
	fmt.Println("go websocket tutorial running..")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// 设置路由
func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}

// home页handler
func homePage(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Home Page")
}

// websocket handler
func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	// 返回true表示接受所有的客户端请求，也可以根据request Origin header信息判断是否建立ws连接
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	// http 升级为 websocket 协议
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client Successfully Connected...")

	// 读取响应websocket请求
	reader(ws)
}

// 读取响应websocket请求
func reader(conn *websocket.Conn) {
	for {
		// 读取信息 messageType( =1 表示TextMessage =2表示BinaryMessage); p表示读取到的信息，类型为[]byte
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		// 日志打印读取的信息
		log.Println(string(p))

		// 向客户端写入信息
		var buffer bytes.Buffer
		buffer.Write(p)
		buffer.Write([]byte(" send from server"))
		p2 := buffer.Bytes()
		if err := conn.WriteMessage(messageType, p2); err != nil {
			log.Println(err)
			return
		}
	}
}
