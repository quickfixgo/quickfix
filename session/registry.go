package session

type Registry struct {
  newSession chan Session
}

var Sessions *Registry

func init() {
  Sessions=new(Registry)
}

func (registry *Registry) Add(session Session) {registry.newSession <- session}
