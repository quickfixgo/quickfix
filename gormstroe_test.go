package quickfix

import (
	"testing"

	"github.com/glebarez/sqlite"
	. "github.com/smartystreets/goconvey/convey"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewGormDB() (*gorm.DB, error) {
	dsn := "host=127.0.0.1 user=postgres dbname=lb_test port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err == nil {
		db.Migrator().DropTable(&GormSessions{})
	}
	return db, err
}

func NewGormSqliteDB() (db *gorm.DB, err error) {
	db, err = gorm.Open(sqlite.Open("test.db?_pragma=busy_timeout(500000)"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	sqldb, _ := db.DB()
	sqldb.SetMaxOpenConns(1)
	return
}

func Test_GromStoreCreate(t *testing.T) {
	Convey(`GromStoreCreate`, t, func() {
		db, err := NewGormDB()
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

func Test_GromStoreSaveMessage(t *testing.T) {
	Convey(`GromStoreSaveMessage`, t, func() {
		db, err := NewGormDB()
		So(err, ShouldBeNil)
		So(db, ShouldNotBeNil)
		sessionID := SessionID{BeginString: "FIX.4.2", TargetCompID: "IB", SenderCompID: "LB"}
		appSettings := NewSettings()
		appSettings.sessionSettings = map[SessionID]*SessionSettings{
			sessionID: {},
		}
		Convey(`duplicate key`, func() {
			f := NewGormStoreFactory(appSettings, db)
			store, err := f.Create(sessionID)
			So(err, ShouldBeNil)
			err = store.SaveMessage(1, []byte(`test msg`))
			err = store.SaveMessage(1, []byte(`test msg`))
			So(err, ShouldBeNil)
		})
	})
}

func Test_SetNextSenderMsgSeqNum(t *testing.T) {
	Convey(`SetNextSenderMsgSeqNum`, t, func() {
		db, err := NewGormSqliteDB()
		So(err, ShouldBeNil)
		So(db, ShouldNotBeNil)
		sessionID := SessionID{BeginString: "FIX.4.2", TargetCompID: "IB", SenderCompID: "LB"}
		sessionID = SessionID{}
		appSettings := NewSettings()
		appSettings.sessionSettings = map[SessionID]*SessionSettings{
			sessionID: {},
		}
		f := NewGormStoreFactory(appSettings, db)
		store, err := f.Create(sessionID)
		So(err, ShouldBeNil)
		err = store.SetNextSenderMsgSeqNum(10)
		So(err, ShouldBeNil)
	})
}
