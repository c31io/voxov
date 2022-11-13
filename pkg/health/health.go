// https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/#define-a-grpc-liveness-probe
package health

import (
	"context"
	"log"

	"google.golang.org/grpc/health/grpc_health_v1"
)

// Service field is not very useful in microservices.
type Health struct {
	Alive bool
	Dead  chan struct{}
}

func (h *Health) NowDead() {
	log.Println("NowDead")
	h.Alive = false
	h.Dead <- struct{}{}
}

func (h *Health) Check(ctx context.Context, _ *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	log.Println("Checking health")
	if h.Alive {
		log.Printf("Health status: %s", grpc_health_v1.HealthCheckResponse_SERVING)
		return &grpc_health_v1.HealthCheckResponse{
			Status: grpc_health_v1.HealthCheckResponse_SERVING,
		}, nil
	} else {
		log.Printf("Health status: %s", grpc_health_v1.HealthCheckResponse_NOT_SERVING)
		return &grpc_health_v1.HealthCheckResponse{
			Status: grpc_health_v1.HealthCheckResponse_NOT_SERVING,
		}, nil
	}
}

func (h *Health) Watch(in *grpc_health_v1.HealthCheckRequest, stream grpc_health_v1.Health_WatchServer) error {
	log.Println("Watching health")
	status, _ := h.Check(context.Background(), in)
	stream.Send(status)
	for range h.Dead {
	}
	log.Printf("Health status: %s", grpc_health_v1.HealthCheckResponse_NOT_SERVING)
	stream.Send(&grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_NOT_SERVING,
	})
	return nil
}
