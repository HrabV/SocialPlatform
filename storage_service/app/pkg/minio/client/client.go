package client

import (
	"context"
	"errors"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"socia/storage/pkg/utils"
)

type MinioClient struct {
	minio *minio.Client
}

func NewMinioClient(accessKey, secretKey, endpoint string, useSSL bool) (*MinioClient, error) {
	mc, err := minio.New(endpoint, &minio.Options{
		Secure: useSSL,
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
	})

	if !utils.MinioLiveCheck(endpoint, useSSL) {
		return nil, errors.New("server is unavailable")
	}

	return &MinioClient{
		minio: mc,
	}, err
}

func (mc *MinioClient) CreateBucket(ctx context.Context, bucketName string, options minio.MakeBucketOptions) error {
	return mc.minio.MakeBucket(ctx, bucketName, options)
}

func (mc *MinioClient) DeleteBucket(ctx context.Context, bucketName string) error {
	return mc.minio.RemoveBucket(ctx, bucketName)
}

func (mc *MinioClient) UploadFile(ctx context.Context, bucketName, fileName string, reader io.Reader, size int64, options minio.PutObjectOptions) (minio.UploadInfo, error) {
	exists, err := mc.minio.BucketExists(ctx, bucketName)
	if !exists {
		err = mc.CreateBucket(ctx, bucketName, minio.MakeBucketOptions{})
	}
	info, err := mc.minio.PutObject(ctx, bucketName, fileName, reader, size, options)
	return info, err
}

func (mc *MinioClient) GetFile(ctx context.Context, bucketName, fileName string, options minio.GetObjectOptions) (*minio.Object, error) {
	return mc.minio.GetObject(ctx, bucketName, fileName, options)
}

func (mc *MinioClient) DeleteFile(ctx context.Context, bucketName, fileName string, options minio.RemoveObjectOptions) error {
	return mc.minio.RemoveObject(ctx, bucketName, fileName, options)
}
