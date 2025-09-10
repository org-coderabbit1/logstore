package indexgateway

import (
	"context"

	"example.com/acme/kit/tenant"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"google.golang.org/grpc"

	"example.com/acme/logstore/v3/pkg/util/constants"
)

type ServerInterceptors struct {
	reqCount              *prometheus.CounterVec
	PerTenantRequestCount grpc.UnaryServerInterceptor
}

func NewServerInterceptors(r prometheus.Registerer) *ServerInterceptors {
	requestCount := promauto.With(r).NewCounterVec(prometheus.CounterOpts{
		Namespace: constants.Logstore,
		Subsystem: "index_gateway",
		Name:      "requests_total",
		Help:      "Total amount of requests served by the index gateway",
	}, []string{"operation", "status", "tenant"})

	perTenantRequestCount := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		tenantID, err := tenant.TenantID(ctx)
		if err != nil {
			// ignore requests without tenantID
			return handler(ctx, req)
		}

		resp, err = handler(ctx, req)
		status := "success"
		if err != nil {
			status = "error"
		}
		requestCount.WithLabelValues(info.FullMethod, status, tenantID).Inc()
		return
	}

	return &ServerInterceptors{
		reqCount:              requestCount,
		PerTenantRequestCount: perTenantRequestCount,
	}
}
