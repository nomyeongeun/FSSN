package main

// 필요한 함수 및 모듈 import
import (
	"context"
	"log"
	"time"
	"io"
	"fmt"
	"google.golang.org/grpc"
	pb "full-stack-networking/noh-myeong-eun/bidirectional-streaming/protos"
)

const (
    address = "localhost:50051"
)

// 프로토에 정의된 Message 형태로 만들어주는 함수
func make_message(message string) (*pb.Message) {
	return &pb.Message{Message :message}
}

// 서버에 streaming하고 나서 서버에서 받은 값을 출력한다
func doStreaming(c pb.BidirectionalClient) error {
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
	stream.CloseSend()
	for {
		content,err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
            log.Fatalf("ListInfo stream - %v", err)
        }
    	fmt.Printf("[server to client] %s\n", content.GetMessage())
	}
	
} 

func main() {
    // gRPC 서버에 연결한다
    conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()

	// gRPC 클라이언트를 생성한다
    c := pb.NewBidirectionalClient(conn) 
    _, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()
	
	// 양방향 스트리밍을 한다
	doStreaming(c)
	
}