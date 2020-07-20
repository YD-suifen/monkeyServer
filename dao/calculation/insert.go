package calculation

import (
	"fmt"
	"monkeyServer/logUtils"
	"monkeyServer/utils"
)

func CreateTable(tableName string) string {

	logUtils.Info("CreateTable start name=%v", tableName)
	db := utils.SqlxCli()
	defer db.Close()

	Table := tableName+ "_" + utils.Todaydate()
	tableStruck := "(id int(100) AUTO_INCREMENT,hostName varchar(100),keyName varchar(50),MaxValue float DEFAULT '0',MinValue float DEFAULT '0',timeUnix int(100),PRIMARY KEY (`id`)) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;"

	sql := "create table " + Table + tableStruck

	logUtils.Debugf("CreateMjPVTable sql=%v", sql)
	if _, err := db.Exec(sql); err != nil {
		logUtils.Errorf("CreateMjPVTable is error=%v", err)
		return ""
	}
	return Table
}

func Insert(tableName, keyName string,max,min float64,timeUnix int64) bool {
	db := utils.SqlxCli()
	defer db.Close()
	
	sql := fmt.Sprintf("insert into %v (hostName,keyName,MaxValue,MinValue,timeUnix) value ('%v',%v,%v,%v,)",tableName,keyName,max,min,timeUnix)
	logUtils.Debugf("Insert sql=%v", sql)
	if _, err := db.Exec(sql); err != nil {
		logUtils.Error("Insert err=%v", err)
		return false
	}
	return true
}

