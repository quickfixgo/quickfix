package session

import(
  "net"
  "bufio"
  "quickfixgo/log"
  "quickfixgo/fix"
  "quickfixgo/message/basic"
    )

type connection struct {
  netConn net.Conn
  session *Session
  nextMessage chan []byte
}

func (c * connection) cleanup() {
  c.netConn.Close()
  Deactivate(c.session.ID)
}

func newConnection(netConn net.Conn, s *Session) *connection {
  c:=new(connection)
  c.netConn=netConn
  c.nextMessage=make(chan []byte)
  c.session = s

  return c
}

func (c * connection) sessionLoop() {
  defer c.cleanup()

  disconnect:=false

  for {
    select {
      case msgBytes,ok:= <-c.nextMessage:
        if ok {
          disconnect=c.session.fixMsgIn(msgBytes)
        } else {
          c.session.onDisconnect()
          return
        }
    }

    if disconnect {return}
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
