package main

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

func NewClient() (*minio.Client, error) {
	// 创建 minio 客户端
	endpoint := "127.0.0.1:9000"
	id := "ricardo"
	secret := "uwilllikeit"
	SessionToken := ""
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(id, secret, SessionToken),
	})
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return minioClient, nil
}
