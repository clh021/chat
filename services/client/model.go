package client

import (
	"github.com/gorilla/websocket"
)

// 客户端 Hub 读写数据到 websocket
type Client struct {
	// hub *hub.Hub
	// User user.User
	Ip string // 上线的IP地址
	Conn *websocket.Conn // 对应 websocket 连接
	Send chan []byte     // 出站消息的缓冲
}
