package main

import (
	"flag"
	"fmt"
	"os"

	"goftp.io/ftpd/web"

	ldbauth "gitea.com/goftp/leveldb-auth"
	ldbperm "gitea.com/goftp/leveldb-perm"
	qiniudriver "gitea.com/goftp/qiniu-driver"
	"github.com/lunny/log"
	_ "github.com/shurcooL/vfsgen"
	"github.com/syndtr/goleveldb/leveldb"
	"goftp.io/server"
)

var (
	version = "v0.2.1027"
	cfgPath string
)

func main() {
	flag.StringVar(&cfgPath, "config", "",
		"config file path, default is config.ini")
	flag.Parse()

	if err := initConfig(); err != nil {
		fmt.Println(err)
		return
	}

	db, err := leveldb.OpenFile("./authperm.db", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	var auth = &ldbauth.LDBAuth{
		DB: db,
	}
	var perm server.Perm
	if permType == "leveldb" {
		perm = ldbperm.NewLDBPerm(db, "root", "root", os.ModePerm)
	} else {
		perm = server.NewSimplePerm("root", "root")
	}

	var factory server.DriverFactory
	if driverType == "file" {
		_, err = os.Lstat(rootPath)
		if os.IsNotExist(err) {
			os.MkdirAll(rootPath, os.ModePerm)
		} else if err != nil {
			fmt.Println(err)
			return
		}
		factory = &server.FileDriverFactory{
			RootPath: rootPath,
			Perm:     perm,
		}
	} else if driverType == "qiniu" {
		factory = qiniudriver.NewQiniuDriverFactory(qiniu.AccessKey,
			qiniu.SecretKey, qiniu.Bucket)
	} else {
		fmt.Println("no driver type input")
		return
	}

	// start web manage UI
	if webCfg.Enabled {
		web.DB = auth
		web.Perm = perm
		web.Factory = factory

		go web.Web(webCfg.Listen, "static", "templates", admin, pass,
			webCfg.TLS, webCfg.CertFile, webCfg.KeyFile)
	}

	opt := &server.ServerOpts{
		Name:    serv.Name,
		Factory: factory,
		Port:    serv.Port,
		Auth:    auth,
	}

	opt.TLS = serv.TLS
	opt.KeyFile = serv.KeyFile
	opt.CertFile = serv.CertFile
	opt.ExplicitFTPS = opt.TLS

	// start ftp server
	ftpServer := server.NewServer(opt)
	log.Info("FTP Server", version)
	err = ftpServer.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
