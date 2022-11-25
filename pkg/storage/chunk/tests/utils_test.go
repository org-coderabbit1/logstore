package tests

import (
	"testing"

	"github.com/stretchr/testify/require"

	"example.com/acme/logstore/pkg/storage/chunk/client"
	"example.com/acme/logstore/pkg/storage/chunk/client/aws"
	"example.com/acme/logstore/pkg/storage/chunk/client/cassandra"
	"example.com/acme/logstore/pkg/storage/chunk/client/gcp"
	"example.com/acme/logstore/pkg/storage/chunk/client/local"
	"example.com/acme/logstore/pkg/storage/chunk/client/testutils"
	"example.com/acme/logstore/pkg/storage/stores/series/index"
)

const (
	userID    = "userID"
	tableName = "test"
)

type storageClientTest func(*testing.T, index.Client, client.Client)

func forAllFixtures(t *testing.T, storageClientTest storageClientTest) {
	var fixtures []testutils.Fixture
	fixtures = append(fixtures, aws.Fixtures...)
	fixtures = append(fixtures, gcp.Fixtures...)
	fixtures = append(fixtures, local.Fixtures...)
	fixtures = append(fixtures, cassandra.Fixtures()...)
	fixtures = append(fixtures, Fixtures...)

	for _, fixture := range fixtures {
		t.Run(fixture.Name(), func(t *testing.T) {
			indexClient, objectClient, closer, err := testutils.Setup(fixture, tableName)
			require.NoError(t, err)
			defer closer.Close()
			storageClientTest(t, indexClient, objectClient)
		})
	}
}
