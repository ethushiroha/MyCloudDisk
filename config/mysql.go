package config

import (
	"MyCloudDisk/global"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func InitMySQL() *sqlx.DB {
	vp := global.VP
	//DBType := vp.GetString("Database.DBType")
	UserName := vp.GetString("database.UserName")
	Password := vp.GetString("database.Password")
	Host := vp.GetString("database.Host")
	DBName := vp.GetString("database.DBName")
	Charset := vp.GetString("database.Charset")
	ParseTime := vp.GetBool("database.ParseTime")
	MaxIdleConns := vp.GetInt("database.MaxIdleConns")
	MaxOpenConns := vp.GetInt("database.MaxOpenConns")
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		UserName,
		Password,
		Host,
		DBName,
		Charset,
		ParseTime,
	)
	maxOpenConns := MaxOpenConns
	maxIdleConns := MaxIdleConns

	MySqlDB, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		panic("初始化失败：连接数据库失败")
	}
	MySqlDB.SetMaxOpenConns(maxOpenConns)
	MySqlDB.SetMaxIdleConns(maxIdleConns)

	if err = MySqlDB.Ping(); err != nil {
		panic("初始化失败：测试数据库连接失败")
		//log.Fatalln("测试数据库连接失败:", err)
	}

	return MySqlDB
}
