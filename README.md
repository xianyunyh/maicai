# 抢菜
电商平台抢菜脚本。
上海疫情期间，由于封控，导致物资严重缺乏，生活所需不得不依靠社区团购和电商平台，在疫情初期开发了抢菜脚本，解决温饱问题。

## 原理
抓取美团买菜小程序接口，程序模拟请求。所以使用前需要会抓包。

需要会用fiddler抓微信小程序的包，具体操作自行搜索（PC微信+fildder非常简单）

## 使用说明
程序使用go语言进行编译，可以下载编译后的文件，也可以自行编译
```
go mod init maicai

go build .

```

## 配置文件

```toml
# 开启后台定时任务
cron_enable=false
# 定时任务执行规则
# 秒 分钟 小时 天 月 星期
cron_rule="50 59 5 * * *"
# 日志级别 trace、debug、info、error、panic
log_level="info"
# 城市id： 1:上海,2:北京,4:广州,7:深圳
city_id=1
# 地址id /api/c/malluser/address/list获取
address_id=23456
# 固定值
utm_term="5.32.5"

#门店信息点 /api/c/poi/location/lbs/v2获取
poi=123
# 精度 
homepage_lng=
# 纬度latitude=
homepage_lat=
# 用户登录token header里的t字段 必须
token=""
# 停顿时间 单位（毫秒）
sleep_ms=400
```

## 注意事项
仅供交流使用，切勿用于商业用途
