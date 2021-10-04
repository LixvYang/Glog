package utils

import (
	"log"
	"gopkg.in/ini.v1"
)

var (
	AppMode string
	HttpPort string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

func init(){
	f, err := ini.Load("config/config.ini")
	if err != nil {
		log.Println("Config  error:",err)
	}

	AppMode = f.Section("server").Key("AppMode").MustString("debug")
	HttpPort = f.Section("server").Key("HttpPort").MustString(":3000")
	
	Db = f.Section("database").Key("Db").MustString("debug")
	DbHost = f.Section("database").Key("DbHost").MustString("localhost")
	DbPort = f.Section("database").Key("DbPort").MustString("3306")
	DbUser = f.Section("database").Key("DbUser").MustString("ginblog")
	DbPassWord = f.Section("database").Key("DbPassWord").MustString("admin123")
	DbName = f.Section("database").Key("DbName").MustString("ginblog")

}