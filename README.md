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

go run .\main.go -h pulsar://xx.test.vpc:6650 -t persistent://meross/iot_raw/q_emqx_online -lc