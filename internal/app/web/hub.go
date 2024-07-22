package web

import (
	"lca/internal/pkg/logging"
	"lca/internal/pkg/util"
)

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client

	brokerEvents *BrokerEventData
	logger       logging.Logger
}

func newHub(brokerEvents *BrokerEventData, logger logging.Logger) *Hub {
	return &Hub{
		broadcast:    make(chan []byte),
		register:     make(chan *Client),
		unregister:   make(chan *Client),
		clients:      make(map[*Client]bool),
		brokerEvents: brokerEvents,
		logger:       logger,
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		case event := <-h.brokerEvents.Channel:
			jsonEvent, err := util.ToJson(event)
			if err != nil {
				h.logger.Errors(err)
			} else {
				for client := range h.clients {
					client.send <- jsonEvent
				}
			}
		}
	}
}
