package main

import proto "github.com/z0sum/jebb/chat/proto"

//Hub ...
// type Hub struct {
// 	register   chan string
// 	unregister chan string
// 	users      map[string]bool
// }

// func newHub() *Hub {
// 	return &Hub{
// 		register:   make(chan string),
// 		unregister: make(chan string),
// 		users:      make(map[string]bool),
// 	}
// }

// //Init the Hub
// func (h *Hub) Init() {
// 	select {
// 	case user := <-h.register:
// 		h.users[user] = true
// 	case user := <-h.unregister:
// 		if _, ok := h.users[user]; ok {
// 			delete(h.users, user)
// 		}
// 	}
// }

type Hub interface {
	register(string) bool
	unregister(string) bool
	list() []string
	link(*proto.ChatService_SocketStream, string) bool
	listStreams() map[string]*proto.ChatService_SocketStream
	unlink(string) 
}

type hub struct {
	users   map[string]bool
	streams map[string]*proto.ChatService_SocketStream
}

func (h *hub) register(user string) bool {
	h.users[user] = true
	return true
}

func (h *hub) unregister(user string) bool {
	if _, ok := h.users[user]; ok {
		delete(h.users, user)
	}
	return true
}

func (h *hub) list() []string {
	res := []string{}
	for k := range h.users {
		res = append(res, k)
	}
	return res
}

func (h *hub) link(stream *proto.ChatService_SocketStream, user string) bool {
	h.streams[user] = stream
	return true
}

func (h *hub) listStreams() map[string]*proto.ChatService_SocketStream {
	return h.streams
}

func (h *hub) unlink(user string) {
	if _, ok := h.streams[user]; ok {
		delete(h.streams, user)
	}
}
