/*
Package quickfix is a full featured messaging engine for the FIX protocol.  It is a 100% Go open source implementation of the popular C++ QuickFIX engine (http://quickfixengine.org).

Creating your QuickFIX Application

Creating a FIX application is as easy as implementing the QuickFIX Application interface. By implementing these interface methods in your derived type, you are requesting to be notified of events that occur on the FIX engine.  The function that you should be most aware of is fromApp.

Here are explanations of what these callback functions provide for you.

 OnCreate(sessionID SessionID)

OnCreate is called when quickfix creates a new session. A session comes into and remains in existence for the life of the application. Sessions exist whether or not a counter party is connected to it. As soon as a session is created, you can begin sending messages to it. If no one is logged on, the messages will be sent at the time a connection is established with the counterparty.

 OnLogon(sessionID SessionID)

OnLogon notifies you when a valid logon has been established with a counter party. This is called when a connection has been established and the FIX logon process has completed with both parties exchanging valid logon messages.

 OnLogout(sessionID SessionID)

OnLogout notifies you when an FIX session is no longer online. This could happen during a normal logout exchange or because of a forced termination or a loss of network connection.

	ToAdmin(message Message, sessionID SessionID)

ToAdmin provides you with a peak at the administrative messages that are being sent from your FIX engine to the counter party. This is normally not useful for an application however it is provided for any logging you may wish to do. Notice that the Message is not const. This allows you to add fields before an administrative message is sent out.

	ToApp(message Message, sessionID SessionID) error

ToApp notifies you of application messages that you are being sent to a counterparty. Notice that the Message is not const. This allows you to add fields before an application message before it is sent out.

	FromAdmin(message Message, sessionID SessionID) MessageRejectError

FromAdmin notifies you when an administrative message is sent from a counterparty to your FIX engine. This can be usefull for doing extra validation on logon messages such as for checking passwords.

	FromApp(msg Message, sessionID SessionID) MessageRejectError

FromApp is one of the core entry points for your FIX application. Every application level request will come through here. If, for example, your application is a sell-side OMS, this is where you will get your new order requests. If you were a buy side, you would get your execution reports here.

The sample code below shows how you might start up a FIX acceptor which listens on a socket. If you wanted an initiator, you would simply replace NewAcceptor in this code fragment with NewInitiator.

 package main

 import (
	"flag"
	"github.com/quickfixgo/quickfix"
	"os"
 )

 func main() {
	flag.Parse()
	fileName := flag.Arg(0)

	//FooApplication is your type that implements the Application interface
	var app FooApplication

	cfg, _ := os.Open(fileName)
	appSettings, _ := quickfix.ParseSettings(cfg)
	storeFactory := quickfix.NewMemoryStoreFactory()
	logFactory, _ := quickfix.NewFileLogFactory(appSettings)
	acceptor, _ := quickfix.NewAcceptor(app, storeFactory, appSettings, logFactory)

	acceptor.Start()
	//for condition == true { do something }
	acceptor.Stop()
 }

Configuring QuickFIX

A QuickFIXGo acceptor or initiator can maintain as many FIX sessions as you would like. A FIX session is identified by a group of settings defined within the configuration section for a session (or inherited from the default section). The identification settings are:

 | Setting      | Required? |
 |--------------|-----------|
 | BeginString  |     Y     |
 | SenderCompID |     Y     |
 | SenderSubID  |     N     |
 | TargetCompID |     Y     |
 | TargetSubID  |     N     |

The sender settings are your identification and the target settings are for the counterparty. A SessionQualifier can also be use to disambiguate otherwise identical sessions.

Each of the sessions can have several settings associated with them. Some of these settings may not be known at compile time and are therefore passed around in a struct called SessionSettings.

SessionSettings can be read in from any input stream such as a file stream. If you decide to write your own components, (storage for a particular database, a new kind of connector etc...), you may also use the session settings to store settings for your custom component.

A settings file is set up with two types of heading, a [DEFAULT] and a [SESSION] heading. [SESSION] tells QuickFIX/Go that a new Session is being defined. [DEFAULT] is a place that you can define settings which will be inherited by sessions that don't explicitly define them. If you do not provide a setting that QuickFIX/Go needs, it will return an error indicating that a setting is missing or improperly formatted.

See the quickfix/config package for settings you can associate with a session based on the default components provided with QuickFIX/Go.

Here is a typical initiator settings file you might find for a firm that wants to connect to several ECNs.

 # default settings for sessions
 [DEFAULT]
 FileLogPath=log
 SenderCompID=TW

 # session definition
 [SESSION]
 # inherit FileLogPath, SenderCompID from default
 BeginString=FIX.4.1
 TargetCompID=ARCA
 HeartBtInt=20
 SocketConnectPort=9823
 SocketConnectHost=123.123.123.123
 DataDictionary=somewhere/FIX41.xml

 [SESSION]
 BeginString=FIX.4.0
 TargetCompID=ISLD
 HeartBtInt=30
 SocketConnectPort=8323
 SocketConnectHost=23.23.23.23
 DataDictionary=somewhere/FIX40.xml

 [SESSION]
 BeginString=FIX.4.2
 TargetCompID=INCA
 HeartBtInt=30
 SocketConnectPort=6523
 SocketConnectHost=3.3.3.3
 DataDictionary=somewhere/FIX42.xml

Receiving Messages

Most of the messages you will be interested in looking at will be arriving in your FromApp function of your application.  All messages have a header and trailer.  If you want to get Header or Trailer fields, you must access those fields from the Header or Trailer embedded Struct.  All other fields are accessible in the Body embedded struct.

QuickFIX/Go has a type for all messages and fields defined in the standard spec.  The easiest and most typesafe method of receiving messages is by using the quickfix MessageRouter generated message types. Any messages you do not establish routes for will by default return an UnsupportedMessageType reject.

 import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix41/newordersingle"
 )

 type MyApplication struct {
	*quickfix.MessageRouter
 }

 func (m *MyApplication) init() {
	m.MessageRouter=quickfix.NewMessageRouter()
	m.AddRoute(newordersingle.Route(m.onNewOrderSingle))
 }

 func (m *MyApplication) FromApp(msg quickfix.Message, sessionID quickfix.SessionID) (err quickfix.MessageRejectError) {
	return m.Route(msg, sessionID)
 }

 func (m *MyApplication) onNewOrderSingle(msg newordersingle.NewOrderSingle, sessionID quickfix.SessionID) (err quickfix.MessageRejectError) {
	var clOrdID field.ClOrdIDField
	if clOrdID, err = msg.GetClOrdID(); err!=nil {
		return
	}

	//compile time error!! field not defined in FIX41
	var clearingAccount field.ClearingAccountField
	clearingAccount, err = msg.GetClearingAccount()

	...
	return
 }

You can also bypass the MessageRouter and type safe classes by inspecting the Message directly.  The preferred way of doing this is to use the quickfix generated Field types.

 func (m *MyApplication) FromApp(msg quickfix.Message, sessionID quickfix.SessionID) (err quickfix.MessageRejectError) {
	var price field.PriceField
	if err = msg.Body.Get(&field); err!=nil {
		return
	}

	...
	return
 }


Or you can go the least type safe route.

 func (m *MyApplication) FromApp(msg quickfix.Message, sessionID quickfix.SessionID) (err quickfix.MessageRejectError) {
	var field quickfix.FIXString
	if err = msg.Body.GetField(quickfix.Tag(44), &field); err!=nil {
		return
	}

	...
	return
 }

Sending Messages

Messages can be sent to the counter party with the Send and SendToTarget functions.

 //Send determines the session to send Messagable using header fields BeginString, TargetCompID, SenderCompID
 func Send(m Messagable) error

 //SendToTarget sends Messagable based on the sessionID. Convenient for use in FromApp since it provides a session ID for incoming messages
 func SendToTarget(m Messagable, sessionID SessionID) error

*/
package quickfix
