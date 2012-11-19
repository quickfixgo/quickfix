package quickfixgo

import(
    "net"
    "errors"
    "strconv"
    "bufio"
    "quickfixgo/log"
    "quickfixgo/fix"
    "quickfixgo/message/basic"
    )

type Acceptor interface {
  //Start Acceptor.
  Start() error

  //Stop Acceptor.
  Stop()
}

type acceptor struct {
  app Application
  settings SessionSettings
  logFactory log.LogFactory
  globalLog log.Log
}

func (a * acceptor) Start() (e error) {
  port,hasPort:=a.settings.GetGlobalSettings().GetInt(SocketAcceptPort)
  if !hasPort {
    return errors.New("Config Error: Must Provide SocketAcceptPort")
  }

  server, err := net.Listen("tcp",":" + strconv.Itoa(port))
  if server == nil {
    return err
  }

  connections := a.listenForConnections(server)
  go func() {
    for {
      cxn := <-connections
      go a.handleConnection(cxn)
    }
  }()

  return
}

func (a acceptor) Stop() {}

func NewAcceptor(app Application, settings SessionSettings, logFactory log.LogFactory) (Acceptor, error) {
  a:=new(acceptor)
  a.app=app
  a.settings=settings
  a.logFactory=logFactory
  a.globalLog=logFactory.Create()

  return a,nil
}

func (a * acceptor) handleConnection(connection net.Conn) {
  defer func() {
     if err := recover(); err != nil {
       connection.Close()
     }
   }()

  reader := bufio.NewReader(connection)
  parser := newParser(reader)

  msgBytes,err:=parser.readMessage()
  if err !=nil {
    connection.Close()
    return
  }

  msg,err:=basic.MessageFromParsedBytes(msgBytes)
  if err != nil {
    a.globalLog.OnEvent("Invalid message: " + string(msgBytes) + err.Error())
    connection.Close()
    return
  }

  var sessID SessionID
  if beginString,err:=msg.Header().StringField(fix.BeginString); err!=nil {
    sessID.BeginString=beginString.Value()
  }

  if senderCompID,err:=msg.Header().StringField(fix.SenderCompID); err!=nil {
    sessID.SenderCompID=senderCompID.Value()
  }

  if targetCompID,err:=msg.Header().StringField(fix.TargetCompID); err!=nil {
    sessID.TargetCompID=targetCompID.Value()
  }

  if defaultApplVerID,err:= msg.Body().StringField(fix.DefaultApplVerID); err!=nil {
    sessID.DefaultApplVerID = defaultApplVerID.Value()
  }




}

func (a * acceptor) listenForConnections(listener net.Listener) (ch chan net.Conn) {
  ch = make(chan net.Conn)

  go func() {
    for {
      netConn, err := listener.Accept()
      if netConn == nil {
        a.globalLog.OnEventf("Couldn't Accept: %v", err.Error())
        continue
      }

      ch <-netConn
    }
  }()

  return ch
}
