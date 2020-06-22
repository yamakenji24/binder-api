package minio

import (
	"io"
	"net/url"
	"time"

	"github.com/minio/minio-go/v6"
)

func MinioUploader(fp string, fr io.Reader) error {
	mn, bucket := NewMinioHandler()

	_, err := mn.PutObject(bucket, fp, fr, -1, minio.PutObjectOptions{ContentType: "application/pdf"})
	if err != nil {
		return err
	}
	return nil
}

func GenerateURL(fp string) (*url.URL, error) {
	mn, bucket := NewMinioHandler()
	expiry := time.Second * 24 * 60 * 60
	reqParams := make(url.Values)

	geneURL, err := mn.PresignedGetObject(bucket, fp, expiry, reqParams)
	if err != nil {
		return nil, err
	}
	return geneURL, nil
}
