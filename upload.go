package main

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"io"
	"log"
)

func Upload(client *minio.Client, ctx context.Context, bucketName, fileName string, reader io.Reader, fileSize int64) string {
	// Upload the file with PutObject
	info, err := client.PutObject(ctx, bucketName, fileName, reader, fileSize, minio.PutObjectOptions{})
	if err != nil {
		log.Fatalln(err)
		return "上传失败"
	}

	return fmt.Sprint(info.Bucket, "/", info.Key)
}
