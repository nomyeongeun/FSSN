package main

// 필요한 모듈 및 함수 import
import (
	"context"
	"log"
	"time"
    "fmt"
	"google.golang.org/grpc"
	pb "full-stack-networking/noh-myeong-eun/clientstreaming/protos"
)

const (
    address = "localhost:50051"
    value    = 5
)

// 프로토에 정의된 Message 형태로 만들어 주는 함수
func make_message(message string) (*pb.Message) {
	return &pb.Message{Message :message}
}

// 메시지를 만들어서 서버에 보내는 함수
func doStreaming(c pb.ClientStreamingClient) error {
	messages := []*pb.Message{
        make_message("message #1"),
        make_message("message #2"),
        make_message("message #3"),
        make_message("message #4"),
        make_message("message #5"),
    }
	stream, err := c.GetServerResponse(context.Background()) // 수정
	if err != nil {
		log.Fatalf("Error while receiving response",err)
	}
	for _, msg := range messages {
		fmt.Printf("[client to server] %s\n", msg.Message)
        stream.Send(msg)
    }
	res, err := stream.CloseAndRecv()
	fmt.Printf("[server to client] %d\n", res.Value)
    
	return nil
} 

func main() {
    // gRPC 서버에 연결한다
    conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()

    // gRPC 클라이언트를 생성한다.
    c := pb.NewClientStreamingClient(conn) 
    _, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

	// 서버로 스트리밍한다
    doStreaming(c)
}