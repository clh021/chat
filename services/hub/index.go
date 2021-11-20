package hub

import (
	"github.com/clh021/chat/services/client"
)

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *client.Client),
		unregister: make(chan *client.Client),
		clients:    make(map[*client.Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		// 上线的客户端记入客户端在线状态列表
		case client := <-h.register:
			h.clients[client] = true
		// 下线的客户端从客户端在线状态列表中删除，并关闭发送 channel
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.Send)
			}
		// 消息广播，给每一个在线客户端都发送消息
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.clients, client)
				}
			}
		}
	}
}
