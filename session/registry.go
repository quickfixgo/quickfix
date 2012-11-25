package session

type sessionActivate struct {
  ID
  reply chan *Session
}

type sessionResource struct {
  session *Session
  active bool
}

type Registry struct {
  newSession chan *Session
  activate chan sessionActivate
  deactivate chan ID
}

var Sessions *Registry

func init() {
  Sessions=new(Registry)
  Sessions.newSession = make(chan *Session)
  Sessions.activate = make(chan sessionActivate)
  Sessions.deactivate = make(chan ID)

  go Sessions.sessionResourceServerLoop()
}

func Activate(sessionID ID) *Session {
  response:=make(chan *Session)
  Sessions.activate<-sessionActivate{sessionID, response}
  return <- response
}

func Deactivate(sessionID ID) {
  Sessions.deactivate<-sessionID
}

func (r *Registry) sessionResourceServerLoop() {
  sessions:=make(map[ID] sessionResource)

  for{
    select {
      case session:=<-r.newSession:
        sessions[session.ID]=sessionResource{session,false}

      case deactivatedID:=<-r.deactivate:
        if resource, ok:=sessions[deactivatedID]; ok {
          resource.active = false
        }

      case request:=<-r.activate:
        resource, ok:=sessions[request.ID]

        switch {
          case !ok:
            request.reply<- nil

          case resource.active:
            request.reply<- nil

          default:
            resource.active = true
            request.reply<-resource.session
        }
    }
  }
}
