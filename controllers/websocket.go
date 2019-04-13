package controllers

import (
	"container/list"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

type WSController struct {
	beego.Controller
}

type EventType int

const (
	EVENT_JOIN = iota
	EVENT_LEAVE
	EVENT_MESSAGE
)

type Event struct {
	Type      EventType // JOIN, LEAVE, MESSAGE
	Timestamp int // Unix timestamp (secs)
	Content   string
	Conn      *websocket.Conn
}

var (
	// Channel for new join users.
	subscribe = make(chan Subscriber, 10)
	// Channel for exit users.
	unsubscribe = make(chan Subscriber, 10)
	// Send events here to publish them.
	publish = make(chan Event, 10)
	//// Long polling waiting list.
	//waitingList = list.New()
	subscribers = list.New()
)

type Subscriber struct {
	Conn *websocket.Conn
}

func Join(ws *websocket.Conn) {
	subscribe <- Subscriber{Conn:ws}
}

func Leave(ws *websocket.Conn) {
	unsubscribe <- Subscriber{Conn:ws}
}

func newEvent(ep EventType, msg string, ws *websocket.Conn) Event {
	return Event{ep, int(time.Now().Unix()), msg, ws}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:    4096,
	WriteBufferSize:   4096,
	EnableCompression: true,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (c *WSController) Get() {
	//TODO: Crash without this
	c.TplName = "null.html"

	ws, err := upgrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(c.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		return
	}
	Join(ws)
	defer Leave(ws)
	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			return
		}
		log.Println("Message:", string(p))
		publish <- newEvent(EVENT_MESSAGE, string(p), ws)
	}
}

func wsroom() {
	for{
		select {
		case sub := <-subscribe:
			if !isWsExist(subscribers, sub.Conn) {
				subscribers.PushBack(sub) // Add user to the end of list.
				// Publish a JOIN event.
				publish <- newEvent(EVENT_JOIN, "", nil)
				beego.Info("New user:", ";WebSocket:", &sub.Conn)
			} else {
				beego.Info("Old user:", ";WebSocket:", &sub.Conn)
			}

		case unsub := <-unsubscribe:
			for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
				if sub.Value.(Subscriber).Conn == unsub.Conn {
					subscribers.Remove(sub)
					// Clone connection.
					ws := sub.Value.(Subscriber).Conn
					if ws != nil {
						ws.Close()
						beego.Error("WebSocket closed:", unsub)
					}
					publish <- newEvent(EVENT_LEAVE, "", nil) // Publish a LEAVE event.
					break
				}
			}

		case event := <-publish:
			if event.Type == EVENT_MESSAGE {
				beego.Info("Message from", ";Content:", event.Content)
				var dat map[string]interface{}
				if err := json.Unmarshal([]byte(event.Content), &dat); err != nil {
					panic(err)
				}
			}
		}
	}
}

func isWsExist(subscribers *list.List, ws *websocket.Conn) bool {
	for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
		if sub.Value.(Subscriber).Conn == ws {
			return true
		}
	}
	return false
}

func init() {
	go wsroom()
}