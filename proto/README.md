# Protocol Buffers🌽
**annotations.proto** and **http.proto** copy from [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway/tree/master/third_party/googleapis/google/api), these two files are required to expose HTTP

# Compile🤖
In the current directory, exec the command to compile **annotations.proto** and **http.proto**:

```
protoc --go_out=plugins=grpc:./ ./http.proto
protoc --go_out=plugins=grpc:./ ./annotations.proto
```

generate these two files:
```
-rw-r--r--  1 kevin  staff   2.4K Feb  3 14:19 annotations.pb.go
-rw-r--r--  1 kevin  staff    20K Feb  3 14:19 http.pb.go
```

continue compiling **rest.proto**:

```
protoc -I/usr/local/include -I. \
    --grpc-gateway_out=. --go_out=plugins=grpc:.\
     rest.proto
```

generate these two files:
```
-rw-r--r--  1 kevin  staff   7.4K Feb  3 14:46 rest.pb.go
-rw-r--r--  1 kevin  staff    10K Feb  3 14:46 rest.pb.gw.go
```

# 👋👋👋
Proto pre work is finished
