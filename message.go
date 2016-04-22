package stomp

import "github.com/gmallard/stompngo"

type MsgHandler interface {
	Ack()
	GetMessage() string
}

type StompMessage struct {
	s       *Stomp
	msg     stompngo.MessageData
	Message string
}

//Ack
func (m *StompMessage) Ack() {
	m.s.conn.Ack(m.msg.Message.Headers)
}
func (m *StompMessage) GetMessage() string {
	return m.Message
}

//NewMessage
func NewMessage(s *Stomp, msg stompngo.MessageData) *StompMessage {
	return &StompMessage{s: s, msg: msg, Message: msg.Message.BodyString()}
}
