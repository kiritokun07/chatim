package hub

import "github.com/zeromicro/go-zero/core/logx"

type (
	Hub struct {
		clients    map[string]*Client // 在线运营
		Register   chan *Client       // 注册消息
		Unregister chan *Client       // 注册消息
		Broadcast  chan []byte        //发消息
	}
)

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[string]*Client),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		//Operators:  make(map[int64]*Client),
		//Services:   make(map[int64]*Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.clients[client.Token] = client
			logx.Infov(client.Token + "上线了")
			println(client.Token + "上线了")
		case client := <-h.Unregister:
			if _, ok := h.clients[client.Token]; ok {
				delete(h.clients, client.Token)
				close(client.Send)
			}
		case message := <-h.Broadcast:
			for _, client := range h.clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.clients, client.Token)
				}
			}
		}
	}
}
