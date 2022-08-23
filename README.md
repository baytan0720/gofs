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
项目开始   
##### Linux/Macos
```Bash
cd sh && sh nn.sh
cd sh && sh dn.sh
cd src/client/cmd && go build -o gofs main.go
./gofs help
```
##### gprc
* 安装grpc和protobuf
* 进入DataNode目录
```bash
cd src
protoc -I ./proto ./proto/*.proto --go_out=.
protoc -I ./proto ./proto/*.proto --go-grpc_out=.
```