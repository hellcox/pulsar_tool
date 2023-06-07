package main

import (
	"flag"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"os"
	"os/signal"
	"pulsar-demo/model"
	"pulsar-demo/tool/toolpulsar"
	"syscall"
)

func main() {
	//action := flag.String("a", "", "Action:\t\t执行动作e")
	host := flag.String("h", "", fmt.Sprintf("%-15s%s", "host", "主机地址"))
	topic := flag.String("t", "", fmt.Sprintf("%-15s%s", "subTopic", "主题"))
	//addr := flag.String("addr", "", fmt.Sprintf("%-15s%s", "localAddress", "指定源IP"))
	//useSsl := flag.Bool("ssl", false, fmt.Sprintf("%-15s%s", "openSsl", "是否启用ssl"))
	//port := flag.Int("p", 1883, fmt.Sprintf("%-15s%s", "port", "端口号"))
	logMsg := flag.Bool("lm", false, fmt.Sprintf("%-15s%s", "logMsg", "打印消息"))
	logCount := flag.Bool("lc", false, fmt.Sprintf("%-15s%s", "logCount", "打印统计信息"))

	//解析命令行参数写入注册的flag里
	flag.Parse()
	request := &model.Request{
		Host:     *host,
		Topic:    *topic,
		LogMsg:   *logMsg,
		LogCount: *logCount,
	}

	fmt.Println(jsoniter.MarshalToString(request))
	go toolpulsar.Start(*request)

	// 优雅关闭，等待关闭信号
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	select {
	case sig := <-c:
		{
			_ = sig
			os.Exit(1)
		}
	}
}
