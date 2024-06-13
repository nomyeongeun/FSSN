package main

// (1) grpc를 포함해서 필요한 모듈 import & (2) protoc가 생성한 클래스도 import
import (
	"context"
	"log"
	"time"
    "fmt"
	"google.golang.org/grpc"
	pb "full-stack-networking/noh-myeong-eun/hello-gRPC/protos"
)

const (
    address = "localhost:50051"
    value    = 4
)

func main() {
    // gRPC 서버에 연결한다
    conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()

    // gRPC 클라이언트를 생성한다
    c := pb.NewMyServiceClient(conn)
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    // 원격 함수를 호출한다
    reply, err := c.MyFunction(ctx, &pb.MyNumber{Value: value})
    if err != nil {
        log.Fatalf("MyFunction error: %v", err)
    }

    // 결과를 출력한다
    fmt.Printf("gRPC result: %v\n", reply.Value)
}