package util

import (
	"bytes"
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
)

type IMinioClient interface {
	Upload(bucketName, objectName string, content []byte) error
	GetObject(bucketName, objectName string) ([]byte, error)
	MakeBucket(bucketName string) error
}

type MinioClient struct {
	Client *minio.Client
}

func InitMinio() *minio.Client {
	endpoint := viper.Get("APP_MINIO_HOST").(string)
	accessKey := viper.Get("APP_MINIO_USERNAME").(string)
	secretKey := viper.Get("APP_MINIO_PASSWORD").(string)

	useSSL := false

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		fmt.Printf("Error creating MinIO client: %s\n", err)
	}
	return minioClient
}

func (Minio *MinioClient) MakeBucket(bucketName string) error {
	ctx := context.Background()
	exists, err := Minio.Client.BucketExists(ctx, bucketName)
	if err != nil {
		return err
	}
	if !exists {
		return Minio.Client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	}
	return nil
}

func (Minio *MinioClient) Upload(bucketName, objectName string, content []byte) error {
	ctx := context.Background()
	reader := bytes.NewReader(content)

	_, err := Minio.Client.PutObject(ctx, bucketName, objectName, reader, int64(len(content)), minio.PutObjectOptions{})
	return err
}

func (Minio *MinioClient) GetObject(bucketName, objectName string) ([]byte, error) {
	ctx := context.Background()
	object, err := Minio.Client.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	defer object.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(object)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
