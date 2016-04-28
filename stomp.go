package stomp

import (
	"fmt"

	"github.com/go-stomp/stomp"

	//"github.com/gmallard/stompngo"
	//"github.com/go-stomp/stomp"

	//"github.com/gmallard/stompngo"
)

//StompMQ manage stomp server
type Stomp struct {
	conn    *stomp.Conn
	address string
}

//NewStompMQ
func NewStomp(address string) (st *Stomp, err error) {
	st = &Stomp{address: address}
	st.conn, err = stomp.Dial("tcp", "192.168.1.1:61613")
	if err != nil {
		return
	}
	return
}

//Send
func (s *Stomp) Send(queue string, msg string) (err error) {
	err = s.conn.Send(
		fmt.Sprintf("/queue/%s", queue), // destination
		"text/plain",                    // content-type
		[]byte(msg),
		stomp.SendOpt.Receipt,
		stomp.SendOpt.Header("expires", "2049-12-31 23:59:59")) // body
	return
}

//Subscribe
func (s *Stomp) Consume(queue string, count int, call func(MsgHandler) bool) (err error) {
	sub, err := s.conn.Subscribe(fmt.Sprintf("/queue/%s", queue), stomp.AckClient)
	if err != nil {
		return err
	}
	for i := 0; i < count; i++ {
		msg := <-sub.C
		handler := NewMessage(s, msg)
		b := call(handler)
		if b {
			handler.Ack()
		}

	}
	err = sub.Unsubscribe()
	return

	return
}

//Close
func (s *Stomp) Close() {
	s.conn.Disconnect()
}

/*
//NewStompMQ
func NewStomp(address string) (mq *Stomp, err error) {
	mq = &Stomp{address: address}
	con, err := net.Dial("tcp", address)
	if err != nil {
		return
	}
	header := stompngo.Headers{"accept-version", "1.1"}
	mq.conn, err = stompngo.Connect(con, header)
	return
}

//Send
func (s *Stomp) Send(queue string, msg string) error {
	header := stompngo.Headers{"destination", queue, "persistent", "true"}
	return s.conn.Send(header, msg)
}

//Subscribe
func (s *Stomp) Consume(queue string, call func(MsgHandler)) (err error) {
	if !s.conn.Connected() {
		err = errors.New("not connect to stomp server")
		return
	}
	subscriberHeader := stompngo.Headers{"destination",
		fmt.Sprintf("/queue/%s", queue), "ack", "client"}
	msgChan, err := s.conn.Subscribe(subscriberHeader)
	if err != nil {
		return
	}
	for {
		msg := <-msgChan
		call(NewMessage(s, msg))
	}
}

//Close
func (s *Stomp) Close() {
	s.conn.Disconnect(stompngo.Headers{})
}
*/
