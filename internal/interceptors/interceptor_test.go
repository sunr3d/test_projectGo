package interceptors_test

import (
	"context"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus/testutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"

	"link_service/internal/interceptors"
	"link_service/internal/server/metrics"
)

func TestMetricsUnaryInterceptor(t *testing.T) {
	// Create the interceptor
	interceptor := interceptors.MetricsUnaryInterceptor()

	// Create a dummy handler
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		time.Sleep(100 * time.Millisecond) // Simulate some processing time
		return nil, status.Error(1, "test error")
	}

	// Call the interceptor
	_, err := interceptor(context.Background(), nil, &grpc.UnaryServerInfo{FullMethod: "/test/method"}, handler)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	// Check the metrics
	if count := testutil.ToFloat64(metrics.RequestCount.WithLabelValues("/test/method")); count != 1 {
		t.Errorf("expected 1 request, got %v", count)
	}
	if count := testutil.ToFloat64(metrics.ErrorCount.WithLabelValues("/test/method")); count != 1 {
		t.Errorf("expected 1 error, got %v", count)
	}
	if duration := testutil.CollectAndCount(metrics.RequestDuration); duration <= 0 {
		t.Errorf("expected positive duration, got %v", duration)
	}
}
