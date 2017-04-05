package msgPusher

import (
	"LemonChat/msg"
	"LemonChat/msgServer"
	"net"
	"encoding/json"
	"log"
)

type MessagePusher struct {
	msgQ      *msg.MessageQueue
	isRunning bool
	clientMan *msgServer.ClientManager
}

func NewMsgPusher() *MessagePusher {
	return &MessagePusher{msgQ: msg.NewMsgQueue(100000), isRunning: false}
}

func (p *MessagePusher) PushMessage(v interface{}) {
	err := p.msgQ.Put(v)
	if err != nil {
		log.Println("Add msg to queue fail.")
		return
	}
	go p.Run()
}

func (p *MessagePusher) Run() {
	if p.isRunning {
		return
	}
	p.isRunning = true
	for {
		v, err := p.msgQ.Get()
		if err != nil {
			return
		}
		msg := &msg.Message(v)
		b, err := json.Marshal(msg)
		if err != nil {
			continue
		}
		c, err := p.clientMan.GetClient(msg.DstId)
		if err != nil {
			log.Println("get client error : ", err.Error())
			continue
		}
		send(c.Conn, b)
	}
}

func send(conn net.Conn, b []byte) error {
	_, err := conn.Write(b)
	if err != nil {
		return err
	}
	return nil
}
