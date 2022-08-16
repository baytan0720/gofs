# gofs
Go语言实现分布式存储系统  
```
  _____  ____  ______ _____    
 / ____|/ __ \|  ____/ ____|   
| |  __| |  | | |__ | (___     
| | |_ | |  | |  __| \___ \    
| |__| | |__| | |    ____) |   
 \_____|\____/|_|   |_____/    
 ```

### 7.29  
项目开始，搭建底层及RPC  
使用说明：  
拉取项目到本地并运行   
##### Linux/Macos
```Bash
cd NameNode/cmd && go run main.go
cd DataNode/cmd && go run main.go
cd Client && go run *.go
```
### 8.3
实现注册中心、心跳机制


### 8.4
使用grpc
* 安装grpc和protobuf
* 进入DataNode目录
```bash
protoc -I .\internal\service\pb .\internal\service\pb\FileReplicate.proto --go-grpc_out="." --go_out="."
```
一般流程如下
1. 编写proto文件，指定rpc的输入和输出并定义rpc服务
2. 在服务端，实现rpcServiceServer方法，并且进行注册
3. 在客户端，编写调用rpcServiceClient方法的函数即可

### 8.5
优化逻辑

### 8.6
进一步优化心跳机制和重连机制
实现DataNode block信息上报

### 8.16
实现Put方法
实现Client端命令行解析

查看命令
```bash
./gofs help
```