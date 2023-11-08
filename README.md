# 说明

pulsar 测试工具

# 参数

```
  -h string                        
        host           主机地址    
  -lc                              
        logCount       打印统计信息
  -lm                              
        logMsg         打印消息    
  -t string                        
        subTopic       主题       
```

# 示例

打印统计信息
go run .\main.go -h pulsar://us-test-pulsar.meross.test.vpc:6650 -t persistent://meross/iot_raw/q_emqx_ctrl_msg -lc

打印消息详情
go run .\main.go -h pulsar://us-test-pulsar.meross.test.vpc:6650 -t persistent://meross/iot_raw/q_emqx_ctrl_msg -lm

消费所有topic
go run main.go -h pulsar://127.0.0.1:6650 -t persistent://meross/iot_raw_normal/q_emqx_online,persistent://meross/iot_raw_normal/q_emqx_ctrl_msg_old,persistent://meross/iot_raw_normal/q_emqx_stat_msg_old,persistent://meross/iot_raw_normal/q_emqx_stat_msg,persistent://meross/iot_raw_normal/q_emqx_metry_msg,persistent://meross/iot_raw_normal/q_emqx_need_ack,persistent://meross/iot_raw_high/q_emqx_major -lm
