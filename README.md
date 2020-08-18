# Golang Gin
動手練習，並製作筆記

# 靜態文件處理

存放目錄
```
./static
```

![image](./images/20200818175658.png)

# / 

**GET**

回傳

```
{"message":"Hi"}
```
    
![image](./images/20200818165824.png)

# /login

**GET**

![image](./images/20200818170136.png)

**POST**

![image](./images/20200818170406.png)

# /json

**GET**

回傳

```
{"age":20,"message":"Hello Golang","name":"IvesShe"}
```

![image](./images/20200818170556.png)

# /more_json

**GET**

回傳

```
{"name":"ChiChi","message":"Hello World","age":18}
```

![image](./images/20200818170640.png)

# /web

**GET**

回傳

```
{"age":"","message":"Hi","name":"18"}
```

![image](./images/20200818170740.png)

# /blog/:name/:age

**GET**

輸入 ex

```
http://localhost:9090/blog/iveshe/20
```

回傳

```
{"age":"20","name":"iveshe"}
```

![image](./images/20200818170920.png)

# /user

**GET**

輸入 ex

```
http://localhost:9090/user?username=IvesShe&password=123456
```

回傳

```
{"status":"ok"}{"username":"IvesShe","password":"123456"}
```

![image](./images/20200818171233.png)

# /form

**POST**

回傳

```
{
    "status": "/form POST ok"
}{
    "username": "ivesshe",
    "password": "111111"
}
```

![image](./images/20200818171450.png)

# /json

**POST**

回傳

```
{
    "status": "/json POST ok"
}{
    "username": "ivesshe",
    "password": "111111"
}
```

![image](./images/20200818171426.png)

# /upload

**GET**

![image](./images/20200818171716.png)

![image](./images/20200818171741.png)

**POST**

![image](./images/20200818171810.png)




# /upload_more

- /upload **GET**
![image](./images/20200818172110.png)

- upload_more **POST**

```
{"message":"3 files uploaded!"}
```

![image](./images/20200818172236.png)

上傳成功

![image](./images/20200818172254.png)

# /turn_this

轉向turn_that

# /turn

輸出
```
{"message":"turn_that"}
```

![image](./images/20200818173025.png)

# /shop

## GET

輸出
```
{"method":"GET"}
```

![image](./images/20200818173320.png)

## POST

輸出
```
{
    "method": "POST"
}
```
![image](./images/20200818173809.png)

## PUT

輸出
```
{
    "method": "PUT"
}
```
![image](./images/20200818173926.png)

## DELETE

輸出
```
{
    "method": "DELETE"
}
```

![image](./images/20200818173942.png)

# /book

同/shop

# /video

## shop 

輸出
```
{"method":"/video/shop"}
```

![image](./images/20200818174718.png)

## login 

輸出
```
{"method":"/video/login"}
```

![image](./images/20200818174914.png)

## user 

輸出
```
{"method":"/video/user"}
```

![image](./images/20200818174932.png)

# /member

輸出
```
{"msg":"使用了ShowHandler中間件..."}{"message":"member"}
```

![image](./images/20200818175114.png)

![image](./images/20200818175157.png)