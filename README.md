# Protocol Buffers🐣

# Dependent🔗
Download HTTP plugin
```go
$ go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
```

# Run⛹️‍♂️
Compile the ProtoBuf file
1. Exec **[http_proto](https://github.com/klbud/grpc-expose-http/blob/main/http_proto/README.md)**

2. Start the gRPC service first
```go
go run -v grpc_server.go
```
3. Then start the HTTP service
```go
go run -v http_server.go
```

# Test it🐞
GET
```
// request
\
curl -X GET \
  http://localhost:8080/get/hello%20gRpc \
  -H 'content-type: application/json' \
  | json_pp
```

```
// response
> curl -X GET \
>   http://localhost:8080/get/hello%20gRpc \
>   -H 'content-type: application/json' \
>   | json_pp
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    26  100    26    0     0   2363      0 --:--:-- --:--:-- --:--:--  2363
{
   "value" : "Get=hello gRpc"
}
```

POST
```
// request
\
curl -X POST \
  http://localhost:8080/post \
  -H 'content-type: application/json' \
  -d '{"value":"hello gRPC"}' \
 |  json_pp
```

```
// response
> curl -X POST \
>   http://localhost:8080/post \
>   -H 'content-type: application/json' \
>   -d '{"value":"hello gRPC"}' \
>  |  json_pp
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    49  100    27  100    22   1500   1222 --:--:-- --:--:-- --:--:--  2722
{
   "value" : "Post=hello gRPC"
}
```
