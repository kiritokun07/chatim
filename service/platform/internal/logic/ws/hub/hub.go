package hub

type (
	Hub struct {
		Clients    map[string]*Client // 在线运营
		Register   chan *Client       // 注册消息
		Unregister chan *Client       // 注册消息
		Broadcast  chan []byte        //发消息
	}
)

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[string]*Client),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan []byte),
		//Operators:  make(map[int64]*Client),
		//Services:   make(map[int64]*Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client.Token] = client
			println(client.Token + "上线了")
		case client := <-h.Unregister:
			if _, ok := h.Clients[client.Token]; ok {
				delete(h.Clients, client.Token)
				close(client.Send)
			}
		case message := <-h.Broadcast:
			println("开始广播了")
			for _, client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, client.Token)
				}
			}
		}
	}
}
