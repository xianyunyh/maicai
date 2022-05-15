package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/robfig/cron/v3"
)

var (
	BuildDate       string
	defaultCronRule = "30 59 5 * * *"
)

func main() {
	fmt.Println("Last Build DateTime ", BuildDate)
	var confFile string
	flag.StringVar(&confFile, "f", "config.toml", "config文件")
	flag.Parse()
	conf := &Config{}
	f, err := os.OpenFile("meituan.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err == nil {
		defer f.Close()
	}
	confData, err := ioutil.ReadFile(confFile)
	if err != nil {
		log.Errorf("读取配置文件出错:%s", err.Error())
		return
	}
	//解析配置
	conf, err = ParseConf(bytes.NewBuffer(confData))
	if err != nil {
		log.Errorf("解析文件遇到错误:%s", err.Error())
		return
	}
	// 初始化日志
	InitLogger(withOutput(f), withLogLevel(conf.LogLevel))
	c := cron.New(cron.WithSeconds())
	job := &MeiTuanJob{conf: conf, notify: TmuxNotify{media: conf.Media}}
	if !conf.CronEnable {
		job.Run()
		return
	}
	rule := defaultCronRule
	if conf.CronRule != "" {
		rule = conf.CronRule
	}
	next, err := getNextTime(rule, cronOpt)
	if err != nil {
		log.Errorf("定时任务[%s]规则不合法:%s", rule, err.Error())
		return
	}
	log.Info("程序将在后台定时运行")
	log.Infof("下一次运行时间:%s", next.Format("2006-01-02 15:04:05"))
	c.AddJob(rule, job)
	c.Start()
	signalChans := make(chan os.Signal, 3)

	signal.Notify(signalChans, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT)

	select {
	case <-signalChans:
		log.Info("stop")
		time.Sleep(1 * time.Second)
		c.Stop()
	}

}
