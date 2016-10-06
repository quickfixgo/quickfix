package quickfix

type messagePool struct {
	m []*Message
}

func (p *messagePool) New() *Message {
	msg := NewMessage()
	return msg
}

func (p *messagePool) Get() (msg *Message) {
	if len(p.m) > 0 {
		msg, p.m = p.m[len(p.m)-1], p.m[:len(p.m)-1]
	} else {
		msg = p.New()
	}

	msg.keepMessage = false

	return
}

func (p *messagePool) Put(msg *Message) {
	p.m = append(p.m, msg)
}
