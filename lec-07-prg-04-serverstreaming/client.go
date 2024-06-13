package main

import (
	"context"
	"log"
	"time"
	"google.golang.org/grpc"
	"io"
    "fmt"
	pb "full-stack-networking/noh-myeong-eun/serverstreaming/protos"
)

const (
    address = "localhost:50051"
    value    = 5
)

func main() {
    // gRPC 서버에 연결한다
    conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()

    // gRPC 클라이언트를 생성한다
    c := pb.NewServerStreamingClient(conn) 
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    // 원격 함수를 호출한다
    stream, err := c.GetServerResponse(ctx, &pb.Number{Value: value})
    if err != nil {
        log.Fatalf("GetServerResponse error: %v", err)
    }

    // 서버에서 보내는 값들을 받아 출력한다
	for {
		content,err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
            log.Fatalf("ListInfo stream - %v", err)
        }
    	fmt.Printf("[server to client] %s\n", content.GetMessage())
	}
    
}