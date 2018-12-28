package main

import (
	"fmt"
	"os"

	"github.com/pengkebao/ftpserver/auth"
	"github.com/pengkebao/ftpserver/conf"
	"github.com/pengkebao/ftpserver/driver"

	"github.com/goftp/server"
	"github.com/lunny/log"
)

func main() {
	_, err := os.Lstat(conf.RootPath)
	if os.IsNotExist(err) {
		os.MkdirAll(conf.RootPath, os.ModePerm)
	} else if err != nil {
		fmt.Println(err)
		return
	}
	//设置权限。
	factory := &filedriver.FileDriverFactory{
		RootPath: "file",
		Perm:     server.NewSimplePerm(conf.PermOwner, conf.PermGroup),
	}

	auth := new(auth.Auth)

	opt := &server.ServerOpts{
		Name:         "go",
		Factory:      factory,
		Port:         conf.Port,
		Auth:         auth,
		PassivePorts: conf.PasvPort,
	}
	// start ftp server
	ftpServer := server.NewServer(opt)
	log.Info("FTP Server", "1.0")
	err = ftpServer.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
