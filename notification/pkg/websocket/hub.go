package websocket

type Hub struct {
	clients map[uint]*Client

	register chan *Client

	unregister chan *Client

	broadcast chan []byte
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[uint]*Client),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte),
	}
}

func (h *Hub) Run() {
	for {
		select {
		// register
		case client := <-h.register:
			h.clients[client.userId] = client
		// unregister
		case client := <-h.unregister:
			delete(h.clients, client.userId)
		// broadcast:
		case message := <-h.broadcast:
			for _, client := range h.clients {
				select {
				case client.send <- message:
				default:
					delete(h.clients, client.userId)
					close(client.send)
				}
			}
		}

	}
}
