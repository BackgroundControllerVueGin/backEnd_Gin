# backEnd_Gin 说明文档

## 接口说明

### user

#### /user_login

>POST

```json
{
  "username":"12345","password":"123456"
}
```



```json
{
  "code": 200,
  "data": [
    {
      "password": "123456",
      "username": "root"
    }
  ]
}
```

#### /user_register

> POST

```json
{
  "username": "11111",
  "password": "123456",
  "passwordAgain": "123456"
}
```

```json
{
  "code": 200,
  "data": {
    "username": "11111",
    "password": "123456"
  },
  "msg": "注册成功"
}
```

#### /user_cancel

> POST

```json
{
  "username": "11111",
  "password": "123456"
}
```

```json
{
  "code": 200,
  "data": [
    {
      "username": "11111",
      "password": "123456"
    }
  ],
  "msg": "注销成功",
  "number": 1
}
```

>number:注销的账号数

### student

#### /student_show

>GET

```
N/A
```

```JSON
{
  "code": 200,
  "data": 
  [
    {
      "studentID": 153,
      "name": "零梦",
      "age": 20,
      "sex": "男",
      "department": "人工智能"
    },
    {
      "studentID": 154,
      "name": "李明",
      "age": 21,
      "sex": "女",
      "department": "人工智能"
    },
    {
      "studentID": 155,
      "name": "委屈",
      "age": 22,
      "sex": "男",
      "department": "人工智能"
    },
    {
      "studentID": 156,
      "name": "林妹妹",
      "age": 22,
      "sex": "女",
      "department": "现代音乐"
    },
    {
      "studentID": 157,
      "name": "0m",
      "age": 22,
      "sex": "男",
      "department": "土木工程"
    }
  ]
}
```

#### /student_show/studentID_search

>POST

```json
{
  "studentID":"153"
}
```

```jSON
{
  "code": 200,
  "data": [
    {
      "studentID": 153,
      "name": "零梦",
      "age": 20,
      "sex": "男",
      "department": "人工智能"
    }
  ]
}
```

