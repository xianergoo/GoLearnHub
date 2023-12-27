// package main

// func main() {
// 	// // 设置日志级别
// 	// // log.SetLevel(log.InfoLevel)
// 	// log.SetLevel(log.DebugLevel)
// 	// // 下面输出，日志级别由低到高，输出情况由上面的日志级别控制
// 	// // 例如：设置日志级别为 ErrorLevel，则 infor、debug、warn不再输出
// 	// log.Info("info")
// 	// log.Debug("debug")
// 	// log.Warn("warn")
// 	// log.Error("error")
// 	// log.Panic("panic")
// 	// log.Fatal("fatal")

// 	// log.SetFormatter(&log.JSONFormatter{})
// 	// log.Info("info")

// 	// log.SetOutput(os.Stdout)
// 	// log.Info("begin")

// 	// f, err := os.OpenFile("test_logrus.log", os.O_CREATE|os.O_WRONLY, 0666)
// 	// if err == nil {
// 	// 	log.SetOutput(f)
// 	// } else {
// 	// 	log.Info("create log file error")
// 	// }

// 	// log.Info("end")

// 	log.SetReportCaller(true)
// 	var event = "do order"
// 	var topic = "order"
// 	var key = 1001
// 	// log.Fatalf("event %s send to topic %s failed, used: %d",
// 	// 	event, topic, key)

// 	log.WithFields(log.Fields{
// 		"event": event,
// 		"topic": topic,
// 		"key":   key,
// 	}).Info("event send failed")

// }

// var log = logrus.New()

// func init() {
// 	log.Out = os.Stdout
// 	log.Formatter = &logrus.JSONFormatter{}

// 	logrus.SetLevel(logrus.InfoLevel)
// }

// func main() {
// 	var event = "do order"
// 	var topic = "order"
// 	var key = 1001
// 	// log.Fatalf("event %s send to topic %s failed, used: %d",
// 	// 	event, topic, key)

// 	log.WithFields(logrus.Fields{
// 		"event": event,
// 		"topic": topic,
// 		"key":   key,
// 	}).Info("event send failed")
// }

package main

import (
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
)

func init() {

	path := "message.log"

	/* 日志轮转相关函数

	   `WithLinkName` 为最新的日志建立软连接

	   `WithRotationTime` 设置日志分割的时间，隔多久分割一次

	    WithMaxAge 和 WithRotationCount二者只能设置一个

	    `WithMaxAge` 设置文件清理前的最长保存时间

	    `WithRotationCount` 设置文件清理前最多保存的个数

	*/

	// 下面配置日志每隔 1 分钟轮转一个新文件，保留最近 3 分钟的日志文件，多余的自动清理掉。

	writer, _ := rotatelogs.New(

		path+".%Y%m%d%H%M",

		rotatelogs.WithLinkName(path),
		rotatelogs.WithMaxAge(time.Duration(180)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(60)*time.Second),
	)

	log.SetOutput(writer)

	//log.SetFormatter(&log.JSONFormatter{})

}

func main() {
	for {
		log.Info("hello, world!")
		time.Sleep(time.Duration(2) * time.Second)
	}
}
