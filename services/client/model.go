package client

import (
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

// 客户端 Hub 读写数据到 websocket
type Client struct {
	// hub *hub.Hub

	// 对应 websocket 连接
	Conn *websocket.Conn

	// 出站消息的缓冲
	Send chan []byte
}
