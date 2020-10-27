# 理解 HTTP
HTTP 请求的结构
```bash
curl -s -o /dev/null -v http://google.com
*   Trying 74.125.24.139...
* TCP_NODELAY set
* Connected to google.com (74.125.24.139) port 80 (#0)
> GET / HTTP/1.1
> Host: google.com
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 301 Moved Permanently
< Location: http://www.google.com/
< Content-Type: text/html; charset=UTF-8
< Date: Tue, 27 Oct 2020 02:58:20 GMT
< Expires: Thu, 26 Nov 2020 02:58:20 GMT
< Cache-Control: public, max-age=2592000
< Server: gws
< Content-Length: 219
< X-XSS-Protection: 0
< X-Frame-Options: SAMEORIGIN
< 
{ [219 bytes data]
* Connection #0 to host google.com left intact
* Closing connection 0

```
- 以字符`>`开头的行描述了客户端发出的请求
- 以字符`<`开头的行描述收到的请求
- 请求的详细信息描述了随请求发送的一些报头，这些报头向服务器提供了一些有关客户端的信息
- 响应详细描述了一些报头，这些报头指出了响应的内容类型、长度和发送时间