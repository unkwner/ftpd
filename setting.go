package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"

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
	var configFile = cfgPath
	if len(configFile) == 0 {
		dir, err := exePath()
		if err != nil {
			return err
		}

		defaultCfgPath := filepath.Join(filepath.Dir(dir), "config.ini")
		_, err = os.Stat(defaultCfgPath)
		if err != nil && !os.IsNotExist(err) {
			return err
		} else if err == nil {
			configFile = defaultCfgPath
		}
	}

	if len(configFile) == 0 {
		return nil
	}

	cfg, err := goconfig.LoadConfigFile(configFile)
	if err != nil {
		return err
	}

	log.Info("Loaded config file:", configFile)

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

func getEnv(name string, defaultValue string) string {
	v, ok := syscall.Getenv(name)
	if !ok {
		return defaultValue
	}
	return v
}

func getEnvBool(name string, defaultValue bool) bool {
	v, ok := syscall.Getenv(name)
	if !ok {
		return defaultValue
	}
	r, _ := strconv.ParseBool(v)
	return r
}

func getEnvInt(name string, defaultValue int) int {
	v, ok := syscall.Getenv(name)
	if !ok {
		return defaultValue
	}
	r, _ := strconv.Atoi(v)
	return r
}

func readConfigFromEnvs() error {
	permType = getEnv("PERM_TYPE", permType)
	driverType = getEnv("DRIVER_TYPE", driverType)
	rootPath = getEnv("FILE_ROOTPATH", rootPath)

	qiniu.AccessKey = getEnv("QINIU_ACCESSKEY", qiniu.AccessKey)
	qiniu.SecretKey = getEnv("QINIU_SECRETKEY", qiniu.SecretKey)
	qiniu.Bucket = getEnv("QINIU_BUCKET", qiniu.Bucket)

	minio.Endpoint = getEnv("MINIO_ENDPOINT", minio.Endpoint)
	minio.AccessKey = getEnv("MINIO_ACCESSKEY", minio.AccessKey)
	minio.SecretKey = getEnv("MINIO_SECRETKEY", minio.SecretKey)
	minio.Bucket = getEnv("MINIO_BUCKET", minio.Bucket)
	minio.UseSSL = getEnvBool("MINIO_USE_SSL", minio.UseSSL)

	webCfg.Enabled = getEnvBool("WEB_ENABLE", webCfg.Enabled)
	webCfg.Listen = getEnv("WEB_LISTEN", webCfg.Listen)
	admin = getEnv("ADMIN_USER", admin)
	pass = getEnv("ADMIN_PASS", pass)
	webCfg.TLS = getEnvBool("WEB_TLS", webCfg.TLS)
	webCfg.CertFile = getEnv("WEB_CERTFILE", webCfg.CertFile)
	webCfg.KeyFile = getEnv("WEB_KEYFILE", webCfg.KeyFile)

	serv.Name = getEnv("SERVER_NAME", serv.Name)
	serv.Port = getEnvInt("SERVER_PORT", serv.Port)
	serv.TLS = getEnvBool("SERVER_TLS", serv.TLS)
	serv.KeyFile = getEnv("SERVER_KEY_FILE", serv.KeyFile)
	serv.CertFile = getEnv("SERVER_CERT_FILE", serv.CertFile)
	return nil
}
