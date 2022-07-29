# gofs
Go语言实现分布式存储系统  
  
### 7.29  
项目开始，搭建底层及RPC  
使用说明：  
拉取项目到本地并运行  
```Bash
go mod tidy
cd NameNode && go run *.go
cd ../DataNode && go run *.go
cd ../Client && go run *.go
```
