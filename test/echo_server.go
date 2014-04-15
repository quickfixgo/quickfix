package main

import (
	"fmt"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/cracker"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/enum"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/fix40"
	"github.com/quickfixgo/quickfix/fix41"
	"github.com/quickfixgo/quickfix/fix42"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/fix50"
	"github.com/quickfixgo/quickfix/fix50sp1"
	"github.com/quickfixgo/quickfix/fix50sp2"
	"github.com/quickfixgo/quickfix/log"
	"github.com/quickfixgo/quickfix/message"
	"github.com/quickfixgo/quickfix/settings"
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

func (e EchoApplication) FromAdmin(msg message.Message, sessionID quickfix.SessionID) (reject message.MessageReject) {
	return
}

func (e *EchoApplication) FromApp(msg message.Message, sessionID quickfix.SessionID) (reject message.MessageReject) {
	fmt.Println("Got Message ", msg)
	return cracker.Crack(msg, sessionID, e)
}

func (e *EchoApplication) processMsg(msg message.Message, sessionID quickfix.SessionID) (reject message.MessageReject) {
	orderID := new(field.ClOrdID)
	if err := msg.Body.Get(orderID); err != nil {
		return
	}

	reply := message.CreateMessageBuilder()
	sessionOrderID := sessionID.String() + orderID.Value
	possResend := new(field.PossResend)
	if err := msg.Header.Get(possResend); err == nil && possResend.Value {
		if e.OrderIds[sessionOrderID] {
			return
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

	return
}

func (e EchoApplication) OnFIX40NewOrderSingle(msg fix40.NewOrderSingle, sessionID quickfix.SessionID) message.MessageReject {
	return e.processMsg(msg.Message, sessionID)
}

func (e EchoApplication) OnFIX41NewOrderSingle(msg fix41.NewOrderSingle, sessionID quickfix.SessionID) message.MessageReject {
	return e.processMsg(msg.Message, sessionID)
}

func (e EchoApplication) OnFIX42NewOrderSingle(msg fix42.NewOrderSingle, sessionID quickfix.SessionID) message.MessageReject {
	return e.processMsg(msg.Message, sessionID)
}

func (e EchoApplication) OnFIX43NewOrderSingle(msg fix43.NewOrderSingle, sessionID quickfix.SessionID) message.MessageReject {
	return e.processMsg(msg.Message, sessionID)
}

func (e EchoApplication) OnFIX44NewOrderSingle(msg fix44.NewOrderSingle, sessionID quickfix.SessionID) message.MessageReject {
	return e.processMsg(msg.Message, sessionID)
}

func (e EchoApplication) OnFIX50NewOrderSingle(msg fix50.NewOrderSingle, sessionID quickfix.SessionID) message.MessageReject {
	return e.processMsg(msg.Message, sessionID)
}

func (e EchoApplication) OnFIX50SP1NewOrderSingle(msg fix50sp1.NewOrderSingle, sessionID quickfix.SessionID) message.MessageReject {
	return e.processMsg(msg.Message, sessionID)
}

func (e EchoApplication) OnFIX50SP2NewOrderSingle(msg fix50sp2.NewOrderSingle, sessionID quickfix.SessionID) message.MessageReject {
	return e.processMsg(msg.Message, sessionID)
}

func (e EchoApplication) OnFIX42SecurityDefinition(msg fix42.SecurityDefinition, sessionID quickfix.SessionID) (reject message.MessageReject) {
	reply := msg.ToBuilder()
	quickfix.SendToTarget(reply, sessionID)
	return
}

func (e EchoApplication) OnFIX43SecurityDefinition(msg fix43.SecurityDefinition, sessionID quickfix.SessionID) (reject message.MessageReject) {
	reply := msg.ToBuilder()
	quickfix.SendToTarget(reply, sessionID)
	return
}

func (e EchoApplication) OnFIX44SecurityDefinition(msg fix44.SecurityDefinition, sessionID quickfix.SessionID) (reject message.MessageReject) {
	reply := msg.ToBuilder()
	quickfix.SendToTarget(reply, sessionID)
	return
}

func (e EchoApplication) OnFIX50SecurityDefinition(msg fix50.SecurityDefinition, sessionID quickfix.SessionID) (reject message.MessageReject) {
	reply := msg.ToBuilder()
	quickfix.SendToTarget(reply, sessionID)
	return
}

func (e EchoApplication) OnFIX50SP1SecurityDefinition(msg fix50sp1.SecurityDefinition, sessionID quickfix.SessionID) (reject message.MessageReject) {
	reply := msg.ToBuilder()
	quickfix.SendToTarget(reply, sessionID)
	return
}

func (e EchoApplication) OnFIX50SP2SecurityDefinition(msg fix50sp2.SecurityDefinition, sessionID quickfix.SessionID) (reject message.MessageReject) {
	reply := msg.ToBuilder()
	quickfix.SendToTarget(reply, sessionID)
	return
}

func main() {
	app := new(EchoApplication)

	globalSettings := settings.NewDictionary()
	globalSettings.SetInt(settings.SocketAcceptPort, 5001)
	globalSettings.SetString(settings.SenderCompID, "ISLD")
	globalSettings.SetString(settings.TargetCompID, "TW")
	globalSettings.SetBool(settings.ResetOnLogon, true)

	appSettings := settings.NewApplicationSettings(globalSettings)

	appSettings.AddSession("FIX40", settings.NewDictionary().
		SetString(settings.BeginString, fix.BeginString_FIX40).
		SetString(settings.DataDictionary, "../spec/FIX40.xml"))

	appSettings.AddSession("FIX41", settings.NewDictionary().
		SetString(settings.BeginString, fix.BeginString_FIX41).
		SetString(settings.DataDictionary, "../spec/FIX41.xml"))

	appSettings.AddSession("FIX42", settings.NewDictionary().
		SetString(settings.BeginString, fix.BeginString_FIX42).
		SetString(settings.DataDictionary, "../spec/FIX42.xml"))

	appSettings.AddSession("FIX43", settings.NewDictionary().
		SetString(settings.BeginString, fix.BeginString_FIX43).
		SetString(settings.DataDictionary, "../spec/FIX43.xml"))

	appSettings.AddSession("FIX44", settings.NewDictionary().
		SetString(settings.BeginString, fix.BeginString_FIX44).
		SetString(settings.DataDictionary, "../spec/FIX44.xml"))

	appSettings.AddSession("FIX50", settings.NewDictionary().
		SetString(settings.BeginString, fix.BeginString_FIXT11).
		SetString(settings.DefaultApplVerID, enum.ApplVerID_FIX50).
		SetString(settings.TransportDataDictionary, "../spec/FIXT11.xml").
		SetString(settings.AppDataDictionary, "../spec/FIX50.xml"))

	appSettings.AddSession("FIX50SP1", settings.NewDictionary().
		SetString(settings.BeginString, fix.BeginString_FIXT11).
		SetString(settings.DefaultApplVerID, enum.ApplVerID_FIX50SP1).
		SetString(settings.TransportDataDictionary, "../spec/FIXT11.xml").
		SetString(settings.AppDataDictionary, "../spec/FIX50SP1.xml"))

	appSettings.AddSession("FIX50SP2", settings.NewDictionary().
		SetString(settings.BeginString, fix.BeginString_FIXT11).
		SetString(settings.DefaultApplVerID, enum.ApplVerID_FIX50SP2).
		SetString(settings.TransportDataDictionary, "../spec/FIXT11.xml").
		SetString(settings.AppDataDictionary, "../spec/FIX50SP2.xml"))

	acceptor, err := quickfix.NewAcceptor(app, appSettings, log.ScreenLogFactory{})
	if err != nil {
		fmt.Printf("Unable to create Acceptor: ", err)
		return
	}

	if err = acceptor.Start(); err != nil {
		fmt.Printf("Unable to start Acceptor: ", err)
		return
	}

	interrupt := make(chan os.Signal)
	signal.Notify(interrupt)
	<-interrupt

	acceptor.Stop()
}
