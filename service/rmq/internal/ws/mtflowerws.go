package ws

import (
	"context"
	"sync"
	"time"

	"chatim/shared/mq"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
)

type (
	MtflowerWs struct {
		conn          *websocket.Conn
		lock          sync.Mutex
		wsUrl         string
		producerTopic string //下行消息生产者topic
		producer      *mq.Producer
	}
)

func NewMtflowerWs(wsUrl, producerTopic string, producer *mq.Producer) (*MtflowerWs, error) {
	ws := &MtflowerWs{
		lock:          sync.Mutex{},
		wsUrl:         wsUrl,
		producerTopic: producerTopic,
		producer:      producer,
	}
	ws.connect()
	threading.GoSafe(func() {
		ws.ping()
	})
	return ws, nil
}

func (ws *MtflowerWs) connect() {
	dialer := websocket.Dialer{}
	conn, _, err := dialer.Dial(ws.wsUrl, nil)
	if err != nil {
		logx.Error(err)
	}
	ws.conn = conn
	logx.Infof("ws connected url is: %s", ws.wsUrl)
	threading.GoSafe(func() {
		ws.Listen()
	})
}

func (ws *MtflowerWs) ping() {
	ticker := time.NewTicker(time.Second * 30)
	for {
		select {
		case <-ticker.C:
			if err := ws.pingMsg(); err != nil {
				logx.Error(err)
				ws.connect()
			}
		}
	}
}

func (ws *MtflowerWs) pingMsg() interface{} {
	ws.lock.Lock()
	defer ws.lock.Unlock()
	return ws.conn.WriteMessage(websocket.PingMessage, []byte{})
}

func (ws *MtflowerWs) Listen() {
	for {
		messageType, messageData, err := ws.conn.ReadMessage()
		if nil != err {
			logx.Error(err)
			ws.connect()
			break
		}
		switch messageType {
		case websocket.TextMessage:
			//TODO 收到美团鲜花平台方消息，解密，发mq 这里要不要转一下格式
			logx.Info(string(messageData))
			err := ws.producer.PushByTopic(context.TODO(), ws.producerTopic, messageData)
			if err != nil {
				logx.Error(err)
			}
		case websocket.BinaryMessage:
			logx.Info(string(messageData))
		case websocket.CloseMessage:
			logx.Error("mtflower ws closed")
			ws.connect()
			break
		case websocket.PingMessage:
			logx.Info("ping")
		case websocket.PongMessage:
			logx.Info("pong")
		default:
		}
	}
}

func (ws *MtflowerWs) Send(bytes []byte) error {
	ws.lock.Lock()
	defer ws.lock.Unlock()
	logx.Info(string(bytes))
	return ws.conn.WriteMessage(websocket.TextMessage, bytes)
}
