package main

// 필요한 모듈 및 함수 import
import (
    "log"
    "net"
	"io"
	"fmt"
	"google.golang.org/grpc"
	pb "full-stack-networking/noh-myeong-eun/clientstreaming/protos"
)

// gRPC 서버 구조체 정의
type server struct {
    pb.UnimplementedClientStreamingServer
}

// 클라이언트로부터 streaming 받아 처리하는 함수
func (s *server) GetServerResponse(stream pb.ClientStreaming_GetServerResponseServer) (error) {
	fmt.Printf("Server processing gRPC client-streaming.\n")
	count := 0
	for {
		_, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Number{Value : int32(count)})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream", err)
		}
		count++
		
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
	pb.RegisterClientStreamingServer(grpcServer, &server{})

    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}