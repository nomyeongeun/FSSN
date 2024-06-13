package main

// 모듈 및 함수 import
import (
    "log"
    "net"
    "fmt"
	"google.golang.org/grpc"
	pb "full-stack-networking/noh-myeong-eun/serverstreaming/protos"
)

// gRPC 서버 구조체 정의
type server struct {
    pb.UnimplementedServerStreamingServer
}

// Message에 담는 함수
func make_message(message string) (*pb.Message) {
	return &pb.Message{Message :message}
}

// Message를 만들어서 보내는 함수
func (s *server) GetServerResponse(req *pb.Number, stream pb.ServerStreaming_GetServerResponseServer) error {
	messages := []*pb.Message{
        make_message("message #1"),
        make_message("message #2"),
        make_message("message #3"),
        make_message("message #4"),
        make_message("message #5"),
    }
	fmt.Printf("Server processing gRPC server-streaming {%d}.\n", req.Value)

	for _, message := range messages {
        if err := stream.Send(message); err != nil {
            return err
        }
    }
    return nil
} 

// gRPC 서버를 열고 위의 함수로 Message를 보낸다.
func main() {
    fmt.Printf("Starting server. Listening on port 50051.\n")
    lis, err := net.Listen("tcp", "localhost:50051")
    if err != nil {
        log.Fatalf("failed to listen %v", err)
    }

    grpcServer := grpc.NewServer()
	pb.RegisterServerStreamingServer(grpcServer, &server{})

    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}