//author: Jay.Yuen
package main

import (
	"io/ioutil"
	"os"
	"os/user"
	"strconv"

	jsoniter "github.com/json-iterator/go"
)

var Uid_ int
var Gid_ int
var Config_ config
var UploadDir_ string

var json jsoniter.API

func init() {
	json = jsoniter.ConfigCompatibleWithStandardLibrary
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
	init_file_type()

	for _, v := range Config_.Server.Upload.Allow {
		s := false
		FileType_.Range(func(_, val interface{}) bool {
			vv := val.(string)
			if v == vv {
				s = true
				return false
			}
			return true
		})
		if !s {
			Exit("file type [" + v + "] undefined!")
		}
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
	UploadDir_ = Config_.Server.Work + "/" + Config_.Server.Upload.Dir + "/"
	err = os.Mkdir(UploadDir_, 644)
	err = os.Chown(UploadDir_, Uid_, Gid_)
	if err != nil {
		Exit(err)
	}

	init_db_connect()
}
