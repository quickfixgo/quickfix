package main

import (
	"fmt"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/cracker"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/fix40"
	"github.com/quickfixgo/quickfix/fix41"
	"github.com/quickfixgo/quickfix/fix42"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/fix50"
	"github.com/quickfixgo/quickfix/fix50sp1"
	"github.com/quickfixgo/quickfix/fix50sp2"
	"github.com/quickfixgo/quickfix/message"
	"os"
	"os/signal"
)

type EchoApplication struct {
	cracker.MessageCracker
	OrderIds map[string]bool
}

func (e EchoApplication) OnCreate(sessionID quickfix.SessionID) {
	fmt.Printf("OnCreate %v\n", sessionID.String())
}
func (e *EchoApplication) OnLogon(sessionID quickfix.SessionID) {
	fmt.Printf("OnLogon %v\n", sessionID.String())
	e.OrderIds = make(map[string]bool)
}
func (e *EchoApplication) OnLogout(sessionID quickfix.SessionID) {
	fmt.Printf("OnLogout %v\n", sessionID.String())
}
func (e EchoApplication) ToAdmin(msgBuilder message.MessageBuilder, sessionID quickfix.SessionID) {
}

func (e EchoApplication) ToApp(msgBuilder message.MessageBuilder, sessionID quickfix.SessionID) (err error) {
	return
}

func (e EchoApplication) FromAdmin(msg message.Message, sessionID quickfix.SessionID) (err errors.MessageRejectError) {
	return
}

func (e *EchoApplication) FromApp(msg message.Message, sessionID quickfix.SessionID) (err errors.MessageRejectError) {
	fmt.Println("Got Message ", msg)
	return cracker.Crack(msg, sessionID, e)
}

func (e *EchoApplication) processMsg(msg message.Message, sessionID quickfix.SessionID) errors.MessageRejectError {
	orderID := new(field.ClOrdID)
	if err := msg.Body.Get(orderID); err != nil {
		return err
	}

	reply := message.CreateMessageBuilder()
	sessionOrderID := sessionID.String() + orderID.Value
	possResend := new(field.PossResend)
	if err := msg.Header.Get(possResend); err == nil && possResend.Value {
		if e.OrderIds[sessionOrderID] {
			return nil
		}

		reply.Header.Set(possResend)
	}

	e.OrderIds[sessionOrderID] = true

	msgType := new(field.MsgType)
	msg.Header.Get(msgType)
	reply.Header.Set(msgType)

	for _, tag := range msg.Body.Tags() {
		field := new(message.StringValue)
		if err := msg.Body.GetField(tag, field); err == nil {
			reply.Body.SetField(tag, field)
		}
	}

	quickfix.SendToTarget(reply, sessionID)

	return nil
}

func (e EchoApplication) OnFIX40NewOrderSingle(msg fix40.NewOrderSingle, sessionID quickfix.SessionID) errors.MessageRejectError {
	return e.processMsg(msg.Message, sessionID)
}

func (e EchoApplication) OnFIX41NewOrderSingle(msg fix41.NewOrderSingle, sessionID quickfix.SessionID) errors.MessageRejectError {
	return e.processMsg(msg.Message, sessionID)
}

func (e EchoApplication) OnFIX42NewOrderSingle(msg fix42.NewOrderSingle, sessionID quickfix.SessionID) errors.MessageRejectError {
	return e.processMsg(msg.Message, sessionID)
}

func (e EchoApplication) OnFIX43NewOrderSingle(msg fix43.NewOrderSingle, sessionID quickfix.SessionID) errors.MessageRejectError {
	return e.processMsg(msg.Message, sessionID)
}

func (e EchoApplication) OnFIX44NewOrderSingle(msg fix44.NewOrderSingle, sessionID quickfix.SessionID) errors.MessageRejectError {
	return e.processMsg(msg.Message, sessionID)
}

func (e EchoApplication) OnFIX50NewOrderSingle(msg fix50.NewOrderSingle, sessionID quickfix.SessionID) errors.MessageRejectError {
	return e.processMsg(msg.Message, sessionID)
}

func (e EchoApplication) OnFIX50SP1NewOrderSingle(msg fix50sp1.NewOrderSingle, sessionID quickfix.SessionID) errors.MessageRejectError {
	return e.processMsg(msg.Message, sessionID)
}

func (e EchoApplication) OnFIX50SP2NewOrderSingle(msg fix50sp2.NewOrderSingle, sessionID quickfix.SessionID) errors.MessageRejectError {
	return e.processMsg(msg.Message, sessionID)
}

func (e EchoApplication) OnFIX42SecurityDefinition(msg fix42.SecurityDefinition, sessionID quickfix.SessionID) (err errors.MessageRejectError) {
	reply := msg.ToBuilder()
	quickfix.SendToTarget(reply, sessionID)
	return
}

func (e EchoApplication) OnFIX43SecurityDefinition(msg fix43.SecurityDefinition, sessionID quickfix.SessionID) (err errors.MessageRejectError) {
	reply := msg.ToBuilder()
	quickfix.SendToTarget(reply, sessionID)
	return
}

func (e EchoApplication) OnFIX44SecurityDefinition(msg fix44.SecurityDefinition, sessionID quickfix.SessionID) (err errors.MessageRejectError) {
	reply := msg.ToBuilder()
	quickfix.SendToTarget(reply, sessionID)
	return
}

func (e EchoApplication) OnFIX50SecurityDefinition(msg fix50.SecurityDefinition, sessionID quickfix.SessionID) (err errors.MessageRejectError) {
	reply := msg.ToBuilder()
	quickfix.SendToTarget(reply, sessionID)
	return
}

func (e EchoApplication) OnFIX50SP1SecurityDefinition(msg fix50sp1.SecurityDefinition, sessionID quickfix.SessionID) (err errors.MessageRejectError) {
	reply := msg.ToBuilder()
	quickfix.SendToTarget(reply, sessionID)
	return
}

func (e EchoApplication) OnFIX50SP2SecurityDefinition(msg fix50sp2.SecurityDefinition, sessionID quickfix.SessionID) (err errors.MessageRejectError) {
	reply := msg.ToBuilder()
	quickfix.SendToTarget(reply, sessionID)
	return
}

func main() {
	app := new(EchoApplication)

	cfg, err := os.Open("session.cfg")
	if err != nil {
		fmt.Println("Error opening cfg:", err)
		return
	}

	appSettings, err := quickfix.ParseSettings(cfg)
	if err != nil {
		fmt.Println("Error reading cfg:", err)
		return
	}

	acceptor, err := quickfix.NewAcceptor(app, appSettings, quickfix.ScreenLogFactory{})
	if err != nil {
		fmt.Println("Unable to create Acceptor: ", err)
		return
	}

	if err = acceptor.Start(); err != nil {
		fmt.Println("Unable to start Acceptor: ", err)
		return
	}

	interrupt := make(chan os.Signal)
	signal.Notify(interrupt)
	<-interrupt

	acceptor.Stop()
}
