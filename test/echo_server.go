package main

import (
	"fmt"
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/cracker"
	"github.com/cbusbey/quickfixgo/field"
	"github.com/cbusbey/quickfixgo/fix40"
	"github.com/cbusbey/quickfixgo/fix41"
	"github.com/cbusbey/quickfixgo/fix42"
	"github.com/cbusbey/quickfixgo/fix43"
	"github.com/cbusbey/quickfixgo/fix44"
	"github.com/cbusbey/quickfixgo/fix50"
	"github.com/cbusbey/quickfixgo/fix50sp1"
	"github.com/cbusbey/quickfixgo/fix50sp2"
	"github.com/cbusbey/quickfixgo/log"
	"github.com/cbusbey/quickfixgo/settings"
	"os"
	"os/signal"
)

type EchoApplication struct {
	cracker.MessageCracker
	OrderIds map[string]bool
}

func (e EchoApplication) OnCreate(sessionID quickfixgo.SessionID) {
	fmt.Printf("OnCreate %v\n", sessionID.String())
}
func (e *EchoApplication) OnLogon(sessionID quickfixgo.SessionID) {
	fmt.Printf("OnLogon %v\n", sessionID.String())
	e.OrderIds = make(map[string]bool)
}
func (e *EchoApplication) OnLogout(sessionID quickfixgo.SessionID) {
	fmt.Printf("OnLogout %v\n", sessionID.String())
}
func (e EchoApplication) ToAdmin(msgBuilder quickfixgo.MessageBuilder, sessionID quickfixgo.SessionID) {
}

func (e EchoApplication) ToApp(msgBuilder quickfixgo.MessageBuilder, sessionID quickfixgo.SessionID) (err error) {
	return
}

func (e EchoApplication) FromAdmin(msg quickfixgo.Message, sessionID quickfixgo.SessionID) (reject quickfixgo.MessageReject) {
	return
}

func (e *EchoApplication) FromApp(msg quickfixgo.Message, sessionID quickfixgo.SessionID) (reject quickfixgo.MessageReject) {
	fmt.Println("Got Message ", msg)
	return cracker.Crack(msg, sessionID, e)
}

func (e *EchoApplication) processMsg(msg quickfixgo.Message, sessionID quickfixgo.SessionID) (reject quickfixgo.MessageReject) {
	orderID := new(field.ClOrdID)
	if err := msg.Body.Get(orderID); err != nil {
		return
	}

	reply := quickfixgo.NewMessageBuilder()
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
		field := new(quickfixgo.StringValue)
		if err := msg.Body.GetField(tag, field); err == nil {
			reply.Body.SetField(tag, field)
		}
	}

	quickfixgo.SendToTarget(reply, sessionID)

	return
}

func (e EchoApplication) OnFIX40NewOrderSingle(msg fix40.NewOrderSingle, sessionID quickfixgo.SessionID) quickfixgo.MessageReject {
	return e.processMsg(msg.Message, sessionID)
}

func (e EchoApplication) OnFIX41NewOrderSingle(msg fix41.NewOrderSingle, sessionID quickfixgo.SessionID) quickfixgo.MessageReject {
	return e.processMsg(msg.Message, sessionID)
}

func (e EchoApplication) OnFIX42NewOrderSingle(msg fix42.NewOrderSingle, sessionID quickfixgo.SessionID) quickfixgo.MessageReject {
	return e.processMsg(msg.Message, sessionID)
}

func (e EchoApplication) OnFIX43NewOrderSingle(msg fix43.NewOrderSingle, sessionID quickfixgo.SessionID) quickfixgo.MessageReject {
	return e.processMsg(msg.Message, sessionID)
}

func (e EchoApplication) OnFIX44NewOrderSingle(msg fix44.NewOrderSingle, sessionID quickfixgo.SessionID) quickfixgo.MessageReject {
	return e.processMsg(msg.Message, sessionID)
}

func (e EchoApplication) OnFIX50NewOrderSingle(msg fix50.NewOrderSingle, sessionID quickfixgo.SessionID) quickfixgo.MessageReject {
	return e.processMsg(msg.Message, sessionID)
}

func (e EchoApplication) OnFIX50SP1NewOrderSingle(msg fix50sp1.NewOrderSingle, sessionID quickfixgo.SessionID) quickfixgo.MessageReject {
	return e.processMsg(msg.Message, sessionID)
}

func (e EchoApplication) OnFIX50SP2NewOrderSingle(msg fix50sp2.NewOrderSingle, sessionID quickfixgo.SessionID) quickfixgo.MessageReject {
	return e.processMsg(msg.Message, sessionID)
}

func (e EchoApplication) OnFIX42SecurityDefinition(msg fix42.SecurityDefinition, sessionID quickfixgo.SessionID) (reject quickfixgo.MessageReject) {
	reply := msg.ToBuilder()
	quickfixgo.SendToTarget(reply, sessionID)
	return
}

func (e EchoApplication) OnFIX43SecurityDefinition(msg fix43.SecurityDefinition, sessionID quickfixgo.SessionID) (reject quickfixgo.MessageReject) {
	reply := msg.ToBuilder()
	quickfixgo.SendToTarget(reply, sessionID)
	return
}

func (e EchoApplication) OnFIX44SecurityDefinition(msg fix44.SecurityDefinition, sessionID quickfixgo.SessionID) (reject quickfixgo.MessageReject) {
	reply := msg.ToBuilder()
	quickfixgo.SendToTarget(reply, sessionID)
	return
}

func (e EchoApplication) OnFIX50SecurityDefinition(msg fix50.SecurityDefinition, sessionID quickfixgo.SessionID) (reject quickfixgo.MessageReject) {
	reply := msg.ToBuilder()
	quickfixgo.SendToTarget(reply, sessionID)
	return
}

func (e EchoApplication) OnFIX50SP1SecurityDefinition(msg fix50sp1.SecurityDefinition, sessionID quickfixgo.SessionID) (reject quickfixgo.MessageReject) {
	reply := msg.ToBuilder()
	quickfixgo.SendToTarget(reply, sessionID)
	return
}

func (e EchoApplication) OnFIX50SP2SecurityDefinition(msg fix50sp2.SecurityDefinition, sessionID quickfixgo.SessionID) (reject quickfixgo.MessageReject) {
	reply := msg.ToBuilder()
	quickfixgo.SendToTarget(reply, sessionID)
	return
}

func main() {
	app := new(EchoApplication)

	globalSettings := settings.NewDictionary()
	globalSettings.SetInt(settings.SocketAcceptPort, 5001)
	globalSettings.SetString(settings.SenderCompID, "ISLD")
	globalSettings.SetString(settings.TargetCompID, "TW")

	appSettings := settings.NewApplicationSettings(globalSettings)

	appSettings.AddSession("FIX40", settings.NewDictionary().
		SetString(settings.BeginString, quickfixgo.BeginString_FIX40).
		SetString(settings.DataDictionary, "../spec/FIX40.xml"))

	appSettings.AddSession("FIX41", settings.NewDictionary().
		SetString(settings.BeginString, quickfixgo.BeginString_FIX41).
		SetString(settings.DataDictionary, "../spec/FIX41.xml"))

	appSettings.AddSession("FIX42", settings.NewDictionary().
		SetString(settings.BeginString, quickfixgo.BeginString_FIX42).
		SetString(settings.DataDictionary, "../spec/FIX42.xml"))

	appSettings.AddSession("FIX43", settings.NewDictionary().
		SetString(settings.BeginString, quickfixgo.BeginString_FIX43).
		SetString(settings.DataDictionary, "../spec/FIX43.xml"))

	appSettings.AddSession("FIX44", settings.NewDictionary().
		SetString(settings.BeginString, quickfixgo.BeginString_FIX44).
		SetString(settings.DataDictionary, "../spec/FIX44.xml"))

	appSettings.AddSession("FIX50", settings.NewDictionary().
		SetString(settings.BeginString, quickfixgo.BeginString_FIXT11).
		SetString(settings.DefaultApplVerID, quickfixgo.ApplVerID_FIX50))

	appSettings.AddSession("FIX50SP1", settings.NewDictionary().
		SetString(settings.BeginString, quickfixgo.BeginString_FIXT11).
		SetString(settings.DefaultApplVerID, quickfixgo.ApplVerID_FIX50SP1))

	appSettings.AddSession("FIX50SP2", settings.NewDictionary().
		SetString(settings.BeginString, quickfixgo.BeginString_FIXT11).
		SetString(settings.DefaultApplVerID, quickfixgo.ApplVerID_FIX50SP2))

	acceptor, err := quickfixgo.NewAcceptor(app, appSettings, log.ScreenLogFactory{})
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
