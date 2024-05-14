package __

import context "context"

type GrpcCallBack interface {
	Send(ctx context.Context, req *GeneralRequest) (*GeneralResponse, error)
}
