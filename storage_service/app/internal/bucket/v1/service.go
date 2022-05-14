package v1

import (
	"context"
	"errors"
	"github.com/minio/minio-go/v7"
	"socia/storage/pkg/minio/client"
)

type BucketService struct {
	minioClient *client.MinioClient
}

func NewBucketService(minioClient *client.MinioClient) *BucketService {
	return &BucketService{
		minioClient: minioClient,
	}
}

func (bs *BucketService) CreateBucket(ctx context.Context, bucketName string, options minio.MakeBucketOptions) error {
	err := bs.minioClient.CreateBucket(ctx, bucketName, options)
	if err != nil {
		return errors.New("cannot create bucket")
	}

	return nil
}

func (bs *BucketService) DeleteBucket(ctx context.Context, bucketName string) error {
	err := bs.minioClient.DeleteBucket(ctx, bucketName)
	if err != nil {
		return errors.New("cannot delete bucket")

	}
	return nil
}
