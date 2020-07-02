package minio

import (
	"os"

	"github.com/minio/minio-go/v6"
)

var minioClient *minio.Client

var (
	endpoint  string
	accessKey string
	secretKey string
	useSSL    bool
	bucket    string
)

func NewMinioHandler() (*minio.Client, string) {
	if minioClient != nil {
		return minioClient, bucket
	}

	endpoint = os.Getenv("ENDPOINT")
	accessKey = os.Getenv("ACCESSKEY")
	secretKey = os.Getenv("SECRETKEY")
	bucket = os.Getenv("BUCKET")
	useSSL = false
	minioClient, err := minio.New(endpoint, accessKey, secretKey, useSSL)

	if err != nil {
		panic(err)
	}
	return minioClient, bucket
}
