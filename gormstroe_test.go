package quickfix

import (
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	. "github.com/smartystreets/goconvey/convey"
)

func NewGormDB() (*gorm.DB, error) {
	return gorm.Open("postgres", "host=127.0.0.1 user=postgres dbname=initiator sslmode=disable password=123456")
}
func Test_GromStoreCreate(t *testing.T) {
	Convey(`GromStoreCreate`, t, func() {
		db, err := NewGormDB()
		db.LogMode(true)
		So(err, ShouldBeNil)
		So(db, ShouldNotBeNil)
		sessionID := SessionID{BeginString: "FIX.4.2", TargetCompID: "IB", SenderCompID: "LB"}
		appSettings := NewSettings()
		appSettings.sessionSettings = map[SessionID]*SessionSettings{
			sessionID: {},
		}
		Convey(`none session`, func() {
			
		})
		Convey(`have session`, func() {
			f := NewGormStoreFactory(appSettings, db)
			f.Create(sessionID)
		})
	})
}
