package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/config"
	field "github.com/quickfixgo/quickfix/gen/field"
	tag "github.com/quickfixgo/quickfix/gen/tag"
	filelog "github.com/quickfixgo/quickfix/log/file"
	"github.com/quickfixgo/quickfix/store/file"
	"github.com/quickfixgo/quickfix/store/mongo"
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

		if bytes.Equal(possResend.Write(), []byte("Y")) {
			if e.OrderIds[sessionOrderID] {
				return nil
			}
		}

		e.OrderIds[sessionOrderID] = true
	}

	reply := copyMessage(msg)
	if bytes.Equal(possResend.Write(), []byte("Y")) {
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
	app.log = log.New(io.Discard, "", log.LstdFlags)
	//app.log = log.New(os.Stdout, "", log.LstdFlags)

	router.AddRoute(quickfix.BeginStringFIX40, "D", app.processMsg)
	router.AddRoute(quickfix.BeginStringFIX41, "D", app.processMsg)
	router.AddRoute(quickfix.BeginStringFIX42, "D", app.processMsg)
	router.AddRoute(quickfix.BeginStringFIX43, "D", app.processMsg)
	router.AddRoute(quickfix.BeginStringFIX44, "D", app.processMsg)
	router.AddRoute(quickfix.ApplVerIDFIX50, "D", app.processMsg)
	router.AddRoute(quickfix.ApplVerIDFIX50SP1, "D", app.processMsg)
	router.AddRoute(quickfix.ApplVerIDFIX50SP2, "D", app.processMsg)

	router.AddRoute(quickfix.BeginStringFIX42, "d", app.processMsg)
	router.AddRoute(quickfix.BeginStringFIX43, "d", app.processMsg)
	router.AddRoute(quickfix.BeginStringFIX44, "d", app.processMsg)
	router.AddRoute(quickfix.ApplVerIDFIX50, "d", app.processMsg)
	router.AddRoute(quickfix.ApplVerIDFIX50SP1, "d", app.processMsg)
	router.AddRoute(quickfix.ApplVerIDFIX50SP2, "d", app.processMsg)

	fmt.Println("starting test server")
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

	fileLogFactory, err := filelog.NewLogFactory(appSettings)
	if err != nil {
		fmt.Println("Error creating file log factory:", err)
		return
	}
	//fileLogFactory := quickfix.NewScreenLogFactory()
	// if err != nil {
	// 	fmt.Println("Error creating file log factory:", err)
	// 	return
	// }

	storeType := os.Args[2]

	var acceptor *quickfix.Acceptor
	switch strings.ToUpper(storeType) {
	case "MONGO":
		mongoDbCxn := "mongodb://localhost:27017"
		mongoDatabase := "automated_testing_database"
		mongoReplicaSet := "replicaset"

		appSettings.GlobalSettings().Set(config.MongoStoreConnection, mongoDbCxn)
		appSettings.GlobalSettings().Set(config.MongoStoreDatabase, mongoDatabase)
		appSettings.GlobalSettings().Set(config.MongoStoreReplicaSet, mongoReplicaSet)
		appSettings.GlobalSettings().Set(config.DynamicSessions, "Y")

		acceptor, err = quickfix.NewAcceptor(app, mongo.NewStoreFactory(appSettings), appSettings, fileLogFactory)
	case "FILE":
		fileStoreRootPath := path.Join(os.TempDir(), fmt.Sprintf("FileStoreTestSuite-%d", os.Getpid()))
		fileStorePath := path.Join(fileStoreRootPath, fmt.Sprintf("%d", time.Now().UnixNano()))
		appSettings.GlobalSettings().Set(config.FileStorePath, fileStorePath)
		appSettings.GlobalSettings().Set(config.DynamicSessions, "Y")

		acceptor, err = quickfix.NewAcceptor(app, file.NewStoreFactory(appSettings), appSettings, fileLogFactory)
	case "MEMORY":
		fallthrough
	default:
		acceptor, err = quickfix.NewAcceptor(app, quickfix.NewMemoryStoreFactory(), appSettings, fileLogFactory)
	}

	if err != nil {
		fmt.Println("Unable to create Acceptor: ", err)
		return
	}

	if err = acceptor.Start(); err != nil {
		fmt.Println("Unable to start Acceptor: ", err)
		return
	}

	// http handler for seqnum resets for specific test defs. not applicable for most tests.
	http.HandleFunc("/seqnum", func(w http.ResponseWriter, r *http.Request) {
		defer w.WriteHeader(http.StatusOK)
		queryParams := r.URL.Query()
		sessionStr, ok := queryParams["SESSION"]
		if !ok {
			fmt.Println("cannot find session")
			os.Exit(1)
		}
		splts := strings.Split(sessionStr[0], ":")
		parties := strings.Split(splts[1], "->")
		sessionID := quickfix.SessionID{
			BeginString:  splts[0],
			SenderCompID: parties[0],
			TargetCompID: parties[1],
		}
		num, ok := queryParams["NEXTTARGETSEQNUM"]
		if ok {
			seqnumInt, cErr := strconv.Atoi(num[0])
			if cErr != nil {
				fmt.Println("cannot find seqnum")
				os.Exit(1)
			}
			setErr := quickfix.SetNextTargetMsgSeqNum(sessionID, seqnumInt)
			if setErr != nil {
				fmt.Println("err setting target seqnum: " + setErr.Error())
				return
			}
			fmt.Println("set target seqnum to: " + num[0])
			return
		}

		num, ok = queryParams["NEXTSENDERSEQNUM"]
		if ok {
			seqnumInt, cErr := strconv.Atoi(num[0])
			if cErr != nil {
				fmt.Println("cannot find seqnum")
				os.Exit(1)
			}
			setErr := quickfix.SetNextSenderMsgSeqNum(sessionID, seqnumInt)
			if setErr != nil {
				fmt.Println("err setting sender seqnum: " + setErr.Error())
				return
			}
			fmt.Println("set sender seqnum to: " + num[0])
			return
		}

		fmt.Println("never found seqnum")
		os.Exit(1)
	})
	websrvr := http.Server{Addr: ":8095"}
	defer func() {
		if err := websrvr.Close(); err != nil {
			fmt.Println("unable to stop srver: ", err)
		}
	}()
	go func() {
		if err := websrvr.ListenAndServe(); err != http.ErrServerClosed {
			fmt.Println("websrvr err: ", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	fmt.Println("stopping test server")

	acceptor.Stop()
}
