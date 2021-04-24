## 接口文档

### 所有接口一律Post请求

#### 1 登陆接口同之前项目

#### 2 获取可预约的days

Path: /v1/available-days/list

Resp: 
```
{
    "data":{
        "days":["2021-10-10","2021-10-10"]
    }
}

```

#### 3获取可预约的小时
#### /v1/avaliable-hours/list
Req:
```
{
    "day" : "2021-10-01"
}
```

Resp
```
{
    "data":{
        "hours":["2021-10-10","2021-10-10"]
    }
}
```

#### 4获取可预约的小时
#### /v1/avaliable-hours/list
Req:
```
{
    "day" : "2021-10-01",
    "hour": 10 
}
```

Resp
```
{
    "data":{
        "miniutes":["2021-10-10","2021-10-10"]
    }
}
```



#### 5进行预约
#### /v1/make-appointment

Req:
```
{
    "day" : "2021-10-01",
    "hour": 10,
    "minute": 0,
    "open_id":"",
    "name" : "",
    "phone_num":"",
}
```


#### 6预约列表
#### /v1/appointment/list
Req:
```bazaar
{
  "open_id" : "", 
}
```

Resp 
```bazaar
{
    "data":[
            {
            "datetime" : "2021-10-23 10:00:00",
            "appointment_id" : 123
        }
      ]
}  
```

#### 7取消预约
#### /v1/cancel-appointment

Req:
```bazaar
{
  "open_id" : "", 
  "appintment_id":123123
}
```



#### 8签到
v1./sign-in
Req:
```bazaar
{
  "appintment_id":123123
}
```



#### 9管理系统预约列表
#### /admin/appointment/list
Req:
```bazaar
{
  "date" : "2019-10-10", //非必填 不填返回最近一个月预约的所有
}
```

Resp
```bazaar
[
"data": [{
    "day" : "2021-10-01",
    "hour": 10,
    "minute": 0,
    "open_id":"",
    "name" : "",
    "phone_num":"",
    "status" : 0 //0正常 1已经取消
}]
}
```




#### 9管理系统签到列表
#### /admin/signin-appointment/list
Req:
```bazaar
{
  "date" : "2019-10-10", //非必填 不填返回最近一个月预约的所有
}
```

Resp
```bazaar
[
"data": [{
    "day" : "2021-10-01",
    "hour": 10,
    "minute": 0,
    "open_id":"",
    "name" : "",
    "phone_num":"",
    "status" : 0 //0正常 1已经取消
}]
}
```


#### 10如果有新签到的通过websocket发送
#### /ws/signin-notify
{
    "appointment_id":123,
    "event_type" : "1"  //1签到 目前来看只有1
}




