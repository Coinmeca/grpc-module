package __

import context "context"

type GrpcCallBack interface {
	AddData(ctx context.Context, req *GeneralRequest) (*GeneralResponse, error)
}
