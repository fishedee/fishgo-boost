package sqlf

import (
	"testing"
	"time"

	. "github.com/fishedee/fishgo-boost/app/log"
	. "github.com/fishedee/fishgo-boost/assert"
)

func initSqliteDatabase2() SqlfDB {
	log, err := NewLog(LogConfig{
		Driver: "console",
	})
	if err != nil {
		panic(err)
	}
	db, err := NewSqlfDB(log, nil, SqlfDBConfig{
		Driver:     "sqlite_localtime",
		SourceName: ":memory:?_inttotime=1",
		Debug:      true,
	})
	if err != nil {
		panic(err)
	}
	db.MustExec(`
	create table t_user(
		userId integer primary key autoincrement,
		name char(32) not null,
		createTime timestamp not null default 0,
		modifyTime timestamp not null default 0
	);
	`)
	return db
}

func initMySqlDatabase2() SqlfDB {
	log, err := NewLog(LogConfig{
		Driver: "console",
	})
	if err != nil {
		panic(err)
	}
	db, err := NewSqlfDB(log, nil, SqlfDBConfig{
		Driver:     "mysql",
		SourceName: "root:123@tcp(localhost:3306)/test?parseTime=true&loc=Local",
		Debug:      true,
	})
	if err != nil {
		panic(err)
	}
	db.MustExec(`
	drop table if exists t_user;
	`)
	db.MustExec(`
	create table t_user(
		userId int not null auto_increment,
		name char(32) not null,
		createTime datetime not null default '1970-01-01 08:00:00',
		modifyTime datetime not null default '1970-01-01 08:00:00',
		primary key(userId)
	)engine=innodb default charset=utf8mb4;`)
	return db
}

type User2 struct {
	UserId     int `sqlf:"autoincr"`
	Name       string
	CreateTime time.Time
	ModifyTime time.Time
}

func parseTime(str string) time.Time {
	today, err := time.ParseInLocation("2006-01-02 15:04:05", str, time.Local)
	if err != nil {
		panic(err)
	}
	return today
}

func testInsertTime(t *testing.T, initDatabase func() SqlfDB) {
	db := initDatabase()

	//用sql插入数据
	db.MustExec(`
		insert into t_user(userId, name , createTime ,modifyTime )values
		(10001,'fish','2020-07-01 00:00:00','2020-07-02 18:00:00')
	`)

	users := []User2{}
	db.MustQuery(&users, "select * from t_user where userId = 10001")

	AssertEqual(t, users, []User2{
		{UserId: 10001, Name: "fish", CreateTime: parseTime("2020-07-01 00:00:00"), ModifyTime: parseTime("2020-07-02 18:00:00")},
	})

	//用api插入数据
	user2 := User2{
		UserId:     10002,
		Name:       "Cat",
		CreateTime: parseTime("2020-08-01 00:00:00"),
		ModifyTime: time.Unix(1, 0),
	}
	db.MustExec("insert into t_user(?.insertColumn) values ?.insertValue", user2, user2)

	db.MustQuery(&users, "select * from t_user where userId = 10002")

	AssertEqual(t, users, []User2{
		user2,
	})

	//测试单个type类型
	db.MustExec("insert into t_user(userId,name,createTime,modifyTime) values(?,?,?,?)", 10003, "dog", parseTime("2020-08-02 00:00:00"), time.Unix(1, 0))

	db.MustQuery(&users, "select * from t_user where userId = 10003", users)

	AssertEqual(t, users, []User2{
		User2{UserId: 10003, Name: "dog", CreateTime: parseTime("2020-08-02 00:00:00"), ModifyTime: time.Unix(1, 0)},
	})

}

func testAll2(t *testing.T, initDatabase func() SqlfDB) {
	testInsertTime(t, initDatabase)
}

func TestAll2(t *testing.T) {
	testAll2(t, initSqliteDatabase2)
	testAll2(t, initMySqlDatabase2)
}
