package core

import (
	"github.com/gogather/com"
	"path/filepath"
	"runtime"
)

var DB *Sqlite
var C *Config

func Judger() {
	C = &Config{}
	C.NewConfig("conf/config.ini")

	DB = &Sqlite{}
	DB.NewSqlite()

	createBuildDir()

	TcpStart()
}

func createBuildDir() error {
	var err error
	err = nil

	buildPath := filepath.Join(C.Get(runtime.GOOS, "buildpath"))
	if !com.PathExist(buildPath) {
		err = com.Mkdir(buildPath)
	}

	return err
}
