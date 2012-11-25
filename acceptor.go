package quickfixgo

import(
    "net"
    "errors"
    "strconv"
    "quickfixgo/log"
    "quickfixgo/session"
    "quickfixgo/settings"
    )

type Acceptor interface {
  //Start Acceptor.
  Start() error

  //Stop Acceptor.
  Stop()
}

type acceptor struct {
  app Application
  settings settings.ApplicationSettings
  logFactory log.LogFactory
  globalLog log.Log
}

func (a * acceptor) Start() (e error) {
  port,hasPort:=a.settings.GetGlobalSettings().GetInt(settings.SocketAcceptPort)
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
      go session.HandleAcceptorConnection(cxn, a.globalLog)
    }
  }()

  return
}

func (a acceptor) Stop() {}

func NewAcceptor(app Application, settings settings.ApplicationSettings, logFactory log.LogFactory) (Acceptor, error) {
  a:=new(acceptor)
  a.app=app
  a.settings=settings
  a.logFactory=logFactory
  a.globalLog=logFactory.Create()

  for _,s:=range settings.GetSessionSettings() {
    if err:=session.Create(s,logFactory); err!=nil {
      return nil, err
    }
  }

  return a,nil
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
