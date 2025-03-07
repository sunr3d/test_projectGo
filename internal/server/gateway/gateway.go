package gateway

import (
	"context"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pbls "link_service/proto/link_service"
)

type Gateway struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) *Gateway {
	return &Gateway{
		logger: logger,
	}
}

func (g *Gateway) Run(ctx context.Context, grpcAddress, httpAddress string) error {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	if err := pbls.RegisterLinkServiceHandlerFromEndpoint(ctx, mux, grpcAddress, opts); err != nil {
		return fmt.Errorf("gateway.Run, failed to register handler: %w", err)
	}

	g.logger.Info("HTTP server started", zap.String("address", httpAddress))

	if err := http.ListenAndServe(httpAddress, mux); err != nil {
		return fmt.Errorf("gateway.Run, failed server HTTP: %w", err)
	}

	return nil
}
