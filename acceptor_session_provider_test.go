package quickfix

import (
	"strings"
	"testing"

	"github.com/quickfixgo/quickfix/config"
	"github.com/stretchr/testify/suite"
)

var _ Application = &noopApp{}

type noopApp struct {
}

func (n *noopApp) FromAdmin(_ *Message, _ SessionID) MessageRejectError {
	return nil
}

func (n *noopApp) FromApp(_ *Message, _ SessionID) MessageRejectError {
	return nil
}

func (n *noopApp) OnCreate(_ SessionID) {
}

func (n *noopApp) OnLogon(_ SessionID) {
}

func (n *noopApp) OnLogout(_ SessionID) {
}

func (n *noopApp) ToAdmin(_ *Message, _ SessionID) {
}

func (n *noopApp) ToApp(_ *Message, _ SessionID) error {
	return nil
}

type DynamicAcceptorSessionProviderTestSuite struct {
	suite.Suite

	dynamicAcceptorSessionProvider *dynamicAcceptorSessionProvider

	settings            *Settings
	messageStoreFactory MessageStoreFactory
	logFactory          LogFactory
	app                 Application
	sessionFactory      *sessionFactory
}

func (suite *DynamicAcceptorSessionProviderTestSuite) TestNewDefaultTemplateIDProvider() {
	cfg := `
[default]
ConnectionType=acceptor
SocketAcceptPort=9878
BeginString=FIX.4.2
TimeZone=America/New_York
StartTime=00:00:01
EndTime=23:59:59
HeartBtInt=30

[session]
AcceptorTemplate=Y
SenderCompID=test1
TargetCompID=*
ResetOnLogon=Y

[session]
AcceptorTemplate=Y
SenderCompID=test2
TargetCompID=*
ResetOnLogon=Y
	`
	stringReader := strings.NewReader(cfg)
	settings, err := ParseSettings(stringReader)
	if err != nil {
		suite.FailNow("parse setting failed", err)
	}

	provider, err := NewDefaultTemplateIDProvider(settings)
	if err != nil {
		suite.FailNow("create TemplateIDProvider failed", err)
	}

	s1 := SessionID{BeginString: "FIX.4.2", SenderCompID: "test1", TargetCompID: "cli"}
	templateID1 := provider.GetTemplateID(s1)
	suite.Require().NotNil(templateID1, "expected template matched")
	suite.Require().Equal(
		*templateID1,
		SessionID{BeginString: "FIX.4.2", SenderCompID: "test1", TargetCompID: "*"},
		"unexpected templateID",
	)

	s2 := SessionID{BeginString: "FIX.4.3", SenderCompID: "test1", TargetCompID: "cli"}
	templateID2 := provider.GetTemplateID(s2)
	suite.Require().Nilf(templateID2, "expected template not matched for %v", s2.String())

	s3 := SessionID{BeginString: "FIX.4.2", SenderCompID: "X", TargetCompID: "cli"}
	templateID3 := provider.GetTemplateID(s3)
	suite.Require().Nilf(templateID3, "expected template not matched for %v", s3.String())
}

func (suite *DynamicAcceptorSessionProviderTestSuite) SetupTest() {
	suite.settings = NewSettings()
	suite.messageStoreFactory = NewMemoryStoreFactory()
	suite.logFactory = nullLogFactory{}
	suite.app = &noopApp{}
	suite.sessionFactory = &sessionFactory{}
	templateMappings := make([]*TemplateMapping, 0)

	templateID1 := SessionID{BeginString: "FIX.4.2", SenderCompID: "ANY", TargetCompID: "ANY"}
	templateMappings = append(
		templateMappings,
		&TemplateMapping{Pattern: SessionID{BeginString: WildcardPattern, SenderCompID: "S1", TargetCompID: WildcardPattern}, TemplateID: templateID1},
	)
	suite.setUpSettings(templateID1, "ResetOnLogout", "Y")

	templateId2 := SessionID{BeginString: "FIX.4.4", SenderCompID: "S1", TargetCompID: "ANY"}
	templateMappings = append(
		templateMappings,
		&TemplateMapping{Pattern: SessionID{BeginString: "FIX.4.4", SenderCompID: WildcardPattern, TargetCompID: WildcardPattern}, TemplateID: templateId2},
	)
	suite.setUpSettings(templateId2, "RefreshOnLogon", "Y")

	templateId3 := SessionID{BeginString: "FIX.4.4", SenderCompID: "ANY", TargetCompID: "ANY"}
	templateMappings = append(
		templateMappings,
		&TemplateMapping{Pattern: SessionID{BeginString: "FIX.4.2", SenderCompID: WildcardPattern, SenderSubID: WildcardPattern, SenderLocationID: WildcardPattern,
			TargetCompID: WildcardPattern, TargetSubID: WildcardPattern, TargetLocationID: WildcardPattern, Qualifier: WildcardPattern,
		}, TemplateID: templateId3},
	)
	suite.setUpSettings(templateId3, "ResetOnDisconnect", "Y")

	templateIDProvider := &DefaultTemplateIDProvider{
		templateMappings: templateMappings,
	}

	suite.dynamicAcceptorSessionProvider = NewDynamicAcceptorSessionProvider(suite.settings, suite.messageStoreFactory,
		suite.logFactory, suite.app, templateIDProvider)
}

func (suite *DynamicAcceptorSessionProviderTestSuite) setUpSettings(TemplateID SessionID, key, value string) {
	sessionSettings := NewSessionSettings()
	sessionSettings.Set(config.BeginString, TemplateID.BeginString)
	sessionSettings.Set(config.SenderCompID, TemplateID.SenderCompID)
	sessionSettings.Set(config.SenderSubID, TemplateID.SenderSubID)
	sessionSettings.Set(config.SenderLocationID, TemplateID.SenderLocationID)
	sessionSettings.Set(config.TargetCompID, TemplateID.TargetCompID)
	sessionSettings.Set(config.TargetSubID, TemplateID.TargetSubID)
	sessionSettings.Set(config.TargetLocationID, TemplateID.TargetLocationID)
	sessionSettings.Set(config.SessionQualifier, TemplateID.Qualifier)

	sessionSettings.Set("StartTime", "00:00:00")
	sessionSettings.Set("EndTime", "00:00:00")
	sessionSettings.Set(key, value)
	suite.settings.AddSession(sessionSettings)
}

func (suite *DynamicAcceptorSessionProviderTestSuite) TestSessionCreation() {
	type expected struct {
		sessionID         SessionID
		resetOnLogout     bool
		refreshOnLogon    bool
		resetOnDisconnect bool
	}
	var tests = []struct {
		name     string
		input    SessionID
		expected expected
	}{
		{
			name: "session created - matched",
			input: SessionID{
				BeginString: "FIX.4.2", SenderCompID: "SENDER", SenderSubID: "SENDERSUB", SenderLocationID: "SENDERLOC",
				TargetCompID: "TARGET", TargetSubID: "TARGETSUB", TargetLocationID: "TARGETLOC", Qualifier: "",
			},
			expected: expected{
				sessionID: SessionID{
					BeginString: "FIX.4.2", SenderCompID: "SENDER", SenderSubID: "SENDERSUB", SenderLocationID: "SENDERLOC",
					TargetCompID: "TARGET", TargetSubID: "TARGETSUB", TargetLocationID: "TARGETLOC", Qualifier: "",
				},
				resetOnLogout:     false,
				refreshOnLogon:    false,
				resetOnDisconnect: true,
			},
		},
		{
			name: "create session - matching the first",
			input: SessionID{
				BeginString: "FIX.4.4", SenderCompID: "S1", TargetCompID: "T",
			},
			expected: expected{
				sessionID: SessionID{
					BeginString: "FIX.4.4", SenderCompID: "S1", TargetCompID: "T",
				},
				resetOnLogout:     true,
				refreshOnLogon:    false,
				resetOnDisconnect: false,
			},
		},
		{
			name: "create session - matching the second",
			input: SessionID{
				BeginString: "FIX.4.4", SenderCompID: "X", TargetCompID: "Y",
			},
			expected: expected{
				sessionID: SessionID{
					BeginString: "FIX.4.4", SenderCompID: "X", TargetCompID: "Y",
				},
				resetOnLogout:     false,
				refreshOnLogon:    true,
				resetOnDisconnect: false,
			},
		},
	}

	for _, test := range tests {
		session, err := suite.dynamicAcceptorSessionProvider.GetSession(test.input)
		suite.NoError(err)
		suite.NotNil(session)
		sessionID := session.sessionID
		suite.Require().Equal(test.expected.sessionID, sessionID, test.name+": created sessionID not expected")
		suite.Require().Equal(test.expected.resetOnLogout, session.ResetOnLogout, test.name+":ResetOnLogout not expected")
		suite.Require().Equal(test.expected.refreshOnLogon, session.RefreshOnLogon, test.name+":RefreshOnLogon not expected")
		suite.Require().Equal(test.expected.resetOnDisconnect, session.ResetOnDisconnect, test.name+":ResetOnDisconnect not expected")
		UnregisterSession(sessionID)
	}
}

func (suite *DynamicAcceptorSessionProviderTestSuite) TestTemplateNotFound() {
	var tests = []struct {
		name  string
		input SessionID
	}{
		{
			name: "template not found",
			input: SessionID{
				BeginString: "FIX.4.3", SenderCompID: "S", TargetCompID: "T",
			},
		},
	}

	for _, test := range tests {
		_, err := suite.dynamicAcceptorSessionProvider.GetSession(test.input)
		suite.Error(err, test.name+": expected error for template not found")
	}
}

func TestDynamicAcceptorSessionProviderTestSuite(t *testing.T) {
	suite.Run(t, new(DynamicAcceptorSessionProviderTestSuite))
}
