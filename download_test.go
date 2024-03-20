package main

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"io"
	"os"
	"testing"
)

func TestDownload(t *testing.T) {
	ctx := context.Background()
	client, err := NewClient()
	if err != nil {
		t.Error(err)
		return
	}

	bucketName := "test"
	fileName := "main.go"

	// 下载文件
	object, err := client.GetObject(ctx, bucketName, fileName, minio.GetObjectOptions{})
	if err != nil {
		t.Error(err)
		return
	}

	// 将文件流写入文件
	file, err := os.OpenFile(fmt.Sprintf("download_%s", fileName), os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		t.Error(err)
		return
	}

	defer file.Close()

	all, err := io.ReadAll(object)
	if err != nil {
		t.Error(err)
		return
	}

	if _, err := file.Write(all); err != nil {
		t.Error(err)
		return
	}
}
