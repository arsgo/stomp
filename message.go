package stomp

import  s "github.com/go-stomp/stomp"

type MsgHandler interface {
	Ack()
	GetMessage() string
}

type StompMessage struct {
	s       *Stomp
	msg     *s.Message
	Message string
}

//Ack
func (m *StompMessage) Ack() {
	m.s.conn.Ack(m.msg)
}
func (m *StompMessage) GetMessage() string {
	return m.Message
}

//NewMessage
func NewMessage(s *Stomp, msg *s.Message) *StompMessage {
	return &StompMessage{s: s, msg: msg, Message: string(msg.Body)}
}
