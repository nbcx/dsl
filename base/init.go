package base

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"io"
	"os"
)

var (
	Gin     = ginRouter()
	LocalIp = getServerIp()
)

func Init() {
	var config = pflag.StringP("config", "c", "app.ini", "input your config path")
	pflag.StringP("log.path", "l", "debug.log", "")
	pflag.IntP("client.port", "p", 8080, "")
	pflag.IntP("server.port", "s", 8089, "")
	pflag.IntP("distributed.port", "r", 9002, "")
	viper.BindPFlags(pflag.CommandLine)
	pflag.Parse()

	viper.SetConfigFile(*config)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	initLog()
}

// 初始化日志
func initLog() {
	var f = os.Stdout
	// Disable Console Color, you don't need console color when writing the logs to file.
	// Logging to a file.
	logFile := viper.GetString("log.path")

	if logFile == "debug" {
		gin.DisableConsoleColor()
		log.SetLevel(log.DebugLevel)
		//log.SetReportCaller(true)
	} else {
		f, _ = os.Create(logFile)
	}

	log.SetFormatter(&log.TextFormatter{
		//DisableColors:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
		//ForceFormatting: true,
	})

	log.SetOutput(f)

	gin.SetMode(gin.DebugMode)
	gin.DefaultWriter = io.MultiWriter(f)
}

// Default returns an Engine instance with the Logger and Recovery middleware already attached.
func ginRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())
	return engine
}
