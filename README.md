## 接口文档

### 所有接口一律Post请求

#### 1 登陆
Path: /v1/login
Req:
```json
{
  "user_id": "",
  "password" : ""
}
```

Resp:
```
{
    },
    "error_msg": "success",
    "error_no": 0,
    "request_id": "fe338962-5fb3-423e-bba1-20f39db0a543"
}
```

#### 2 退出

Path: /v1/logout

Req:无
```
{
    },
    "error_msg": "success",
    "error_no": 0,
    "request_id": "fe338962-5fb3-423e-bba1-20f39db0a543"
}
```