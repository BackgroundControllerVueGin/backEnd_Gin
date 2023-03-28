# backEnd_Gin 说明文档

* [backEnd\_Gin 说明文档](#backend_gin-说明文档)
  * [接口说明](#接口说明)
    * [user](#user)
      * [/user\_login](#user_login)
      * [/user\_register](#user_register)
      * [/user\_cancel](#user_cancel)
    * [student](#student)
      * [/student\_list](#student_list)
      * [/student\_searchID](#student_searchid)
      * [/student\_delete](#student_delete)
      * [/student\_add](#student_add)
      * [/student\_modify](#student_modify)
    * [功能类接口](#功能类接口)
      * [/next\_page](#next_page)

## 接口说明

### user

#### /user_login

>POST

```json
{
  "username": "admin",
  "password": "123456"
}
```



```json
{
  "code": 200,
  "data": [
    {
      "no": "2",
      "username": "admin",
      "password": "123456"
    }
  ],
  "message": "登录成功"
}
```

#### /user_register

> POST

```json
{
  "username": "123456",
  "password": "123456",
  "passwordAgain": "123456"
}
```

```json
{
  "code": 200,
  "data": {
    "no": "9",
    "username": "123456",
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

#### /student_list

>GET

```
N/A
```

```JSON
{
  "code": 200,
  "data": [
    {
      "studentID": "1234567153",
      "name": "零梦",
      "age": "20",
      "sex": "男",
      "department": "人工智能",
      "studentclass": "软件工程一班"
    },
    {
      "studentID": "1234567154",
      "name": "李明",
      "age": "21",
      "sex": "女",
      "department": "人工智能",
      "studentclass": "软件工程二班"
    },
    {
      "studentID": "1234567155",
      "name": "委屈",
      "age": "22",
      "sex": "男",
      "department": "人工智能",
      "studentclass": "软件工程一班"
    },
    {
      "studentID": "1234567156",
      "name": "林妹妹",
      "age": "22",
      "sex": "女",
      "department": "现代音乐",
      "studentclass": "现代音乐一班"
    },
    {
      "studentID": "1234567157",
      "name": "0m",
      "age": "22",
      "sex": "男",
      "department": "土木工程",
      "studentclass": "和水泥一班"
    },
    {
      "studentID": "1234567158",
      "name": "yw",
      "age": "23",
      "sex": "男",
      "department": "继续教育",
      "studentclass": "人工智能一班"
    }
  ]
}
```

#### /student_searchID

>POST

```json
{
  "studentID":"1234567153"
}
```

```jSON
{
  "code": 200,
  "data": [
    {
      "studentID": "1234567153",
      "name": "零梦",
      "age": "20",
      "sex": "男",
      "department": "人工智能",
      "studentclass": "软件工程一班"
    }
  ]
}
```

#### /student_delete

> POST

```json
{
  "studentID": "12344444"
}
```

```json
{
  "code": 200,
  "msg": "删除成功",
  "number": 1
}
```

#### /student_add

> POST

```json
{
  "studentID": "12344444",
  "name": "零梦",
  "age": "20",
  "sex": "男",
  "department": "人工智能",
  "studentclass":"软件工程一班"
}
```

```json
{
  "code": 200,
  "data": {
    "studentID": "12344444",
    "name": "零梦",
    "age": "20",
    "sex": "男",
    "department": "人工智能",
    "studentclass": "软件工程一班"
  },
  "msg": "注册成功"
}
```

#### /student_modify

> POST

```JSON
//其中的学号是必须数据库内所有的数据，建议与/student_searchID联动使用
{
  "studentID": "123",
  "name": "大猛",
  "age": "20" ,
  "sex": "男",
  "department": "人工智能"
}
```

```json
{
  "code": 200,
  "data": {
    "studentID": "123",
    "name": "大猛",
    "age": "20",
    "sex": "男",
    "department": "人工智能"
  },
  "msg": "修改成功"
}
```

### 功能类接口

#### /next_page

> /next_page：暂定该名，会修改
>
> POST:用于实现student表的翻页功能

```json
//数据类型为INT
//data_number:为一页有多少数据
//current_page:为当前页码数
{
  "data_number": 4,
  "current_page": 2
}
```

```json
//total为当前数据库内有的数据条数
{
  "code": 200,
  "data": [
    {
      "studentID": "1234567157",
      "name": "0m",
      "age": "22",
      "sex": "男",
      "department": "土木工程",
      "studentclass": "和水泥一班"
    },
    {
      "studentID": "1234567158",
      "name": "yw",
      "age": "23",
      "sex": "男",
      "department": "继续教育",
      "studentclass": "人工智能一班"
    }
  ],
  "total": 6
}
```

