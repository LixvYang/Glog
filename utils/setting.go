package utils

import (
	"gopkg.in/ini.v1"
	"log"
)

var (
	AppMode  string
	HttpPort string
	JwtKey	 string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

func init() {
	f, err := ini.Load("config/config.ini")
	if err != nil {
		log.Println("Config  error:", err)
	}

	AppMode = f.Section("server").Key("AppMode").MustString("debug")
	HttpPort = f.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = f.Section("server").Key("JwtKey").MustString("89js82js72")

	Db = f.Section("database").Key("Db").MustString("mysql")
	DbHost = f.Section("database").Key("DbHost").MustString("localhost")
	DbPort = f.Section("database").Key("DbPort").MustString("3307")
	DbUser = f.Section("database").Key("DbUser").MustString("root")
	DbPassWord = f.Section("database").Key("DbPassWord").MustString("admin123")
	DbName = f.Section("database").Key("DbName").MustString("ginblog")

}
