package toolpulsar

import (
	"context"
	"encoding/binary"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/go-restruct/restruct"
	"log"
	"pulsar-demo/model"
	"strings"
	"time"
)

var client pulsar.Client

var IsStop bool

func init() {
	IsStop = false
}

func Consume(req model.Request) {
	//使用client对象实例化consumer
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            req.Topic,
		SubscriptionName: "pulsar_tool",
		Type:             pulsar.Shared,
	})
	if err != nil {
		panic(err)
	}
	topicArr := strings.Split(req.Topic, "/")
	last := topicArr[len(topicArr)-1]
	fmt.Println("\n=====> READ QUEUE <=====", last, time.Now().Format("2006-01-02 15:04:05"))

	ctx := context.Background()
	defer consumer.Close()
	//无限循环监听topic
	for {
		if IsStop {
			break
		}
		msg, err := consumer.Receive(ctx)
		if err != nil {
			log.Fatal(err)
		} else {
			// fmt.Printf("Received message :   %s \n", time.Now().Format("2006-01-02 15:04:05"))
		}

		Count++
		now := time.Now()
		LastTime = now
		if Count == 1 {
			StartTime = LastTime
			fmt.Printf("%-8s\t%v\t%d\t\n", "[开始消费]", LastTime.Format("2006-01-02 15:04:05"), Count)
		}

		ms := msg.Payload()
		for i := 0; i < 1; i++ {
			if last == "q_emqx_online" { // 上下线消息
				resStruct := frestructOnline(ms)
				_ = resStruct
				if req.LogMsg {
					fmt.Printf("===>%02d [%s] %+v\n", time.Now().Second(), last, resStruct)
				}
			} else { // 常规消息
				resStruct := frestruct(ms)
				if req.LogMsg {
					fmt.Printf("===>%02d [%s] %+v\n", time.Now().Second(), last, resStruct)
				}
			}
		}
		_ = consumer.Ack(msg)
	}
}

func Start(req model.Request) {
	if req.Host == "" {
		panic("参数错误：host")
	}
	if req.Topic == "" {
		panic("参数错误：topic")
	}
	Init(req.Host)
	if req.LogCount {
		LogCount(1000)
	}
	topics := strings.Split(req.Topic, ",")
	for _, topic := range topics {
		req.Topic = topic
		go Consume(req)
	}

}

func Init(host string) {
	c, err := pulsar.NewClient(pulsar.ClientOptions{
		MaxConnectionsPerBroker: 5,
		URL:                     host,
	})
	if err != nil {
		panic(err)
	}
	client = c
}

func Producer(req model.Request) {
	Init(req.Host)
	ctx := context.Background()
	if req.Topic == "" {
		req.Topic = "test"
	}
	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: req.Topic,
	})

	if err != nil {
		log.Fatalf(" producer:%v", err)
	}

	defer producer.Close()

	msg := `{"header":{"messageId":"f612fb49845bee6191ea05e1548aa7a2","namespace":"Appliance.Control.ToggleX","triggerSrc":"CloudAlexa","method":"PUSH","payloadVersion":1,"from":"/appliance/2201208098807451860148e1e986b2fb/publish","uuid":"2201208098807451860148e1e986b2fb","timestamp":1673925167,"timestampMs":749,"sign":"2e4375b4631d573499dd0b0585cee295"},"payload":{"channel":0,"togglex":{"channel":0,"onoff":1,"lmTime":1673911325}}}`
	clientId := "2201208098807451860148e1e986b2fb"
	MsgStruct := model.NormalMsg{
		Flags:        1,
		Version:      1,
		Cluster:      1,
		QOS:          1,
		IPV4:         2130706433,
		RevTime:      uint64(time.Now().UnixMilli()),
		ClientIdSize: uint32(len(clientId)),
		ClientId:     clientId,
		TopicSize:    uint32(len(clientId)),
		Topic:        clientId,
		PayloadSize:  uint32(len(msg)),
		Payload:      msg,
	}

	Data, _ := restruct.Pack(binary.BigEndian, &MsgStruct)
	sendMsg := &pulsar.ProducerMessage{
		Payload: Data,
	}

	sd, err := producer.Send(ctx, sendMsg)
	if err != nil {
		log.Fatalf("Producer could not send message:%v", err)
	}
	fmt.Println(sd)

}

var Count int64         // 所有消息总数
var StartTime time.Time //所有消息开始时间
var LastTime time.Time  //所有消息结束时间
var LastCount int64
var TestCount int64     //测试消息总数
var TestStartTime int64 //测试消息开始时间
var TestLastTime int64  //测试消息结束时间
var TestLastCount int64
var specialCount int64 //处理特殊消息条数
var specialTime int64  //特殊消息总耗时

func LogCount(millsec int64) {
	go func() {
		for {
			time.Sleep(time.Duration(millsec) * time.Millisecond)
			Now := Count
			if Now == LastCount {
				continue
			}
			if Now > 0 {
				fmt.Printf("%-8s\t[%v ~ %v]\t总耗时:%v (%.2fs)\t消费速率:%d/秒\t消费总数:%v\n", "[统计信息]",
					StartTime.Format("2006-01-02 15:04:05"),
					LastTime.Format("2006-01-02 15:04:05"),
					LastTime.Sub(StartTime), LastTime.Sub(StartTime).Seconds(),
					Now-LastCount,
					Now,
				)
				LastCount = Now
			}
			if IsStop {
				break
			}
		}
	}()
}

func frestruct(bts []byte) *model.NormalMsg {
	c := model.NormalMsg{}
	_ = restruct.Unpack(bts, binary.BigEndian, &c)
	return &c
}

func frestructOnline(bts []byte) *model.OnlineMsg {
	c := model.OnlineMsg{}
	_ = restruct.Unpack(bts, binary.BigEndian, &c)
	return &c
}
