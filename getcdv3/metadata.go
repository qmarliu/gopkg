package getcdv3

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func WithRequestMateData() grpc.DialOption {
	return grpc.WithUnaryInterceptor(setMetadata)
}

func setMetadata(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) (err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		md = metadata.Pairs()
	}
	if operationID, ok := ctx.Value("operationID").(string); ok {
		md.Set("operationID", operationID)
	}
	if opUserID, ok := ctx.Value("opUserID").(string); ok {
		md.Set("opUserID", opUserID)
	}
	return invoker(metadata.NewOutgoingContext(ctx, md), method, req, reply, cc, opts...)
}
