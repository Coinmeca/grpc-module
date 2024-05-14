package __

import (
	context "context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"
)

type server struct {
	UnimplementedCoinmecaGrpcModuleServer
	grpcCallBack GrpcCallBack
}

type GRPCClientManager struct {
	Client CoinmecaGrpcModuleClient
}

func NewServer(grpcCallBack GrpcCallBack) *server {
	return &server{
		grpcCallBack: grpcCallBack,
	}
}

func InitGrpcClient(address string) (CoinmecaGrpcModuleClient, *grpc.ClientConn) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to gRPC server: %v", err)
	}
	client := NewCoinmecaGrpcModuleClient(conn)
	return client, conn
}

func StartGrpcServerV2(grpcCallBack GrpcCallBack, port string) {
	go func() {
		lis, err := net.Listen("tcp", port)
		if err != nil {
			log.Printf("Failed to listen: %v", err)
			return
		}
		defer lis.Close()

		s := grpc.NewServer()
		RegisterCoinmecaGrpcModuleServer(s, NewServer(grpcCallBack))

		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

		go func() {
			<-stop
			log.Println("Shutting down gRPC server...")
			s.GracefulStop()
			lis.Close()
		}()

		log.Printf("gRPC server listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Printf("Failed to serve: %v", err) // 종료하지 않고 오류 로깅
		}
	}()
}

func StartGrpcServer(grpcCallBack GrpcCallBack, port string) {
	go func() {
		lis, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}
		defer lis.Close()

		s := grpc.NewServer()
		RegisterCoinmecaGrpcModuleServer(s, NewServer(grpcCallBack))

		log.Printf("gRPC server listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
}

func (s *server) Send(ctx context.Context, req *GeneralRequest) (*GeneralResponse, error) {

	return s.grpcCallBack.Send(ctx, req)
}
