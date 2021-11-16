package chat

import (
	"flag"
	"log"
	"net/http"

	"github.com/clh021/chat/services/client"
	"github.com/clh021/chat/services/hub"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}

// 处理 websocket 请求
func serveWs(hub *hub.Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// 实例化当前客户端并记入在线列表
	client := &client.Client{Conn: conn, Send: make(chan []byte, 256)}
	// hub: hub,
	// client.hub.register <- client

	// 通过在新的 goroutine 中完成所有工作，允许收集调用者引用的内存。
	go client.WritePump()
	go client.ReadPump()
}

func Run() {
	addr := flag.String("addr", ":8080", "http service address")
	flag.Parse()
	hub := hub.NewHub()
	go hub.Run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	log.Println("ListenAndServe:", *addr)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("Error: ", err)
	}
}
