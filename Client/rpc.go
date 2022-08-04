package main

//rpc调用的参数，注意结构体字段首字母必须大写
//结构体命名标准:方法名+Args/Reply

type Args struct {
}

type Reply struct {
}

type HelloArgs struct{}

type HelloReply struct {
	S string
}
