package client

import (
	"context"

	"example.com/acme/logstore/v3/pkg/compactor/client/grpc"
	"example.com/acme/logstore/v3/pkg/compactor/deletion/deletionproto"
)

type CompactorClient interface {
	GetAllDeleteRequestsForUser(ctx context.Context, userID string) ([]deletionproto.DeleteRequest, error)
	GetCacheGenerationNumber(ctx context.Context, userID string) (string, error)

	JobQueueClient() grpc.JobQueueClient

	Name() string
	Stop()
}
