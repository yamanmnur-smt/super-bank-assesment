package injectors

import (
	"yamanmnur/simple-dashboard/pkg/util"

	"github.com/minio/minio-go/v7"
)

func InjectorMinioDI(minioClient *minio.Client, container *Container) {
	container.MinioClient = &util.MinioClient{Client: minioClient}
}
