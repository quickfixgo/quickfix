package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/tag"
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
func (e EchoApplication) ToAdmin(msgBuilder *quickfix.Message, sessionID quickfix.SessionID) {
}

func (e EchoApplication) ToApp(msgBuilder *quickfix.Message, sessionID quickfix.SessionID) (err error) {
	return
}

func (e EchoApplication) FromAdmin(msg *quickfix.Message, sessionID quickfix.SessionID) (err quickfix.MessageRejectError) {
	return
}

func (e *EchoApplication) FromApp(msg *quickfix.Message, sessionID quickfix.SessionID) (err quickfix.MessageRejectError) {
	e.log.Println("Got Message ", msg)
	return router.Route(msg, sessionID)
}

func (e *EchoApplication) processMsg(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
	var possResend field.PossResendField
	msg.Header.Get(&possResend)

	if msg.Body.Has(tag.ClOrdID) {
		var orderID field.ClOrdIDField

		if err := msg.Body.Get(&orderID); err != nil {
			return err
		}

		sessionOrderID := sessionID.String() + orderID.String()
		if possResend.FIXBoolean {
			if e.OrderIds[sessionOrderID] {
				return nil
			}
		}

		e.OrderIds[sessionOrderID] = true
	}

	reply := copyMessage(msg)
	if possResend.FIXBoolean {
		reply.Header.Set(possResend)
	}

	quickfix.SendToTarget(reply, sessionID)

	return nil
}

func copyMessage(msg *quickfix.Message) *quickfix.Message {
	msgType := new(field.MsgTypeField)
	msg.Header.Get(msgType)

	msg.Header.Clear()
	msg.Trailer.Clear()

	msg.Header.Set(msgType)

	return msg
}

func main() {
	app := &EchoApplication{}
	app.log = log.New(ioutil.Discard, "", log.LstdFlags)
	//app.log = log.New(os.Stdout, "", log.LstdFlags)

	router.AddRoute(enum.BeginStringFIX40, "D", app.processMsg)
	router.AddRoute(enum.BeginStringFIX41, "D", app.processMsg)
	router.AddRoute(enum.BeginStringFIX42, "D", app.processMsg)
	router.AddRoute(enum.BeginStringFIX43, "D", app.processMsg)
	router.AddRoute(enum.BeginStringFIX44, "D", app.processMsg)
	router.AddRoute(string(enum.ApplVerID_FIX50), "D", app.processMsg)
	router.AddRoute(string(enum.ApplVerID_FIX50SP1), "D", app.processMsg)
	router.AddRoute(string(enum.ApplVerID_FIX50SP2), "D", app.processMsg)

	router.AddRoute(enum.BeginStringFIX42, "d", app.processMsg)
	router.AddRoute(enum.BeginStringFIX43, "d", app.processMsg)
	router.AddRoute(enum.BeginStringFIX44, "d", app.processMsg)
	router.AddRoute(string(enum.ApplVerID_FIX50), "d", app.processMsg)
	router.AddRoute(string(enum.ApplVerID_FIX50SP1), "d", app.processMsg)
	router.AddRoute(string(enum.ApplVerID_FIX50SP2), "d", app.processMsg)

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
