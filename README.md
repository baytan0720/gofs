# gofs
Go语言实现分布式存储系统  
  
### 7.29  
项目开始，搭建底层及RPC  
使用说明：  
拉取项目到本地并运行 

Linux系统下
```Bash
go mod tidy
cd NameNode && go run *.go
cd DataNode && go run *.go
cd Client && go run *.go
```

Windows系统下
```Bash
# 根目录
go mod tidy
# NameNode 目录
cd NameNode
go build
.\NameNode.exe
# DataNode 目录
cd DataNode
go build
.\DataNode.exe
# Client 目录
cd Client
go build
.\Client.exe
```
