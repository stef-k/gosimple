package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//_ "github.com/go-sql-driver/mysql" // uncomment appropriately
	//_ "github.com/lib/pq" // uncomment appropriately
	_ "github.com/mattn/go-sqlite3"
	"fmt"
)

// sync database (makes migrations)
func generateTables()  {
	name := "default"
	force := true
	verbose := false
	orm.RunSyncdb(name, force, verbose)
}

// to use your app's models with the ORM, you must first
// register them. Put all your app's models here
func registerModels() {
	orm.RegisterModel(new(User))
}

// check and set which database to use
func dbInit() {
	database := beego.AppConfig.DefaultString("database", "sqlite3")
	dbname := beego.AppConfig.DefaultString("dbname", "database")
	if database == "sqlite3" {
		dbname += ".db"
	}
	if database == "sqlite3" {
		orm.RegisterDataBase("default", "sqlite3", "./" + dbname)
	} else {
		dbuser := beego.AppConfig.String("dbuser")
		dbpass := beego.AppConfig.String("dbpass")
		dbaddress := beego.AppConfig.String("dbaddress")
		dbport := beego.AppConfig.String("dbport")
		dbMaxIdleCons, _ := beego.AppConfig.Int("dbMaxIdleCons")
		if database == "mysql" {
			orm.RegisterDataBase("default", "mysql",
				fmt.Sprintf("%v:%v@tcp/%v", dbuser, dbpass, dbname),
				dbMaxIdleCons)
		} else if database == "postgres" {
			orm.RegisterDataBase("default", "postgres",
				fmt.Sprintf("postgresql://%v:%v@%v:%v", dbuser, dbpass, dbaddress, dbport),
				dbMaxIdleCons)
		}
	}
}

// initialize ORM and app models
func init() {
	dbInit()
	registerModels()
	generateTables()
}
