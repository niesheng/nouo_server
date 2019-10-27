//author: Jay.Yuen
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"strconv"
)

var Uid_ int
var Gid_ int
var Config_ config

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
	fmt.Println(Config_)
}
