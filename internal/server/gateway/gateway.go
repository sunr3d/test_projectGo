package gateway

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

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

const ReadHeaderTimeoutDuration = 5 * time.Second

func (g *Gateway) Run(ctx context.Context, grpcAddress, httpAddress string) error {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	if err := pbls.RegisterLinkServiceHandlerFromEndpoint(ctx, mux, grpcAddress, opts); err != nil {
		return fmt.Errorf("gateway.Run, failed to register handler: %w", err)
	}

	server := &http.Server{
		Addr:              httpAddress,
		Handler:           mux,
		ReadHeaderTimeout: ReadHeaderTimeoutDuration,
	}

	go func() {
		g.logger.Info("gateway.Run: Gateway HTTP server started", zap.String("address", httpAddress))

		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			g.logger.Error("gateway.Run: ", zap.Error(err))
		}
	}()

	/// Блок остановки сервера по сигналу отмены контекста.
	<-ctx.Done()

	g.logger.Info("gateway.Run: context canceled")

	if err := server.Shutdown(ctx); err != nil {
		return fmt.Errorf("gateway.Run: failed to shutdown: %w", err)
	}

	g.logger.Info("gateway.Run: gateway server shutdown")

	return nil
}
