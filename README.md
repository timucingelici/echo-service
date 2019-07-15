# PURPOSE

This is a small demo project, written in Go to containerise and deploy as a micro-service.

It accepts any kind of request over HTTP and echoes the request body back to the client.

# EXAMPLE

```
$ http POST localhost:8888 hello=world

HTTP/1.1 200 OK
Content-Length: 18
Content-Type: text/plain; charset=utf-8
Date: Mon, 15 Jul 2019 19:52:19 GMT

{
    "hello": "world"
}
```

```
$ curl -X PUT localhost:8888/some-endpoint --data "Hello, 世界"

Hello, 世界%
```