package hub

type (
	Hub struct {
		Register   chan *Client //注册消息
		UnRegister chan *Client //注销消息
		Broadcast  chan []byte  //广播消息
	}
)

func NewHub() *Hub {
	return &Hub{
		Register:   make(chan *Client),
		UnRegister: make(chan *Client),
		Broadcast:  make(chan []byte),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			//h.Clients[client.Token] = client
			println(client.Token + "上线了")
		case client := <-h.UnRegister:
			//if _, ok := h.Clients[client.Token]; ok {
			//	delete(h.Clients, client.Token)
			close(client.Send)
			//}
		case _ = <-h.Broadcast:
			println("开始广播了")
			//for _, client := range h.Clients {
			//	select {
			//	case client.Send <- message:
			//	default:
			//		close(client.Send)
			//		delete(h.Clients, client.Token)
			//	}
			//}
		}
	}
}
