package main

import (
	"context"
	"github.com/minio/minio-go/v7"
	"testing"
)

func TestCreateBucket(t *testing.T) {
	ctx := context.Background()
	client, err := NewClient()
	if err != nil {
		t.Error(err)
		return
	}

	bucketName := "test"

	err = client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	if err != nil {
		t.Error(err)
		return
	}
}
