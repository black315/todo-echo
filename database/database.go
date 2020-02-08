package database

import (
   "github.com/jinzhu/gorm"
   _ "github.com/go-sql-driver/mysql"
)

func GetDBConnection() *gorm.DB {
   db, err := gorm.Open(GetDBConfig())
   if err != nil {
      panic(err)
   }
   return db
}

func GetDBConfig() (string, string) {
   DBMS := "mysql"
   USER := "root"
   PASS := ""
   PROTOCOL := ""
   DBNAME := "todo_echo_gorm"
   OPTION := "charset=utf8&parseTime=True&loc=Local"

   CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?" + OPTION

   return DBMS, CONNECT
 }
