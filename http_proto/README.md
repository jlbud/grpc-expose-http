# Protocol Buffers🌽
**annotations.proto** and **http.proto** copy from [[grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway)
](https://github.com/grpc-ecosystem/grpc-gateway/tree/master/third_party/googleapis/google/api), these two files are required to expose HTTP

# Compile🤖
In the current directory, exec the command to compile **annotations.proto** and **http.proto**:

```
protoc --go_out=plugins=grpc:./ ./http.proto
protoc --go_out=plugins=grpc:./ ./annotations.proto
```

generate these two files:
![pb.go](https://upload-images.jianshu.io/upload_images/2059314-c69bea6cf02027cd.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

continue compiling **rest.proto**:

```
protoc -I/usr/local/include -I. \
    -I${GOPATH}/src \
    -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --grpc-gateway_out=. --go_out=plugins=grpc:.\
     rest.proto
```

generate these two files:
![pb.go](https://upload-images.jianshu.io/upload_images/2059314-e6d37955d9aca592.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

# 👋👋👋
Proto pre work is finished