# 程序说明
接收客户端 request，并将 request 中带的 header 写入 response header。

# 使用方法
1. 对required的参数name赋值
```
> curl -v http://localhost:8080/hello?name=john -H "phone:123" -H "email:abc@def.com"
* Mark bundle as not supporting multiuse
  < HTTP/1.1 200 OK
  < Accept: */*
  < Content-Type: application/json
  < Email: abc@def.com
  < Phone: 123
  < User-Agent: curl/7.79.1
  < Date: Sun, 24 Jul 2022 07:45:36 GMT
  < Content-Length: 25
  <
* Connection #0 to host localhost left intact
  {"message":"hello, john"}%
```
2. required的参数name缺失
```
> curl -v http://localhost:8080/hello?name1=john -H "phone:123" -H "email:abc@def.com"
* Mark bundle as not supporting multiuse
  < HTTP/1.1 400 Bad Request
  < Date: Sun, 24 Jul 2022 07:47:45 GMT
  < Content-Length: 47
  < Content-Type: text/plain; charset=utf-8
  <
* Connection #0 to host localhost left intact
  {"message":"missing required parameter [name]"}%
```
