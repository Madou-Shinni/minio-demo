package main

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"io"
	"os"
	"testing"
)

// TestUpload 模拟文件流上传
func TestUpload(t *testing.T) {
	ctx := context.Background()
	client, err := NewClient()
	if err != nil {
		t.Error(err)
		return
	}

	bucketName := "test"
	fileName := "dapa.jpg"

	open, err := os.Open(fileName)
	if err != nil {
		t.Error(err)
		return
	}

	stat, err := open.Stat()
	if err != nil {
		t.Error(err)
		return
	}

	defer open.Close()

	// 模拟文件流上传
	reader := io.Reader(open)

	// Upload the file with PutObject
	info, err := client.PutObject(ctx, bucketName, fileName, reader, stat.Size(), minio.PutObjectOptions{})
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(info)
	t.Logf("上传成功 filePath: %s", fmt.Sprint(bucketName, "/", fileName))
}
