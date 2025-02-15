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

func InitGrpcClientV2(address string) (CoinmecaGrpcModuleClient, *grpc.ClientConn) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithDefaultCallOptions(
		grpc.MaxCallRecvMsgSize(1024*1024*16), // 16MB
		grpc.MaxCallSendMsgSize(1024*1024*16), // 16MB
	))

	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		log.Fatalf("Could not connect to gRPC server: %v", err)
	}
	client := NewCoinmecaGrpcModuleClient(conn)
	return client, conn
}

func InitGrpcClient(address string) (CoinmecaGrpcModuleClient, *grpc.ClientConn) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(1024*1024*16)))
	opts = append(opts, grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(1024*1024*16)))

	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		log.Fatalf("Could not connect to gRPC server: %v", err)
	}
	client := NewCoinmecaGrpcModuleClient(conn)
	return client, conn

	//conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	//if err != nil {
	//	log.Fatalf("Could not connect to gRPC server: %v", err)
	//}
	//client := NewCoinmecaGrpcModuleClient(conn)
	//return client, conn
}

func StartGrpcServerV2(grpcCallBack GrpcCallBack, port string) {
	go func() {
		listen, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatalf("Failed to listen : %v", err)
		}
		defer listen.Close()

		opts := []grpc.ServerOption{
			grpc.UnaryInterceptor(loggingInterceptor),
			grpc.MaxRecvMsgSize(1024 * 1024 * 16), // 16MB
			grpc.MaxSendMsgSize(1024 * 1024 * 16), // 16MB
		}

		s := grpc.NewServer(opts...)
		RegisterCoinmecaGrpcModuleServer(s, NewServer(grpcCallBack))

		log.Printf("gRPC server listening at %v", listen.Addr())
		if err := s.Serve(listen); err != nil {
			log.Fatalf("Failed to serve : %v", err)
		}
	}()
}

func loggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// request log
	log.Printf("Received request: %v", req)
	resp, err := handler(ctx, req)
	// answer log
	log.Printf("Sent response: %v", resp)
	return resp, err
}

func StartGrpcServer(grpcCallBack GrpcCallBack, port string) {
	go func() {
		lis, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}
		defer lis.Close()

		//s := grpc.NewServer(grpc.UnaryInterceptor(loggingInterceptor))
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
