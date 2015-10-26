package main

import (
	"fmt"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/fix/field"
	fix40nos "github.com/quickfixgo/quickfix/fix40/newordersingle"
	fix41nos "github.com/quickfixgo/quickfix/fix41/newordersingle"
	fix42nos "github.com/quickfixgo/quickfix/fix42/newordersingle"
	fix42secdef "github.com/quickfixgo/quickfix/fix42/securitydefinition"
	fix43nos "github.com/quickfixgo/quickfix/fix43/newordersingle"
	fix43secdef "github.com/quickfixgo/quickfix/fix43/securitydefinition"
	fix44nos "github.com/quickfixgo/quickfix/fix44/newordersingle"
	fix44secdef "github.com/quickfixgo/quickfix/fix44/securitydefinition"
	fix50nos "github.com/quickfixgo/quickfix/fix50/newordersingle"
	fix50secdef "github.com/quickfixgo/quickfix/fix50/securitydefinition"
	fix50sp1nos "github.com/quickfixgo/quickfix/fix50sp1/newordersingle"
	fix50sp1secdef "github.com/quickfixgo/quickfix/fix50sp1/securitydefinition"
	fix50sp2nos "github.com/quickfixgo/quickfix/fix50sp2/newordersingle"
	fix50sp2secdef "github.com/quickfixgo/quickfix/fix50sp2/securitydefinition"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
)

var router *quickfix.MessageRouter = quickfix.NewMessageRouter()

type EchoApplication struct {
	log      *log.Logger
	OrderIds map[string]bool
}

func (e EchoApplication) OnCreate(sessionID quickfix.SessionID) {
	e.log.Printf("OnCreate %v\n", sessionID.String())
}
func (e *EchoApplication) OnLogon(sessionID quickfix.SessionID) {
	e.log.Printf("OnLogon %v\n", sessionID.String())
	e.OrderIds = make(map[string]bool)
}
func (e *EchoApplication) OnLogout(sessionID quickfix.SessionID) {
	e.log.Printf("OnLogout %v\n", sessionID.String())
}
func (e EchoApplication) ToAdmin(msgBuilder quickfix.Message, sessionID quickfix.SessionID) {
}

func (e EchoApplication) ToApp(msgBuilder quickfix.Message, sessionID quickfix.SessionID) (err error) {
	return
}

func (e EchoApplication) FromAdmin(msg quickfix.Message, sessionID quickfix.SessionID) (err quickfix.MessageRejectError) {
	return
}

func (e *EchoApplication) FromApp(msg quickfix.Message, sessionID quickfix.SessionID) (err quickfix.MessageRejectError) {
	e.log.Println("Got Message ", msg)
	return router.Route(msg, sessionID)
}

func (e *EchoApplication) processMsg(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {

	orderID := new(field.ClOrdIDField)
	if err := msg.Body.Get(orderID); err != nil {
		return err
	}

	reply := copyMessageToBuilder(msg)
	sessionOrderID := sessionID.String() + orderID.Value
	possResend := new(field.PossResendField)
	if err := msg.Header.Get(possResend); err == nil && possResend.Value {
		if e.OrderIds[sessionOrderID] {
			return nil
		}

		reply.Header.Set(possResend)
	}

	e.OrderIds[sessionOrderID] = true

	quickfix.SendToTarget(reply, sessionID)

	return nil
}

func copyMessageToBuilder(msg quickfix.Message) quickfix.Message {
	reply := quickfix.NewMessage()

	msgType := new(field.MsgTypeField)
	msg.Header.Get(msgType)
	reply.Header.Set(msgType)

	for _, tag := range msg.Body.Tags() {
		field := new(quickfix.StringValue)
		if err := msg.Body.GetField(tag, field); err == nil {
			reply.Body.SetField(tag, field)
		}
	}

	return reply
}

func (e *EchoApplication) OnFIX40NewOrderSingle(msg fix40nos.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
	return e.processMsg(msg.Message, sessionID)
}

func (e *EchoApplication) OnFIX41NewOrderSingle(msg fix41nos.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
	return e.processMsg(msg.Message, sessionID)
}

func (e *EchoApplication) OnFIX42NewOrderSingle(msg fix42nos.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
	return e.processMsg(msg.Message, sessionID)
}

func (e *EchoApplication) OnFIX43NewOrderSingle(msg fix43nos.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
	return e.processMsg(msg.Message, sessionID)
}

func (e *EchoApplication) OnFIX44NewOrderSingle(msg fix44nos.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
	return e.processMsg(msg.Message, sessionID)
}

func (e *EchoApplication) OnFIX50NewOrderSingle(msg fix50nos.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
	return e.processMsg(msg.Message, sessionID)
}

func (e *EchoApplication) OnFIX50SP1NewOrderSingle(msg fix50sp1nos.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
	return e.processMsg(msg.Message, sessionID)
}

func (e *EchoApplication) OnFIX50SP2NewOrderSingle(msg fix50sp2nos.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
	return e.processMsg(msg.Message, sessionID)
}

func (e *EchoApplication) OnFIX42SecurityDefinition(msg fix42secdef.Message, sessionID quickfix.SessionID) (err quickfix.MessageRejectError) {
	reply := copyMessageToBuilder(msg.Message)
	quickfix.SendToTarget(reply, sessionID)
	return
}

func (e *EchoApplication) OnFIX43SecurityDefinition(msg fix43secdef.Message, sessionID quickfix.SessionID) (err quickfix.MessageRejectError) {
	reply := copyMessageToBuilder(msg.Message)
	quickfix.SendToTarget(reply, sessionID)
	return
}

func (e *EchoApplication) OnFIX44SecurityDefinition(msg fix44secdef.Message, sessionID quickfix.SessionID) (err quickfix.MessageRejectError) {
	reply := copyMessageToBuilder(msg.Message)
	quickfix.SendToTarget(reply, sessionID)
	return
}

func (e *EchoApplication) OnFIX50SecurityDefinition(msg fix50secdef.Message, sessionID quickfix.SessionID) (err quickfix.MessageRejectError) {
	reply := copyMessageToBuilder(msg.Message)
	quickfix.SendToTarget(reply, sessionID)
	return
}

func (e *EchoApplication) OnFIX50SP1SecurityDefinition(msg fix50sp1secdef.Message, sessionID quickfix.SessionID) (err quickfix.MessageRejectError) {
	reply := copyMessageToBuilder(msg.Message)
	quickfix.SendToTarget(reply, sessionID)
	return
}

func (e *EchoApplication) OnFIX50SP2SecurityDefinition(msg fix50sp2secdef.Message, sessionID quickfix.SessionID) (err quickfix.MessageRejectError) {
	reply := copyMessageToBuilder(msg.Message)
	quickfix.SendToTarget(reply, sessionID)
	return
}

func main() {
	app := &EchoApplication{}
	app.log = log.New(ioutil.Discard, "", log.LstdFlags)
	//app.log = log.New(os.Stdout, "", log.LstdFlags)

	router.AddRoute(fix40nos.Route(app.OnFIX40NewOrderSingle))
	router.AddRoute(fix41nos.Route(app.OnFIX41NewOrderSingle))
	router.AddRoute(fix42nos.Route(app.OnFIX42NewOrderSingle))
	router.AddRoute(fix43nos.Route(app.OnFIX43NewOrderSingle))
	router.AddRoute(fix44nos.Route(app.OnFIX44NewOrderSingle))
	router.AddRoute(fix50nos.Route(app.OnFIX50NewOrderSingle))
	router.AddRoute(fix50sp1nos.Route(app.OnFIX50SP1NewOrderSingle))
	router.AddRoute(fix50sp2nos.Route(app.OnFIX50SP2NewOrderSingle))

	router.AddRoute(fix42secdef.Route(app.OnFIX42SecurityDefinition))
	router.AddRoute(fix43secdef.Route(app.OnFIX43SecurityDefinition))
	router.AddRoute(fix44secdef.Route(app.OnFIX44SecurityDefinition))
	router.AddRoute(fix50secdef.Route(app.OnFIX50SecurityDefinition))
	router.AddRoute(fix50sp1secdef.Route(app.OnFIX50SP1SecurityDefinition))
	router.AddRoute(fix50sp2secdef.Route(app.OnFIX50SP2SecurityDefinition))

	cfg, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Error opening %v, %v\n", os.Args[1], err)
		return
	}

	appSettings, err := quickfix.ParseSettings(cfg)
	if err != nil {
		fmt.Println("Error reading cfg:", err)
		return
	}

	fileLogFactory, err := quickfix.NewFileLogFactory(appSettings)

	if err != nil {
		fmt.Println("Error creating file log factory:", err)
		return
	}

	acceptor, err := quickfix.NewAcceptor(app, quickfix.NewMemoryStoreFactory(), appSettings, fileLogFactory)
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
