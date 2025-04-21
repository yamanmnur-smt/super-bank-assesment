package util_test

import (
	"testing"
	"yamanmnur/simple-dashboard/pkg/util"

	"github.com/spf13/viper"
)

func TestMinioUpload(t *testing.T) {
	viper.SetConfigFile("../../.env")
	viper.ReadInConfig()

	client := util.InitMinio()
	minio := util.MinioClient{Client: client}

	bucket := "test-bucket"
	err := minio.MakeBucket(bucket)
	if err != nil {
		t.Fatal(err)
	}

	err = minio.Upload(bucket, "test.txt", []byte("unit test content"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestMinioDownload(t *testing.T) {
	viper.SetConfigFile("../../.env")
	viper.ReadInConfig()
	client := util.InitMinio()

	minio := util.MinioClient{Client: client}

	bucket := "test-bucket"
	fileName := "download_test.txt"
	content := []byte("hello download!")

	minio.MakeBucket(bucket)

	err := minio.Upload(bucket, fileName, content)
	if err != nil {
		t.Fatal(err)
	}

	result, err := minio.GetObject(bucket, fileName)
	if err != nil {
		t.Fatal(err)
	}

	if string(result) != string(content) {
		t.Errorf("Expected '%s', got '%s'", content, result)
	}

}
