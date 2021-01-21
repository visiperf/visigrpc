package main

import (
	"fmt"
	"log"

	"github.com/visiperf/visigrpc/v3"
	"google.golang.org/grpc"
)

const target = "YOUR_GRPC_SERVER_TARGET"
const maxMsgSize = 1024 * 1024 * 20

func main() {
	opts := []grpc.DialOption{
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(maxMsgSize),
		),
	}

	conn, err := visigrpc.NewSecuredDialer().Dial(target, opts...)
	if err != nil {
		log.Fatalf("failed to dial with grpc server: %v", err)
	}
	defer conn.Close()

	fmt.Println(conn.GetState())
}
