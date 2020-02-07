package main

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/lunny/log"
	"github.com/unknwon/goconfig"
)

var (
	permType   = "simple"
	driverType = "file"
	rootPath   string

	qiniu = struct {
		Bucket    string
		AccessKey string
		SecretKey string
	}{}

	minio = struct {
		Endpoint  string
		Bucket    string
		AccessKey string
		SecretKey string
		UseSSL    bool
	}{}

	webCfg = struct {
		Enabled  bool
		Listen   string
		TLS      bool
		CertFile string
		KeyFile  string
	}{
		Enabled: true,
		Listen:  ":8181",
		TLS:     false,
	}

	admin = "admin"
	pass  = "123456"

	serv = struct {
		Name     string
		Port     int
		TLS      bool
		KeyFile  string
		CertFile string
	}{
		Name:     "Go Ftp Server",
		Port:     2121,
		TLS:      false,
		KeyFile:  "key.pem",
		CertFile: "cert.pem",
	}
)

// exePath returns the executable path.
func exePath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	return filepath.Abs(file)
}

func initConfig() error {
	var cfgFiles []string

	if len(cfgPath) > 0 {
		f, _ := filepath.Abs(cfgPath)
		cfgFiles = append(cfgFiles, f)
	} else {
		dir, err := exePath()
		if err != nil {
			return err
		}

		cfgPath = filepath.Join(filepath.Dir(dir), "config.ini")
		_, err = os.Stat(cfgPath)
		if err != nil && !os.IsNotExist(err) {
			return err
		} else if err == nil {
			cfgFiles = append(cfgFiles, cfgPath)
		}

		customPath := filepath.Join(filepath.Dir(dir), "custom.ini")
		_, err = os.Stat(customPath)
		if err != nil && !os.IsNotExist(err) {
			return err
		} else if err == nil {
			cfgFiles = append(cfgFiles, customPath)
		}
	}

	if len(cfgFiles) == 0 {
		return nil
	}

	cfg, err := goconfig.LoadConfigFile(cfgFiles[0], cfgFiles[1:]...)
	if err != nil {
		return err
	}

	log.Info("Loaded config files:", cfgFiles)

	permType = cfg.MustValue("perm", "type", permType)
	driverType = cfg.MustValue("driver", "type", driverType)
	rootPath = cfg.MustValue("file", "rootpath", rootPath)

	qiniu.AccessKey, _ = cfg.GetValue("qiniu", "accessKey")
	qiniu.SecretKey, _ = cfg.GetValue("qiniu", "secretKey")
	qiniu.Bucket, _ = cfg.GetValue("qiniu", "bucket")

	minio.Endpoint, _ = cfg.GetValue("minio", "endpoint")
	minio.AccessKey, _ = cfg.GetValue("minio", "accessKey")
	minio.SecretKey, _ = cfg.GetValue("minio", "secretKey")
	minio.Bucket, _ = cfg.GetValue("minio", "bucket")
	minio.UseSSL = cfg.MustBool("minio", "use_ssl", false)

	webCfg.Enabled = cfg.MustBool("web", "enable", webCfg.Enabled)
	webCfg.Listen = cfg.MustValue("web", "listen", webCfg.Listen)
	admin = cfg.MustValue("admin", "user", admin)
	pass = cfg.MustValue("admin", "pass", pass)
	webCfg.TLS = cfg.MustBool("web", "tls", webCfg.TLS)
	webCfg.CertFile = cfg.MustValue("web", "certFile", webCfg.CertFile)
	webCfg.KeyFile = cfg.MustValue("web", "keyFile", webCfg.KeyFile)

	serv.Name = cfg.MustValue("server", "name", serv.Name)
	serv.Port = cfg.MustInt("server", "port", serv.Port)
	serv.TLS = cfg.MustBool("server", "tls", false)
	serv.KeyFile = cfg.MustValue("server", "key_file", "")
	serv.CertFile = cfg.MustValue("server", "cert_file", "")

	return nil
}
