package main

// (1) grpc를 포함해서 필요한 모듈 import
// (2) 원격 호출될 함수를 import
// (3) protoc가 생성한 클래스 import
import (
	"context"
    "log"
    "net"
    "fmt"
	"google.golang.org/grpc"
    function "full-stack-networking/noh-myeong-eun/hello-gRPC/myfunc"
	pb "full-stack-networking/noh-myeong-eun/hello-gRPC/protos"
)

// gRPC 서버 구조체 정의
type server struct {
    pb.UnimplementedMyServiceServer
}

// 원격 호출될 함수에 대한 서버용 핸들러 함수 작성
// request는 MyNumber 형태, 수행 결과도 MyNumber 형태
func (s *server) MyFunction(ctx context.Context, in *pb.MyNumber) (*pb.MyNumber, error) {
	// 원격 호출할 함수에 client로부터 받은 입력 파라미터를 전달하고 결과를 가져온다
    return &pb.MyNumber{Value: function.MyFunc(in.GetValue())},nil
	 
}

func main() {
    // 통신 포트 열기
    fmt.Printf("Starting server. Listening on port 50051.\n")
    lis, err := net.Listen("tcp", "localhost:50051")
    if err != nil {
        log.Fatalf("failed to listen %v", err)
    }

    // gRPC 서버를 생성한다
    grpcServer := grpc.NewServer()

    // gRPC 서버에 위에서 정의핸 핸들러를 추가한다
	pb.RegisterMyServiceServer(grpcServer, &server{})

    // 서버에서 클라이언트의 요청을 처리한다
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}