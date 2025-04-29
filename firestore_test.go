package store_test

import (
	"context"
	"testing"

	"cloud.google.com/go/firestore"

	"github.com/stretchr/testify/require"
)

func TestTimestampReadInTxn(t *testing.T) {
	client, err := firestore.NewClient(context.Background(), "test-db")
	require.NoError(t, err)
	defer client.Close()

	doc := client.Collection("jobs").Doc("43a9d36f-3451-4134-9d2e-6b97ed3cb9ef")
	result, err := doc.Create(context.Background(), map[string]string{
		"field1": "value1",
		"field2": "value2",
	})
	require.NoError(t, err)

	_, err = doc.Get(context.Background())
	require.NoError(t, err)

	_, err = doc.WithReadOptions(firestore.ReadTime(result.UpdateTime)).Get(context.Background())
	require.NoError(t, err)
}
