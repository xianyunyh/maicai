# 抢菜
电商平台抢菜脚本。
上海疫情期间，由于封控，导致物资严重缺乏，生活所需不得不依靠社区团购和电商平台，在疫情初期开发了抢菜脚本，解决温饱问题。

## 美团脚本原理
抓取美团买菜小程序接口，程序模拟请求。所以使用前需要会抓包。

需要会用fiddler抓微信小程序的包，具体操作自行搜索（PC微信+fildder非常简单）

## 叮咚脚本原理
分析叮咚H5页面，程序模拟请求。使用前，请使用PC微信客户端，将https://wx.m.ddxq.mobi/ 发送给好友或者文件传输助手。
然后在微信中打开[叮咚H5网页](https://wx.m.ddxq.mobi/)。并进行授权登录。然后点击 **我的**->**收获的地址** 获取配置文件相关信息。

## 使用说明
程序使用go语言进行编译，可以下载编译后的文件: [下载地址](https://github.com/xianyunyh/maicai/releases/tag/0.0.2)，也可以自行编译
```
$ go mod init maicai
 
$ go build .

# 运行编译后的文件 配置文件见下方
$ ./maicai.exe -f config.toml
```

## 配置文件

配置文件使用toml格式，里面的信息需要抓包填写。

```toml
# 日志级别 trace、debug、info、error、panic
log_level="info"

# 停顿时间 单位（毫秒）
sleep_ms=500
[meituan]
#开启美团脚本
enable=true
# 城市id： 1:上海,2:北京,4:广州,7:深圳
city_id=1
# 地址id /api/c/malluser/address/list获取
address_id=1212121
# 固定值
utm_term="5.32.5"
#门店信息点 /api/c/poi/location/lbs/v2获取
poi=111
# 精度 
homepage_lng=122.2222
# 纬度
homepage_lat=32.2222
# 用户登录token header里的t字段
token="232131231"
# 停顿时间 单位（毫秒）
sleep_ms=500
[dingdong]
#开启叮咚脚本
enable=true
#用户UID
uid="1234567"
# 点击个人中心收货地址 /api/v1/user/address/ 中获取以下信息
city_id="0101" #0101上海
# 收货地址id
address_id="1212121"
# station_id 门店id
poi="12121" 
# 用户token header中Cookie中的DDXQSESSID值
#比如cookie: DDXQSESSID=1212121,那么token=1212121
token="121212121"
```

## 注意事项
仅供交流使用，切勿用于商业用途
