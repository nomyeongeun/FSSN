package main

// 필요한 함수 및 모듈 import
import (
    "log"
    "net"
	"io"
	"fmt"
	"google.golang.org/grpc"
	pb "full-stack-networking/noh-myeong-eun/bidirectional-streaming/protos"
)

// gRPC 서버 구조체 정의
type server struct {
    pb.UnimplementedBidirectionalServer
}

// client로부터 streaming을 받아서 그대로 다시 client에게 streaming하는 함수
func (s *server) GetServerResponse(stream pb.Bidirectional_GetServerResponseServer) (error) {
	fmt.Printf("Server processing gRPC bidirectional streaming.\n")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		stream.Send(req)
		if err != nil {
			log.Fatalf("Error while reading client stream", err)
		}

	}
} 

// gRPC 서버를 열고 핸들러를 추가한다
func main() {
	fmt.Printf("Starting server. Listening on port 50051.\n")
    lis, err := net.Listen("tcp", "localhost:50051")
    if err != nil {
        log.Fatalf("failed to listen %v", err)
    }

    grpcServer := grpc.NewServer()
	pb.RegisterBidirectionalServer(grpcServer, &server{})

    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}