package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	A, B int
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing error: ", err)
	}

	args := Args{4, 5}
	var reply int

	// 调用Multiply方法
	err = client.Call("Arith.Multiply", &args, &reply)
	if err != nil {
		log.Fatal("arith error: ", err)
	}

	fmt.Printf("Multiply: %d * %d = %d\n", args.A, args.B, reply)

	args = Args{10, 0}
	var fReply float64

	// 调用Divide方法
	err = client.Call("Arith.Divide", &args, &fReply)
	if err != nil {
		log.Fatal("arith error: ", err)
	}

	fmt.Printf("Divide: %d / %d = %f\n", args.A, args.B, fReply)
}
