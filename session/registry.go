package session

type sessionActivate struct {
  ID
  reply chan *session
}

type sessionResource struct {
  session *session
  active bool
}

type registry struct {
  newSession chan *session
  activate chan sessionActivate
  deactivate chan ID
}

var sessions *registry

func init() {
  sessions=new(registry)
  sessions.newSession = make(chan *session)
  sessions.activate = make(chan sessionActivate)
  sessions.deactivate = make(chan ID)

  go sessions.sessionResourceServerLoop()
}

func activate(sessionID ID) *session {
  response:=make(chan *session)
  sessions.activate<-sessionActivate{sessionID, response}
  return <- response
}

func deactivate(sessionID ID) {
  sessions.deactivate<-sessionID
}

func (r *registry) sessionResourceServerLoop() {
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
