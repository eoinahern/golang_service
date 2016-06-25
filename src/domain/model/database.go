package model

import (
	"fmt"
	"log"
	"database/sql"
	  _ "github.com/go-sql-driver/mysql"
)


func newDatabase(name string, pass string, dbname string) *Database {
	db := new(Database)
	
	dbstring := fmt.Sprintf("%s:%s@/%s",name,pass,dbname)
	db.mydbconn, err := sql.Open("mysql", dbstring)
	
	
	if(err != nil){
		log.Fatal("couldnt connect", "error")
		return nil
	}
	
	return db
}


type Database struct {
	 mydbconn *sql.DB
}







