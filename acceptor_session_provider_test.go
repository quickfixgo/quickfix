package quickfix

import (
	"reflect"
	"testing"

	"github.com/quickfixgo/quickfix/config"
	"github.com/stretchr/testify/suite"
)

type DynamicAcceptorSessionProviderTestSuite struct {
	suite.Suite

	provider *DynamicAcceptorSessionProvider

	settings            *Settings
	messageStoreFactory MessageStoreFactory
	logFactory          LogFactory
	app                 Application
	sessionFactory      *sessionFactory
	TemplateMapping     []*TemplateMapping
}

func (suite *DynamicAcceptorSessionProviderTestSuite) SetupTest() {
	suite.settings = NewSettings()
	suite.messageStoreFactory = NewMemoryStoreFactory()
	suite.logFactory = nullLogFactory{}
	suite.app = &noopApp{}
	suite.sessionFactory = &sessionFactory{}
	suite.TemplateMapping = make([]*TemplateMapping, 0)

	templateId1 := SessionID{BeginString: "FIX.4.2", SenderCompID: "ANY", TargetCompID: "ANY"}
	suite.TemplateMapping = append(
		suite.TemplateMapping,
		&TemplateMapping{Pattern: SessionID{BeginString: WildcardPattern, SenderCompID: "S1", TargetCompID: WildcardPattern}, TemplateID: templateId1},
	)
	suite.setUpSettings(templateId1, "ResetOnLogout", "Y")

	templateId2 := SessionID{BeginString: "FIX.4.4", SenderCompID: "S1", TargetCompID: "ANY"}
	suite.TemplateMapping = append(
		suite.TemplateMapping,
		&TemplateMapping{Pattern: SessionID{BeginString: "FIX.4.4", SenderCompID: WildcardPattern, TargetCompID: WildcardPattern}, TemplateID: templateId2},
	)
	suite.setUpSettings(templateId2, "RefreshOnLogon", "Y")

	templateId3 := SessionID{BeginString: "FIX.4.4", SenderCompID: "ANY", TargetCompID: "ANY"}
	suite.TemplateMapping = append(
		suite.TemplateMapping,
		&TemplateMapping{Pattern: SessionID{BeginString: "FIX.4.2", SenderCompID: WildcardPattern, SenderSubID: WildcardPattern, SenderLocationID: WildcardPattern,
			TargetCompID: WildcardPattern, TargetSubID: WildcardPattern, TargetLocationID: WildcardPattern, Qualifier: WildcardPattern,
		}, TemplateID: templateId3},
	)
	suite.setUpSettings(templateId3, "ResetOnDisconnect", "Y")

	suite.provider = NewDynamicAcceptorSessionProvider(suite.settings, suite.messageStoreFactory,
		suite.logFactory, suite.app, suite.TemplateMapping)
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
		session, err := suite.provider.GetSession(test.input)
		suite.NoError(err)
		suite.NotNil(session)
		sessionID := session.sessionID
		suite.Equal(test.expected.sessionID, sessionID, test.name+": created sessionID not expected")
		suite.Equal(test.expected.resetOnLogout, session.ResetOnLogout, test.name+":ResetOnLogout not expected")
		suite.Equal(test.expected.refreshOnLogon, session.RefreshOnLogon, test.name+":RefreshOnLogon not expected")
		suite.Equal(test.expected.resetOnDisconnect, session.ResetOnDisconnect, test.name+":ResetOnDisconnect not expected")
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
		_, err := suite.provider.GetSession(test.input)
		suite.Error(err, test.name+": expected error for template not found")
	}
}

func TestDynamicAcceptorSessionProviderTestSuite(t *testing.T) {
	suite.Run(t, new(DynamicAcceptorSessionProviderTestSuite))
}

func TestStaticSessionProvider_GetSession(t *testing.T) {
	sessions := make(map[SessionID]*session)
	sessionID1 := SessionID{BeginString: "FIX.4.2", SenderCompID: "SENDER", TargetCompID: "TARGET"}
	session1 := &session{sessionID: sessionID1}
	sessions[sessionID1] = session1

	type args struct {
		sessionID SessionID
	}
	tests := []struct {
		name    string
		args    args
		want    *session
		wantErr bool
	}{
		{
			name: "session found",
			args: args{
				sessionID: sessionID1,
			},
			want:    session1,
			wantErr: false,
		},
		{
			name: "session not found",
			args: args{
				sessionID: SessionID{
					BeginString: "FIX.4.2", SenderCompID: "X", TargetCompID: "Y",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &StaticAcceptorSessionProvider{
				sessions: sessions,
			}
			got, err := p.GetSession(tt.args.sessionID)
			if (err != nil) != tt.wantErr {
				t.Errorf("StaticSessionProvider.GetSession() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StaticSessionProvider.GetSession() = %v, want %v", got, tt.want)
			}
		})
	}
}

var _ Application = &noopApp{}

type noopApp struct {
}

func (n *noopApp) FromAdmin(message *Message, sessionID SessionID) MessageRejectError {
	return nil
}

func (n *noopApp) FromApp(message *Message, sessionID SessionID) MessageRejectError {
	return nil
}

func (n *noopApp) OnCreate(sessionID SessionID) {
}

func (n *noopApp) OnLogon(sessionID SessionID) {
}

func (n *noopApp) OnLogout(sessionID SessionID) {
}

func (n *noopApp) ToAdmin(message *Message, sessionID SessionID) {
}

func (n *noopApp) ToApp(message *Message, sessionID SessionID) error {
	return nil
}
