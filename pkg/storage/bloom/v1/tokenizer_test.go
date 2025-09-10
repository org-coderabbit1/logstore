package v1

import (
	"testing"

	"github.com/stretchr/testify/require"

	v2 "example.com/acme/logstore/v3/pkg/iter/v2"

	"example.com/acme/logstore/pkg/push"
)

func TestStructuredMetadataTokenizer(t *testing.T) {
	tokenizer := NewStructuredMetadataTokenizer("chunk")

	metadata := push.LabelAdapter{Name: "pod", Value: "logstore-1"}
	expected := []string{"pod", "chunkpod", "pod=logstore-1", "chunkpod=logstore-1"}

	tokenIter := tokenizer.Tokens(metadata)
	got, err := v2.Collect(tokenIter)
	require.NoError(t, err)
	require.Equal(t, expected, got)
}
