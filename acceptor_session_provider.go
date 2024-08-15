package quickfix

import "github.com/quickfixgo/quickfix/config"

const (
	WildcardPattern string = "*"
)

// TemplateIDProvider is an interface for obtaining templateIDs for inbound sessions.
//
//	The SessionSettings for the template SessionID must be configured in the Acceptor.
//	The sessionSettings of an inbound session inherits from the template SessionID.
//	If no matching template is found, return nil, and no session will be created for the inbound logon request.
type TemplateIDProvider interface {
	GetTemplateID(inbound SessionID) (templateID *SessionID)
}

type TemplateIDProviderSetter interface {
	SetTemplateIDProvider(TemplateIDProvider)
}

type DefaultTemplateIDProvider struct {
	templateMappings []*TemplateMapping
}

func NewDefaultTemplateIDProvider(settings *Settings) (*DefaultTemplateIDProvider, error) {
	templateMappings := make([]*TemplateMapping, 0)
	for sid, ss := range settings.SessionSettings() {
		var (
			acceptorTemplate bool
			err              error
		)
		if ss.HasSetting(config.AcceptorTemplate) {
			acceptorTemplate, err = ss.BoolSetting(config.AcceptorTemplate)
			if err != nil {
				return nil, err
			}
		}
		if acceptorTemplate {
			templateMappings = append(templateMappings, &TemplateMapping{
				Pattern:    sid,
				TemplateID: sid,
			})
		}
	}
	return &DefaultTemplateIDProvider{templateMappings: templateMappings}, nil
}

func (p *DefaultTemplateIDProvider) GetTemplateID(inbound SessionID) (templateID *SessionID) {
	return p.lookupTemplateID(inbound)
}

func (p *DefaultTemplateIDProvider) lookupTemplateID(sessionID SessionID) *SessionID {
	for _, mapping := range p.templateMappings {
		if isTemplateMatching(mapping.Pattern, sessionID) {
			return &mapping.TemplateID
		}
	}
	return nil
}

type dynamicAcceptorSessionProvider struct {
	settings            *Settings
	messageStoreFactory MessageStoreFactory
	logFactory          LogFactory
	sessionFactory      *sessionFactory
	application         Application

	templateIDProvider TemplateIDProvider
}

func NewDynamicAcceptorSessionProvider(settings *Settings, messageStoreFactory MessageStoreFactory, logFactory LogFactory,
	application Application, templateIDProvider TemplateIDProvider,
) *dynamicAcceptorSessionProvider {
	return &dynamicAcceptorSessionProvider{
		settings:            settings,
		messageStoreFactory: messageStoreFactory,
		logFactory:          logFactory,
		sessionFactory:      &sessionFactory{},
		application:         application,
		templateIDProvider:  templateIDProvider,
	}
}

func (p *dynamicAcceptorSessionProvider) GetSession(sessionID SessionID) (*session, error) {
	s, ok := lookupSession(sessionID)
	if !ok && p.templateIDProvider != nil {
		templateID := p.templateIDProvider.GetTemplateID(sessionID)
		if templateID == nil {
			return nil, errUnknownSession
		}
		dynamicSessionSettings := p.settings.globalSettings.clone()
		templateSettings, ok := p.settings.sessionSettings[*templateID]
		if !ok {
			return nil, errUnknownSession
		}
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
	if s == nil {
		return nil, errUnknownSession
	}
	return s, nil
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
