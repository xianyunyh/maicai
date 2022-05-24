package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"time"
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
	f, err := os.OpenFile("maicai.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
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
	sleepMs := time.Duration(conf.SleepMs)
	jobs := make([]Job, 0, 2)
	if conf.DingDong.Enable {
		ddJob := NewDDJob(conf.DingDong, time.Duration(conf.SleepMs), conf.Debug)
		jobs = append(jobs, ddJob)
	}
	if conf.Meituan.Enable {
		meituanJob := &MeiTuanJob{conf: conf.Meituan, SleepMs: sleepMs}
		jobs = append(jobs, meituanJob)
	}
	if len(jobs) == 0 {
		log.Info("没有配置可执行的job")
		return
	}
	wg := sync.WaitGroup{}
	for _, job := range jobs {
		wg.Add(1)
		go func(j Job) {
			defer wg.Done()
			j.Run()
		}(job)
	}
	wg.Wait()
}
