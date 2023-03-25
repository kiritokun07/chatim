package hub

import (
	"bytes"
	"time"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 4096
	bufSize        = 256
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

type Client struct {
	Hub   *Hub
	Conn  *websocket.Conn
	Send  chan []byte
	Read  chan []byte
	Token string //放客户端的各种信息
}

func NewClient(hub *Hub, conn *websocket.Conn, token string) *Client {
	return &Client{
		Hub:   hub,
		Conn:  conn,
		Send:  make(chan []byte),
		Read:  make(chan []byte),
		Token: token,
	}
}

func (c *Client) Run() {
	//不能开协程
	c.register()
	threading.GoSafe(func() {
		c.readPump()
	})
	threading.GoSafe(func() {
		c.writePump()
	})
}

func (c *Client) register() {
	c.Hub.Register <- c
}

func (c *Client) readPump() {
	defer func() {
		c.Hub.UnRegister <- c
		_ = c.Conn.Close()
	}()
	c.Conn.SetReadLimit(maxMessageSize)
	_ = c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error { _ = c.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure, websocket.CloseNoStatusReceived) {
				logx.Errorf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		//var d msg.UpMessage
		//_ = jsonx.Unmarshal(message, &d)
		c.Read <- message
	}
}

func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		_ = c.Conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			_ = c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				_ = c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			_ = c.Conn.WriteMessage(websocket.TextMessage, message)
		case <-ticker.C:
			_ = c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
