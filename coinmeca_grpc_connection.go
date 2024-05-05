package __

import (
	context "context"
	"log"
	"net"

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

func StartGrpcServer(grpcCallBack GrpcCallBack, port string) {
	go func() {
		lis, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		defer lis.Close()

		s := grpc.NewServer()
		RegisterCoinmecaGrpcModuleServer(s, NewServer(grpcCallBack))

		log.Printf("gRPC server listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
}

func (s *server) AddData(ctx context.Context, req *GeneralRequest) (*GeneralResponse, error) {

	return s.grpcCallBack.AddData(ctx, req)
}
