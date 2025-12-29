package main


type room struct {
	// forward is a channel that holds incoming messages
	// that should be forwarded to the other clients.
	forward chan []byte
	// join is a channel for clients wishing to join the room.
	join chan *client
	// leave is a channel for clients wishing ro leave the room.
	leave chan *client
	// clients holds all current clients in the room
	clients map[*client]bool
}

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			// joining
			r.clients[client] = true
		case client := <-r.leave:
			// leaving
			delete(r.clients, client)
			close(client.send)
		case msg := <-r.forward:
			// forward message to all clients
			for client := range r.clients {
			select {
			case client.send <- msg:
				// send message
			default:
				// failed to send
				delete(r.clients, client)
				close(client.send)
			}
			}
		}
	}
}