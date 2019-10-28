//author: Jay.Yuen
package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"os"
	"os/user"
	"strconv"
)

var Uid_ int
var Gid_ int
var Config_ config

var db *sql.DB

func init() {
	//load config
	if len(os.Args) != 2 {
		Exit("please input config file's path")
	}
	file_info, err := os.Stat(os.Args[1])
	if err != nil {
		Exit(err) //power error
	}
	if file_info.IsDir() {
		Exit("please input config file's path")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		Exit(err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		Exit(err)
	}
	err = json.Unmarshal(data, &Config_)
	if err != nil {
		Exit(err)
	}
	//create upload folder
	usr, err := user.Lookup(Config_.Postgres.Admin)
	if err != nil {
		Exit(err)
	}
	Uid_, err = strconv.Atoi(usr.Uid)
	if err != nil {
		Exit(err)
	}
	Gid_, err = strconv.Atoi(usr.Gid)
	if err != nil {
		Exit(err)
	}
	err = os.Mkdir(Config_.Upload, 644)
	err = os.Chown(Config_.Upload, Uid_, Gid_)
	if err != nil {
		Exit(err)
	}
	//使用数据库连接池
	ConnectString := "user=" + Config_.Postgres.Username + " password=" + Config_.Postgres.Password + " host=" + Config_.Postgres.Host + " port=" + Config_.Postgres.Port + " dbname=" + Config_.Postgres.Database + " sslmode=disable"
	db, _ = sql.Open("postgres", ConnectString)
	db.SetConnMaxLifetime(3600000)
	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(100)
}
