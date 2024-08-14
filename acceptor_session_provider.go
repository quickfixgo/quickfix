package quickfix

import "github.com/quickfixgo/quickfix/config"

const (
	WildcardPattern string = "*"
)

type AcceptorSessionProvider interface {
	GetSession(SessionID) (*session, error)
}

type StaticAcceptorSessionProvider struct {
	sessions map[SessionID]*session
}

func (p *StaticAcceptorSessionProvider) GetSession(sessionID SessionID) (*session, error) {
	s, ok := p.sessions[sessionID]
	if !ok {
		return nil, errUnknownSession
	}
	return s, nil
}

// DynamicAcceptorSessionProvider dynamically defines sessions for an acceptor. This can be useful for
// applications like simulators that want to accept any connection and
// dynamically create an associated session.
//
// For more complex situations, you can use this class as a starting
// point for implementing your own AcceptorSessionProvider.
type DynamicAcceptorSessionProvider struct {
	settings            *Settings
	messageStoreFactory MessageStoreFactory
	logFactory          LogFactory
	sessionFactory      *sessionFactory
	application         Application
	templateMappings    []*TemplateMapping
}

func NewDynamicAcceptorSessionProvider(settings *Settings, messageStoreFactory MessageStoreFactory, logFactory LogFactory,
	application Application, templateMappings []*TemplateMapping,
) *DynamicAcceptorSessionProvider {
	return &DynamicAcceptorSessionProvider{
		settings:            settings,
		messageStoreFactory: messageStoreFactory,
		logFactory:          logFactory,
		sessionFactory:      &sessionFactory{},
		application:         application,
		templateMappings:    templateMappings,
	}
}

func (p *DynamicAcceptorSessionProvider) FindTemplateID(sessionID SessionID) *SessionID {
	return p.lookupTemplateID(sessionID)
}

func (p *DynamicAcceptorSessionProvider) GetSession(sessionID SessionID) (*session, error) {
	s, ok := lookupSession(sessionID)
	if !ok {
		templateID := p.lookupTemplateID(sessionID)
		if templateID == nil {
			return nil, errUnknownSession
		}
		dynamicSessionSettings := p.settings.globalSettings.clone()
		templateSettings := p.settings.sessionSettings[*templateID]
		dynamicSessionSettings.overlay(templateSettings)
		dynamicSessionSettings.Set(config.BeginString, sessionID.BeginString)
		dynamicSessionSettings.Set(config.SenderCompID, sessionID.SenderCompID)
		dynamicSessionSettings.Set(config.SenderSubID, sessionID.SenderSubID)
		dynamicSessionSettings.Set(config.SenderLocationID, sessionID.SenderLocationID)
		dynamicSessionSettings.Set(config.TargetCompID, sessionID.TargetCompID)
		dynamicSessionSettings.Set(config.TargetSubID, sessionID.TargetSubID)
		dynamicSessionSettings.Set(config.TargetLocationID, sessionID.TargetLocationID)
		var err error
		s, err = p.sessionFactory.createSession(sessionID,
			p.messageStoreFactory,
			dynamicSessionSettings,
			p.logFactory,
			p.application,
		)
		if err != nil {
			return nil, err
		}
	}
	return s, nil
}

func (provider *DynamicAcceptorSessionProvider) lookupTemplateID(sessionID SessionID) *SessionID {
	for _, mapping := range provider.templateMappings {
		if isTemplateMatching(mapping.Pattern, sessionID) {
			return &mapping.TemplateID
		}
	}
	return nil
}

func isTemplateMatching(pattern SessionID, sessionID SessionID) bool {
	return matches(pattern.BeginString, sessionID.BeginString) &&
		matches(pattern.SenderCompID, sessionID.SenderCompID) &&
		matches(pattern.SenderSubID, sessionID.SenderSubID) &&
		matches(pattern.SenderLocationID, sessionID.SenderLocationID) &&
		matches(pattern.TargetCompID, sessionID.TargetCompID) &&
		matches(pattern.TargetSubID, sessionID.TargetSubID) &&
		matches(pattern.TargetLocationID, sessionID.TargetLocationID)
}

func matches(pattern string, value string) bool {
	return WildcardPattern == pattern || pattern == value
}

// TemplateMapping mapping from a sessionID pattern to a session template ID.
type TemplateMapping struct {
	Pattern    SessionID
	TemplateID SessionID
}
