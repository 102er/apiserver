package main

import (
	"fmt"
	v1 "github.com/102er/apiserver/api/v1"
)

func main() {

	reply := v1.HelloReply{Message: "102er"}
	fmt.Println("start api server,", reply.Message)
}
