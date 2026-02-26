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
		Driver:     "sqlite",
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

func initPostgresqlDatabase2() SqlfDB {
	log, err := NewLog(LogConfig{
		Driver: "console",
	})
	if err != nil {
		panic(err)
	}
	db, err := NewSqlfDB(log, nil, SqlfDBConfig{
		Driver:     "postgres_fix",
		SourceName: "user=postgres password=postgres host=127.0.0.1 port=5432 dbname=test sslmode=disable timezone=Asia/Shanghai",
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
		userId integer not null GENERATED ALWAYS AS IDENTITY (start with 10001),
		name varchar(32) not null,
		createTime timestamptz not null default '1970-01-01 08:00:00',
		modifyTime timestamptz not null default '1970-01-01 08:00:00',
		primary key(userId)
	);`)
	/* 等效上述代码
	db.MustExec(`
	CREATE SEQUENCE t_user_userid_seq;
	create table t_user(
		userId integer not null DEFAULT nextval('t_user_userid_seq'),
		name varchar(32) not null,
		createTime timestamp not null default '1970-01-01 08:00:00',
		modifyTime timestamp not null default '1970-01-01 08:00:00'
	);`)
	*/
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

	/*
		sqlite的时间字段有较多的问题，它的特点是，时间字段只是字符串类型，没有时区概念，也不会自动进行归一化，所以

		* 最佳实践就是时间字段只能存UTC时区的规范化格式的字符串，形如2020-07-01 00:00:00，注意不含时区部分
		* 如果存LOCAL时区的字符串，就会导致0时间无法读出来，而且绝大部分的sqlite驱动都遵循这个标准，违反这个标准就容易导致换个驱动就读不出来
		* 如果时间带时区部分，那么搜索就难以匹配，虽然'2020-07-01 08:00:00+08:00'和'2020-07-01 00:00:00+00:00'是同一个时刻，但是sqlite没有对所有时间规范化，它们在字符串形式上不一致，就会导致失败。所以时间不能含时区部分。
		* 使用SQL来插入数据的时候，我们要使用datetime函数来做时间数据的归一化。使用SQLF API来插入数据的时候，SQLF已经默认做了归一化
	*/
	//用sql插入数据
	if db.Engine() == SQLITE {
		db.MustExec(`
		insert into t_user(userId, name , createTime ,modifyTime )values
		(10001,'fish',datetime('2020-07-01 00:00:00+08:00'),datetime('2020-07-02 18:00:00+08:00'))
	`)
	} else if db.Engine() == MYSQL {
		db.MustExec(`
		insert into t_user(userId, name , createTime ,modifyTime )values
		(10001,'fish','2020-07-01 00:00:00','2020-07-02 18:00:00')
	`)
	} else {
		db.MustExec(`
		insert into t_user( name , createTime ,modifyTime )values
		('fish','2020-07-01 00:00:00','2020-07-02 18:00:00')
	`)
	}

	users := []User2{}
	db.MustQuery(&users, "select * from t_user where userId = 10001")

	user1 := User2{
		UserId:     10001,
		Name:       "fish",
		CreateTime: parseTime("2020-07-01 00:00:00"),
		ModifyTime: parseTime("2020-07-02 18:00:00"),
	}

	AssertEqual(t, users, []User2{
		user1,
	})

	db.MustQuery(&users, "select * from t_user where createTime = ?", parseTime("2020-07-01 00:00:00"))

	AssertEqual(t, users, []User2{
		user1,
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

	if db.Engine() == SQLITE {
		db.MustQuery(&users, "select * from t_user where createTime = datetime('2020-08-01 00:00:00+08:00')")
	} else {
		db.MustQuery(&users, "select * from t_user where createTime = '2020-08-01 00:00:00'")
	}

	AssertEqual(t, users, []User2{
		user2,
	})

	//测试单个type类型
	insertResult := db.MustExec("insert into t_user(name,createTime,modifyTime) values(?,?,?)", "dog", parseTime("2020-08-02 00:00:00"), time.Unix(1, 0))
	AssertEqual(t, insertResult.MustLastInsertId(), int64(10003))

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
	testAll2(t, initPostgresqlDatabase2)
}
