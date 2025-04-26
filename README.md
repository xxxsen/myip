myip
===

获取本机ip

## 基本配置

```jsonc
{
    "bind": ":5578",
    "headers": [ //从哪些header提取ip, 如果header都没有包含ip, 则直接从对端的conn获取ip地址
        "Cf-Connecting-Ip",
        "X-Real-IP",
        "X-Forwarded-For"
    ],
    "log_config": {
        "level": "info",
        "console": true
    }
}
```

## 运行

### 命令行运行

```shell
myip --config=./config.json
```

### docker运行

```yaml
services:
  myip:
    image: xxxsen/myip:latest
    container_name: "myip"
    restart: always
    volumes:
      - ./config:/config
    ports:
      - "5578:5578"
    command: "--config=/config/config.json"
```
