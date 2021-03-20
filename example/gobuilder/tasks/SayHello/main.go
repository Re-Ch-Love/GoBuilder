package main

import (
	"fmt"
	caller2 "gitee.com/KongchengPro/GoBuilder/pkg/tdk/caller"
	"os"
)

func main() {
	// os.Args[1]是经过json序列化的，GoBuilder传递来的信息
	// 这里把他反序列化以便操作
	caller := caller2.UnmarshalArgs(os.Args[1])
	// 可以输出来查看caller
	//fmt.Printf("%+v", caller)
	// caller.Args是调用task时传入的参数
	if len(caller.Args) > 0 {
		fmt.Println("hello, " + caller.Args[0])
	} else {
		fmt.Println("hello, world")
	}
}
