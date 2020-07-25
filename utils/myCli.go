package utils

import (
	"fmt"
	"monkeyServer/logUtils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func SqlxCli() *sqlx.DB {

	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", Config.DB.DbUser,Config.DB.DbPass,Config.DB.DbHost,Config.DB.DbName)
	if dbClient, err := sqlx.Connect("mysql", dns); err != nil{
		logUtils.Errorf("mysql connect error=%v",err)
	}else {
		dbClient.SetMaxOpenConns(200)
		dbClient.SetMaxIdleConns(20)
		return dbClient
	}
	return nil
}
