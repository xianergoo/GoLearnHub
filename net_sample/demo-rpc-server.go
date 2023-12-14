package main

import (
	"errors"
	"log"
	"net"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Arith struct{}

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, reply *float64) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	*reply = float64(args.A) / float64(args.B)
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)

	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error: ", err)
	}

	log.Println("rpc server listening on :1234 ...")
	rpc.Accept(l)
}
