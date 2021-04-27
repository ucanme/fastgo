## 接口文档

### 所有接口一律Post请求

#### 1 登陆接口同之前项目

#### 2 获取可预约的days

Path: /v1/available-days/list

Resp: 
```
{
    "data": {
        "days": [
            "2021-05-01",
            "2021-05-02",
            "2021-05-03",
            "2021-05-05",
            "2021-05-08",
            "2021-05-09"
        ]
    },
    "error_msg": "success",
    "error_no": 0,
    "request_id": "fe338962-5fb3-423e-bba1-20f39db0a543"
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
    "data": [
        {
            "hour": 9,
            "status": 0
        },
        {
            "hour": 10,
            "status": 0
        },
        {
            "hour": 13,
            "status": 0
        },
        {
            "hour": 14,
            "status": 0
        },
        {
            "hour": 15,
            "status": 0
        }
    ],
    "error_msg": "success",
    "error_no": 0,
    "request_id": "152b4815-57d4-49e8-8630-401e128d9b79"
}
```

#### 4获取可预约的小时
#### /avaliable-minutes/list
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
    "data": {
        "minutes": [
            {
                "minute": 0,
                "status": 0
            },
            {
                "minute": 10,
                "status": 0
            },
            {
                "minute": 20,
                "status": 0
            },
            {
                "minute": 30,
                "status": 0
            },
            {
                "minute": 40,
                "status": 0
            },
            {
                "minute": 50,
                "status": 0
            }
        ]
    },
    "error_msg": "success",
    "error_no": 0,
    "request_id": "0791d96d-fc24-4910-b29c-7fc483127c15"
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




