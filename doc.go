/*
Package quickfix is a full featured messaging engine for the FIX protocol.  It is a 100% Go open source implementaiton of the popular C++ QuickFIX engine (http://quickfixengine.org).

Creating your QuickFIX Application

Creating a FIX application is as easy as implementing the QuickFIX Application interface. By implementing these interface methods in your derived class, you are requesting to be notified of events that occur on the FIX engine.  The function that you should be most aware of is fromApp.

Here are explanations of what these functions provide for you.

 OnCreate(sessionID SessionID)

This method is called when quickfix creates a new session. A session comes into and remains in existence for the life of the application. Sessions exist whether or not a counter party is connected to it. As soon as a session is created, you can begin sending messages to it. If no one is logged on, the messages will be sent at the time a connection is established with the counterparty.

 OnLogon(sessionID SessionID)

This callback notifies you when a valid logon has been established with a counter party. This is called when a connection has been established and the FIX logon process has completed with both parties exchanging valid logon messages.

 OnLogout(sessionID SessionID)

This callback notifies you when an FIX session is no longer online. This could happen during a normal logout exchange or because of a forced termination or a loss of network connection.

	ToAdmin(message Message, sessionID SessionID)

This callback provides you with a peak at the administrative messages that are being sent from your FIX engine to the counter party. This is normally not useful for an application however it is provided for any logging you may wish to do. Notice that the Message is not const. This allows you to add fields before an adminstrative message before it is sent out.

	ToApp(message Message, sessionID SessionID) error

This is a callback for application messages that you are being sent to a counterparty. Notice that the Message is not const. This allows you to add fields before an application message before it is sent out.

	FromAdmin(message Message, sessionID SessionID) MessageRejectError

This callback notifies you when an administrative message is sent from a counterparty to your FIX engine. This can be usefull for doing extra validation on logon messages such as for checking passwords.

	FromApp(msg Message, sessionID SessionID) MessageRejectError

This is one of the core entry points for your FIX application. Every application level request will come through here. If, for example, your application is a sell-side OMS, this is where you will get your new order requests. If you were a buy side, you would get your execution reports here.

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

	//FooApplication is your class that implements the Application interface
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
*/
package quickfix
