package session

import(
  "net"
  "bufio"
  "quickfixgo/log"
  "quickfixgo/fix"
  "quickfixgo/message"
  "quickfixgo/message/basic"
    )

type connection struct {
  netConn net.Conn
  session *session
  nextMessage chan []byte
  writerShutDown chan int
}

func (c * connection) cleanup() {
  c.netConn.Close()
  c.writerShutDown<-1

  deactivate(c.session.ID)
}

func newConnection(netConn net.Conn, s *session) *connection {
  c:=new(connection)
  c.netConn=netConn
  c.nextMessage=make(chan []byte)
  c.session = s

  c.writerShutDown=writeLoop(netConn, s.MessageOut)

  return c
}

func (c * connection) sessionLoop() {
  defer c.cleanup()

  for {
    select {
      case msgBytes,ok:= <-c.nextMessage:
        if ok {
          if disconnect:=c.session.fixMsgIn(msgBytes); disconnect {
            return
          }
        } else {
          c.session.onDisconnect()
          return
        }
    }
  }
}



func HandleAcceptorConnection(netConn net.Conn, log log.Log) {
  defer func() {
     if err := recover(); err != nil {
       netConn.Close()
     }
   }()

  reader := bufio.NewReader(netConn)
  parser := newParser(reader)

  msgBytes,err:=parser.readMessage()
  if err !=nil {
    netConn.Close()
    return
  }

  msg,err:=basic.MessageFromParsedBytes(msgBytes)
  if err != nil {
    log.OnEvent("Invalid message: " + string(msgBytes) + err.Error())
    netConn.Close()
    return
  }

  var sessID ID
  if beginString,err:=msg.Header().StringField(fix.BeginString); err==nil {
    sessID.BeginString=beginString.Value()
  }

  if senderCompID,err:=msg.Header().StringField(fix.SenderCompID); err==nil {
    sessID.TargetCompID=senderCompID.Value()
  }

  if targetCompID,err:=msg.Header().StringField(fix.TargetCompID); err==nil {
    sessID.SenderCompID=targetCompID.Value()
  }

  if defaultApplVerID,err:= msg.Body().StringField(fix.DefaultApplVerID); err==nil {
    sessID.DefaultApplVerID = defaultApplVerID.Value()
  }

  s:=Activate(sessID)

  if s == nil {
    log.OnEvent("Session not found for incoming message: " + msg.String())
    netConn.Close()
    return
  }

  connection:=newConnection(netConn, s)
  go func() { connection.nextMessage <- msgBytes}()

  connection.sessionLoop()
}

func writeLoop(connection net.Conn, messageOut chan message.Buffer) (shutDown chan int) {
  shutDown=make(chan int,1)
  go func() {
    for {
      select {
        case <-shutDown:
          return
        case msg:=<-messageOut:
          connection.Write(msg.Bytes())
          msg.Free()
      }
    }
  }()

  return
}
